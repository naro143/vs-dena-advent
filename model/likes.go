package model

import "time"

type Likes struct {
	_kind     string    `boom:"kind" json:"-"`
	ID        string    `boom:"id" json:"-"`
	Naitei    int64     `json:"naitei"`
	Shinsotsu int64     `json:"shinsotsu"`
	General   int64     `json:"general"`
	UpdatedAt time.Time `json:"updated_at"`
}
