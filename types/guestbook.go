package types

type GuestbookPostForGet struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp uint64 `json:"timestamp"`
}

type GuestbookPostForCreate struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Password string `json:"password"`
}

type GuestbookPostForDelete struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type GuestbookGetResponse struct {
	Posts []GuestbookPostForGet `json:"posts"`
	Total int                   `json:"total"`
}
