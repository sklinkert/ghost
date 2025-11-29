package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pages struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	ID                 string `json:"id,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	Title              string `json:"title,omitempty"`
	MobileDoc          string `json:"mobiledoc,omitempty"`
	Lexical            string `json:"lexical,omitempty"`
	Slug               string `json:"slug,omitempty"`
	HTML               string `json:"html,omitempty"`
	CommentID          string `json:"comment_id,omitempty"`
	FeatureImage       string `json:"feature_image,omitempty"`
	Featured           bool   `json:"featured,omitempty"`
	Page               bool   `json:"page,omitempty"`
	MetaTitle          string `json:"meta_title,omitempty"`
	MetaDescription    string `json:"meta_description,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	PublishedAt        string `json:"published_at,omitempty"`
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
	Status             string `json:"status,omitempty"`
	Visibility         string `json:"visibility,omitempty"`
	Tags               []Tag  `json:"tags,omitempty"`
}

func (g *Ghost) GetPages() (Pages, error) {
	const ghostPagesURLSuffix = "%s/ghost/api/v2/content/pages/?key=%s&limit=all"
	var pages Pages
	var url = fmt.Sprintf(ghostPagesURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}

func (g *Ghost) AdminGetPages() (Pages, error) {
	const ghostPagesURLSuffix = "%s/ghost/api/v3/admin/pages/?key=%s&limit=all&include=tags&formats=html,lexical,mobiledoc"
	var pages Pages
	var url = fmt.Sprintf(ghostPagesURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}

func (g *Ghost) AdminGetPage(pageId string) (Pages, error) {
	const ghostPagesURLSuffix = "%s/ghost/api/v3/admin/pages/%s/?key=%s&include=tags&formats=html,lexical,mobiledoc"
	var pages Pages
	var url = fmt.Sprintf(ghostPagesURLSuffix, g.url, pageId, g.contentAPIToken)

	if err := g.getJson(url, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}

func (g *Ghost) AdminCreatePage(page Page) (Pages, error) {
	var pages Pages

	newPage := Pages{Pages: []Page{page}}
	createData, err := json.Marshal(&newPage)
	if err != nil {
		return pages, err
	}

	pageURL := fmt.Sprintf("%s/ghost/api/v3/admin/pages/", g.url)
	if page.HTML != "" {
		pageURL = pageURL + "?source=html"
	}

	if err := g.postJson(pageURL, createData, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}

func (g *Ghost) AdminUpdatePage(page Page, sourceType SourceType) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	newPage := Pages{Pages: []Page{page}}
	updateData, err := json.Marshal(&newPage)
	if err != nil {
		return err
	}

	pageUpdateURL := fmt.Sprintf("%s/ghost/api/v3/admin/pages/%s?save_revision=1", g.url, page.ID)
	if sourceType != "" {
		pageUpdateURL = pageUpdateURL + "&source=" + string(sourceType)
	}

	req, err := http.NewRequest(http.MethodPut, pageUpdateURL, bytes.NewBuffer(updateData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost "+g.jwtToken)
	resp, err := g.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	content, _ := io.ReadAll(resp.Body)
	responseBody := string(content)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}
	return nil
}
