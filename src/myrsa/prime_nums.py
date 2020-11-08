import math
import random
import threading
from myrsa.exp import quick_exp, quick_exp_mod
from queue import Queue
from typing import Tuple

prime_nums_lt_10k = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]

for i in range(30, 10240):
    flag = True
    for j in range(2, int(math.sqrt(i)) + 1):
        if i % j == 0:
            flag = False
            break
    if flag:
        prime_nums_lt_10k.append(i)

len_of_pretest = len(prime_nums_lt_10k)


def miller_rabin_test(p: int, p_bit_size: int) -> bool:
    times = 10
    if p_bit_size >= 1024:
        times = 3
    elif p_bit_size >= 512:
        times = 4
    elif p_bit_size >= 384:
        times = 5
    elif p_bit_size >= 256:
        times = 7
    n_minus = p - 1
    s = 0
    while n_minus % 2 == 0:
        s += 1
        n_minus //= 2
    d = n_minus

    for _ in range(times):
        # a = prime_nums_lt_10k[random.randint(0, len_of_pretest - 1)]
        a = random.randint(1023, p-1)
        r1 = quick_exp_mod(a, d, p)
        r_count = 0
        for r in range(s):
            r2 = quick_exp_mod(a, quick_exp(2, r) * d, p)
            if r1 != 1 and r2 != p-1:
                r_count += 1
            else:
                break
        if r_count == s:
            return False
        continue
    return True


def produce_pq(n_bit_size: int, n_threads: int) -> Tuple[int, int]:
    res_queue = Queue()
    first_prime = None
    for idx in range(n_threads):
        new_thread = threading.Thread(target=random_big_prime_num, args=(n_bit_size // 2, res_queue))
        new_thread.setDaemon(True)
        new_thread.start()

    while True:
        p = res_queue.get()
        if first_prime is not None:
            if p != first_prime:
                res_queue.task_done()
                return first_prime, p
        else:
            first_prime = p
        res_queue.task_done()


def random_big_prime_num(p_bit_size: int, res_queue: Queue) -> None:
    while True:
        p = random.randint(2 ** (p_bit_size // 2) + 1, 2 ** p_bit_size - 1)
        if p % 2 == 0:
            p += 1
        while True:
            # pre-check
            pre_check_flag = True
            for n in prime_nums_lt_10k:
                if p % n == 0:
                    pre_check_flag = False
                    break
            if not pre_check_flag:
                p += 2
                continue
            # Miller-Rabin test
            # 检验次数越多，概率越大
            if not miller_rabin_test(p, p_bit_size):
                p += 2
                continue
            else:
                res_queue.put(p)
