package domain

type Todo struct {
	Content   string `json:content`
	CreatedAt int    `json:createdAt`
}
