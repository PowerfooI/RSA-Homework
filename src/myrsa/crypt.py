from typing import Tuple
from myrsa.exp import quick_exp_mod
from myrsa.types import *


def encrypt_msg(key: PublicKey, msg: int) -> int:
    n, e = key.get_values()
    return quick_exp_mod(msg, e, n)


def decrypt_msg(key: PrivateKey, c_msg: int) -> int:
    n, d, p, q = key.get_values()
    return quick_exp_mod(c_msg, d, p*q)
