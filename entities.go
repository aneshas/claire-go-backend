package main

type (
	// Car make entity
	Make struct {
		Id          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
	}

	// Car model entity
	Model struct {
		Id          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
		MakeId      string `json:"make_id"`
	}

	// Car tags
	Tag struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}
)
