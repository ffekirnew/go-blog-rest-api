package dtos

type BlogDto struct {
	BlogId int    `json:"blogId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
