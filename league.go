package sc2

import (
	"fmt"
)

type LeagueData struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`

	Key struct {
		LeagueId int `json:"league_id"`
		SeasonId int `json:"season_id"`
		QueueId  int `json:"queue_id"`
		TeamType int `json:"team_type"`
	} `json:"key"`

	Tier []struct {
		Id       int `json:"id"`
		Division []struct {
			Id          int `json:"id"`
			LadderId    int `json:"ladder_id"`
			MemberCount int `json:"member_count"`
		} `json:"division"`
	} `json:"tier"`
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
