package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (g *Ghost) AdminGetTags() (Tags, error) {
	var ghostPostsURLSuffix = "%s/ghost/api/v3/admin/tags/?key=%s&limit=all&include=count.posts"
	var tags Tags
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &tags); err != nil {
		return tags, err
	}

	return tags, nil
}

func (g *Ghost) AdminCreateTags(tags Tags) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	var ghostPostsURLSuffix = "%s/ghost/api/v3/admin/tags/?key=%s"
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	updateData, _ := json.Marshal(&tags)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(updateData))
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
			fmt.Printf("Error closing Body: %v", err)
		}
	}(resp.Body)

	content, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}

func (g *Ghost) AdminDeleteTag(tag Tag) error {
	var ghostPostsURLSuffix = "%s/ghost/api/v3/admin/tags/%s/?key=%s"
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, tag.Id, g.contentAPIToken)

	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
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
			fmt.Printf("Error closing Body: %v", err)
		}
	}(resp.Body)

	content, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}
