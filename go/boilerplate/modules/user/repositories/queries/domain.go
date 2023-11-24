package queries

type Users struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdat"`
}
