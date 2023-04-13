package model

import "time"

type Activity struct {
	Activity_id int       `json:"id"`
	Title       string    `json:"title"`
	Email       string    `json:"email"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}
