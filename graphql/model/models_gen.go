// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Camera struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Video struct {
	ID         string    `json:"id"`
	CameraName string    `json:"cameraName"`
	Date       time.Time `json:"date"`
	Duration   int64     `json:"duration"`
}
