package server

import (
	"fmt"
	"io"
	
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
		
		var urlString string
        // fixs here needed 
		// err handling through out can be done more neatly
		for _,i := range resources{
			if r.URL.Path == i.Endpoint {
				urlString = i.Url
				break
			}
		}

		if urlString == "" {
			log.Fatal("url not found")
		}
	
		req,err := http.NewRequest(r.Method,urlString,r.Body)
		if err != nil{
			log.Fatal(err)
		}
		req.Header = r.Header

		resp, err := http.DefaultClient.Do(req)
        if err != nil {
			log.Fatal(err)
        }
        defer resp.Body.Close()
		body,err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w,string(body))
	}
}


func Server(resources []Resources){
	log.Println("starting server at port 8000")
	const port string = ":8000"
	// handel all requests 
	// hit the url and send back the response
	http.HandleFunc("/",Home(resources))
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Fatal(err)
	}
}