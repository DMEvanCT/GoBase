package Database

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)


func OpenDB(dbusername, dbpass, serverip string ) (*sql.DB) {
	db, err := sql.Open("mysql", dbusername + ":" + dbpass +  "@tcp(" + serverip + ")" + "/")
	if err != nil {
		log.Fatal("Sorry there was a problem connecting to the database with user " + dbusername + " host " + serverip +  " pass " + dbpass + " Please check /etc/commservice/credentials.yaml")
		log.Fatal(err)

	}
	return db
}

// Specific for the auth db
func DatabaseInitAuth() (*sql.DB)  {
	viper.AddConfigPath("/etc/commservice/")
	viper.SetConfigName("comconfig")
	viper.ReadInConfig()
	dbusername := viper.GetString("authdb.username")
	dbpass := viper.GetString("authdb.password")
	serverip := viper.GetString("authdb.dbhost")
	databaseconn := OpenDB(dbusername, dbpass, serverip)


	return databaseconn

}

// Generic database connection configpath ex: /etc/comservice/ configname ex commservice username db.username password db.password host db.host
func DatabaseInitAll(configpath, configname, usernanme, password, host string) (*sql.DB)  {
	viper.AddConfigPath(configpath)
	viper.SetConfigName(configname)
	viper.ReadInConfig()
	dbusername := viper.GetString(usernanme)
	dbpass := viper.GetString(password)
	serverip := viper.GetString(host)
	databaseconn := OpenDB(dbusername, dbpass, serverip)


	return databaseconn
}


// For use with the DBByEnvFunction
func DatabaseInitAllHost(configpath, configname, usernanme, password, host string) (*sql.DB)  {
	viper.AddConfigPath(configpath)
	viper.SetConfigName(configname)
	viper.ReadInConfig()
	dbusername := viper.GetString(usernanme)
	dbpass := viper.GetString(password)
	serverip := host

	db, err := sql.Open("mysql", dbusername + ":" + dbpass +  "@tcp(" + serverip + ")" + "/")
	if err != nil {
		log.Fatal("Sorry there was a problem connecting to the database with user " + dbusername + " host " + serverip +  " pass " + dbpass + " Please check /etc/commservice/credentials.yaml")
		log.Fatal(err)

	}
	return db

}

// Used to connect to Redis
func RedisInit(redishost, redispass string, db int)  (*redis.Client){
	client := redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: redispass, // no password set
		DB:       db,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}



