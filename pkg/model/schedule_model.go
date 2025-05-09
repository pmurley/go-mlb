package model

import "time"

// ScheduleResponse represents the top-level response for MLB schedule data
type ScheduleResponse struct {
	Copyright            string         `json:"copyright"`
	TotalItems           int            `json:"totalItems"`
	TotalEvents          int            `json:"totalEvents"`
	TotalGames           int            `json:"totalGames"`
	TotalGamesInProgress int            `json:"totalGamesInProgress"`
	Dates                []ScheduleDate `json:"dates"`
}

// ScheduleDate represents a single date in the schedule
type ScheduleDate struct {
	Date                 string  `json:"date"`
	TotalItems           int     `json:"totalItems"`
	TotalEvents          int     `json:"totalEvents"`
	TotalGames           int     `json:"totalGames"`
	TotalGamesInProgress int     `json:"totalGamesInProgress"`
	Games                []Game  `json:"games"`
	Events               []Event `json:"events"`
}

// Game represents a single game in the schedule
type Game struct {
	GamePk                 int        `json:"gamePk"`
	GameGuid               string     `json:"gameGuid"`
	Link                   string     `json:"link"`
	GameType               string     `json:"gameType"`
	Season                 string     `json:"season"`
	GameDate               time.Time  `json:"gameDate"`
	OfficialDate           string     `json:"officialDate"`
	Status                 GameStatus `json:"status"`
	Teams                  GameTeams  `json:"teams"`
	Venue                  Venue      `json:"venue"`
	Content                Content    `json:"content"`
	IsTie                  bool       `json:"isTie"`
	GameNumber             int        `json:"gameNumber"`
	PublicFacing           bool       `json:"publicFacing"`
	DoubleHeader           string     `json:"doubleHeader"`
	GamedayType            string     `json:"gamedayType"`
	Tiebreaker             string     `json:"tiebreaker"`
	CalendarEventID        string     `json:"calendarEventID"`
	SeasonDisplay          string     `json:"seasonDisplay"`
	DayNight               string     `json:"dayNight"`
	ScheduledInnings       int        `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool       `json:"reverseHomeAwayStatus"`
	InningBreakLength      int        `json:"inningBreakLength"`
	GamesInSeries          int        `json:"gamesInSeries"`
	SeriesGameNumber       int        `json:"seriesGameNumber"`
	SeriesDescription      string     `json:"seriesDescription"`
	RecordSource           string     `json:"recordSource"`
	IfNecessary            string     `json:"ifNecessary"`
	IfNecessaryDescription string     `json:"ifNecessaryDescription"`
}

// GameStatus represents the status of a game
type GameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
	AbstractGameCode  string `json:"abstractGameCode"`
}

// GameTeams represents the teams playing in a game
type GameTeams struct {
	Away TeamInfo `json:"away"`
	Home TeamInfo `json:"home"`
}

// TeamInfo represents information about a team in a game
type TeamInfo struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int          `json:"score"`
	Team         ScheduleTeam `json:"team"`
	IsWinner     bool         `json:"isWinner"`
	SplitSquad   bool         `json:"splitSquad"`
	SeriesNumber int          `json:"seriesNumber"`
}

// LeagueRecord represents a team's record
type LeagueRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Pct    string `json:"pct"`
}

// ScheduleTeam represents a baseball team
type ScheduleTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Venue represents a game venue
type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Content represents content-related information
type Content struct {
	Link string `json:"link"`
}

// Event is a placeholder for the events array
type Event struct {
	// Fields would be added here if events data were provided
}

// PostseasonScheduleResponse represents the response for postseason schedule data
// PostseasonScheduleResponse represents the top-level response
type PostseasonScheduleResponse struct {
	Copyright            string           `json:"copyright"`
	TotalItems           int              `json:"totalItems"`
	TotalEvents          int              `json:"totalEvents"`
	TotalGames           int              `json:"totalGames"`
	TotalGamesInProgress int              `json:"totalGamesInProgress"`
	Wait                 int              `json:"wait"`
	Dates                []PostseasonDate `json:"dates"`
}

// PostseasonDate represents a date with games
type PostseasonDate struct {
	Date                 string           `json:"date"`
	TotalItems           int              `json:"totalItems"`
	TotalEvents          int              `json:"totalEvents"`
	TotalGames           int              `json:"totalGames"`
	TotalGamesInProgress int              `json:"totalGamesInProgress"`
	Games                []PostseasonGame `json:"games"`
	Events               []interface{}    `json:"events"`
}

// PostseasonGame represents a single postseason game
type PostseasonGame struct {
	GamePk                 int             `json:"gamePk"`
	GameGuid               string          `json:"gameGuid"`
	Link                   string          `json:"link"`
	GameType               string          `json:"gameType"`
	Season                 string          `json:"season"`
	GameDate               time.Time       `json:"gameDate"`
	OfficialDate           string          `json:"officialDate"`
	Status                 GameStatus      `json:"status"`
	Teams                  PostseasonTeams `json:"teams"`
	Venue                  Venue           `json:"venue"`
	Content                Content         `json:"content"`
	IsTie                  bool            `json:"isTie"`
	IsFeaturedGame         bool            `json:"isFeaturedGame"`
	GameNumber             int             `json:"gameNumber"`
	PublicFacing           bool            `json:"publicFacing"`
	DoubleHeader           string          `json:"doubleHeader"`
	GamedayType            string          `json:"gamedayType"`
	Tiebreaker             string          `json:"tiebreaker"`
	CalendarEventID        string          `json:"calendarEventID"`
	SeasonDisplay          string          `json:"seasonDisplay"`
	DayNight               string          `json:"dayNight"`
	Description            string          `json:"description"`
	ScheduledInnings       int             `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool            `json:"reverseHomeAwayStatus"`
	InningBreakLength      int             `json:"inningBreakLength"`
	GamesInSeries          int             `json:"gamesInSeries"`
	SeriesGameNumber       int             `json:"seriesGameNumber"`
	SeriesDescription      string          `json:"seriesDescription"`
	RecordSource           string          `json:"recordSource"`
	IfNecessary            string          `json:"ifNecessary"`
	IfNecessaryDescription string          `json:"ifNecessaryDescription"`
}

// PostseasonTeams contains home and away team information
type PostseasonTeams struct {
	Away TeamInfo `json:"away"`
	Home TeamInfo `json:"home"`
}

// Was unable to get a response for these that contained anything for the present or past couple years,
// but the structure of the response appeared the same as the postseason schedule response
type TuneInScheduleResponse PostseasonScheduleResponse

// PostseasonSeriesScheduleResponse represents the top-level response for postseason series
type PostseasonSeriesScheduleResponse struct {
	Copyright            string         `json:"copyright"`
	TotalItems           int            `json:"totalItems"`
	TotalEvents          int            `json:"totalEvents"`
	TotalGames           int            `json:"totalGames"`
	TotalGamesInProgress int            `json:"totalGamesInProgress"`
	Wait                 int            `json:"wait"`
	Series               []SeriesDetail `json:"series"`
}

// SeriesDetail represents each series in the postseason
type SeriesDetail struct {
	Series               SeriesInfo       `json:"series"`
	TotalItems           int              `json:"totalItems"`
	TotalGames           int              `json:"totalGames"`
	TotalGamesInProgress int              `json:"totalGamesInProgress"`
	Games                []PostseasonGame `json:"games"`
	SortOrder            int              `json:"sortOrder"`
}

// SeriesInfo contains the identifier information for a series
type SeriesInfo struct {
	ID         string `json:"id"`
	SortNumber int    `json:"sortNumber"`
	IsDefault  bool   `json:"isDefault"`
	GameType   string `json:"gameType"`
}

type TieGamesResponse ScheduleResponse
