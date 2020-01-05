package domain

type Activity struct {
	ID           string
	Category     int
	Content      Content
	Rank         int
	Tags         []string
	Favorites    []string
	Gifts        []string
	UserTiny     UserTiny
}

type Content struct {
	Subject    string
	Url        string
	Comment    string
}
