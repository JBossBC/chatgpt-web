package middlerware

import (
	"chatgpt-web/internal/app/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const extraInterface = "user/"

const QueryKey = "Authorization"

var (
	claimsElems = []string{"username", "level", "levelDeadline"}
)

func JWT(context *gin.Context) {
	path := context.Request.URL.Path
	if strings.HasPrefix(path, extraInterface) {
		context.Next()
		return
	}
	key := context.Request.Header.Get(QueryKey)
	if key == "" {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("未通过验证"))
		return
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(key, claims, nil)
	if err != nil {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("未通过验证"))
		return
	}
	// whether the jwt is invalid
	customVerify(claims)
	// give the claims to gin.context
	if claimsToContext(map[string]interface{}(claims), context) {
		context.Next()
	} else {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("未通过验证"))
		return
	}
}
func customVerify(claims jwt.Claims) bool {
	mapping := (map[string]interface{})(claims.(jwt.MapClaims))
	value, ok := mapping["exd"].(string)
	if !ok {
		return false
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return false
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", value, location)
	if err != nil {
		return false
	}
	if t.After(time.Now()) {
		return true
	}
	return false
}
func claimsToContext(data map[string]interface{}, context *gin.Context) bool {
	for _, value := range claimsElems {
		claimsValue, ok := data[value]
		if !ok {
			return false
		}
		context.Keys[value] = claimsValue
	}
	return true
}
