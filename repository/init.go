package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RezaOptic/fiber-project-structure/config"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"time"
)

// dbs struct for managing database connections
type dbs struct {
	Redis         redis.UniversalClient
	SqlConnection *sql.Conn
}

var DBS dbs

// Init function for init databases
func Init() {
	redisConnection()
	sqlConnection()
}

// redisConnection function for connecting to redis server
func redisConnection() {
	opt := redis.UniversalOptions{
		Addrs: config.Configs.Redis.Addresses,
	}
	DBS.Redis = redis.NewUniversalClient(&opt)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := DBS.Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis connected with result :%s \n", result)
}

func sqlConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Configs.Psql.Host, config.Configs.Psql.Port, config.Configs.Psql.User, config.Configs.Psql.Password,
		config.Configs.Psql.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	DBS.SqlConnection, err = db.Conn(ctx)
	if err != nil {
		panic(err)
	}

}
