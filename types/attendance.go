package types

type AttendanceCreate struct {
	Side  string `json:"side"`
	Name  string `json:"name"`
	Meal  string `json:"meal"`
	Count int    `json:"count"`
}
