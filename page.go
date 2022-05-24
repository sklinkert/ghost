package ghost

import "fmt"

type Pages struct {
	Pages []Page `json:"pages"`
}

type Page struct {
	ID                 string `json:"id,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	Title              string `json:"title,omitempty"`
	MobileDoc          string `json:"mobiledoc"`
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
}

func (g *Ghost) GetPages() (Pages, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v2/content/pages/?key=%s&limit=all"
	var pages Pages
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}
