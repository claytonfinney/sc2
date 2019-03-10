package sc2

import (
	"fmt"
)

type LadderSummary struct {
	ShowCaseEntries []struct {
		LadderId string `json:"ladderId"`
		Team     struct {
			LocalizedGameMode string `json:"localizedGameMode"`
			Members           []struct {
				FavoriteRace string `json:favoriteRace"`
				Name         string `json:"name"`
				PlayerId     string `json:"playerId"`
				Region       int    `json:"region"`
			} `json:"members"`
		} `json:"team"`
		LeagueName            string `json:"leagueName"`
		LocalizedDivisionName string `json:"localizedDivisionName"`
		Rank                  int    `json:"rank"`
		Wins                  int    `json:"wins"`
		Losses                int    `json:"losses"`
	} `json:"showCaseEntries"`

	PlacementMatches struct {
		LocalizedGameMode string `json:"localizedGameMode"`
		Members           []struct {
			FavoriteRace string `json:favoriteRace"`
			Name         string `json:"name"`
			PlayerId     string `json:"playerId"`
			Region       int    `json:"region"`
		} `json:"members"`
		GamesRemaining int `json:"gamesRemaining"`
	} `json:"placementMatches"`

	AllLadderMemberships struct {
		LadderId          string `json:"ladderId"`
		LocalizedGameMode string `json:"localizedGameMode"`
		Rank              int    `json:"rank"`
	} `json:"allLadderMemberships"`
}

type Ladder struct {
	LadderTeams []struct {
		TeamMembers []struct {
			Id           string `json:"id"`
			Realm        int    `json:"realm"`
			Region       int    `json:"region"`
			DisplayName  string `json:"displayName"`
			FavoriteRace string `json:"favoriteRace"`
		}
		PreviousRank  int `json:"previousRank"`
		Points        int `json:"points"`
		Wins          int `json:"wins"`
		Losses        int `json:"losses"`
		Mmr           int `json:"mmr"`
		JoinTimestamp int `json:"joinTimestamp"`
	} `json:"ladderTeams"`
}

type Season struct {
	SeasonId  int    `json:"seasonId"`
	Number    int    `json:"number"`
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
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
