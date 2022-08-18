package types

import "time"

type Todo struct {
	ID          int64
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
