package book

type Book2 struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Pages  int    `json:"pages"`
	Link   string `json:"link"`
}
