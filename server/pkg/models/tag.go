package models

import "time"

type Tag struct {
	ID        int64     `json:"id"`
	Creator   int64     `json:"creator,omitempty"`
	Title     string    `json:"title"`
	Name      string    `json:"name,omitempty"`
	Md        string    `json:"md,omitempty"`
	Cover     string    `json:"cover,omitempty"`
	Icon      string    `json:"icon"`
	PostCount int       `json:"postCount,omitempty"`
	Created   time.Time `json:"created,omitempty"`
	Updated   time.Time `json:"updated,omitempty"`
}

func (t *Tag) SetCover() {
	if t.Cover == "" {
		t.Cover = DefaultTagCover
	}
}

type Tags []*Tag

func (t Tags) Len() int      { return len(t) }
func (t Tags) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Tags) Less(i, j int) bool {
	return t[i].Created.Unix() > t[j].Created.Unix()
}
