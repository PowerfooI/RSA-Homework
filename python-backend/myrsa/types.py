from typing import Tuple


class PublicKey:
    def __init__(self, size: int, n: int, e: int):
        self.size = size
        self.n_value = n
        self.e_value = e

    def __str__(self):
        return '[PublicKey Info]\nsize: {}\nn: {}\ne: {}\n'.format(self.size, hex(self.n_value), hex(self.e_value))

    def get_values(self) -> Tuple[int, int]:
        return self.n_value, self.e_value

    def get_size(self) -> int:
        return self.size


class PrivateKey:
    def __init__(self, size: int, n: int, d: int, p: int, q: int):
        self.size = size
        self.n_value = n
        self.d_value = d
        self.p_value = p
        self.q_value = q

    def __str__(self):
        return '[PrivateKey Info]\nsize: {}\nn: {}\nd: {}\np: {}\nq: {}\n'.format(self.size, hex(self.n_value),
                                                                                  hex(self.d_value),
                                                                                  hex(self.p_value),
                                                                                  hex(self.q_value))

    def get_values(self) -> Tuple[int, int, int, int]:
        return self.n_value, self.d_value, self.p_value, self.q_value

    def get_size(self) -> int:
        return self.size
