package model

type LeagueId int32

const (
	AmericanLeague LeagueId = 103
	NationalLeague LeagueId = 104
)

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

// AllStarWriteInsResponse represents the response for All Star write-in candidates
type AllStarWriteInsResponse struct {
	Copyright string                   `json:"copyright"`
	People    []PotentialAllStarPlayer `json:"people"`
}

// PotentialAllStarPlayer represents a player eligible for All Star write-in
type PotentialAllStarPlayer struct {
	ID                 int       `json:"id"`
	FullName           string    `json:"fullName"`
	Link               string    `json:"link"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	PrimaryNumber      string    `json:"primaryNumber,omitempty"`
	BirthDate          string    `json:"birthDate"`
	CurrentAge         int       `json:"currentAge"`
	BirthCity          string    `json:"birthCity"`
	BirthStateProvince string    `json:"birthStateProvince,omitempty"`
	BirthCountry       string    `json:"birthCountry"`
	Height             string    `json:"height"`
	Weight             int       `json:"weight"`
	Active             bool      `json:"active"`
	UseName            string    `json:"useName"`
	UseLastName        string    `json:"useLastName"`
	MiddleName         string    `json:"middleName,omitempty"`
	BoxscoreName       string    `json:"boxscoreName"`
	NickName           string    `json:"nickName,omitempty"`
	Gender             string    `json:"gender"`
	NameMatrilineal    string    `json:"nameMatrilineal,omitempty"`
	IsPlayer           bool      `json:"isPlayer"`
	IsVerified         bool      `json:"isVerified"`
	DraftYear          int       `json:"draftYear,omitempty"`
	Pronunciation      string    `json:"pronunciation,omitempty"`
	MlbDebutDate       string    `json:"mlbDebutDate,omitempty"`
	BatSide            BatSide   `json:"batSide"`
	PitchHand          PitchHand `json:"pitchHand"`
	NameFirstLast      string    `json:"nameFirstLast"`
	NameSlug           string    `json:"nameSlug"`
	FirstLastName      string    `json:"firstLastName"`
	LastFirstName      string    `json:"lastFirstName"`
	LastInitName       string    `json:"lastInitName"`
	InitLastName       string    `json:"initLastName"`
	FullFMLName        string    `json:"fullFMLName"`
	FullLFMName        string    `json:"fullLFMName"`
	StrikeZoneTop      float64   `json:"strikeZoneTop"`
	StrikeZoneBottom   float64   `json:"strikeZoneBottom"`
}

// BatSide represents which side the player bats from
type BatSide struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// PitchHand represents which hand the player pitches with
type PitchHand struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// It's just a list of players with attributes -- No vote data the responses at all at time of writing
type AllStarFinalVoteResponse AllStarWriteInsResponse
type AllStarBallotResponse AllStarWriteInsResponse
