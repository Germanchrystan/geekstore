package domain

type Subgenre struct {
	Id    int    `json:"id"`
	Name  string `json:"subgenre_name"`
	Genre string `json:"genre_name"`
}
