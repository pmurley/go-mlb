package client

// LeaguesResponse represents the full API response for leagues endpoint
type LeaguesResponse struct {
	Copyright string   `json:"copyright"`
	Leagues   []League `json:"leagues"`
}

// League represents a baseball league
type League struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	Link             string         `json:"link"`
	Abbreviation     string         `json:"abbreviation"`
	NameShort        string         `json:"nameShort"`
	SeasonState      string         `json:"seasonState"`
	HasWildCard      bool           `json:"hasWildCard"`
	HasSplitSeason   bool           `json:"hasSplitSeason,omitempty"`
	NumGames         int            `json:"numGames,omitempty"`
	HasPlayoffPoints bool           `json:"hasPlayoffPoints,omitempty"`
	NumTeams         int            `json:"numTeams,omitempty"`
	NumWildcardTeams int            `json:"numWildcardTeams,omitempty"`
	SeasonDateInfo   SeasonDateInfo `json:"seasonDateInfo,omitempty"`
	Season           string         `json:"season,omitempty"`
	OrgCode          string         `json:"orgCode,omitempty"`
	ConferencesInUse bool           `json:"conferencesInUse,omitempty"`
	DivisionsInUse   bool           `json:"divisionsInUse,omitempty"`
	Sport            struct {
		ID   int    `json:"id,omitempty"`
		Link string `json:"link,omitempty"`
	} `json:"sport,omitempty"`
	SortOrder int  `json:"sortOrder,omitempty"`
	Active    bool `json:"active,omitempty"`
}

// SeasonDateInfo contains dates for the league's season
type SeasonDateInfo struct {
	SeasonID                  string  `json:"seasonId,omitempty"`
	PreSeasonStartDate        string  `json:"preSeasonStartDate,omitempty"`
	PreSeasonEndDate          string  `json:"preSeasonEndDate,omitempty"`
	SeasonStartDate           string  `json:"seasonStartDate,omitempty"`
	SpringStartDate           string  `json:"springStartDate,omitempty"`
	SpringEndDate             string  `json:"springEndDate,omitempty"`
	RegularSeasonStartDate    string  `json:"regularSeasonStartDate,omitempty"`
	LastDate1stHalf           string  `json:"lastDate1stHalf,omitempty"`
	AllStarDate               string  `json:"allStarDate,omitempty"`
	FirstDate2ndHalf          string  `json:"firstDate2ndHalf,omitempty"`
	RegularSeasonEndDate      string  `json:"regularSeasonEndDate,omitempty"`
	PostSeasonStartDate       string  `json:"postSeasonStartDate,omitempty"`
	PostSeasonEndDate         string  `json:"postSeasonEndDate,omitempty"`
	SeasonEndDate             string  `json:"seasonEndDate,omitempty"`
	OffseasonStartDate        string  `json:"offseasonStartDate,omitempty"`
	OffSeasonEndDate          string  `json:"offSeasonEndDate,omitempty"`
	SeasonLevelGamedayType    string  `json:"seasonLevelGamedayType,omitempty"`
	GameLevelGamedayType      string  `json:"gameLevelGamedayType,omitempty"`
	QualifierPlateAppearances float64 `json:"qualifierPlateAppearances,omitempty"`
	QualifierOutsPitched      float64 `json:"qualifierOutsPitched,omitempty"`
}
