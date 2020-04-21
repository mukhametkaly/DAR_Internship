package Account

type Account struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type RedisConfig struct {
	Addr string
	Password string
	DB int
}


