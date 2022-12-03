package handlers

type Genre struct {
	GenreId int32  `json:"GenreId,omitempty"`
	Name    string `json:"Name,omitempty"`
}

type GenreResponse struct {
	Genres []Genre `json:"genres,omitempty"`
}
