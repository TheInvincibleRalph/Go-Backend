Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself.

Example:
            var posts []Post
            posts = append(posts, newItem)


`Vars(r)`

Vars is a function provided by the gorilla/mux package. It extracts the route variables 
(e.g posts/id: posts and id are route variables) from the request. 
These variables are the dynamic segments of the URL defined in the route pattern

`mux.Vars(r)`

This function call extracts the route variables from the HTTP request r and returns them as a map[string]string. 
In this map, the keys are the names of the variables defined in the route pattern, and the values are the corresponding 
values extracted from the URL.