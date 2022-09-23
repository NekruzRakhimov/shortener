package db

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

var redisConn redis.Conn

func initRedis() redis.Conn {
	c, err := redis.DialURL(os.Getenv("REDIS_URL"), redis.DialTLSSkipVerify(true))
	if err != nil {
		// Handle error
	}
	defer c.Close()

	//settingParams := utils.AppSettings.PostgresParams

	//connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	//	settingParams.Server,
	//	settingParams.Port,
	//	settingParams.User,
	//	settingParams.DataBase,
	//	settingParams.Password)
	//
	//db, err := gorm.Open("postgres", connString)
	//if err != nil {
	//	logger.Error.Fatal("Couldn't connect to database", err.Error())
	//}
	//
	//// enabling gorm log mode, used for debugging
	//db.LogMode(true)
	//
	//db.SingularTable(true)

	return c
}

// StartRedisConnection Creates connection to database
func StartRedisConnection() {
	redisConn = initRedis()
}

// GetRedisConn func for getting db conn globally
func GetRedisConn() redis.Conn {
	return redisConn
}
