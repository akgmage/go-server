package main

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lasstname"`
}

