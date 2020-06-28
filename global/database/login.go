package db

// login redis(Redis)
type login struct {}

func (l login) RedisKey(item string) string {
	return "login_" + item
}