package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	ID                 string `json:"id,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	Title              string `json:"title,omitempty"`
	MobileDoc          string `json:"mobiledoc,omitempty"`
	Slug               string `json:"slug,omitempty"`
	HTML               string `json:"html,omitempty"`
	CommentID          string `json:"comment_id,omitempty"`
	FeatureImage       string `json:"feature_image,omitempty"`
	Featured           bool   `json:"featured,omitempty"`
	Page               bool   `json:"page,omitempty"`
	MetaTitle          string `json:"meta_title,omitempty"`
	MetaDescription    string `json:"meta_description,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`   // "2022-01-05T22:39:28.000Z"
	UpdatedAt          string `json:"updated_at,omitempty"`   // "2022-04-02T16:01:24.000Z"
	PublishedAt        string `json:"published_at,omitempty"` // "2022-01-19T06:31:00.000Z"
	CustomExcerpt      string `json:"custom_excerpt,omitempty"`
	OGImage            string `json:"og_image,omitempty"`
	OGTitle            string `json:"og_title,omitempty"`
	OGDescription      string `json:"og_description,omitempty"`
	TwitterImage       string `json:"twitter_image,omitempty"`
	TwitterTitle       string `json:"twitter_title,omitempty"`
	TwitterDescription string `json:"twitter_description,omitempty"`
	CustomTemplate     string `json:"custom_template,omitempty"`
	URL                string `json:"url,omitempty"`
	Excerpt            string `json:"excerpt,omitempty"`
	Tags               []Tag  `json:"tags,omitempty"`
	Status             string `json:"status,omitempty"`     // "published"
	Visibility         string `json:"visibility,omitempty"` // "public"
}

type Posts struct {
	Posts []Post `json:"posts"`
}

func (g *Ghost) AdminGetPosts() (Posts, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v3/admin/posts/?key=%s&limit=all"
	var posts Posts
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) AdminGetPostsByTag(tag string) (Posts, error) {
	var ghostPostsURLSuffix = "%s/ghost/api/v3/admin/posts/?key=%s&limit=all&formats=html,mobiledoc&filter=tag:" + tag
	var posts Posts
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) GetPosts() (Posts, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v2/content/posts/?key=%s&limit=all&include=tags"
	var posts Posts
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) GetPostsByTag(tag string) (Posts, error) {
	var ghostPostsURLSuffix = "%s/ghost/api/v2/content/posts/?key=%s&include=tags&limit=all&filter=tag:" + tag
	var posts Posts
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) AdminCreatePost(post Post) (Posts, error) {
	var posts Posts

	newPost := Posts{Posts: []Post{post}}
	updateData, err := json.Marshal(&newPost)
	if err != nil {
		return posts, err
	}

	postURL := fmt.Sprintf("%s/ghost/api/v3/admin/posts/", g.url)
	if post.HTML != "" {
		postURL = postURL + "?source=html"
	}

	if err := g.checkAndRenewJWT(); err != nil {
		return posts, err
	}

	if err := g.postJson(postURL, updateData, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) AdminUpdatePost(post Post) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	newPost := Posts{Posts: []Post{post}}
	updateData, _ := json.Marshal(&newPost)
	postUpdateURL := fmt.Sprintf("%s/ghost/api/v3/admin/posts/%s", g.url, post.ID)
	req, err := http.NewRequest(http.MethodPut, postUpdateURL, bytes.NewBuffer(updateData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost"+" "+g.jwtToken)
	resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %s", err)
		}
	}(resp.Body)

	content, _ := io.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}
