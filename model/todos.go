package model

type Todos struct {
	todo_id           int    `json:"id"`
	activity_group_id int    `json:"activity_group_id"`
	title             string `json:"title"`
	is_active         bool   `json:"is_active"`
	priority          string `json:"priority"`
	created_at        string `json:"createdAt"`
	updated_at        string `json:"updatedAt"`
}
