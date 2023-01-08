package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"main/myTest/http"
	"strings"
	"time"
)

const (
	CachePrefixTokenJWT    = "token:jwt:"
	cacheKeyTokenJWTSystem = "token:jwt:system"
)

func JWT(cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			err := errors.New("Authorization为空，请重新登录")
			http.ErrJson(c, err)
			c.Abort()
			return
		}

		auths := strings.Split(bearerToken, " ")
		if len(auths) != 2 || auths[0] != "Bearer" {
			err := errors.New("Authorization格式错误，请重新登录")
			http.ErrJson(c, err)
			c.Abort()
			return
		}

		token := auths[1]
		id, newtoken, err := verifyToken(token, cache)
		if err != nil {
			http.ErrJson(c, err)
			c.Abort()
			return
		}

		if newtoken != "" {
			c.Header("Authorization", newtoken)
		}
		c.Set("id", id)
		c.Next()
	}
}

func verifyToken(token string, cache *redis.Client) (id uint, newtoken string, err error) {
	claim, err := parseToken(token)
	if err != nil {
		return
	}
	key := CachePrefixTokenJWT + claim.Id
	cacheToken, err := cache.Get(key).Result()
	if err == redis.Nil { //？？？？
		err = errors.New("已在其它设备登录，请重新登陆")
		return
	}
	if err != nil {
		return
	}
	if cacheToken != token {
		err = errors.New("已在其它设备登录，请重新登陆")
		return
	}

	id = cast.ToUint(claim.Id) //？？？？
	if claim.ExpiresAt > time.Now().Unix() {
		return
	}
	newtoken, err = GenToken(id)
	return
}

func GenToken(id uint) (newtoken string, err error) {
	now := time.Now().Unix()
	if newtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        cast.ToString(id),
		IssuedAt:  now,
		ExpiresAt: now + int64(2*time.Hour/time.Second),
	}).SignedString([]byte("JWT1OGHNS0MSE9AUKDCDOMB65BNLVSA0")); err != nil {
		return
	}
	return
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte("JWT1OGHNS0MSE9AUKDCDOMB65BNLVSA0"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaim.Claims.(*jwt.StandardClaims); ok && tokenClaim.Valid {
		return claims, nil
	}
	return nil, err
}
