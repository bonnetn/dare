package entity

type Configuration struct {
	MongoDB MongoDBConfiguration `json:"mongodb,required"`
}

type MongoDBConfiguration struct {
	Host     string `json:"host,required"`
	Database string `json:"database,required"`
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

