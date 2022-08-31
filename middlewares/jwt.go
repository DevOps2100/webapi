package middlewares

import (
	"errors"
	"time"
	"webapi/global"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// token过期时间
const TokenExpireDuration = time.Hour * 2

// token加密密钥
var TokenSecret = []byte(global.Config.JWTKey.SigningKey)

// 生成JWT
func GenToken(username string) (int64, string, error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "webapi",
		},
	}

	// 使用指定签名方法创建对象
	tokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := tokenGen.SignedString(TokenSecret)
	// 使用指定的secret签名并获得完整编码后的token字符串
	return c.ExpiresAt, token, err
}

// 解析JWT
func ParseToken(TokenString string) (*MyClaims, error) {
	// 开始解析token
	token, err := jwt.ParseWithClaims(TokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	// 错误检查
	if err != nil {
		return nil, err
	}
	// 校验
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 验证JWT
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		// token为空
		if token == "" {
			ctx.JSON(401, "认证错误")
			ctx.Abort()
			return
		}

		// 解析验证token正确性
		jwt := NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					ctx.JSON(401, gin.H{
						"msg": "授权已过期",
					})
					ctx.Abort()
					return
				}
			}
			ctx.JSON(401, gin.H{
				"msg": "Token验证失败",
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		TokenSecret,
	}
}

type CustomClaims struct {
	ID       uint
	Username string
	Password string
	jwt.StandardClaims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired 令牌过期
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
