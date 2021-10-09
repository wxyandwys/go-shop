<template>
  <div>
    <van-nav-bar
      title="地址"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />

    <van-address-list
      v-model="chosenAddressId"
      :list="list"
      @add="onAdd"
      @edit="onEdit"
      
    />
  </div>
</template>
<script>
import api from "./../../api/API"
export default {
  data() {
    return {
      chosenAddressId: '1',
      list: [
        {
          id: '1',
          name: '张三',
          tel: '13000000000',
          address: '浙江省杭州市西湖区文三路 138 号东方通信大厦 7 楼 501 室',
          isDefault: true,
        },
        {
          id: '2',
          name: '李四',
          tel: '1310000000',
          address: '浙江省杭州市拱墅区莫干山路 50 号',
        },
      ],
    }
  },
  methods: {
    onClickLeft() {
      this.$router.push({name: "Persion"})
    },
    onAdd() {
      this.$router.push({name: "AddressAdd"})
    },
    onEdit(e) {
      this.$router.push({name:"AddressEdit", query:{id: e.id}})
    },
    ShowAddress() {
      this.$API.get(api.SHOP_ADDRESS).then(res => {
        let addresses = res.data.address
        let list = []
        for (var i = 0; i < addresses.length; i++) {
          let add = {
            id: addresses[i].id,
            name: addresses[i].name,
            tel: addresses[i].tel,
            address: addresses[i].province + addresses[i].city  + addresses[i].county,
            isDefault: addresses[i].isDefault
          }
          if (add.isDefault) {
            this.chosenAddressId = add.id
          }
          list.push(add)
        }
        this.list = list
      })
    },
    
  },
  watch: {
    chosenAddressId() {
      console.log(this.chosenAddressId)
    }
  },
  created() {
    this.ShowAddress()
  }
}
</script>