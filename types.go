package ghost

import "time"

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID                 string    `json:"id,omitempty"`
	UUID               string    `json:"uuid,omitempty"`
	Title              string    `json:"title,omitempty"`
	Slug               string    `json:"slug,omitempty"`
	HTML               string    `json:"html,omitempty"`
	CommentID          string    `json:"comment_id,omitempty"`
	FeatureImage       string    `json:"feature_image,omitempty"`
	Featured           bool      `json:"featured,omitempty"`
	Page               bool      `json:"page,omitempty"`
	MetaTitle          string    `json:"meta_title,omitempty"`
	MetaDescription    string    `json:"meta_description,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          string    `json:"updated_at,omitempty"`
	PublishedAt        time.Time `json:"published_at,omitempty"`
	CustomExcerpt      string    `json:"custom_excerpt,omitempty"`
	OGImage            string    `json:"og_image,omitempty"`
	OGTitle            string    `json:"og_title,omitempty"`
	OGDescription      string    `json:"og_description,omitempty"`
	TwitterImage       string    `json:"twitter_image,omitempty"`
	TwitterTitle       string    `json:"twitter_title,omitempty"`
	TwitterDescription string    `json:"twitter_description,omitempty"`
	CustomTemplate     string    `json:"custom_template,omitempty"`
	URL                string    `json:"url,omitempty"`
	Excerpt            string    `json:"excerpt,omitempty"`
}

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}
