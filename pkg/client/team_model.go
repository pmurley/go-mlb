package client

// TeamsResponse is the response to the /teams endpoint
type TeamsResponse struct {
	Copyright string `json:"copyright"`
	Teams     []Team `json:"teams"`
}

type Team struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Link            string `json:"link"`
	Season          int    `json:"season"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	ShortName       string `json:"shortName"`
	ClubName        string `json:"clubName"`
	FranchiseName   string `json:"franchiseName"`
	Active          bool   `json:"active"`
	AllStarStatus   string `json:"allStarStatus"`
	FileCode        string `json:"fileCode"`
	TeamCode        string `json:"teamCode"`
	Division        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"division"`
	League struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"league"`
	Sport struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"sport"`
	SpringLeague struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
	} `json:"springLeague"`
	SpringVenue struct {
		ID   int    `json:"id"`
		Link string `json:"link"`
	} `json:"springVenue"`
	Venue struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`
}
