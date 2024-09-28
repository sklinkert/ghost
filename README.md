# Unofficial Go Client for Ghost Blogs 

Not affiliated in any way with Ghost.org.

[Ghost](https://ghost.org/) Client (ContentAPI + AdminAPI)

## Supported features

### Generic

* [x] Upload images

### Posts
* [x] Add post
* [x] Get posts
* [x] Update post
* [ ] Delete post

### Pages
* [ ] Add page
* [x] Get pages
* [ ] Update page
* [ ] Delete page

### Tags
* [x] Add tag
* [x] Get tags
* [x] Update tag
* [x] Delete tag

### Members
* [x] Add member
* [x] Get members
* [ ] Update member

### Images
* [x] Add image

```go
package main

import (
	"fmt"
	"github.com/sklinkert/ghost"
)

func main() {
	contentAPIToken := "837484..."
	adminAPIToken := "90968696..."
	ghostAPI := ghost.New("https://example.com", contentAPIToken, adminAPIToken)

	posts, err := ghostAPI.GetPosts()
	if err != nil {
		fmt.Printf("cannot get posts from ghost api: %v\n", err)
		return
	}

	for _, post := range posts.Posts {
		fmt.Println(post.Title)

		// Update existing post
		post.Title = "new title"
		if err := ghostAPI.AdminUpdatePost(post); err != nil {
			fmt.Printf("update failed: %v\n", err)
			break
		}
	}
	
	// Upload new image
	imageURL, err := ghostAPI.AdminUploadImage("./myimage.jpg")
	if err != nil {
		fmt.Printf("Image upload failed: %v\n", err)
	}
	fmt.Println(imageURL)
}
```
