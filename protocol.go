package main

import "time"

type Msg struct {
	Msg  string    `json:"msg"`
	Time time.Time `json:"time"`
}

type Error struct {
	Error string    `json:"error"`
	Time  time.Time `json:"time"`
}
