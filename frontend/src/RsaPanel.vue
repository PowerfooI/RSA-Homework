<template>
  <div>
    <div class="title">RSA工具栏</div>
    <el-row>
      <el-col :span="10"
              :offset="1">
        <el-row class="panel-col">
          <el-col class="op-unit">
            <el-form label-position="left"
                     label-width="120px"
                     :model="keysForm">
              <el-form-item label="密钥位数">
                <el-input v-model="keysForm.n_len"
                          disabled></el-input>
              </el-form-item>
              <el-form-item label="密钥N值">
                <el-input v-model="keysForm.n_val"
                          type="textarea"
                          :autosize="{ minRows: 2}"
                          disabled></el-input>
              </el-form-item>
              <el-form-item label="密钥P值">
                <el-input v-model="keysForm.p_val"
                          type="textarea"
                          :autosize="{ minRows: 2}"
                          disabled></el-input>
              </el-form-item>
              <el-form-item label="密钥Q值">
                <el-input v-model="keysForm.q_val"
                          type="textarea"
                          :autosize="{ minRows: 2}"
                          disabled></el-input>
              </el-form-item>
              <el-form-item label="公钥E值">
                <el-input v-model="keysForm.e_val"
                          type="textarea"
                          :autosize="{ minRows: 2}"
                          disabled></el-input>
              </el-form-item>
              <el-form-item label="私钥D值">
                <el-input v-model="keysForm.d_val"
                          type="textarea"
                          :autosize="{ minRows: 2}"
                          disabled></el-input>
              </el-form-item>
            </el-form>
          </el-col>
          <el-col class="op-unit">
            <el-form label-position="left"
                     label-width="120px"
                     :model="gen_key_form">
              <el-form-item label="生成密钥位数">
                <el-select v-model="gen_key_form.n_len"
                           class="gen-select"
                           placeholder="请选择生成密钥的位数">
                  <el-option
                    v-for="item in gen_options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                  </el-option>
                </el-select>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="primary"
                           @click="onRequestGenKey">生成密钥
                </el-button>
                <el-button type="success"
                           @click="onCopy"
                >复制密钥
                </el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </el-col>
      <el-col :span="10"
              :offset="2">
        <el-row class="panel-col">
          <el-col class="op-unit">
            <el-form label-position="left"
                     label-width="120px"
                     :model="gen_key_form">
              <el-form-item label="明文消息">
                <el-input v-model="msgForm.msg"
                          type="textarea"
                          :autosize="{ minRows: 3}"
                          placeholder="请在此输入需要加密/签名的明文消息">
                </el-input>
              </el-form-item>
              <el-form-item label="加密消息">
                <el-input v-model="msgForm.encodedMsg"
                          type="textarea"
                          disabled
                          :autosize="{ minRows: 3}"
                          placeholder="加密后的消息">
                </el-input>
              </el-form-item>
              <el-form-item label="解密消息">
                <el-input v-model="msgForm.decodedMsg"
                          type="textarea"
                          disabled
                          :autosize="{ minRows: 3}"
                          placeholder="解密后的消息">
                </el-input>
              </el-form-item>
              <el-form-item label="数字签名">
                <el-input v-model="msgForm.signature"
                          type="textarea"
                          disabled
                          :autosize="{ minRows: 2}"
                          placeholder="数字签名">
                </el-input>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="primary"
                           @click="encodeWithPublicKey">公钥加密
                </el-button>
                <el-button type="primary"
                           @click="decodeWithPrivateKey">私钥解密
                </el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="success"
                           @click="encodeWithPrivateKey">私钥加密
                </el-button>
                <el-button type="success"
                           @click="decodeWithPublicKey">公钥解密
                </el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="info"
                           @click="signWithPublicKey">公钥签名
                </el-button>
                <el-button type="info"
                           @click="verifyWithPrivateKey">私钥验证
                </el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="warning"
                           @click="signWithPrivateKey">私钥签名
                </el-button>
                <el-button type="warning"
                           @click="verifyWithPublicKey">公钥验证
                </el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </el-col>
    </el-row>

  </div>

</template>

<script>
  const baseUrl = 'http://localhost:8080/';
  // import qs from 'qs'
  export default {
    name: "RsaPanel",
    computed: {},
    data() {
      return {
        keyStr: '',
        keysForm: {
          n_len: 0,
          n_val: '',
          e_val: '',
          d_val: '',
          p_val: '',
          q_val: '',
        },
        msgForm: {
          msg: '',
          encodedMsg: '',
          decodedMsg: '',
          signature: '',
        },
        gen_key_form: {
          n_len: 256,
        },
        gen_options: [
          {
            value: 256,
            label: '256位',
          }, {
            value: 512,
            label: '512位',
          }, {
            value: 768,
            label: '768位',
          }, {
            value: 1024,
            label: '1024位',
          }, {
            value: 2048,
            label: '2048位',
          },
        ],
      };
    },
    methods: {
      onRequestGenKey() {
        this.$http.get(baseUrl + `keys?n_byte=${this.gen_key_form.n_len}`)
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              Object.assign(this.keysForm, res.data.privateKey);
              this.keysForm.e_val = res.data.publicKey.e_val;
              this.$message.success(`密钥对生成成功，用时 ${res.data.costTime} ms`);
              this.keyStr = `n_len=${this.keysForm.n_len}\nn_val=${this.keysForm.n_val}\ne_val=${this.keysForm.e_val}\nd_val=${this.keysForm.d_val}\np_val=${this.keysForm.p_val}\nq_val=${this.keysForm.q_val}`;
            } else {
              this.$message.warning('生成密钥失败，不知道原因是什么');
            }
          })
          .catch(err => {
            console.log(err);
          });
      },
      encodeWithPublicKey() {
        this.$http.post(baseUrl + 'encode', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            e_val: this.keysForm.e_val,
          },
          msg: this.msgForm.msg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success("使用公钥加密成功");
              this.msgForm.encodedMsg = res.data.encodedMsg;
            } else {
              this.$message.error('加密失败');
            }
          })
          .catch(err => {
            this.$message.error(err);
          });
      },
      encodeWithPrivateKey() {
        this.$http.post(baseUrl + 'encode', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            d_val: this.keysForm.d_val,
            p_val: this.keysForm.p_val,
            q_val: this.keysForm.q_val,
          },
          msg: this.msgForm.msg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success("使用私钥加密成功");
              this.msgForm.encodedMsg = res.data.encodedMsg;
            } else {
              this.$message.error('加密失败');
            }
          })
          .catch(err => {
            this.$message.error(err);
          });
      },
      decodeWithPublicKey() {
        this.$http.post(baseUrl + 'decode', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            e_val: this.keysForm.e_val,
          },
          msg: this.msgForm.encodedMsg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success("使用公钥解密成功");
              this.msgForm.decodedMsg = res.data.decodedMsg;
            } else {
              this.$message.error('解密失败');
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.error('解密失败，请使用配对密钥进行解密');
          });
      },
      decodeWithPrivateKey() {
        this.$http.post(baseUrl + 'decode', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            d_val: this.keysForm.d_val,
            p_val: this.keysForm.p_val,
            q_val: this.keysForm.q_val,
          },
          msg: this.msgForm.encodedMsg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success("使用私钥解密成功");
              this.msgForm.decodedMsg = res.data.decodedMsg;
            } else {
              this.$message.error('解密失败');
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.error('解密失败，请使用配对密钥进行解密');
          });
      },
      signWithPublicKey() {
        this.$http.post(baseUrl + 'sign', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            e_val: this.keysForm.e_val,
          },
          msg: this.msgForm.msg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success('使用公钥签名成功');
              this.msgForm.signature = res.data.signature;
            } else {
              console.log('error encountered');
              this.$message.success('签名错误');
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.success('签名错误');
          });
      },
      verifyWithPrivateKey() {
        this.$http.post(baseUrl + 'verify', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            d_val: this.keysForm.d_val,
            p_val: this.keysForm.p_val,
            q_val: this.keysForm.q_val,
          },
          msg: this.msgForm.msg,
          signature: this.msgForm.signature,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              if (res.data.verified) {
                this.$message.success('验证成功！签名与文本正确对应！');
              } else {
                this.$message.warning('验证失败！消息已经被篡改！');
              }
            } else {
              this.$message.error("验证错误");
              console.log('error encountered');
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.error("验证错误，请使用配对密钥进行签名验证！");
          });
      },
      signWithPrivateKey() {
        this.$http.post(baseUrl + 'sign', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            d_val: this.keysForm.d_val,
            p_val: this.keysForm.p_val,
            q_val: this.keysForm.q_val,
          },
          msg: this.msgForm.msg,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              this.$message.success('使用私钥签名成功');
              this.msgForm.signature = res.data.signature;
            } else {
              this.$message.success('签名错误');
              console.log('error encountered');
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.success('签名错误');
          });
      },
      verifyWithPublicKey() {
        this.$http.post(baseUrl + 'verify', {
          key: {
            n_len: this.keysForm.n_len,
            n_val: this.keysForm.n_val,
            e_val: this.keysForm.e_val,
          },
          msg: this.msgForm.msg,
          signature: this.msgForm.signature,
        })
          .then(res => {
            console.log(res);
            if (res.data.ok) {
              if (res.data.verified) {
                this.$message.success('验证成功！签名与文本正确对应！');
              } else {
                this.$message.warning('验证失败！消息已经被篡改！');
              }
            } else {
              console.log('error encountered');
              this.$message.error("验证错误");
            }
          })
          .catch(err => {
            console.log(err);
            this.$message.error("验证错误，请使用配对密钥进行签名验证！");
          });
      },
      onCopy() {
        if (this.keysForm.n_len === 0) {
          this.$message.info('密钥值为空！');
        } else {
          this.$copyText(this.keyStr);
          this.$message.success('密钥值已经复制到剪切板');
        }
      },
      onCopyFail() {
        this.$message.error('复制密钥值到剪切板中出现错误');
      },
    },
  };
</script>

<style lang="scss"
       scoped>
  .title {
    font-size: 40px;
    font-weight: bold;
    margin-bottom: 30px;
  }

  .gen-select {
    width: 100%;
  }

</style>