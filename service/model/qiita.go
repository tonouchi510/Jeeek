package model

type QiitaPost struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Url           string    `json:"url"`
	LikesCount    int64     `json:"likes_count"`
	Tags          []Tag     `json:"tags"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

type Tag struct {
	Name        string     `json:"name"`
	Versions    []string   `json:"versions"`
}
