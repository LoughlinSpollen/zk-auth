from abc import ABC, abstractmethod
from random import randint
from typing import Optional, Tuple
import grpc
import time
import logging

LOGGER = logging.getLogger()

class SleepingPolicy(ABC):
    @abstractmethod
    def sleep(self, retry_count: int):
        assert retry_count >= 0

class ExponentialBackoff(SleepingPolicy):
    def __init__(self, *, initial_backoff_ms: int, max_backoff_ms: int, multiplier: int):
        LOGGER.debug("ExponentialBackoff init")
        self._initial_backoff = randint(0, initial_backoff_ms)
        self._max_backoff = max_backoff_ms
        self._multiplier = multiplier

    def sleep(self, retry_count: int):
        LOGGER.debug("ExponentialBackoff sleep")
        sleep_range = min(
            self._initial_backoff * self._multiplier ** retry_count, self._max_backoff
        )
        sleep_ms = randint(0, sleep_range)
        LOGGER.debug("Sleeping for %d", sleep_ms)
        time.sleep(sleep_ms / 1000)

class ConnInterceptor(grpc.UnaryUnaryClientInterceptor,
                      grpc.StreamUnaryClientInterceptor):
    def __init__(self, *, max_attempts: int, sleeping_policy: SleepingPolicy,
                 status_for_retry: Optional[Tuple[grpc.StatusCode]] = None):
        LOGGER.debug("ConnInterceptor init")
        self._max_attempts = max_attempts
        self._sleeping_policy = sleeping_policy
        self._status_for_retry = status_for_retry

    def _intercept_call(self, continuation, client_call_details, request_or_iterator):
        for retry_count in range(self._max_attempts):
            response = continuation(client_call_details, request_or_iterator)
            if isinstance(response, grpc.RpcError):
                # last attempt
                if retry_count == (self._max_attempts - 1):
                    LOGGER.warning("ConnInterceptor intercept_stream_unary: retry connection attempts " + str(retry_count))
                    return response
                # not in retryable status codes
                if (self._status_for_retry and response.code() not in self._status_for_retry):
                    return response
                self._sleeping_policy.sleep(retry_count)
            else:
                return response

    def intercept_unary_unary(self, continuation, client_call_details, request):
        return self._intercept_call(continuation, client_call_details, request)

    def intercept_stream_unary(self, continuation, client_call_details, request_iterator):
        return self._intercept_call(continuation, client_call_details, request_iterator)
