package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Post is a struct that that represents a post
type Post struct {
	Data   string `json:"data"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// User is a struct that represent a user in our application
type User struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post //a slice that takes in every post.

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", addItems).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getApost).Methods("GET")
	router.HandleFunc("/posts/{id}", updateApost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":3000", router)

}

func getApost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	post := posts[id]
	json.NewEncoder(w).Encode(post)

}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addItems(w http.ResponseWriter, r *http.Request) { //r accepts an http request while w is responsible for writing a response
	w.Header().Set("Content-Type", "application/json")

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost) //decodes the json data from the request body into the newPost variable(which is an instance of Post)

	posts = append(posts, newPost) //appends the newPost to the posts slice
	json.NewEncoder(w).Encode(posts)

}

func updateApost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	//get the Id of the post from the root parameters
	var idParam string = mux.Vars(r)["id"] //mux.Vars takes in the incoming request and returns the route variables, if any. However, a specific variable can be wrapped in a square bracket for retrieval.
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//get the value from the JSON body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost
	json.NewEncoder(w).Encode(updatedPost)
}

func patchPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get the Id of the post from the root parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//get the current value
	post := posts[id]
	json.NewDecoder(r.Body).Decode(&post)

	posts[id] = post //a pointer can be used above to replace this (such that post := &posts[id] can be used instead of using post := posts[id])
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get the Id of the post from the root parameters
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	//delete the post from the slice
	//https://go.dev/wiki/SliceTricks
	posts = append(posts[:id], posts[id+1:]...)
	w.WriteHeader(200)

}
