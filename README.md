# RSA Midterm Assignment of *Applied Cryptography*

### Implementation Details

* Modular Exponentiation - Montgomery Modular Exponentiation Algorithm
* Prime Number Generation - Miller-Rabin Primality Test
* Padding Method - PKCS #1 v1.5

![Frontend page](/images/frontend.png)

### Functionality Description

* Provides key generation for 256, 512, 768, 1024, and 2048-bit RSA keys, with the ability to copy the keys with one click.
* Public key encryption + private key decryption, private key encryption + public key decryption, public key signing + private key verification, and private key signing + public key verification.


### Structure

- `backend`: Golang backend with complete functionality
- `python-backend`: Python backend, only implements key generation and encryption/decryption
- `frontend`: Vue frontend
