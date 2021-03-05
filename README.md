# ghost
Ghost CMS API (ContentAPI + AdminAPI)

## Supported features

* [x] Get posts
* [x] Update posts
* [ ] Get articles
* [ ] Update articles
* [x] Upload images

```go
package main

import (
	"fmt"
	"github.com/sklinkert/ghost"
)

func main() {
	contentAPIToken := "837484..."
	adminAPIToken := "90968696..."
	ghostAPI, err := ghost.New("https://example.com", contentAPIToken, adminAPIToken)
	if err != nil {
		fmt.Printf("ghost.New() failed: %v\n", err)
		return
	}

	posts, err := ghostAPI.GetPosts()
	if err != nil {
		fmt.Printf("cannot get posts from ghost api: %v\n", err)
		return
	}

	for _, post := range posts.Posts {
		fmt.Println(post.Title)

		// Update existing post
		post.Title = "new title"
		if err := ghostAPI.UpdatePost(post); err != nil {
			fmt.Printf("update failed: %v\n", err)
			break
		}
	}
	
	// Upload new image
	imageURL, err := ghostAPI.UploadImage("./myimage.jpg")
	if err != nil {
		fmt.Printf("Image upload failed: %v\n", err)
	}
	fmt.Println(imageURL)
}
```