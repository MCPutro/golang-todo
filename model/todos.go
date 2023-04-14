package model

import "time"

type Todo struct {
	Todo_id           int       `json:"id"`
	Activity_group_id int       `json:"activity_group_id"`
	Title             string    `json:"title"`
	Is_active         bool      `json:"is_active"`
	Priority          string    `json:"priority"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
