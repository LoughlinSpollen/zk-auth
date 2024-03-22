#!/usr/bin/env python3.7
# coding=utf-8

import logging
import random
import sys
import os
import traceback

import usecase.auth_usecase as usecase
LOGGER = logging.getLogger()

def client():
    LOGGER.info("zk_auth client")
    try:
        authUsecase = usecase.ZKAuthUsecase()
        userId = "user"
        password = random.randint(1,1000) 
        authUsecase.register(userId, password)            
        sessionId = authUsecase.auth(userId)        
        if sessionId != None:
            LOGGER.info("Successful authentication with sessionId: " + sessionId)
    except :
        LOGGER.error("Exception running client" + traceback.format_exc())


if __name__ == '__main__':
    logging.basicConfig(
        level=logging.DEBUG if os.getenv('DEBUG') == 'True' else logging.INFO,
        format="%(asctime)s [%(threadName)-12.12s] [%(levelname)-5.5s]  %(message)s",
        handlers=[
            logging.StreamHandler(sys.stdout)
        ])
    client()
