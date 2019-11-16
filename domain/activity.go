package domain

type Activity struct {
	ID           string
	Category     int
	Content      Content
	Rank         int
	Tags         []string
	User         UserTiny
}

type Content struct {
	Subject    string
	Url        string
	Comment    string
}
