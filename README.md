# gomall
字节跳动青训营抖音商城项目

## 功能
### 认证中心
- [x] 分发身份令牌
- [x] 续期身份令牌
- [x] 校验身份令牌

### 用户服务
- [x] 创建用户(注册)
- [x] 登录
- [x] 用户登出
- [x] 删除用户
- [x] 更新用户(改变用户权限、 更新密码、 启用用户、 禁用用户)
- [x] 获取用户身份信息(单个用户、批量用户)

### 商品服务
- [x] 创建商品
- [x] 修改商品信息
- [x] 删除商品
- [x] 查询商品信息(单个商品、批量商品)

### 购物车服务
- [x] 创建购物车
- [x] 清空购物车
- [x] 获取购物车信息

### 订单服务
- [x] 创建订单
- [x] 修改订单信息
- [x] 订单定时取消

### 结算
- [x] 订单结算

### 支付
- [x] 取消支付
- [x] 定时取消支付
- [x] 支付

### AI大模型
- [x] 订单查询
- [x] 模拟自动下单(从购物车中下单)

## 数据库设计
### 用户表 user
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | uint | 用户id | pk |
| Email | varchar(255) | 邮箱 | not null |
| PasswordHashed | varchar(255) | 密码 | not null |
| Role | char(10) | 用户角色 | admin, user; 默认值:user|
| Status | char(10) | 用户状态 | disabled, enabled|


### 商品表 product
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 商品id | pk |
| Name |    | 商品名 | |
| Description| | 商品描述 ||
| Picture | | 商品图片| |
| Price |  | 商品价格 ||
| Categoreis | | 商品种类| fk |

### 商品种类表 category
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 商品id | pk |
| Name |    | 商品名 | |
| Description| | 商品描述 ||
| Products | | 商品种类|fk|

### 购物车表 cart
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 购物车id | pk |
| UserId |    | 用户id | |
| ProductId| | 商品id ||
| Qty | | 商品数量||

### 订单表 order
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 订单id | pk |
| OrderId |    | 订单号 | |
| UserId| | 用户id ||
| UserCurrency | | 使用货币种类||
| Consignee | | 收货人信息| fk |
| OrderItems | |订单明细| fk|
| OrderState | | 订单状态 | placed, paid, canceled |

### 订单明细表 order_item
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 订单id | pk |
| ProductId | uint | 商品id | |
| OrderIdRefer| | 用户id |fk|
| Quantity | |商品数量| fk|
| Cost | | 总价格 | |

### 支付记录表 payment
| 字段 | 类型 | 描述 | 备注 |
| --  | -- | -- | --|
| ID | int | 订单id | pk |
| UserId| | 用户id ||
| OrderId |    | 订单号 | |
| TransactionId | | 交易id||
| Amount | |总金额|  |
| PayAt | | 支付时间 | |

## 项目代码结构
```
. 
├── app
│   ├── ai          AI大模型服务
│   ├── auth        认证服务
│   ├── cart        购物车服务
│   ├── checkout    结算服务
│   ├── email       邮件服务
│   ├── frontend    前端
│   ├── order       订单服务
│   ├── payment     支付服务
│   ├── product     商品服务
│   └── user        用户服务
├── common           公共函数和工具类
├── db               数据库初始化文件
├── deploy           接口定义文件
├── idl              接口定义文件
└── rpc_gen
```