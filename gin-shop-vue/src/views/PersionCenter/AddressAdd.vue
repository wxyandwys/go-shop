<template>
  <div>
    <van-nav-bar
      title="添加地址"
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
    }
  },
  methods: {
    onClickLeft() {
      this.$router.push({name: "Address"})
    },
    onSave(e) {
      this.$API.post(api.SHOP_INSERT_ADDRESS, e).then(res => {
        this.$toast('添加地址');
        this.onClickLeft()
      })
      
    },
    onChange(e) {
      console.log(e)
    },
  }
}
</script>