package postgresql

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
