package middlewares

import (
	"coltstail-server/conf"
	"coltstail-server/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	UserID   int64
	Username string `json:"username" comment:"手机号"`
	RoleID   uint
	jwt.StandardClaims
}

func CreateJWT(user *models.TblAccount) (string, error) {
	expiresTime := time.Now().Add(7 * 24 * time.Hour).Unix() //有效期为7天
	claims := CustomClaims{
		UserID:   user.ID,
		Username: user.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime,
			NotBefore: time.Now().Unix(),
			Issuer:    conf.Conf().JWT_ISSUER,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(conf.Conf().JWT_SECRET))
}

func ParseJWT(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 只接受 HS256，拒绝 none / 其他算法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(conf.Conf().JWT_SECRET), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			if !claim.VerifyIssuer(conf.Conf().JWT_ISSUER, true) || claim.UserID <= 0 {
				return nil, errors.New("invalid token")
			}
			return claim, nil
		}
	}
	if err == nil {
		err = errors.New("invalid token")
	}
	return nil, err
}

// validate token and set the jwt info into ctx (ctx.Get("jwt") can get the info)
func AuthenticateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    403,
				"message": "登录过期,请重新登录",
				"data":    nil,
				"success": false,
			})
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		claim, err := ParseJWT(token)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    403,
				"message": "登录过期,请重新登录",
				"data":    nil,
				"success": false,
			})
			return
		}
		ctx.Set("jwt", claim)
		ctx.Next()
	}
}

// validate token and set the jwt info into ctx (ctx.Get("jwt") can get the info)
func GetUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth != "" && strings.HasPrefix(auth, "Bearer ") {
			token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
			if claim, err := ParseJWT(token); err == nil {
				ctx.Set("jwt", claim)
			}
		}
		ctx.Next()
	}
}
