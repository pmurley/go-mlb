package model

import "fmt"

// SportsResponse represents the top-level JSON response for sports data
type SportsResponse struct {
	Copyright string  `json:"copyright"`
	Sports    []Sport `json:"sports"`
}

// Sport represents an individual sport entry in the sports array
type Sport struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Link         string `json:"link"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	SortOrder    int    `json:"sortOrder"`
	ActiveStatus bool   `json:"activeStatus"`
}

// Sport ID constants as int32
const (
	SportMLB                int32 = 1
	SportTripleA            int32 = 11
	SportDoubleA            int32 = 12
	SportHighA              int32 = 13
	SportSingleA            int32 = 14
	SportRookie             int32 = 16
	SportWinterLeagues      int32 = 17
	SportMinorLeague        int32 = 21
	SportIndependentLeagues int32 = 23
	SportNegroLeague        int32 = 61
	SportKBO                int32 = 32
	SportNPB                int32 = 31
	SportInternational      int32 = 51
	SportInternational18U   int32 = 509
	SportInternational16U   int32 = 510
	SportInternationalAmat  int32 = 6005
	SportCollege            int32 = 22
	SportHighSchool         int32 = 586
)

// GetSportName returns the name for a sport ID
func GetSportName(id int32) string {
	switch id {
	case SportMLB:
		return "Major League Baseball"
	case SportTripleA:
		return "Triple-A"
	case SportDoubleA:
		return "Double-A"
	case SportHighA:
		return "High-A"
	case SportSingleA:
		return "Single-A"
	case SportRookie:
		return "Rookie"
	case SportWinterLeagues:
		return "Winter Leagues"
	case SportMinorLeague:
		return "Minor League Baseball"
	case SportIndependentLeagues:
		return "Independent Leagues"
	case SportNegroLeague:
		return "Negro League Baseball"
	case SportKBO:
		return "Korean Baseball Organization"
	case SportNPB:
		return "Nippon Professional Baseball"
	case SportInternational:
		return "International Baseball"
	case SportInternational18U:
		return "International Baseball (18U)"
	case SportInternational16U:
		return "International Baseball (16 and under)"
	case SportInternationalAmat:
		return "International Baseball (amateur)"
	case SportCollege:
		return "College Baseball"
	case SportHighSchool:
		return "High School Baseball"
	default:
		return fmt.Sprintf("Unknown Sport (%d)", id)
	}
}

// IsSportIDValid returns whether the ID is a recognized sport ID
func IsSportIDValid(id int32) bool {
	switch id {
	case SportMLB, SportTripleA, SportDoubleA, SportHighA, SportSingleA,
		SportRookie, SportWinterLeagues, SportMinorLeague, SportIndependentLeagues,
		SportNegroLeague, SportKBO, SportNPB, SportInternational,
		SportInternational18U, SportInternational16U, SportInternationalAmat,
		SportCollege, SportHighSchool:
		return true
	default:
		return false
	}
}

type SportsPlayerResponse struct {
	Copyright string         `json:"copyright"`
	People    []SportsPlayer `json:"people"`
}

type SportsPlayer struct {
	ID                 int              `json:"id"`
	FullName           string           `json:"fullName"`
	Link               string           `json:"link"`
	FirstName          string           `json:"firstName"`
	LastName           string           `json:"lastName"`
	PrimaryNumber      string           `json:"primaryNumber"`
	BirthDate          string           `json:"birthDate"`
	CurrentAge         int              `json:"currentAge"`
	BirthCity          string           `json:"birthCity"`
	BirthStateProvince string           `json:"birthStateProvince,omitempty"`
	BirthCountry       string           `json:"birthCountry"`
	Height             string           `json:"height"`
	Weight             int              `json:"weight"`
	Active             bool             `json:"active"`
	CurrentTeam        SportsPlayerTeam `json:"currentTeam"`
	PrimaryPosition    Position         `json:"primaryPosition"`
	UseName            string           `json:"useName"`
	UseLastName        string           `json:"useLastName"`
	MiddleName         string           `json:"middleName"`
	BoxscoreName       string           `json:"boxscoreName"`
	Gender             string           `json:"gender"`
	NameMatrilineal    string           `json:"nameMatrilineal,omitempty"`
	IsPlayer           bool             `json:"isPlayer"`
	IsVerified         bool             `json:"isVerified"`
	DraftYear          int              `json:"draftYear,omitempty"`
	Pronunciation      string           `json:"pronunciation,omitempty"`
	BatSide            Side             `json:"batSide"`
	PitchHand          Side             `json:"pitchHand"`
	NameFirstLast      string           `json:"nameFirstLast"`
	NameSlug           string           `json:"nameSlug"`
	FirstLastName      string           `json:"firstLastName"`
	LastFirstName      string           `json:"lastFirstName"`
	LastInitName       string           `json:"lastInitName"`
	InitLastName       string           `json:"initLastName"`
	FullFMLName        string           `json:"fullFMLName"`
	FullLFMName        string           `json:"fullLFMName"`
	StrikeZoneTop      float64          `json:"strikeZoneTop"`
	StrikeZoneBottom   float64          `json:"strikeZoneBottom"`
}

type SportsPlayerTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type Side struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
