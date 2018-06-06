

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

const (
	DISTANCE = "200km"
)


type Location struct {
      Lat float64 `json:"lat"`
      Lon float64 `json:"lon"`
}

type Post struct {
      // `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
      User     string `json:"user"`
      Message  string  `json:"message"`
      Location Location `json:"location"`
}

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
      lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
      // range is optional 
      ran := DISTANCE 
      if val := r.URL.Query().Get("range"); val != "" { 
         ran = val + "km" 
	  }
	  fmt.Println("range is ", ran)

      // Return a fake post
      p := &Post{
             User:"1111",
             Message:"一生必去的100个地方",
             Location: Location{
                    Lat:lat,
                    Lon:lon,
             },
      }

      js, err := json.Marshal(p)
      if err != nil {
             panic(err)
       //      return
      }

      w.Header().Set("Content-Type", "application/json")
      w.Write(js)


}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		   panic(err)
	//	   return
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message)
}



