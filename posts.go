package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *Ghost) GetPosts() (Posts, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v2/content/posts/?key=%s&limit=all"
	var posts Posts
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &posts); err != nil {
		return posts, err
	}

	return posts, nil
}

func (g *Ghost) UpdatePost(post Post) error {
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
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}
