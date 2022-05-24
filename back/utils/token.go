package utils

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

// GenerateToken
// 生成token
func GenerateToken(id int) (string, time.Time) {
	configKey, _ := web.AppConfig.String("jwtKey")
	expirationTime, _ := web.AppConfig.Int("jwtExpirationTime")
	logs.Debug(configKey, "jwtKey")
	logs.Debug(expirationTime, "jwtExpirationTime")
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
	return tokenStr, time.Unix(expirAt, 0)
}

// ValidateToken 验证token
func ValidateToken(tokenStr string) (int, error) {
	var customErr error
	var userId interface{}
	tokenResult, rError := ParserToken(tokenStr)
	if rError != nil {
		return 0, rError
	}
	configKey, _ := web.AppConfig.String("jwtKey")
	token, err := jwt.Parse(tokenResult, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configKey), nil
	})
	customErr = err

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logs.Debug(claims["UserId"], "UserId")
		logs.Debug(claims["UserAt"], "UserAt")
		logs.Debug(claims["ExpiresAt"], "ExpiresAt")
		userId = claims["UserId"]
		customErr = nil
	} else {
		logs.Error(err, "token验证失败")
		return 0, errors.New(err.Error() + " token验证失败")
	}
	return int(userId.(float64)), customErr
}

// RefreshToken
// 只有在过期之前才能刷新
func RefreshToken(tokenStr string) (string, int64, error) {
	tokenResult, tokenRErr := ParserToken(tokenStr)
	if tokenRErr != nil {
		return "", 0, tokenRErr
	}
	token, err := jwt.ParseWithClaims(tokenResult, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		configKey, _ := web.AppConfig.String("jwtKey")
		return []byte(configKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", 0, err
	}
	configKey, _ := web.AppConfig.String("jwtKey")
	expirationTime, _ := web.AppConfig.Int("jwtExpirationTime")
	expirAt := time.Now().Add(time.Second * time.Duration(expirationTime)).Unix()

	// 添加令牌时限
	newClaims := jwt.MapClaims{
		"UserId":    claims["UserId"],
		"UserAt":    time.Now().Unix(),
		"ExpiresAt": expirAt,
	}
	// 生成新的token
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newTokenStr, newTokenErr := newToken.SignedString([]byte(configKey))
	if newTokenErr != nil {
		logs.Error("生成新的token失败", newTokenErr.Error())
		return "", 0, newTokenErr
	}
	return newTokenStr, expirAt, err
}

// ParserToken token 获取
func ParserToken(tokenStr string) (string, error) {
	kv := strings.Split(tokenStr, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logs.Error("AuthString invalid:", tokenStr)
		return "", errors.New("Authorization 不存在 Bearer")
	}
	tokenResult := kv[1]
	return tokenResult, nil
}
