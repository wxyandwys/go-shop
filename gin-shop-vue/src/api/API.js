
const HOME = "/home"
const CAPTCHA = "http://127.0.0.1:8901/captcha"

var api = {
  // 首页
  HOME: HOME,
  // 首页获取推荐分类
  SHOP_CATEGORY_LIST: HOME +  "/GetShopCategoryList",
  // 首页获取推荐商品
  SHOP_SHOPS_LIST: HOME + "/GetShopList",
  // 商品数据
  SHOP: HOME + "/GetShop",
  // 商品规格选择
  SHOP_TREE: HOME + "/GetShopTree",
  // 商品规格组合
  SHOP_LIST_SKU: HOME + "/GetShopListSku",
  // 商品参数
  SHOP_PARAMETER: HOME + "/GetShopParameters",
  // 父分类
  SHOP_CATEGORY_LIST_PARENT: HOME + "/GetShopCategoryListParent",
  // 子分类
  SHOP_CATEGORY_LIST_CHILDREN: HOME + "/GetShopCategoryListChildren",
  // 子分类的集合商品
  SHOP_CATEGORY_SHOP_LIST_BYID: HOME + "/GetShopCategoryListChildrenById",

  // 登录
  SHOP_LOGIN: HOME + "/ShopLogin",
  // 验证码
  CAPTCHA: CAPTCHA,
  // 获取用户
  SHOP_USER: HOME + "/GetUser"

}

export default api