package types

type PostGet struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp uint64 `json:"timestamp"`
}

type PostCreate struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Password string `json:"password"`
}

type PostDelete struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type PostsGetResponse struct {
	Posts []PostGet `json:"posts"`
	Total int       `json:"total"`
}
