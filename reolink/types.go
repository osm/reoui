package reolink

import (
	"time"
)

type Time struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"min"`
	Month  int `json:"mon"`
	Second int `json:"sec"`
	Year   int `json:"year"`
}

type SearchResultFile struct {
	RawEndTime   Time   `json:"EndTime"`
	RawStartTime Time   `json:"StartTime"`
	FrameRate    int    `json:"frameRate"`
	Height       int    `json:"height"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	Type         string `json:"type"`
	Width        int    `json:"width"`
}

type SearchResultStatus struct {
	Month int    `json:"mon"`
	Table string `json:"table"`
	Year  int    `json:"year"`
}

type SearchResult struct {
	File    []SearchResultFile   `json:"File,omitempty"`
	Status  []SearchResultStatus `json:"Status"`
	Channel int                  `json:"channel"`
}

type User struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type Search struct {
	Channel    int    `json:"channel"`
	OnlyStatus int    `json:"onlyStatus"`
	StreamType string `json:"streamType"`
	StartTime  Time   `json:"StartTime"`
	EndTime    Time   `json:"EndTime"`
}

type Param struct {
	User   *User   `json:"User,omitempty"`
	Search *Search `json:"Search,omitempty"`
}

type Request struct {
	Command string `json:"cmd"`
	Action  int    `json:"action"`
	Param   Param  `json:"param"`
}

type Token struct {
	LeaseTime time.Duration `json:"leaseTime"`
	Name      string        `json:"name"`
}

type Value struct {
	Token        *Token        `json:"Token,omitempty"`
	SearchResult *SearchResult `json:"SearchResult,omitempty"`
}

type Error struct {
	Detail  string `json:"detail"`
	RSPCode int    `json:"rspCode"`
}

type Response struct {
	Command string `json:"cmd"`
	Code    int    `json:"code"`
	Value   *Value `json:"value,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
