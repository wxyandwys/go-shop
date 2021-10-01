<template>
  <div>
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
      :sku="sku"
      :goods="goods"
      :goods-id="goodsId"
      :quota="quota"
      :quota-used="quotaUsed"
      :hide-stock="sku.hide_stock"
      :custom-stepper-config="customStepperConfig"
      @buy-clicked="onBuyClicked"
      @add-cart="onAddCartClicked"
    />
    <van-cell title="选择产品" is-link  @click="show = true"/>
    

    <van-goods-action>
      <van-goods-action-icon icon="chat-o" text="客服" @click="onClickIcon" />
      <van-goods-action-icon icon="cart-o" text="购物车" @click="onClickIcon" />
      <van-goods-action-icon icon="star" text="已收藏" color="#ff5000" />
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
                imgUrl: 'https://img01.yzcdn.cn/2.jpg',
                previewImgUrl: 'https://img01.yzcdn.cn/2.jpg',
              },
              {
                id: '3',
                name: '粉色',
                imgUrl: 'https://img01.yzcdn.cn/2.jpg',
                previewImgUrl: 'https://img01.yzcdn.cn/2.jpg',
              }
            ],
            largeImageMode: true, //  是否展示大图模式
          }
        ],
        // 所有 sku 的组合列表，比如红色、M 码为一个 sku 组合，红色、S 码为另一个组合
        list: [
          {
            id: 2259, // skuId
            s1: '1', // 规格类目 k_s 为 s1 的对应规格值 id
            //s2: '1', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 100, // 价格（单位分）
            stock_num: 110 // 当前 sku 组合对应的库存
          },
          {
            id: 2260, // skuId
            s1: '2', // 规格类目 k_s 为 s1 的对应规格值 id
            //s2: '1', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 100, // 价格（单位分）
            stock_num: 110 // 当前 sku 组合对应的库存
          },
          {
            id: 2261, // skuId
            s1: '3', // 规格类目 k_s 为 s1 的对应规格值 id
            //s2: '3', // 规格类目 k_s 为 s2 的对应规格值 id
            price: 100, // 价格（单位分）
            stock_num: 110 // 当前 sku 组合对应的库存
          },
        ],
        price: '1.00', // 默认价格（单位元）
        stock_num: 227, // 商品总库存
        collection_id: 2261, // 无规格商品 skuId 取 collection_id，否则取所选 sku 组合对应的 id
        none_sku: false, // 是否无规格商品
        /*
        messages: [
          {
            // 商品留言
            datetime: '0', // 留言类型为 time 时，是否含日期。'1' 表示包含
            multiple: '0', // 留言类型为 text 时，是否多行文本。'1' 表示多行
            name: '留言', // 留言名称
            type: 'text', // 留言类型，可选: id_no（身份证）, text, tel, date, time, email
            required: '1', // 是否必填 '1' 表示必填
            placeholder: '' // 可选值，占位文本
          },
          {
            // 商品留言
            datetime: '0', // 留言类型为 time 时，是否含日期。'1' 表示包含
            multiple: '0', // 留言类型为 text 时，是否多行文本。'1' 表示多行
            name: '电话', // 留言名称
            type: 'tel', // 留言类型，可选: id_no（身份证）, text, tel, date, time, email
            required: '1', // 是否必填 '1' 表示必填
            placeholder: '' // 可选值，占位文本
          }
        ],
        */
        hide_stock: false // 是否隐藏剩余库存
      },
      goods: {
        // 默认商品 sku 缩略图
        picture: 'https://img01.yzcdn.cn/2.jpg'
      },
      customStepperConfig: {
        // 自定义限购文案
        quotaText: '每次限购xxx件',
        // 自定义步进器超过限制时的回调
        handleOverLimit: (data) => {
          const { action, limitType, quota, quotaUsed, startSaleNum } = data;

          if (action === 'minus') {
            this.$toast(startSaleNum > 1  ? `${startSaleNum}件起售` : '至少选择一件商品');
          } else if (action === 'plus') {
            // const { LIMIT_TYPE } = Sku.skuConstants;
            if (limitType === LIMIT_TYPE.QUOTA_LIMIT) {
              let msg = `单次限购${quota}件`;
              if (quotaUsed > 0) msg += `，你已购买${quotaUsed}`;
              this.$toast(msg);
            } else {
              this.$toast('库存不够了');
            }
          }
        },
        // 步进器变化的回调
        handleStepperChange: currentValue => {},
        // 库存
        stockNum: 1999,
        // 格式化库存
        stockFormatter: stockNum => {},
      },
      goodsId: '1',
      show: false,
      picture: undefined,
      quota: 0,
      quotaUsed: 1
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
      this.$toast('点击按钮');
    },
    getShop() {
      let shop = {
        id: this.data
      }
      this.$API.get(api.SHOP, shop).then(res => {
        this.shop = res.data.shop
        this.imgs = res.data.imgs
      })
    },
    onChange(index) {
      this.current = index;
    },
    onBuyClicked(e) {

    },
    onAddCartClicked(e) {

    }
  },
  mounted() {
    this.getShop()
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
</style>
