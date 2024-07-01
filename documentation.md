Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself.

Example:
            var posts []Post
            posts = append(posts, newItem)