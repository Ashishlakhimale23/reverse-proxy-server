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
		
		for _,i := range resources{
			if r.Host == i.Endpoint {
				urlString = i.Url
				break
			}
		}

		if urlString == "" {
			log.Println("required mapping not ")
			fmt.Println(w,"required mapping not found ")
			return
		}
	
		req,err := http.NewRequest(r.Method,urlString,r.Body)
		if err != nil{
			log.Println(err)
			fmt.Fprintln(w,err.Error())
			return
		}
		req.URL.Path = r.URL.Path
		req.Header = r.Header
		
		resp, err := http.DefaultClient.Do(req)
        if err != nil {
			log.Println(err)
			fmt.Fprintln(w,err.Error())
			return
        }
	
        defer resp.Body.Close()
		body,err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		    fmt.Fprintln(w,err.Error())
			return
		}
		fmt.Fprintln(w,string(body))
	}
}



func Server(resources []Resources){
	log.Println("starting server at port 8o")
	const port string = ":80"
	// handel all requests 
	// hit the url and send back the response
	http.HandleFunc("/",Home(resources))
	err := http.ListenAndServe(port,nil)
	if err != nil {
		log.Println(err)
	}
}