import logging
import random

LOGGER = logging.getLogger()

class ChaumPedersenData():
    def __init__(self, data: object):
        LOGGER.debug("ChaumPedersenData init ")
        self._data = data
        self._data.A = pow(data.g, data.a, data.q) # g^a 
        self._data.B = pow(data.g, data.b, data.q) # g^b 
        self._data.C = pow(data.g, (data.a * data.b), data.q) # g^ab
        LOGGER.debug("Public agree (g,g^a, g^b and g^ab) \
                     =(%d,%d,%d,%d)", 
                     self._data.g,
                     self._data.A,
                     self._data.B,
                     self._data.C)

class ChaumPedersenProver(ChaumPedersenData):

    def __init__(self, data: object):
        super().__init__(data=data)

    def register(self, r: int) -> int:
        LOGGER.debug("ChaumPedersenProver register")        
        LOGGER.debug("Use Password (r) %d",r)
        self._data.r = r
        y1 = pow(self._data.g, self._data.r, self._data.q) # g^r 
        y2 = pow(self._data.B, self._data.r, self._data.q) # (g^b)^r
        LOGGER.debug("Prover sends y1 y2 (g^r, B^r)=(%d,%d)", y1, y2)
        return y1, y2
    
    def response(self, s: int) -> int:
        LOGGER.debug("ChaumPedersenProver response")                
        z = (self._data.r + self._data.a * s) % self._data.q # (r + a*s)
        LOGGER.debug("Prover computes z=r+as (mod q)=%d", z)
        return z
    

class ChaumPedersenVerifier(ChaumPedersenData):
    def __init__(self, data: object):
        super().__init__(data=data)
        # for verification
        if hasattr(data, 'r'):
            self._data.y1 = pow(self._data.g, data.r, self._data.q)
            self._data.y2 = pow(self._data.B, data.r, self._data.q)

    def challenge(self) -> int:  
        LOGGER.debug("ChaumPedersenVerifier challenge")      
        self._data.s = random.randint(1, 1000)
        LOGGER.debug("Verifier sends a challenge (s)=%d", self._data.s)
        return self._data.s

    def verify(self, z: int) -> bool:
        LOGGER.debug("ChaumPedersenVerifier verify")        
        LOGGER.debug("Verifier makes two checks using the answer (z)=%d against the challenge s", z)

        a1 = pow(self._data.g, z, self._data.q)
        LOGGER.debug("Verifier calculates a1 g^z mod q=%d", a1)

        a2 = (self._data.A**self._data.s * self._data.y1) % self._data.q
        LOGGER.debug("Verifier calculates a2: A^s y1=%d",a2)
        if a1 != a2:
            return False

        LOGGER.debug("Verifier now checks these are the same")
        LOGGER.debug("Verifier calculates B^z=%d", pow(self._data.B, z, self._data.q))
        b1 = pow(self._data.B, z, self._data.q)

        b2 = (self._data.C**self._data.s * self._data.y2) % self._data.q        
        LOGGER.debug("Verifier calculates C^s y2=%d", b2)
        if b1 != b2:
            return False
        
        return True
    

        
