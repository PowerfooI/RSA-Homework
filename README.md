# 《应用密码学》 RSA 期中作业

### 实现说明

* 幂模运算 - Montgomery 幂模算法
* 素数生成 - Miller-Rabin 素性检测
* 填充方法 - PKCS #1 v1.5

![前端页面](/images/frontend.png)

### 目录结构

- `backend`: golang 后端，功能完整
- `python-backend`: python 后端，只实现到密钥生成和加解密
- `frontend`: vue 前端

### 功能说明

* 提供了 256，512，768，1024 和 2048 位 RSA 密钥的生成，能够一键复制；
* 除了密钥的生成，还提供公钥加密+私钥解密、私钥加密+公钥解密、公钥签名+私钥验证和私钥签名+公钥验证这些功能。

