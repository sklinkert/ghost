package ghost

import "time"

type Posts struct {
	Posts []Post `json:"posts"`
}

// StatusPublished indicates if a post or pages is already published
const StatusPublished = "published"

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

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}

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
	} `json:"count,omitempty"`
}
