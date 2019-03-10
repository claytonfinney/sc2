package sc2

import (
	"fmt"
)

type Static struct {
	Achievements []Achievement `json:"achievements"`
	Criteria     []Criteria    `json:"criteria"`
	Categories   []Category    `json:"categories"`
	Rewards      []Reward      `json:"rewards"`
}

type Achievement struct {
	CategoryId          string   `json:"categoryId"`
	ChainAchievementIds []string `json:"chainAchievementIds"`
	ChainRewardSize     int      `json:"chainRewardSize"`
	CriteriaIds         []string `json:"criteriaIds"`
	Description         string   `json:"description"`
	Flags               int      `json:"flags"`
	Id                  string   `json:"id"`
	ImageUrl            string   `json:"imageUrl"`
	IsChained           bool     `json:"isChained"`
	Points              int      `json:"points"`
	Title               string   `json:"title"`
	UiOrderHint         int      `json:"uiOrderHint"`
}

type Criteria struct {
	AchievementId     string `json:"achievementId"`
	Description       string `json:"description"`
	EvaluationClass   string `json:"evaluationClass"`
	Flags             int    `json:"flags"`
	Id                string `json:"id"`
	NecessaryQuantity int    `json:"necessaryQuantity"`
	UiOrderHint       int    `json:"uiOrderHint"`
}

type Category struct {
	ChildrenCategoryIds   []string `json:"childrenCategoryIds"`
	FeaturedAchievementId string   `json:"featuredAchievementId"`
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	ParentCategoryId      string   `json:"parentCategoryId"`
	Points                int      `json:"points"`
	UiOrderHint           int      `json:"uiOrderHint"`
}

type Reward struct {
	Flags          int    `json:"flags"`
	Id             string `json:"id"`
	AchievementId  string `json:"achievementId"`
	Name           string `json:"name"`
	UnlockableType string `json:"unlockableType"`
	IsSkin         bool   `json:"isSkin"`
	ImageUrl       string `json:"imageUrl"`
	UiOrderHint    int    `json:"uiOrderHint"`
}

type Metadata struct {
	Name       string `json:"name"`
	ProfileUrl string `json:"profileUrl"`
	AvatarUrl  string `json:"avatarUrl"`
	ProfileId  string `json:"profileId"`
	RegionId   int    `json:"regionId"`
	RealmId    int    `json:"realmId"`
}

type Profile struct {
	Summary struct {
		Id                     string `json:"id"`
		Realm                  int    `json:"realm"`
		DisplayName            string `json:"displayName"`
		Portrait               string `json:"portrait"`
		DecalTerran            string `json:"decalTerran"`
		DecalZerg              string `json:"decalZerg"`
		DecalProtoss           string `json"decalProtoss"`
		TotalSwarmLevel        int    `json:"totalSwarmLevel"`
		TotalAchievementPoints int    `json:"totalAchievementPoints"`
	} `json:"summary"`

	Snapshot struct {
		SeasonSnapshot struct {
			Solo struct {
				Rank       int    `json:"rank"`
				LeagueName string `json:"leagueName"`
				TotalGames int    `json:"totalGames"`
				TotalWins  int    `json:"totalWins"`
			} `json:"1v1"`

			Twos struct {
				Rank       int    `json:"rank"`
				LeagueName string `json:"leagueName"`
				TotalGames int    `json:"totalGames"`
				TotalWins  int    `json:"totalWins"`
			} `json:"2v2"`

			Threes struct {
				Rank       int    `json:"rank"`
				LeagueName string `json:"leagueName"`
				TotalGames int    `json:"totalGames"`
				TotalWins  int    `json:"totalWins"`
			} `json:"3v3"`

			Fours struct {
				Rank       int    `json:"rank"`
				LeagueName string `json:"leagueName"`
				TotalGames int    `json:"totalGames"`
				TotalWins  int    `json:"totalWins"`
			} `json:"4v4"`
		}
	} `json:"snapshot"`

	Career struct {
		TerranWins                int    `json:"terranWins"`
		ZergWins                  int    `json:"zergWins"`
		ProtossWins               int    `json:"protossWins"`
		TotalCareerGames          int    `json:"totalCareerGames"`
		TotalGamesThisSeason      int    `json:"totalGamesThisSeason"`
		Current1v1LeagueName      string `json:"current1v1LeagueName"`
		CurrentBestTeamLeagueName string `json:"currentBestTeamLeagueName"`
		Best1v1Finish             struct {
			LeagueName    string `json:"leagueName"`
			TimesAchieved int    `json:"timesAchieved"`
		} `json:"best1v1Finish"`
		BestTeamFInish struct {
			LeagueName    string `json:"leagueName"`
			TimesAchieved int    `json:"timesAchieved"`
		} `json:"bestTeamFinish"`
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

	CategoryPointProgress []struct {
		RewardId      string `json:"rewardId"`
		AchievementId string `json:"achievementId"`
		Selected      bool   `json:"selected"`
	} `json:"categoryPointProgress"`

	EarnedAchievements []struct {
		AchievementId                    string `json:"achievementId"`
		CompletionDate                   int    `json:"completionDate"`
		NumCompletedAchievementsInSeries int    `json:"numCompletedAchievementsInSeries"`
		IsComplete                       bool   `json:"isComplete"`
		InProgress                       bool   `json:"inProgress"`
		Criteria                         []struct {
			CriterionId string `json:"criterionId"`
		} `json:"criteria"`
	}
}

func (c *Client) GetStatic(region int) (*Static, error) {
	uri := fmt.Sprintf("%s/sc2/static/profile/%v", c.BaseUrl, region)
	s := &Static{}
	err := c.Get(uri, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *Client) GetMetadata(region int, realm int, profile int) (*Metadata, error) {
	uri := fmt.Sprintf("%s/sc2/metadata/profile/%v/%v/%v", c.BaseUrl, region, realm, profile)
	m := &Metadata{}
	err := c.Get(uri, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) GetProfile(region int, realm int, profile int) (*Profile, error) {
	uri := fmt.Sprintf("%s/sc2/profile/%v/%v/%v", c.BaseUrl, region, realm, profile)
	p := &Profile{}
	err := c.Get(uri, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
