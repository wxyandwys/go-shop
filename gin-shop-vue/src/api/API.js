
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
  SHOP_USER: HOME + "/GetUser",

  //添加收藏
  SHOP_INSERT_COLLECTION: HOME + "/InsertCollection",
  // 查询用户是否收藏商品
  SHOP_COLLECTION_BYID: HOME + "/ShowCollectionById",
  // 查询用户所有收藏商品
  SHOP_COLLECTION: HOME + "/ShowCollection",
  // 删除收藏
  SHOP_DELETE_COLLECTOPN: HOME + "/DeleteCollection",

  // 添加地址
  SHOP_INSERT_ADDRESS: HOME + "/InsertAddress",
  // 查询地址
  SHOP_ADDRESS: HOME + "/ShowAddress",
  // 修改地址
  SHOP_UPDATE_ADDRESS: HOME + "/UpdateAddress",
  // 查询地址id
  SHOP_ADDRESS_BYID: HOME + "/ShowAddressById"
}

export default api