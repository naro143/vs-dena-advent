package model

import "time"

type Likes struct {
	_kind     string    `boom:"kind" json:"-"`
	Key       string    `boom:"id" json:"-"`
	Shinsotsu int64     `json:"shinsotsu"`
	General   int64     `json:"general"`
	UpdatedAt time.Time `json:"updated_at"`
}
