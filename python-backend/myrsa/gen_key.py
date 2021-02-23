from typing import Tuple
from myrsa.gcd import ext_euclid
from myrsa.prime_nums import produce_pq
from myrsa.types import *


def gen_keys(n_bit_size: int, n_threads: int = 4) -> Tuple[PublicKey, PrivateKey]:
    p, q = produce_pq(n_bit_size, n_threads)
    n = p * q
    phi_n = (p - 1) * (q - 1)
    e = 65537
    d, _, _ = ext_euclid(e, phi_n)
    if d < 0:
        d += phi_n

    return PublicKey(n_bit_size//2, n, e), PrivateKey(n_bit_size//2, n, d, p, q)
