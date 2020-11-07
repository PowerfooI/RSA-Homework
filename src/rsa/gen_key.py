from gcd import * 
from prime_nums import *
from exp import * 

p, q = produce_pq(768)

n = p * q 
phi_n = (p-1)*(q-1)

e = 65537

d, y, _ = ext_euclid(e, phi_n)
if d < 0:
    d += phi_n

print("d = {}, y = {}".format(d, y))

m = 3127

# c = (m ** e) % n 
c = quick_exp_mod(m, e, n)
print("C = {}".format(c))

# md = (c ** d) % n 
md = quick_exp_mod(c, d, n)
print("D = {}".format(md))
