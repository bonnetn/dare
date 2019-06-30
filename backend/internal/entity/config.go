package entity

type Configuration struct {
	Database DatabaseConfiguration `json:"database,required"`
	Redis    RedisConfiguration    `json:"redis,required"`
}

type DatabaseConfiguration struct {
	Host     string `json:"host,required"`
	Database string `json:"database,required"`
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type RedisConfiguration struct {
	Host     string `json:"host,required"`
	Password string `json:"password,required"`
	DB       int    `json:"db,required"`
}
