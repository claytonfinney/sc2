package sc2

import (
	"fmt"
)

// Target struct for Ladder Summary resource:
// /sc2/profile/:regionId/:realmId/:profileId/ladder/summary
type LadderSummary struct {
	ShowCaseEntries      []LadderSummaryShowCaseEntry    `json:"showCaseEntries"`
	PlacementMatches     []LadderSummaryPlacementMatches `json:"placementMatches"`
	AllLadderMemberships []LadderAllMembership           `json:"allLadderMemberships"`
}

type LadderSummaryShowCaseEntry struct {
	LadderId              string            `json:"ladderId"`
	Team                  LadderSummaryTeam `json:"team"`
	LeagueName            string            `json:"leagueName"`
	LocalizedDivisionName string            `json:"localizedDivisionName"`
	Rank                  int               `json:"rank"`
	Wins                  int               `json:"wins"`
	Losses                int               `json:"losses"`
}

type LadderSummaryTeam struct {
	LocalizedGameMode string                    `json:"localizedGameMode"`
	Members           []LadderSummaryTeamMember `json:"members"`
}

type LadderSummaryTeamMember struct {
	FavoriteRace string `json:"favoriteRace"`
	Name         string `json:"name"`
	PlayerId     string `json:"playerId"`
	Region       int    `json:"region"`
}

type LadderSummaryPlacementMatches struct {
	LocalizedGameMode string                    `json:"localizedGameMode"`
	Members           []LadderSummaryTeamMember `json:"members"`
	GamesRemaining    int                       `json:"gamesRemaining"`
}

// Target struct for Ladder resource:
// /sc2/profile/:regionId/:realmId/:profileId/ladder/:ladderId"
type Ladder struct {
	LadderTeams             []LadderTeam           `json:"ladderTeams"`
	AllLadderMemberships    []LadderAllMembership  `json:"allLadderMemberships"`
	LocalizedDivision       string                 `json:"localizedDivision"`
	League                  string                 `json:"league"`
	CurrentLadderMembership LadderLadderMembership `json:"currentLadderMembership"`
	RanksAndPools           []LadderRankAndPool    `json:"ranksAndPools"`
}

type LadderTeam struct {
	TeamMembers   []LadderTeamMember `json:"teamMembers"`
	PreviousRank  int                `json:"previousRank"`
	Points        int                `json:"points"`
	Wins          int                `json:"wins"`
	Losses        int                `json:"losses"`
	Mmr           int                `json:"mmr"`
	JoinTimestamp int                `json:"joinTimestamp"`
}

type LadderTeamMember struct {
	Id           string `json:"id"`
	Realm        int    `json:"realm"`
	Region       int    `json:"region"`
	DisplayName  string `json:"displayName"`
	ClanTag      string `json:"clanTag"`
	FavoriteRace string `json:"favoriteRace"`
}

type LadderLadderMembership struct {
	LadderId          string `json:"ladderId"`
	LocalizedGameMode string `json:"localizedGameMode"`
}

type LadderRankAndPool struct {
	Rank      int `json:"rank"`
	Mmr       int `json:"mmr"`
	BonusPool int `json:"bonusPool"`
}

// Target struct for Season resource:
// /sc2/ladder/season/:regionId
type Season struct {
	SeasonId  int    `json:"seasonId"`
	Number    int    `json:"number"`
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// Shared structs
type LadderAllMembership struct {
	LadderId          string `json:"ladderId"`
	LocalizedGameMode string `json:"localizedGameMode"`
	Rank              int    `json:"rank"`
}

func (c *Client) GetLadderSummary(region int, realm int, profile int) (*LadderSummary, error) {
	uri := fmt.Sprintf("%s/sc2/profile/%v/%v/%v/ladder/summary", c.BaseUrl, region, realm, profile)
	l := &LadderSummary{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (c *Client) GetLadder(region int, realm int, profile int, ladder int) (*Ladder, error) {
	uri := fmt.Sprintf("%s/sc2/profile/%v/%v/%v/ladder/%v", c.BaseUrl, region, realm, profile, ladder)
	l := &Ladder{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetGrandmaster(region int) (*Ladder, error) {
	uri := fmt.Sprintf("%s/sc2/ladder/grandmaster/%v", c.BaseUrl, region)
	l := &Ladder{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetSeason(region int) (*Season, error) {
	uri := fmt.Sprintf("%s/sc2/ladder/season/%v", c.BaseUrl, region)
	s := &Season{}
	err := c.Get(uri, s)
	if err != nil {
		return nil, err
	}

	return s, err
}
