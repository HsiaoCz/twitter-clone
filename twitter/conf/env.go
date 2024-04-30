package conf

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return ":8000"
	}
	return port
}

func GetMongoDBUrl(key string) string {
	url := os.Getenv(key)
	if url == "" {
		return "mongodb://localhost:27017"
	}
	return url
}

func GetMongoDBName(key string) string {
	dbname := os.Getenv(key)
	if dbname == "" {
		return "twitter"
	}
	return dbname
}

func GetMongoDBUserColl(key string) string {
	userColl := os.Getenv(key)
	if userColl == "" {
		return "users"
	}
	return userColl
}
