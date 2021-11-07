package model

type GetBooksRequest struct {
	Name string `json:"name"`
}

type GetBooksResponse struct {
	Data []Book `json:"data"`
}
type Book struct {
	Name     string `json:"full_name"`
	Author   string `json:"author"`
	ImageUrl string `json:"imageurl"`
}
