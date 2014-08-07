package person

import "time"

type Person struct {
	Id        int64
	Name      string
	Age       int
	Email     string
	Job       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
