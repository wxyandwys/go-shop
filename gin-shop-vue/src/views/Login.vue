<template>
  <div>
    <van-nav-bar
      title="登录" style="background: #f2f3f5; "
     ></van-nav-bar>
    <br><br>
    <van-form @submit="onSubmit">
      <van-field
        v-model="login.username"
        name="用户名"
        label="用户名"
        placeholder="用户名"
        :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
        v-model="login.password"
        type="password"
        name="密码"
        label="密码"
        placeholder="密码"
        :rules="[{ required: true, message: '请填写密码' }]"
      />
      <van-field
        v-model="login.value"
        name="验证码"
        label="验证码"
        placeholder="验证码"
        :rules="[{ required: true, message: '请填写验证码' }]"
      />
      <img :src="img" alt="" @click="imgSrc">
      <div style="margin: 16px;">
        <van-button round block type="info" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import api from './../api/API'
import md5 from 'js-md5'
export default {
  data() {
    return {
      login: {
        username: '',
        password: '',
        value: undefined
      },
      img: api.CAPTCHA ,
      random: undefined
    }
  },
  methods: {
    onSubmit(values) {
      this.login.password = md5(this.login.password)
      
      this.$API.post(api.SHOP_LOGIN + "/" + this.random, this.login).then(res => {
        this.login.password = ''
        if (res.code == 201) {
          this.imgSrc()
        } else {
          this.$store.dispatch('GetUser', res.data.userToken)
          sessionStorage.setItem('user', res.data.userToken)
        }
      })
      
    },
    imgSrc() {
      this.random = Math.floor(Math.random()*1000000)
      this.img = api.CAPTCHA + '/' + this.random
    },
    cc() {
      this.$API.get("/home/captcha/verify/" + this.random).then(res => {})
    }
  },
  created() {
    this.imgSrc()
  }
}
</script>