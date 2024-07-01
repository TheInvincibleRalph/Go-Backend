package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Post is a struct that that represents a post
type Post struct {
	Data   string `json:"data"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

var posts []Post

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", addItems).Methods("POST")

	http.ListenAndServe(":3000", router)

}

func addItems(w http.ResponseWriter, r *http.Request) { //r accepts an http request while w is responsible for writing a response

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost) //decodes the json data from the request body into the newPost variable(which is an instance of Post)

	posts = append(posts, newPost) //appends the newPost to the posts slice us

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)

}
