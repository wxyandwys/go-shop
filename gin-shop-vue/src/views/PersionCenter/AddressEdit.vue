<template>
  <div>
    <van-nav-bar
      title="修改地址"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-address-edit
      :area-list="areaList"
      show-postal
      show-set-default
      :area-columns-placeholder="['请选择', '请选择', '请选择']"
      @save="onSave"
      :address-info="address"
    />
  </div>
</template>

<script>
import { areaList } from '@vant/area-data'
import api from "./../../api/API"
export default {
  data() {
    return{
      areaList,
      address: {},
      id: this.$route.query.id
    }
  },
  methods: {
    onClickLeft() {
      this.$router.push({name: "Address"})
    },
    onSave(e) {
      this.$API.post(api.SHOP_UPDATE_ADDRESS, e).then(res => {
        this.$toast('已保存')
      })
      
    },
    onChange(e) {
      console.log(e)
    },
    ShowAddressById() {
      let user = {
        id: this.id
      }
      this.$API.get(api.SHOP_ADDRESS_BYID, user).then(res => {
        this.address = res.data.address
      })
    }
  },
  created() {
    this.ShowAddressById()
  }
}
</script>