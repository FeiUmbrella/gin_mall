# Record
* api: Controller调用服务层逻辑
* cmd: 存放主程序
* conf: 项目所需的配置
* dao: 数据库相关的持久层操作
* middleware: 项目所需中间件
* model: 数据库表相关定义
* pkg: pkg/e 自定义错误; pkg/util 放一些项目相关组件(eg.负责加密的组件)
* routes: 路由，定义项目业务/功能 的路由
* serializer: 输出序列化
* service: 服务层，业务逻辑编写
* static: 存放业务所需的静态文件(eg.图片文件)
## 新增一个业务/功能的步骤(以"用户注册"为例)
1. 在routes中定义相关路由, 请求的类型(eg.GET/POST)
```go
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static")) // 加载静态文件
	v1 := r.Group("/api/v1")                    // 路由组放在 /api/v1 下
	{
		// 先ping一下是否连通
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("user/register", api.UserRegister)
	}
	return r
}
```
2. 在api/user.go中定义UserRegister的Controller
```go
// Controler
func UserRegister(c *gin.Context) {
	var userRegister service.UserService // 请求参数
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
```
3. 在服务层编写相关逻辑
```go
func (service UserService) Register(ctx context.Context) serializer.Response{}
```
## pkg/util
### AES加密算法
AES 加密的基本概念
* 对称加密：AES 是对称加密算法，意味着加密和解密使用相同的密钥。发送方和接收方必须共享相同的密钥才能安全通信。

* 分组加密：AES 将数据分成固定大小的块进行加密。AES 的分组大小为 128 位（16 字节），无论密钥长度是 128 位、192 位还是 256 位。

* 密钥长度：AES 支持三种密钥长度：

> AES-128：128 位密钥
> 
> AES-192：192 位密钥 
>
>AES-256：256 位密钥
>>密钥越长，安全性越高，但计算开销也越大。

AES 加密过程
AES 加密过程包括多个步骤，主要分为以下几个阶段：

密钥扩展：根据初始密钥生成一系列轮密钥（Round Keys），用于每一轮的加密操作。

初始轮（Initial Round）： AddRoundKey：将明文块与第一轮密钥进行异或操作。

主轮（Main Rounds）： AES 加密的核心部分，每一轮都包含以下四个步骤：

SubBytes：使用 S 盒（Substitution Box）对每个字节进行非线性替换。

ShiftRows：对状态矩阵的每一行进行循环移位。

MixColumns：对状态矩阵的每一列进行线性变换。

AddRoundKey：将当前状态与轮密钥进行异或操作。

轮数取决于密钥长度：

AES-128：10 轮

AES-192：12 轮

AES-256：14 轮

最终轮（Final Round）：

SubBytes：与主轮相同。

ShiftRows：与主轮相同。

AddRoundKey：与主轮相同。

MixColumns：在最终轮中省略。

AES 解密过程
AES 解密过程是加密过程的逆过程，使用相同的轮密钥，但操作顺序相反。解密过程也包括初始轮、主轮和最终轮，每个步骤都是加密步骤的逆操作。
