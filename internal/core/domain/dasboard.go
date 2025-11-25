package domain

type Dashboard struct {
	ID          int
	Title       string //`json:"title"`
	Description string //`json:"description"`
	CreatedAt   int64  //`json:"created_at"`
}
