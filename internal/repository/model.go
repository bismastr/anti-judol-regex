package repository

import "time"

type Regex struct {
	Id        int
	Word      string
	Regex     string
	CreatedAt time.Time
}
