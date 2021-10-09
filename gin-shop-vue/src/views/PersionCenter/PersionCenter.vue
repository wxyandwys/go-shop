<template>
  <div>
    <van-nav-bar
      title="个人中心" style="background: #f2f3f5; "
     ></van-nav-bar>
     <div style="padding:15px;display: flex;">
       <van-image
        round
        width="5rem"
        height="5rem"
        :src="user.img"
        />
      <div style="font-size:24px;padding-left: 25px; margin-top:20px;">{{ user.username }}</div>
     </div>
     

    <van-cell-group inset>
      <van-cell title="个人设置" is-link  />
      <van-cell title="余额" is-link  />
      <van-cell title="订单" is-link  />
      <van-cell title="足迹" is-link  />
      <van-cell title="地址" is-link @click="Address" />
      <van-cell title="收藏" is-link @click="Collection" />
    </van-cell-group>
    <br>
    <van-cell-group inset>
      <van-cell title="反馈" is-link  />
      <van-cell title="关于" is-link  />
    </van-cell-group>
    <br>
    <van-cell-group inset>
      <van-button round type="info" block>退出</van-button>
    </van-cell-group>
  </div>
</template>

<script>
import api from "./../../api/API"
export default {
  data() {
    return {
      user: {}
    }
  },
  methods: {
    GetUser() {
      let user = {
        userToken: this.$store.state.user
      }
      this.$API.post(api.SHOP_USER, user).then(res => {
        this.user = res.data.user
      })
    },
    Collection() {
      this.$router.push({name: "Collection"})
    },
    Address() {
      this.$router.push({name: "Address"})
    }
  },
  created() {
    this.GetUser()
  }
}
</script>