import math 
import random 
import exp

prime_nums_lt_10k = [2,3,5,7,11,13,17,19,23,29]

for i in range(30, 10240):
    flag = True
    for j in range(2, int(math.sqrt(i))+1):
        if i % j == 0:
            flag = False
            break 
    if flag:
        prime_nums_lt_10k.append(i)


# print(prime_nums_lt_10k)

def produce_pq(n_size=512):
    p = random_big_prime_num(n_size//2)
    q = random_big_prime_num(n_size//2, p)
    return p, q 

def random_big_prime_num(size, skip_value=0):
    p = random.randint(2**(size//2)+1, 2**size-1)
    if p % 2 == 0:
        p += 1
    while True:
        if p == skip_value:
            continue
        # pre-check 
        for n in prime_nums_lt_10k:
            if p % n == 0:
                continue 
        # Miller-Rabin test 
        n_minus = p - 1
        s = 0
        while n_minus % 2 == 0:
            s += 1
            n_minus //= 2
        d = n_minus
        outer_loop_flag = True
        # 检验次数越多，概率越大
        for _ in range(3):
            a = random.randint(2, p-1)
            for r in range(s):
                r1 = exp.quick_exp_mod(a, d, p)
                r2 = exp.quick_exp_mod(a, exp.quick_exp(2, r)*d, p)
                if r1 != 1 and r2 != 1:
                    outer_loop_flag = False 
                    break 
            if not outer_loop_flag:
                break
        if not outer_loop_flag:
            p += 2
            continue 
        else:
            return p 
