package model

// AttendanceResponse represents the top-level response from the attendance API
type AttendanceResponse struct {
	Copyright       string              `json:"copyright"`
	Records         []AttendanceRecord  `json:"records"`
	AggregateTotals AttendanceAggregate `json:"aggregateTotals"`
}

// AttendanceRecord represents the attendance record for a single team
type AttendanceRecord struct {
	OpeningsTotal            int            `json:"openingsTotal"`
	OpeningsTotalAway        int            `json:"openingsTotalAway"`
	OpeningsTotalHome        int            `json:"openingsTotalHome"`
	OpeningsTotalLost        int            `json:"openingsTotalLost"`
	GamesTotal               int            `json:"gamesTotal"`
	GamesAwayTotal           int            `json:"gamesAwayTotal"`
	GamesHomeTotal           int            `json:"gamesHomeTotal"`
	Year                     string         `json:"year"`
	AttendanceAverageAway    int            `json:"attendanceAverageAway"`
	AttendanceAverageHome    int            `json:"attendanceAverageHome"`
	AttendanceAverageYtd     int            `json:"attendanceAverageYtd"`
	AttendanceHigh           int            `json:"attendanceHigh"`
	AttendanceHighDate       string         `json:"attendanceHighDate"`
	AttendanceHighGame       GameInfo       `json:"attendanceHighGame"`
	AttendanceLow            int            `json:"attendanceLow"`
	AttendanceLowDate        string         `json:"attendanceLowDate"`
	AttendanceLowGame        GameInfo       `json:"attendanceLowGame"`
	AttendanceOpeningAverage int            `json:"attendanceOpeningAverage"`
	AttendanceTotal          int            `json:"attendanceTotal"`
	AttendanceTotalAway      int            `json:"attendanceTotalAway"`
	AttendanceTotalHome      int            `json:"attendanceTotalHome"`
	GameType                 GameType       `json:"gameType"`
	Team                     AttendanceTeam `json:"team"`
}

// GameInfo represents information about a specific game
type GameInfo struct {
	GamePk   int     `json:"gamePk"`
	Link     string  `json:"link"`
	Content  Content `json:"content"`
	DayNight string  `json:"dayNight"`
}

// GameType represents the type of game
type GameType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// AttendanceTeam represents a baseball team
type AttendanceTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// AttendanceAggregate represents aggregate attendance totals across teams
type AttendanceAggregate struct {
	OpeningsTotalAway     int    `json:"openingsTotalAway"`
	OpeningsTotalHome     int    `json:"openingsTotalHome"`
	OpeningsTotalLost     int    `json:"openingsTotalLost"`
	OpeningsTotalYtd      int    `json:"openingsTotalYtd"`
	AttendanceAverageAway int    `json:"attendanceAverageAway"`
	AttendanceAverageHome int    `json:"attendanceAverageHome"`
	AttendanceAverageYtd  int    `json:"attendanceAverageYtd"`
	AttendanceHigh        int    `json:"attendanceHigh"`
	AttendanceHighDate    string `json:"attendanceHighDate"`
	AttendanceTotal       int    `json:"attendanceTotal"`
	AttendanceTotalAway   int    `json:"attendanceTotalAway"`
	AttendanceTotalHome   int    `json:"attendanceTotalHome"`
}
