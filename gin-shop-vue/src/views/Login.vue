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
      <button @click="cc">cc</button>
      <div style="margin: 16px;">
        <van-button round block type="info" native-type="submit">提交</van-button>
      </div>
    </van-form>
  </div>
</template>

<script>
import api from './../api/API'
export default {
  data() {
    return {
      login: {
        username: '',
        password: '',
        value: undefined
      },
      img: api.CAPTCHA
    };
  },
  methods: {
    onSubmit(values) {
      console.log('submit', values);
      this.$API.post(api.SHOP_LOGIN , this.login).then(res => {

      })
    },
    imgSrc() {
      this.img = api.CAPTCHA + '?v=' + Math.floor(Math.random()*1000000)
    },
    cc() {
      this.$API.get("/home/captcha/verify/" + this.login.value).then(res => {})
    }
  },
}
</script>