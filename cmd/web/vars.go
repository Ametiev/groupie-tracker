package web

type Artist struct {
	Solo          bool
	ID            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  int                 `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Locations     string              `json:"locations"`
	ConcertDates  string              `json:"concertDates"`
	DatesLocation map[string][]string `json:"datesLocations"`
	Relations     string              `json:"relations"`
}

type Relations struct {
	Index []struct {
		ID            int                 `json:"id"`
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ErrorStr struct {
	Status  int
	Message string
}

var (
	Artists  []Artist
	Relation Relations
)
