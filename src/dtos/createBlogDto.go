package dtos

type CreateBlogDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
