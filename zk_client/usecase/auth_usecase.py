#!/usr/bin/env python3.7
# coding=utf-8

import logging
import os

import sys
sys.path.insert(0, '..')
import zk_auth_lib.py.chaum_pedersen as cp

import zk_client.infra.network.rpc.zk_auth_api as rpc

Q = os.getenv("ZK_AUTH_PRIME", default="10009") # import prime
G = os.getenv("ZK_AUTH_G", default="3") # import key g
A = os.getenv("ZK_AUTH_A", default="10") # import key a
B = os.getenv("ZK_AUTH_B", default="13") # import key b

LOGGER = logging.getLogger()

class ZKAuthUsecase():
    def __init__(self):
        LOGGER.debug("ZKAuthUsecase")
        data = type("", (), {
            'q': int(Q),
            'g': int(G),
            'a': int(A),
            'b': int(B)
        })()
        self._verifier_service = rpc.ZKAuthAPI()
        self._prover = cp.ChaumPedersenProver(data)


    def register(self, userId: str, password: int):
        LOGGER.debug("Registering user")
        y1, y2 = self._prover.register(password)
        self._verifier_service.register(userId, y1, y2)


    def auth(self, userId: str) -> str:
        LOGGER.debug("Authenticating user")

        authId, challenge = self._verifier_service.challenge(userId, 0, 0) # 0, 0 are not used
        ans = self._prover.response(challenge)
        sessionId = self._verifier_service.answer(authId, ans)
        
        LOGGER.debug("SessionId: " + sessionId)
        return sessionId
        

