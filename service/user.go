package service

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	// todo:前端+后端双重验证
	Key string `json:"key" form:"key"` // 加密的密钥，一开始就要定义好的密钥，这里采用简洁形式只在前端进行验证
}
type SendEmailService struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// 1.绑定邮箱 2.解绑邮箱 3.改密码
	OperationType uint `json:"operation_type" form:"operation_type"`
}

type ValidEmailService struct {
}

// Register 用户注册
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}

	// 对称加密，密文存储初始金额1w
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	// 看username是否存在
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Avatar:   "user.jpg",
		Status:   model.Active,
		Money:    util.Encrypt.AesEncoding("10000"), // 初始金额的加密
	}
	// 密码加密
	// todo: 规范来说前端传来的密码也是加密状态，要先解密再加密存储
	if err := user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error

	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登录
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	// 判断UserName是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil || !exist {
		code = e.ErrorUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}

	// 校验密码
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}

	// http是无状态的，不知道某次请求是谁访问的
	// 签发一个token，来表示访问身份
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code := e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// Update 用户信息修改
func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	// 根据 uid 找到数据库对应 record
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)

	// 修改昵称nickname
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Post 头像更新
func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user) // 更新数据库相关record
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// Send 发送邮件
func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice // 绑定邮箱 修改密码 模板通知
	// 将 http中传来的参数进行jwt加密，其中uId参数用来辨识用户
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	address = conf.ValidEmail + token // 发送内容
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1) // 将mailStr中的“Email”全部替换为address
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail) // 发送方邮箱
	m.SetHeader("To", service.Email)    // 接收方邮箱
	m.SetHeader("Subject", "gin_mall test")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Valid 验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint

	// 验证token
	code := e.Success
	claims, err := util.ParseEmailToken(token)
	if token == "" {
		code = e.InvalidParams
	} else {
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeOut
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 验证成功，获取该用户信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 修改用户信息
	if operationType == 1 { // 绑定邮箱
		user.Email = email
	} else if operationType == 2 { // 解绑邮箱
		user.Email = ""
	} else { // 修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	// 更新用户信息
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
