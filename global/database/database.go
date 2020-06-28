package db

var (
	user_mongodb = &user{}
	registry_redis = &registry{}
	login_mongodb = &login{}
)

func User() *user {
	return user_mongodb
}

func Registry() *registry {
	return registry_redis
}

func Login() *login {
	return login_mongodb
}