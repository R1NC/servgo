package servgo

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

type Redis struct {}

func (rds *Redis) QueryCityIdByName(name string) string {
	conn := getConn()
	exists, err := redis.Bool(conn.Do("EXISTS", name))
	CheckErr(err)
	fmt.Printf("CityName(%s) exists:%#v\n", name, exists)
	if exists {
		cityId, err := redis.String(conn.Do("GET", name))
		CheckErr(err)
		fmt.Printf("CityId:%s\n", cityId)
		return cityId
	}
	return ""
}

func getConn() redis.Conn {
	conn, err := redis.Dial("tcp", "localhost:6379")
	CheckErr(err)
	defer conn.Close()
	return conn
} 
