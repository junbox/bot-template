package models

import "time"

type Tag struct {
	Caption string
}

type Profile struct {
	Active     bool
	Registered int64
	Statuses   []string
	LastVisit  int64
	IsAdmin    bool
	Achives    []Tag
}

func (p *Profile) Init() {
	p.Active = true
	p.Registered = time.Now().Unix()
	p.Statuses = append(make([]string, 1), "user")
	p.LastVisit = time.Now().Unix()
	p.IsAdmin = false
	p.Achives = make([]Tag, 0)
}
