package model

type Articles struct {
	_kind     string    `boom:"kind" json:"-"`
	ID        int64     `boom:"id" json:"-"`
	Shinsotsu []Article `json:"shinsotsu"`
	General   []Article `json:"general"`
}

func (as *Articles) SetOpens() {
	for i, a := range as.General {
		as.General[i].Opened = a.URL != ""
	}
	for i, a := range as.Shinsotsu {
		as.Shinsotsu[i].Opened = a.URL != ""
	}
}

func (as *Articles) SetDays() {
	for i := range as.General {
		as.General[i].Day = int64(i + 1)
	}
	for i := range as.Shinsotsu {
		as.Shinsotsu[i].Day = int64(i + 1)
	}
}

type Article struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
	Likes  int64  `json:"likes"`
	Opened bool   `json:"opened"`
	Day    int64  `json:"day"`
}
