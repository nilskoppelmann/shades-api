package main

type Shade struct {
	Id   int    `json:"id"`
	Grey bool   `json:"grey"`
	Hex  string `json:"hex"`
}

type Shades []Shade
