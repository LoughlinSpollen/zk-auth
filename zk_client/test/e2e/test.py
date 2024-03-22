#!/usr/bin/env python3.7
# coding=utf-8

import random
import sys
import unittest
import usecase.auth_usecase as usecase

class ZKClientTestCase(unittest.TestCase):

    def test_zk_auth_e2e(self):
        try: 
            password = random.randint(1,1000)
            userId = "e2e test"
            authUsecase = usecase.ZKAuthUsecase()
            authUsecase.register(userId, password)            
            sessionId = authUsecase.auth(userId)
            self.assertIsNotNone(sessionId)
            
        except Exception:
            print(sys.exc_info())

if __name__ == "__main__":
    unittest.main()
