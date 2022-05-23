package utils

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateToken
// 生成token
func GenerateToken(id int) string {
	configKey, _ := web.AppConfig.String("jwtKey")
	expirationTime, _ := web.AppConfig.Int("jwtExpirationTime")
	mySigningKey := []byte(configKey)
	expirAt := time.Now().Add(time.Second * time.Duration(expirationTime)).Unix()
	logs.Debug("token 过期时间", time.Unix(expirAt, 0))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":    id,
		"UserAt":    time.Now().Unix(),
		"ExpiresAt": expirAt,
	})

	tokenStr, tokenErr := token.SignedString(mySigningKey)
	if tokenErr != nil {
		logs.Error(tokenErr, "生成token失败")
	}
	return tokenStr
}

// ValidateToken 验证token
func ValidateToken(tokenStr string) error {
	configKey, _ := web.AppConfig.String("jwtKey")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logs.Debug(claims["UserId"], "UserId")
		logs.Debug(claims["UserAt"], "UserAt")
		logs.Debug(claims["ExpiresAt"], "ExpiresAt")
	} else {
		logs.Error(err, "token验证失败")
	}
	return err
}
