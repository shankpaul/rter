package data

import (
	"time"
)

type Term struct {
	Term string `json:",omitempty"`

	Automated bool
	Author    string

	UpdateTime time.Time `json:",omitempty"`
}

type TermRanking struct {
	Term    string `json:",omitempty"`
	Ranking string `json:",omitempty"`

	UpdateTime time.Time `json:",omitempty"`
}
