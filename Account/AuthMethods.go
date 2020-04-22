package Account

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"net/http"
	"strings"
	"time"
)

var (
	client *redis.Client
)
func AuthRedisConnection (conf RedisConfig) {
	client = RedisConnection(conf)
}

func RedisConnection(conf RedisConfig) *redis.Client{
	newclient := redis.NewClient(&redis.Options{
		Addr: conf.Addr,
		Password: conf.Password,
		DB: conf.DB,
	})
	return newclient
}


func CreateToken() string{
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}
func CustomAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), " ")

		tokenFromRedis, _ := client.Get(token[1]).Result()
		if len(tokenFromRedis) != 0 {
			next.ServeHTTP(w, r)
		} else{
			http.Error(w, "error Ros9 hello", http.StatusBadRequest)
			return
		}
	})
}
