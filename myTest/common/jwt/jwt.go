package jwt

import (
	"errors"
	"main/myTest/common/http"
	"main/myTest/domain/entity"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
)

const (
	CachePrefixTokenJWT = "token:jwt:"
)

type customClaims struct {
	jwt.StandardClaims
	UserId   uint
	UserType entity.UserType
}

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
		id, err := verifyToken(token, cache)
		if err != nil {
			http.ErrJson(c, err)
			c.Abort()
			return
		}

		//if newtoken != "" {
		//	c.Header("Authorization", newtoken)
		//}
		c.Set("id", id)
		c.Next()
	}
}

func verifyToken(token string, cache *redis.Client) (id uint, err error) {
	claim, err := parseToken(token)
	if err != nil {
		return
	}
	key := CachePrefixTokenJWT + claim.Id
	cacheToken, err := cache.Get(key).Result()
	if err == redis.Nil { //？？？？
		err = errors.New("登录过期，请重新登录")
		return
	}
	if err != nil {
		return
	}
	if cacheToken != token {
		err = errors.New("已在其它设备登录，请重新登陆")
		return
	}

	id = cast.ToUint(claim.Id)
	//if claim.ExpiresAt > time.Now().Unix() {
	//	return
	//}
	//newtoken, err = GenToken(id, store.CACHE)
	return
}

func GenToken(userId uint, userType entity.UserType, cache *redis.Client) (newtoken string, err error) {
	now := time.Now().Unix()
	if newtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now,
			ExpiresAt: now + int64(2*time.Hour/time.Second),
		},
		UserId:   userId,
		UserType: userType,
	}).SignedString([]byte("JWT1OGHNS0MSE9AUKDCDOMB65BNLVSA0")); err != nil {
		return
	}
	key := CachePrefixTokenJWT + cast.ToString(userId)
	if err = cache.Set(key, newtoken, 2*time.Hour).Err(); err != nil {
		return
	}
	return
}

func parseToken(token string) (*customClaims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &customClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte("JWT1OGHNS0MSE9AUKDCDOMB65BNLVSA0"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenClaim.Claims.(*customClaims); ok && tokenClaim.Valid {
		return claims, nil
	}
	return nil, err
}
