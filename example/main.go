package main

import (
	"log"

	"github.com/SergeyDavidenko/configurations/client"
)

func init() {
	url := "http://localhost:9971/auth2"
	config, err := client.GetConfigFrom(url, "someuser", "supersecretpassword")
	if err != nil {
		log.Fatal(err)
	}
	client.Config = config

}

func main() {
	prop := client.GetBoolOrDefault("kafka.enabled", true)
	prop2 := client.GetIntOrDefault("fixed.delay.ms", 111)
	prop3 := client.GetStringOrDefault("test", "net")
	prop4 := client.GetIntOrDefault("redis.port", 6380)
	log.Println(prop)
	log.Println(prop2)
	log.Println(prop3)
	log.Println(prop4)

}
