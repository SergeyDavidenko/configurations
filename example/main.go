package main

import (
	"log"

	"github.com/SergeyDavidenko/configurations/client"
)

func init() {
	url := "http://localhost:9971/auth2"
	config, err := client.GetConfigFrom(url, "admin", "supersecret")
	if err != nil {
		log.Fatal(err)
	}
	client.Config = config

}

func main() {
	prop := client.GetBoolOrDefault("blcache.kafka.enabled", true)
	prop2 := client.GetIntOrDefault("blcache.sync.fixed.delay.ms", 111)
	prop3 := client.GetStringOrDefault("test", "net")
	prop4 := client.GetIntOrDefault("spring.redis.blcache.port", 6380)
	log.Println(prop)
	log.Println(prop2)
	log.Println(prop3)
	log.Println(prop4)

}
