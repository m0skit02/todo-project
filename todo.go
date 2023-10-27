package todo

type TodoList struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
