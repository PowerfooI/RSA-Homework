def quick_exp(base, exponent):
    ''' Montgomery reduction algorithm
        Compute the base ^ exponent with complexity O(logn)
    '''
    res = 1 
    while exponent > 0:
        if exponent % 2 == 1:
            res *= base
        base *= base
        exponent //= 2 
    return res 

def quick_exp_mod(base, exponent, modulo):
    ''' Montgomery reduction algorithm
        Compute the (base ^ exponent) % modulo with complexity O(logn)
    '''
    res = 1 
    while exponent > 0:
        if exponent % 2 == 1:
            res = res * base % modulo
        base = base * base % modulo 
        exponent //= 2 
    return res 