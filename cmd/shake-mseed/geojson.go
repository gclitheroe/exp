package main

type location struct {
	Network  string `json:"network"`
	Station  string `json:"station"`
	Location string `json:"location"`
	PGAV float64 `json:"pga_v"`
	PGAH float64 `json:"pga_h"`
	PGVV float64 `json:"pgv_v"`
	PGVH float64 `json:"pgv_h"`
}

type point struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

type locationFeature struct {
	Type       string   `json:"type"`
	Properties location `json:"properties"`
	Geometry   point    `json:"geometry"`
}

type locationFeatures struct {
	Type     string    `json:"type"`
	Features []locationFeature `json:"features"`
}