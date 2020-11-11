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
                <el-button type="success">下载密钥</el-button>
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
              <el-form-item label="加密后的消息">
                <el-input v-model="msgForm.encodedMsg"
                          type="textarea"
                          disabled
                          :autosize="{ minRows: 3}"
                          placeholder="加密后的消息">
                </el-input>
              </el-form-item>
              <el-form-item label="解密后的消息">
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
                          :autosize="{ minRows: 2}"
                          placeholder="数字签名">
                </el-input>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="primary">公钥加密</el-button>
                <el-button type="primary">私钥解密</el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="success">公钥加密</el-button>
                <el-button type="success">私钥解密</el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="info">公钥签名</el-button>
                <el-button type="info">私钥验证</el-button>
              </el-form-item>
              <el-form-item style="display: flex; justify-content: left;">
                <el-button type="warning">私钥签名</el-button>
                <el-button type="warning">公钥验证</el-button>
              </el-form-item>
            </el-form>
          </el-col>
        </el-row>
      </el-col>
    </el-row>

  </div>

</template>

<script>
  const baseUrl = 'http://localhost:8080/'
  // import qs from 'qs'
  export default {
    name: "RsaPanel",
    data() {
      return {
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
        gen_options: [{
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
        }],
      };
    },
    methods: {
      onRequestGenKey() {
        this.$http.get(baseUrl + `keys?n_byte=${this.gen_key_form.n_len}`)
        .then(res => {
          console.log(res)
          if (res.data.ok) {
            Object.assign(this.keysForm, res.data.privateKey)
            this.keysForm.e_val = res.data.publicKey.e_val
          } else {
            console.log("不ok")
          }
        })
        .catch(err => {
          console.log(err)
        })
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