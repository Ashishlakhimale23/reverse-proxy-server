package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Resources struct {
	Name string 
	Endpoint string
	Url string
}
type Server struct{
	Host string
	Port string
}

type Configuration struct{
	Server Server
	Resources []Resources
}

func main(){
	var res1 Configuration

	f,err :=os.ReadFile("config.yaml")

    if err!=nil {
		log.Fatal(err)
	}
	
	yaml.Unmarshal(f,&res1)

	fmt.Printf("%v",res1)
}