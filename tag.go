package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Tags struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	CreatedAt          time.Time `json:"created_at,omitempty"`
	Description        string    `json:"description,omitempty"`
	FeatureImage       string    `json:"feature_image,omitempty"`
	Id                 string    `json:"id,omitempty"`
	MetaDescription    string    `json:"meta_description,omitempty"`
	MetaTitle          string    `json:"meta_title,omitempty"`
	Name               string    `json:"name,omitempty"`
	Slug               string    `json:"slug,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	Url                string    `json:"url,omitempty"`
	Visibility         string    `json:"visibility,omitempty"`
	TwitterImage       string    `json:"twitter_image,omitempty"`
	TwitterTitle       string    `json:"twitter_title,omitempty"`
	TwitterDescription string    `json:"twitter_description,omitempty"`
	CodeInjectionHead  string    `json:"codeinjection_head,omitempty"`
	CodeInjectionFoot  string    `json:"codeinjection_foot,omitempty"`
	CanonicalURL       string    `json:"canonical_url,omitempty"`
	AccentColor        string    `json:"accent_color,omitempty"`
	Count              struct {
		Posts int `json:"posts,omitempty"`
	} `json:"count,omitempty,skip"`
}

type NewTags struct {
	Tags []NewTag `json:"tags"`
}

// NewTag - two struct because field "Count" must not exist when creating a new tag
type NewTag struct {
	CreatedAt          time.Time `json:"created_at,omitempty"`
	Description        string    `json:"description,omitempty"`
	FeatureImage       string    `json:"feature_image,omitempty"`
	Id                 string    `json:"id,omitempty"`
	MetaDescription    string    `json:"meta_description,omitempty"`
	MetaTitle          string    `json:"meta_title,omitempty"`
	Name               string    `json:"name,omitempty"`
	Slug               string    `json:"slug,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	Url                string    `json:"url,omitempty"`
	Visibility         string    `json:"visibility,omitempty"`
	TwitterImage       string    `json:"twitter_image,omitempty"`
	TwitterTitle       string    `json:"twitter_title,omitempty"`
	TwitterDescription string    `json:"twitter_description,omitempty"`
	CodeInjectionHead  string    `json:"codeinjection_head,omitempty"`
	CodeInjectionFoot  string    `json:"codeinjection_foot,omitempty"`
	CanonicalURL       string    `json:"canonical_url,omitempty"`
	AccentColor        string    `json:"accent_color,omitempty"`
}

func (g *Ghost) AdminGetTags() (Tags, error) {
	var ghostTagsURLSuffix = "%s/ghost/api/v3/admin/tags/?key=%s&limit=all&include=count.posts"
	var tags Tags
	var url = fmt.Sprintf(ghostTagsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &tags); err != nil {
		return tags, err
	}

	return tags, nil
}

func (g *Ghost) AdminCreateTags(tags NewTags) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	var ghostTagsURLSuffix = "%s/ghost/api/v3/admin/tags/?key=%s"
	var url = fmt.Sprintf(ghostTagsURLSuffix, g.url, g.contentAPIToken)

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

	content, _ := io.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}

func (g *Ghost) AdminDeleteTag(tag Tag) error {
	var ghostTagsURLSuffix = "%s/ghost/api/v3/admin/tags/%s/?key=%s"
	var url = fmt.Sprintf(ghostTagsURLSuffix, g.url, tag.Id, g.contentAPIToken)

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

	content, _ := io.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}

func (g *Ghost) AdminUpdateTag(tag Tag) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	newTag := Tags{Tags: []Tag{tag}}
	updateData, _ := json.Marshal(&newTag)
	tagUpdateURL := fmt.Sprintf("%s/ghost/api/v3/admin/tags/%s", g.url, tag.Id)
	req, err := http.NewRequest(http.MethodPut, tagUpdateURL, bytes.NewBuffer(updateData))
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
