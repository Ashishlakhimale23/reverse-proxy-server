package main

import (
	"log"
	"os"
	"example.com/reverse-proxy/server"
	"gopkg.in/yaml.v3"
)

type Configuration struct{
	Resources []server.Resources
}

func main(){
	var res1 Configuration

	f,err :=os.ReadFile("config.yaml")

    if err!=nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(f,&res1)
	resource := res1.Resources

	server.Server(resource)

}