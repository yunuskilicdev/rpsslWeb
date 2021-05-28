package main

type Choice struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
}

type Choices []Choice