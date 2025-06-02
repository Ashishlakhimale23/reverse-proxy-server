package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type Resources struct {
	Name string 
	Endpoint string
	Url string
}

func Home(resources []Resources) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		urlString:= r.URL.Path
		var url string
		for _,i := range resources{
			if i.Endpoint == urlString{
				url = i.Url
				break
			} 
		}

		fmt.Println(url)
		resp,err := http.Get(url)
		if err != nil{
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(w,string(body))

	}
}


func Server(resources []Resources){
	log.Println("starting server at port 80")
	const port string = ":8000"
	// handel all requests 
	// hit the url and send back the response
	http.HandleFunc("/",Home(resources))
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal(err)
	}
}