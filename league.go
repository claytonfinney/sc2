package sc2

import (
	"fmt"
)

// Target struct for the LeagueData resource:
// /data/sc2/league/:seasonId/:queueId/:teamId/:leagueId
type LeagueData struct {
	Links LeagueDataLinks  `json:"links"`
	Key   LeagueDataKey    `json:"key"`
	Tier  []LeagueDataTier `json:"tier"`
}

type LeagueDataLinks struct {
	Self LeagueDataSelf `json:"self"`
}

type LeagueDataSelf struct {
	Href string `json:"href"`
}

type LeagueDataTier struct {
	Id       int                  `json:"id"`
	Division []LeagueDataDivision `json:"division"`
}

type LeagueDataDivision struct {
	Id          int `json:"id"`
	LadderId    int `json:"ladder_id"`
	MemberCount int `json:"member_count"`
}

type LeagueDataKey struct {
	LeagueId int `json:"league_id"`
	SeasonId int `json:"season_id"`
	QueueId  int `json:"queue_id"`
	TeamType int `json:"team_type"`
}

func (c *Client) GetLeagueData(season int, queue int, team int, league int) (*LeagueData, error) {
	uri := fmt.Sprintf("%s/data/sc2/league/%v/%v/%v/%v", c.BaseUrl, season, queue, team, league)
	l := &LeagueData{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}
