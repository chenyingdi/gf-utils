package gfUtils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"time"
)

func GeneToken(args g.Map, key string, duration time.Duration) (string, error) {
	claim := make(jwt.MapClaims)
	claim["exp"] = gtime.Now().Add(duration).Unix()
	claim["iat"] = gtime.Now().Unix()

	if args != nil {
		for k, v := range args {
			claim[k] = v
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = claim

	return token.SignedString([]byte(key))
}

func ParseToken(tokenStr, key string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token无效！")
	}

	return token.Claims.(jwt.MapClaims), nil
}
