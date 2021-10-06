<template>
  <div style="margin-bottom:50px">
    <van-nav-bar
      title="分类" style="background: #f2f3f5; "
     >
    </van-nav-bar>
    <van-tree-select height="100%" :items="pCategory" :main-active-index.sync="active" @click-nav="cnav">
      <template #content>

          <van-grid :border="false" :column-num="3" :gutter="10" >
            <van-grid-item v-for="(item, k) in cCategory" :key="k" @click="clickView(item)">
              <van-image src="https://img01.yzcdn.cn/vant/apple-1.jpg" />
              <div>
                {{item.name}}
              </div>              
            </van-grid-item>
          </van-grid> 
        
      </template>
    </van-tree-select>
  </div>
</template>

<script>
import api from './../api/API'
export default {
  data() {
    return {
      active: 0,
      items: [{ text: '分组 1' }, { text: '分组 2' }],
      pCategory: [],
      cCategory: []
    }
  },
  methods: {
    GetShopCategoryListChildren(pid) {
      let cid = {
        id: pid
      }
      this.$API.get(api.SHOP_CATEGORY_LIST_CHILDREN, cid).then(res => {
        this.cCategory = res.data.categoryChildren
      })
    },
    GetShopCategoryListParent() {
      this.$API.get(api.SHOP_CATEGORY_LIST_PARENT).then(res => {
        var data = res.data.categoryParent
        for (let j = 0; j < data.length; j++) {
          let category = {
            id: data[j].id,
            text: data[j].name,
          }
          this.pCategory.push(category)
        }

        this.GetShopCategoryListChildren(this.pCategory[0].id)
      })
    },
    cnav(e) {
      var id = this.pCategory[e].id
      this.GetShopCategoryListChildren(id)
    },
    clickView(data) {
      this.$router.push({name: "CategoryChildren", query: {id: data.id}})
    }
  },
  created() {
    this.GetShopCategoryListParent()
  }
}
</script>