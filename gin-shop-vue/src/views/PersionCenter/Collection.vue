<template>
  <div>
    <van-nav-bar
      title="收藏"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />
    <van-card v-for="(item, k) in shops" :key="k"
      :price="item.price"
      :desc="item.name"
      :thumb="item.img"
    >
      <template #footer>
        <van-button size="mini" @click="ShowShop(item)">查看商品</van-button>
        <van-button size="mini" @click="DelCollection(item)">删除</van-button>
      </template>
    </van-card> 
    
  </div>
</template>

<script>
import api from "./../../api/API"
export default {
  data() {
    return {
      shops: []
    }
  },
  methods: {
    onClickLeft() {
      this.$router.push({name: "Persion"})
    },
    ShowCollection() {
      this.$API.get(api.SHOP_COLLECTION).then(res => {
        this.shops = res.data.shops
      })
    },
    ShowShop(e) {
      this.$router.push({path: "/shopData",query:{data: e.id}})
    },
    DelCollection(e) {
      let co = {
        id: e.id
      }
      this.$API.post(api.SHOP_DELETE_COLLECTOPN, co).then(res => {
        this.$toast("已删除")
        this.ShowCollection()
      })
    }
  },
  created() {
    this.ShowCollection()
  }
}
</script>