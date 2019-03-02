package sc2

import (
	"fmt"
)

type LegacyProfile struct {
	Id          string `json:"id"`
	Realm       int    `json:"realm"`
	DisplayName string `json:"displayName"`
	ClanName    string `json:"clanName"`
	ClanTag     string `json:"clanTag"`
	ProfilePath string `json:"profilePath"`
	Portrait    struct {
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
		Offset int    `json:"offset"`
		Url    string `json:"url"`
	} `json:"portrait"`

	Career struct {
		PrimaryRace      string `json:"primaryRace"`
		TerranWins       int    `json:"terranWins"`
		ProtossWins      int    `json:"protossWins"`
		ZergWins         int    `json:"zergWins"`
		Highest1v1Rank   string `json:"highest1v1Rank"`
		HighestTeamRank  string `json:"highestTeamRank"`
		SeasonTotalGames int    `json:"seasonTotalGames"`
		CareerTotalGames int    `json:"careerTotalGames"`
	} `json:"career"`

	SwarmLevels struct {
		Level  int `json:"level"`
		Terran struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"terran"`
		Zerg struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"zerg"`
		Protoss struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"protoss"`
	} `json:"swarmLevels"`

	Campaign struct {
	} `json:"campaign"`

	Season struct {
		SeasonId             int `json:"seasonId"`
		SeasonNumber         int `json:"seasonNumber"`
		SeasonYear           int `json:"seasonYear"`
		TotalGamesThisSeason int `json:"totalGamesThisSeason"`
		Stats                []struct {
			Type  string `json:"type"`
			Wins  int    `json:"wins"`
			Games int    `json:"games"`
		} `json:"stats"`
	} `json:"season"`

	Rewards struct {
		Selected []string `json:"selected"`
		Earned   []string `json:"earned"`
	} `json:"rewards"`

	Achievements struct {
		Points struct {
			TotalPoints    int            `json:"totalPoints"`
			CategoryPoints map[string]int `json:"categoryPoints"`
		} `json:"points"`
		Achievements []struct {
			AchievementId  string `json:"achievementId"`
			CompletionDate int    `json:"completionDate"`
		} `json:"achievements"`
	} `json:"achievements"`
}

type LegacyRewards struct {
	Portraits []LegacyReward `json:"portraits"`
	Skins     []LegacyReward `json:"skins"`
}

type LegacyReward struct {
	Title         string `json:"title"`
	Id            string `json:"id"`
	AchievementId string `json:"achievementId"`
	Icon          struct {
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
		Offset int    `json:"offset"`
		Url    string `json:"url"`
	} `json:"icon"`
}

type LegacyAchievements struct {
	Achievements []struct {
		Title         string `json:"title"`
		Description   string `json:"description"`
		AchievementId string `json:"achievementId"`
		CategoryId    string `json:"categoryId"`
		Points        int    `json:"points"`
		Icon          struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			Url    string `json:"url"`
		} `json:"icon"`
	} `json:"achievements"`

	Categories []struct {
		Title                 string `json:"title"`
		CategoryId            string `json:"categoryId"`
		FeaturedAchievementId string `json:"featuredAchievementId"`
		Children              []struct {
			Title                 string `json:"title"`
			CategoryId            string `json:"categoryId"`
			FeaturedAchievementId string `json:"featuredAchievementId"`
		} `json:"children"`
	} `json:"categories"`
}

type LegacyLadders struct {
	CurrentSeason  LegacySeason `json:"currentSeason"`
	PreviousSeason LegacySeason `json:"previousSeason"`
}

type LegacyLadder struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`

	League struct {
		LeagueKey struct {
			LeagueId int `json:"league_id"`
			SeasonId int `json:"season_id"`
			QueueId  int `json:"queue_id"`
			TeamType int `json:"team_type"`
		} `json:"league_key"`
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"league"`

	Team []struct {
		Id                  int `json:"id"`
		Rating              int `json:"rating"`
		Wins                int `json:"wins"`
		Losses              int `json:"losses"`
		Ties                int `json:"ties"`
		Points              int `json:"points"`
		LongestWinStreak    int `json:"longest_win_streak"`
		CurrentWinStreak    int `json:"current_win_streak"`
		CurrentRank         int `json:"current_rank"`
		HighestRank         int `json:"highest_rank"`
		PreviousRank        int `json:"previous_rank"`
		JoinTimeStamp       int `json:"join_time_stamp"`
		LastPlayedTimeStamp int `json:"last_played_time_stamp"`
		Member              []struct {
			LegacyLink struct {
				Id    int    `json:"id"`
				Realm int    `json:"realm"`
				Name  string `json:"name"`
				Path  string `json:"path"`
			} `json:"legacy_link"`
			PlayedRaceCount []struct {
				Race map[string]string `json:"race"`
			} `json:"played_race_count"`
			Count         int `json:"count"`
			CharacterLink struct {
				Id        int    `json:"id"`
				BattleTag string `json:"battle_tag"`
				Key       struct {
					Href string `json:"href"`
				} `json:"character_link"`
			}
		} `json:"member"`
	} `json:"team"`
}

type LegacySeason struct {
	Ladder []struct {
		LadderName       string `json:"ladderName"`
		LadderId         string `json:"ladderId"`
		Division         int    `json:"division"`
		Rank             int    `json:"rank"`
		League           string `json:"league"`
		MatchMakingQueue string `json:"matchMakingQueue"`
		Wins             int    `json:"wins"`
		Losses           int    `json:"losses"`
		Showcase         bool   `json:"showcase"`
	} `json:"ladder"`

	Characters []struct {
		Id          string `json:"id"`
		Realm       int    `json:"realm"`
		Region      int    `json:"region"`
		DisplayName string `json:"displayName"`
		ClanName    string `json:"clanName"`
		ClanTag     string `json:"clanTag"`
		ProfilePath string `json:"profilePath"`
	} `json:"characters"`
}

type LegacyMatches struct {
	Matches []struct {
		Map      string `json:"map"`
		Type     string `json:"type"`
		Decision string `json:"decision"`
		Speed    string `json:"speed"`
		Date     string `json:"date"`
	} `json:"matches"`
}

func (c *Client) GetLegacyProfile(region int, realm int, profile int) {
	uri := fmt.Sprintf("%s/sc2/legacy/profile/%s/%s/%s", c.BaseUrl, region, realm, profile)
	p := LegacyProfile{}
	err := c.Get(uri, &p)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (c *Client) GetLegacyLadders(region int, realm int, profile int) (*LegacyLadders, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/%s/%s/%s/ladders", c.BaseUrl, region, realm, profile)
	l := &LegacyLadders{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetLegacyMatches(region int, realm int, profile int) (*LegacyMatches, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/%s/%s/%s/matches", c.BaseUrl, region, realm, profile)
	m := &LegacyMatches{}
	err := c.Get(uri, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *Client) GetLegacyAchievements(region int) (*LegacyAchievements, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/data/achievements/%s", c.BaseUrl, region)
	a := &LegacyAchievements{}
	err := c.Get(uri, a)
	if err != nil {
		return nil, err
	}

	return a, err
}

func (c *Client) GetLegacyLadder(region int, id int) (*LegacyLadder, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/ladder/%s/%s", c.BaseUrl, region, id)
	l := &LegacyLadder{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetLegacyRewards(region) (*LegacyRewards, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/data/rewards/%s", c.BaseUrl, region)
	r := &LegacyRewards{}
	err := c.Get(uri, r)
	if err != nil {
		return nil, err
	}

	return r, err
}
