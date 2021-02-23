import time
from myrsa.gen_key import gen_keys
from myrsa.crypt import decrypt_msg, encrypt_msg

if __name__ == '__main__':
    # test for performance
    start = time.time()
    publicKey, privateKey = gen_keys(1024, 4)
    end = time.time()
    print('total time = {:.2f}s'.format(end-start))

    # test for encryption & decryption
    msg = 24234
    c_msg = encrypt_msg(publicKey, msg)
    print('encrypt msg = {}'.format(c_msg))
    d_msg = decrypt_msg(privateKey, c_msg)
    print('decrypt msg = {}'.format(d_msg))
    if msg != d_msg:
        print("Wrong results!")
