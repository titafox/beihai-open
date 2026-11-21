package middlewares

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

func I18n() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("language")
		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
		bundle.LoadMessageFile("es.toml")

		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":  "401",
				"error": "invalid token",
			})
			return
		}
		token := strings.Fields(auth)[1]
		claim, err := ParseJWT(token)
		ctx.Set("jwt", claim)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":  "403",
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
