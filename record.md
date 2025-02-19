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

##  HTTP请求的类型

在后端开发中，常见的HTTP请求类型主要有以下几种，它们各自有不同的用途和特点：

1.**GET**

- **用途**: 用于从服务器获取资源。
- **特点**:
  - 请求参数通常附加在URL后。
  - 数据在URL中可见，安全性较低。
  - 可以被缓存，支持书签。
  - 不应用于修改数据的操作。

 2.**POST**

- **用途**: 用于向服务器提交数据，通常用于创建新资源或提交表单。
- **特点**:
  - 请求参数包含在请求体中，适合传输大量数据。
  - 数据在URL中不可见，安全性较高。
  - 不可缓存，不支持书签。
  - 常用于创建或更新资源。

3. **PUT**

- **用途**: 用于更新服务器上的资源，或创建指定资源（如果不存在）。
- **特点**:
  - 请求参数包含在请求体中。
  - 通常用于更新整个资源。
  - 幂等性（多次请求效果相同）。

4. **DELETE**

- **用途**: 用于删除服务器上的资源。
- **特点**:
  - 请求参数通常附加在URL后。
  - 幂等性（多次请求效果相同）。

5. **PATCH**

- **用途**: 用于部分更新服务器上的资源。
- **特点**:
  - 请求参数包含在请求体中。
  - 仅更新指定字段，而非整个资源。
  - 非幂等性（多次请求效果可能不同）。

6. **HEAD**

- **用途**: 类似于GET，但只获取响应头，不返回响应体。
- **特点**:
  - 用于检查资源是否存在或获取元数据。
  - 不返回实际数据。

7. **OPTIONS**

- **用途**: 用于获取服务器支持的HTTP方法。
- **特点**:
  - 常用于CORS预检请求。
  - 不返回实际数据。

8. **TRACE**

- **用途**: 用于回显服务器收到的请求，主要用于测试或诊断。
- **特点**:
  - 不常用，主要用于调试。

9. **CONNECT**

- **用途**: 用于建立与资源的双向通信，通常用于SSL隧道。
- **特点**:
  - 主要用于代理服务器。

> 总结：
>- **GET** 和 **POST** 是最常用的请求类型，分别用于获取和提交数据。
>- **PUT** 和 **PATCH** 用于更新资源，PUT更新整个资源，PATCH更新部分资源。
>- **DELETE** 用于删除资源。
>- **HEAD**、**OPTIONS**、**TRACE** 和 **CONNECT** 则用于特定场景，如获取元数据、调试或建立隧道
