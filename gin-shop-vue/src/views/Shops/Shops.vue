<template>
  <div style="margin-bottom:50px">
    <van-pull-refresh v-model="isLoading" @refresh="onRefresh">
    <van-nav-bar
      title="标题"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />
    <van-swipe @change="onChange">
      <van-swipe-item v-for="(item, k) in imgs" :key="k">
        <van-image
        
          fit="contain"
          :src="item.img"
        />
      </van-swipe-item>
      <template #indicator>
        <div class="custom-indicator">{{ current + 1 }}/{{ imgs.length }}</div>
      </template>
    </van-swipe>
    <div style="padding:0 20px;">
      <h2>{{ shop.name }}</h2>
      <h3 style="color:red;">￥{{ shop.price }}</h3>
    </div>
    

    <van-sku
      v-model="show"
      :sku="tmp"
      :goods="goods"
      :goods-id="data"
      :quota="quota"
      :quota-used="quotaUsed"
      :hide-stock="tmp.hide_stock"
      ref="skuv"
      @buy-clicked="onBuyClicked"
      @add-cart="onAddCartClicked"
    />
    <van-cell title="选择产品" is-link  @click="show = true"/>

    <van-cell title="显示分享面板" @click="showShare = true"  is-link  />
    <van-share-sheet
      v-model="showShare"
      title="立即分享给好友"
      :options="options"
      @select="onSelect"
    />

    <van-cell is-link title="产品参数" @click="show1 = true" />
    <van-action-sheet v-model="show1" title="产品参数">
      <van-row class="hr" v-for="(item, k) in parameters" :key="k">
        <van-col span="8" class="spansheet">{{ item.k }}</van-col>
        <van-col span="16" class="spansheet">{{ item.v }}</van-col>
      </van-row>
    </van-action-sheet>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
      <p>刷新次数: {{ count }}</p>
    </van-pull-refresh>
    <van-goods-action>
      <van-goods-action-icon icon="chat-o" text="客服" @click="onClickIcon" />
      <van-goods-action-icon icon="cart-o" text="购物车" @click="onClickIcon" />
      <van-goods-action-icon icon="star" text="收藏" color="#ff5000" @click="InsertCollection"/>
      <van-goods-action-button type="warning" text="加入购物车" />
      <van-goods-action-button
        type="danger"
        text="立即购买"
        @click="onClickButton"
      />
    </van-goods-action>
    
  </div>
</template>

<script>
import api from "./../../api/API"
export default {
  data () {

    return {
      show1: false,
      data: this.$route.query.data,
      current: 0,
      shop: {},
      imgs: [],
      sku: {
        // 所有sku规格类目与其值的从属关系，比如商品有颜色和尺码两大类规格，颜色下面又有红色和蓝色两个规格值。
        // 可以理解为一个商品可以有多个规格类目，一个规格类目下可以有多个规格值。
        tree: [
          {
            k: '颜色', // skuKeyName：规格类目名称
            k_s: 's1', // skuKeyStr：sku 组合列表（下方 list）中当前类目对应的 key 值，value 值会是从属于当前类目的一个规格值 id
            v: [
              {
                id: '1', // skuValueId：规格值 id
                name: '红色', // skuValueName：规格值名称
                imgUrl: 'https://img01.yzcdn.cn/2.jpg', // 规格类目图片，只有第一个规格类目可以定义图片
                previewImgUrl: 'https://img01.yzcdn.cn/2.jpg', // 用于预览显示的规格类目图片
              },
              {
                id: '2',
                name: '蓝色',
                imgUrl: "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d3a36269683536a9d46a044392ed37ab.png",
                previewImgUrl: "https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/d3a36269683536a9d46a044392ed37ab.png",
              },
              {
                id: '3',
                name: '粉色',
                imgUrl: 'https://img01.yzcdn.cn/2.jpg',
                previewImgUrl: 'https://img01.yzcdn.cn/2.jpg',
              }
            ],
          largeImageMode: true, //  是否展示大图模式
        },
        {
          k: '尺寸',
          k_s: 's2',
          v: [
            {
              id: '1',
              name: 'S'
            },
            {
              id: '2',
              name: 'L'
            },
            {
              id: '3',
              name: 'X'
            }
          ]
        }
        ],
        // 所有 sku 的组合列表，比如红色、M 码为一个 sku 组合，红色、S 码为另一个组合
        list: [
          {
            id: 2259, // skuId
            s1: '1', // 规格类目 k_s 为 s1 的对应规格值 id
            s2: '1', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 1000, // 价格（单位分）
            stock_num: 100 // 当前 sku 组合对应的库存
          },
          {
            id: 2260, // skuId
            s1: '2', // 规格类目 k_s 为 s1 的对应规格值 id
            s2: '2', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 1800, // 价格（单位分）
            stock_num: 10 // 当前 sku 组合对应的库存
          },
          {
            id: 2261, // skuId
            s1: '3', // 规格类目 k_s 为 s1 的对应规格值 id
            s2: '3', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 100, // 价格（单位分）
            stock_num: 10 // 当前 sku 组合对应的库存
          },
          {
            id: 2262,
            s1: '1',
            s2: '3',
            price: 2000000,
            stock_num: 99
          },
          {
            id: 2262,
            s1: '2',
            s2: '3',
            price: 2000000,
            stock_num: 99
          }
        ],
        price: '1.00', // 默认价格（单位元）
        stock_num: 120, // 商品总库存
        none_sku: false, // 是否无规格商品
        hide_stock: false, // 是否隐藏剩余库存
      },
      goods: {
        // 默认商品 sku 缩略图
        picture: 'https://img01.yzcdn.cn/2.jpg'
      },
      show: false,
      quota: 0,
      quotaUsed: 0,
      tmp: {
        tree: [],
        list: [],
        price: undefined,
        stock_num: undefined,
        none_sku: false,
        hide_stock: false,
      },
      showShare: false,
      options: [
        { name: '微信', icon: 'wechat' },
        { name: '微博', icon: 'weibo' },
        { name: '复制链接', icon: 'link' },
        { name: '分享海报', icon: 'poster' },
        { name: '二维码', icon: 'qrcode' },
      ],
      count: 0,
      isLoading: false,
      parameters: []
    }
  },
  methods: {
    onClickLeft() {
      this.$router.push({name: "Home"})
    },
    onClickIcon() {
      this.$toast('点击图标');
    },
    onClickButton() {
      console.log(this.$refs.skuv.getSkuData())
      this.$toast('点击按钮');
    },
    getShop() {
      let shop = {
        id: this.data
      }
      this.$API.get(api.SHOP, shop).then(res => {
        this.shop = res.data.shop
        this.imgs = res.data.imgs
        this.quota = res.data.shop.quote
        this.tmp.price = res.data.shop.price
        this.tmp.stock_num = res.data.shop.num
        this.tmp.hide_stock = res.data.shop.hide_stock
        this.tmp.none_sku = res.data.shop.none_sku
        this.goods.picture = res.data.shop.picture
      })
    },
    GetShopTree() {
      let shop = {
        id: this.data
      }
      this.$API.get(api.SHOP_TREE, shop ).then(res => {
        
        this.tmp.tree = res.data.trees
      })
    },
    GetShopListSku() {
      let shop = {
        id: this.data
      }
      this.$API.get(api.SHOP_LIST_SKU, shop).then(res => {
        
        var list = res.data.list
        
        for (let i = 0; i < list.length; i++) {
          var z = new Map()
          z.set('id', list[i].id)
          z.set('price', list[i].price)
          z.set('stock_num', list[i].stock_num)
          for (let j = 0;  j < list[i].slsu.length; j++) {
            z.set(list[i].slsu[j].k_s, list[i].slsu[j].v)
          }
          const obj = [...z.entries()].reduce((obj, [key, value]) => (obj[key] = value, obj), {})
          list[i] = obj
        }
        
        this.tmp.list = list
        
      })
    },
    GetShopParameters() {
      let shop = {
        id: this.data
      }
      this.$API.get(api.SHOP_PARAMETER, shop).then(res => {
        this.parameters = res.data.parameters
      })
    },
    onChange(index) {
      this.current = index;
    },
    onBuyClicked(e) {
      console.log(e)
    },
    onAddCartClicked(e) {
      console.log(e)
    },
    onSelect(option) {
      this.$toast(option.name);
      this.showShare = false;
    },
    onRefresh() {
      this.getShop()
      this.GetShopTree()
      this.GetShopListSku()
      this.GetShopParameters()

      setTimeout(() => {
        this.$toast('刷新成功');
        this.isLoading = false;
        this.count++;
      }, 1000);
    },
    InsertCollection() {
      let co = {
        shop_id: this.data
      }
      this.$API.get(api.SHOP_COLLECTION_BYID, co).then(res => {
        if (res.data.count == 0) {
          this.$API.post(api.SHOP_INSERT_COLLECTION, co).then(res => {
            this.$toast("收藏成功")
          })
        } else {
          this.$toast("收藏成功")
        }
      })
      
    }

  },
  mounted() {
    this.getShop()
    this.GetShopTree()
    this.GetShopListSku()
    this.GetShopParameters()
  }
}
</script>

<style scoped>
 .custom-indicator {
    position: absolute;
    right: 5px;
    bottom: 5px;
    padding: 2px 5px;
    font-size: 12px;
    background: rgba(0, 0, 0, 0.7);
    color: #fff;
  }
  .spansheet{
    padding:12px 8px
  }
  .hr {
    border-bottom: 1px solid #eee;
    padding:0 10px;
  }
</style>
