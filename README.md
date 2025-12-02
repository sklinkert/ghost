# Unofficial Go Client for Ghost Blogs

Not affiliated in any way with Ghost.org.

[Ghost](https://ghost.org/) Client (ContentAPI + AdminAPI)

## Installation

```bash
go get github.com/sklinkert/ghost
```

## Supported features

### Posts
* [x] Get posts (Content API + Admin API)
* [x] Get post by ID
* [x] Get posts by tag
* [x] Add post
* [x] Update post
* [x] Delete post
* [x] Search posts

### Pages
* [x] Get pages (Content API + Admin API)
* [x] Get page by ID
* [x] Add page
* [x] Update page
* [x] Delete page

### Tags
* [x] Get tags
* [x] Add tags
* [x] Update tag
* [x] Delete tag

### Members
* [x] Get members (with pagination)
* [x] Get member by ID
* [x] Add member
* [x] Delete member

### Images
* [x] Upload image

## Usage

### Initialization

```go
package main

import (
	"github.com/sklinkert/ghost"
	"net/http"
	"time"
)

func main() {
	contentAPIToken := "837484..."
	adminAPIToken := "90968696..."

	// Default usage
	ghostAPI := ghost.New("https://example.com", contentAPIToken, adminAPIToken)

	// With custom HTTP client
	customClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	ghostAPI = ghost.New("https://example.com", contentAPIToken, adminAPIToken, customClient)
}
```

### Posts

```go
// Get all posts (Content API)
posts, err := ghostAPI.GetPosts()

// Get all posts (Admin API)
posts, err := ghostAPI.AdminGetPosts()

// Get post by ID
posts, err := ghostAPI.AdminGetPost("628f557f0a8ce9486eb37623")

// Get posts by tag
posts, err := ghostAPI.AdminGetPostsByTag("news")

// Search posts
posts, err := ghostAPI.AdminSearchPosts("search query")

// Create a new post
newPost := ghost.Post{
	Title:  "My New Post",
	HTML:   "<p>Post content here</p>",
	Status: ghost.StatusPublished,
}
posts, err := ghostAPI.AdminCreatePost(newPost)

// Update a post
post.Title = "Updated Title"
err := ghostAPI.AdminUpdatePost(post, ghost.SourceHTML)

// Delete a post
err := ghostAPI.AdminDeletePost("628f557f0a8ce9486eb37623")
```

### Pages

```go
// Get all pages (Content API)
pages, err := ghostAPI.GetPages()

// Get all pages (Admin API)
pages, err := ghostAPI.AdminGetPages()

// Get page by ID
pages, err := ghostAPI.AdminGetPage("628f557f0a8ce9486eb37623")

// Create a new page
newPage := ghost.Page{
	Title:  "My New Page",
	HTML:   "<p>Page content here</p>",
	Status: ghost.StatusPublished,
}
pages, err := ghostAPI.AdminCreatePage(newPage)

// Update a page
page.Title = "Updated Title"
err := ghostAPI.AdminUpdatePage(page, ghost.SourceHTML)

// Delete a page
err := ghostAPI.AdminDeletePage("628f557f0a8ce9486eb37623")
```

### Tags

```go
// Get all tags
tags, err := ghostAPI.AdminGetTags()

// Create new tags
newTags := ghost.NewTags{
	Tags: []ghost.NewTag{
		{Name: "Technology", Slug: "technology"},
		{Name: "News", Slug: "news"},
	},
}
err := ghostAPI.AdminCreateTags(newTags)

// Update a tag
tag.Name = "Updated Name"
err := ghostAPI.AdminUpdateTag(tag)

// Delete a tag
err := ghostAPI.AdminDeleteTag(tag)
```

### Members

```go
// Get all members (with automatic pagination)
members, err := ghostAPI.AdminGetMembers()

// Get member by ID
members, err := ghostAPI.AdminGetMember("691ca681b7c6ec3a01a2ba81")
if err == nil && len(members.Members) > 0 {
	member := members.Members[0]
	fmt.Printf("Member: %s (%s)\n", member.Name, member.Email)
}

// Create a new member
newMember := ghost.NewMember{
	Name:  "John Doe",
	Email: "john@example.com",
}
members, err := ghostAPI.AdminCreateMember(newMember)

// Delete a member
err := ghostAPI.AdminDeleteMember("691ca681b7c6ec3a01a2ba81")
```

### Images

```go
// Upload an image
imageURL, err := ghostAPI.AdminUploadImage("./myimage.jpg")
if err != nil {
	fmt.Printf("Image upload failed: %v\n", err)
}
fmt.Println(imageURL)
```

## API Reference

### Client Initialization

```go
func New(url, contentAPIToken, adminAPIToken string, client ...*http.Client) *Ghost
```

### Posts

| Method | Description |
|--------|-------------|
| `GetPosts()` | Get all posts via Content API |
| `GetPost(postId)` | Get a single post via Content API |
| `GetPostsByTag(tag)` | Get posts by tag via Content API |
| `AdminGetPosts()` | Get all posts via Admin API |
| `AdminGetPost(postId)` | Get a single post via Admin API |
| `AdminGetPostsByTag(tag)` | Get posts by tag via Admin API |
| `AdminCreatePost(post)` | Create a new post |
| `AdminUpdatePost(post, sourceType)` | Update an existing post |
| `AdminDeletePost(postId)` | Delete a post |
| `AdminSearchPosts(query)` | Search posts by title or excerpt |

### Pages

| Method | Description |
|--------|-------------|
| `GetPages()` | Get all pages via Content API |
| `AdminGetPages()` | Get all pages via Admin API |
| `AdminGetPage(pageId)` | Get a single page via Admin API |
| `AdminCreatePage(page)` | Create a new page |
| `AdminUpdatePage(page, sourceType)` | Update an existing page |
| `AdminDeletePage(pageId)` | Delete a page |

### Tags

| Method | Description |
|--------|-------------|
| `AdminGetTags()` | Get all tags (with pagination) |
| `AdminCreateTags(tags)` | Create new tags |
| `AdminUpdateTag(tag)` | Update an existing tag |
| `AdminDeleteTag(tag)` | Delete a tag |

### Members

| Method | Description |
|--------|-------------|
| `AdminGetMembers()` | Get all members (with pagination) |
| `AdminGetMember(memberId)` | Get a single member by ID |
| `AdminCreateMember(member)` | Create a new member |
| `AdminDeleteMember(memberId)` | Delete a member |

### Images

| Method | Description |
|--------|-------------|
| `AdminUploadImage(path)` | Upload an image file |
