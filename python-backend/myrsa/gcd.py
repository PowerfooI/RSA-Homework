from typing import Tuple


def gcd(a: int, b: int) -> int:
    if b == 0:
        return a
    else:
        return gcd(b, a % b)


def b_gcd(a: int, b: int) -> int:
    ''' Balanced Euclidean algorithm
    '''
    if b == 0:
        return a
    else:
        r = a % b
        if 2 * r > b:
            r = b - r
        return b_gcd(b, r)


def ext_euclid(a: int, b: int) -> Tuple[int, int, int]:
    ''' Extend Euclid algorithm
        return (u, v, q) s.t. ua + bv = q = gcd(a,b)
    '''
    if b == 0:
        return 1, 0, a
    else:
        u, v, q = ext_euclid(b, a % b)
        u, v = v, (u - (a // b) * v)
        return u, v, q
