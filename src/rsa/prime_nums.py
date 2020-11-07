import math
import random
import exp
import threading

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

pq_arr = [0, 0, 0, 0, 0, 0, 0, 0]
# print(prime_nums_lt_10k)

def check_pq():
    if pq_arr[0] == pq_arr[1]:
        return False
    return True

def miller_rabin_test(s, p, d, times=3):
    for _ in range(times):
        a = prime_nums_lt_10k[random.randint(0, len_of_pretest-1)]
        for r in range(s):
            r1 = exp.quick_exp_mod(a, d, p)
            r2 = exp.quick_exp_mod(a, exp.quick_exp(2, r) * d, p)
            if r1 != 1 and r2 != 1:
                return False
    return True
    

def produce_pq(n_size=768):
    # threads = []
    # for i in range(2):
    #     new_thread = threading.Thread(target=random_big_prime_num(n_size // 2, i))
    #     new_thread.start()
    #     threads.append(new_thread)

    # for i in range(2):
    #     threads[i].join()
    
    p_thread = threading.Thread(target=random_big_prime_num(n_size//2, 0))
    q_thread = threading.Thread(target=random_big_prime_num(n_size//2, 1))
    # random_big_prime_num(n_size//2, 1)
    p_thread.start()
    q_thread.start()
    p_thread.join()
    q_thread.join()

    if not check_pq():
        print('whoops, p is equal to q!')
        random_big_prime_num(n_size // 2, 1)
    
    return pq_arr[0], pq_arr[1]

def random_big_prime_num(size, pos):
    print('begin compute at pos {}'.format(pos))
    p = random.randint(2 ** (size // 2) + 1, 2 ** size - 1)
    if p % 2 == 0:
        p += 1
    while True:
        # print('get one random p')
        # pre-check 
        pre_check_flag = True 
        for n in prime_nums_lt_10k:
            if p % n == 0:
                pre_check_flag = False
                break
                # Miller-Rabin test
        if not pre_check_flag:
            p+=2
            continue
        n_minus = p - 1
        s = 0
        while n_minus % 2 == 0:
            s += 1
            n_minus //= 2
        d = n_minus
        # 检验次数越多，概率越大
        if not miller_rabin_test(s, p, d):
            p+=2
            continue
        else:
            pq_arr[pos] = p
            return 

print(produce_pq())