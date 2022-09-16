//
// Author xiexu
// Date 2020-06-05
//

package mid

import (
	"fmt"
	"project/pkg/web"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

// TokenHandle 识别 token
func TokenHandle(tokenKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("authorization")
		// token 不在头文件, 也许在url参数中
		if len(tokenString) == 0 {
			tokenString = c.Query("token")
		}
		data := strings.Split(tokenString, " ")
		if len(data) >= 2 {
			tokenString = strings.TrimSpace(data[1])
		}

		token, err := parseToken(tokenString, tokenKey)
		if err != nil {
			web.Fail(c, web.ErrUnauthorizedToken.With(err.Error()))
			c.Abort()
			return
		}

		if token.ExpiresAt.Before(time.Now()) {
			web.Fail(c, web.ErrUnauthorizedToken.With(err.Error()))
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("tokenStr", tokenString)
		c.Set("uid", token.UID)

		c.Next()
	}
}

// parseToken 解析 token
func parseToken(token string, key string) (*Claims, error) {
	fun := func(token *jwt.Token) (i any, e error) {
		return []byte(key), nil
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, fun)

	var (
		claim *Claims
		ok    bool
	)
	if tokenClaims != nil {
		claim, ok = tokenClaims.Claims.(*Claims)
		if !ok {
			return nil, fmt.Errorf("令牌类型错误")
		}
	}

	if err == nil {
		if !tokenClaims.Valid {
			return claim, fmt.Errorf("令牌无效")
		}
		return claim, nil
	}

	return claim, err

}

// Claims ...
type Claims struct {
	UID int
	jwt.StandardClaims
}

// NewToken 创建 token
func NewToken(id int, key string) string {
	endTime := 2 * time.Hour

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(endTime)), // 失效时间
			IssuedAt:  jwt.Now(),                       // 签发时间
			Issuer:    "easynts",                       // 签发人
			NotBefore: jwt.Now(),                       // 生效时间
			Subject:   "login",                         // 主题
		},
	}
	// claims.Audience = append(claims.Audience, "easyntd")
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	val, err := tokenClaims.SignedString([]byte(key)) //密钥
	if err != nil {
		return ""
	}
	return val

}
