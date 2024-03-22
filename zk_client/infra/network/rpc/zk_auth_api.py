import os
import logging
import grpc

import build.protos.zk_auth.zk_auth_api_pb2 as zk_auth
import build.protos.zk_auth.zk_auth_api_pb2_grpc as service
from infra.network.rpc.conn_interceptor import ConnInterceptor, ExponentialBackoff

ZK_AUTH_SERVICE_HOST = os.getenv("ZK_AUTH_SERVICE_HOST", default="0.0.0.0")
ZK_AUTH_SERVICE_PORT = os.getenv("ZK_AUTH_SERVICE_PORT", default="1025")
ZK_AUTH_SERVICE_INITIAL_BACKOFF = os.getenv("ZK_AUTH_SERVICE_INITIAL_BACKOFF", default="100")
ZK_AUTH_SERVICE_MAX_BACKOFF = os.getenv("ZK_AUTH_SERVICE_MAX_BACKOFF", default="2000")
ZK_AUTH_SERVICE_BACKOFF_MULTIPLIER = os.getenv("ZK_AUTH_SERVICE_BACKOFF_MULTIPLIER", default="2")

LOGGER = logging.getLogger()

class ZKAuthAPI():
    def __init__(self):
        LOGGER.debug("ZKAuthAPI")

        # The underlying http/2 connection will automatically retry when
        # the connection is closed or dead. This adds an additional gRPC retry mechanism.
        # This is useful for cloud deployments when client/server orchestration is not sync'd.
        # Each of the retry calls (including the initial one) has a http/2 retry deadline. 
        # The client will retry 24 times with an exponential backoff policy.
        interceptors = (
            ConnInterceptor(
                max_attempts=24,
                sleeping_policy=ExponentialBackoff(initial_backoff_ms=int(ZK_AUTH_SERVICE_INITIAL_BACKOFF),
                                                   max_backoff_ms=int(ZK_AUTH_SERVICE_MAX_BACKOFF),
                                                   multiplier=int(ZK_AUTH_SERVICE_BACKOFF_MULTIPLIER)),
                status_for_retry=(grpc.StatusCode.UNAVAILABLE,),
            ),
        )

        channel = grpc.insecure_channel(ZK_AUTH_SERVICE_HOST + ':' + ZK_AUTH_SERVICE_PORT)
        self._intercept_channel = grpc.intercept_channel(channel, *interceptors)
        self._client = service.AuthStub(self._intercept_channel)

    def register(self, user: str, y1: int, y2: int):
        LOGGER.debug("ZKAuthAPI Register")
        self._client.Register(zk_auth.RegisterRequest(
            user=user,
            y1=y1,
            y2=y2
        ))
    
    def challenge(self, user: str, r1: int, r2: int):
        LOGGER.debug("ZKAuthAPI Authentication Challenge")
        response = self._client.CreateAuthenticationChallenge(zk_auth.AuthenticationChallengeRequest(
            user=user,
            r1=r1,
            r2=r2
        ))
        return response.auth_id, response.c

    def answer(self, auth_id: str, s: int) -> str:
        LOGGER.debug("ZKAuthAPI Authentication Answer")
        response = self._client.VerifyAuthentication(zk_auth.AuthenticationAnswerRequest(
            auth_id=auth_id,
            s=s
        ))
        return response.session_id



