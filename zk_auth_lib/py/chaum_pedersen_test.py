#!/usr/bin/env python3.7
# coding=utf-8

import random
import unittest

import logging
import sys
import os
LOGGER = logging.getLogger()

from chaum_pedersen import ChaumPedersenData, ChaumPedersenProver, ChaumPedersenVerifier



class ChaumPedersenTest(unittest.TestCase):
    def __init__(self, methodName: str = "runTest") -> None:
        super().__init__(methodName)
        logging.basicConfig(
            level=logging.DEBUG if os.getenv('DEBUG') == 'True' else logging.INFO,
            format="%(asctime)s [%(threadName)-12.12s] [%(levelname)-5.5s]  %(message)s",
            handlers=[
                logging.StreamHandler(sys.stdout)
            ])
        LOGGER.info("ChaumPedersenTest")


    def test_functional(self):        
        T = type("", (), {})()
        T.q = 10003
        T.g = 3
        T.a = 10
        T.b = 13
        p = ChaumPedersenProver(T)

        # r is not shared with the prover on initialization
        r = random.randint(1,1000)
        T.r = r
        v = ChaumPedersenVerifier(T)

        p.register(r)
        s = v.challenge()
        z = p.response(s)
        res = v.verify(z)
        
        self.assertEqual(res, True)


