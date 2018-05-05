package main

import (
    "encoding/json"
//	"fmt"
    "io/ioutil"
	"gopkg.in/mgo.v2"
)


func main() {
    data, err := ioutil.ReadFile("movies.json")
    if err != nil {
        return
	}
	session, err := mgo.Dial("localhost")
	if err != nil {
			panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("Moviedata").C("movie")
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {   
	}
	if err := c.Insert(v...); err != nil {
	}
	
}

