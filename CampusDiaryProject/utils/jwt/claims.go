package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int `json:"userid"`
	jwt.StandardClaims
}

//TokenExpireDuration 定义JWT的过期时间 一个月
const TokenExpireDuration = time.Hour * 720

//MySecret 定义Secret
var MySecret = []byte("来自M78的token签名")
