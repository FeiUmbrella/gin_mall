package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400 // 非法参数

	// User 模块错误
	ErrorExistUser             = 30001 // 用户已存在
	ErrorFailEncryption        = 30002
	ErrorUserNotFound          = 30003 // 用户不存在
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeOut = 30006
	ErrorAuthCheckTokenFail    = 3007
	ErrorUploadFail            = 3008
	ErrorSendEmail             = 3009

	// Product 模块错误
	ErrorProductImgUpload = 40001

	// 收藏夹错误
	ErrorFavoriteExist = 50001
)
