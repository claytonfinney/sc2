package sc2

import (
	"fmt"
)

// Target struct for the Static resource:
// /sc2/static/profile/:regionId
type Static struct {
	Achievements []StaticAchievement `json:"achievements"`
	Criteria     []StaticCriteria    `json:"criteria"`
	Categories   []StaticCategory    `json:"categories"`
	Rewards      []StaticReward      `json:"rewards"`
}

type StaticAchievement struct {
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

type StaticCriteria struct {
	AchievementId     string `json:"achievementId"`
	Description       string `json:"description"`
	EvaluationClass   string `json:"evaluationClass"`
	Flags             int    `json:"flags"`
	Id                string `json:"id"`
	NecessaryQuantity int    `json:"necessaryQuantity"`
	UiOrderHint       int    `json:"uiOrderHint"`
}

type StaticCategory struct {
	ChildrenCategoryIds   []string `json:"childrenCategoryIds"`
	FeaturedAchievementId string   `json:"featuredAchievementId"`
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	ParentCategoryId      string   `json:"parentCategoryId"`
	Points                int      `json:"points"`
	UiOrderHint           int      `json:"uiOrderHint"`
}

type StaticReward struct {
	Flags          int    `json:"flags"`
	Id             string `json:"id"`
	AchievementId  string `json:"achievementId"`
	Name           string `json:"name"`
	UnlockableType string `json:"unlockableType"`
	IsSkin         bool   `json:"isSkin"`
	ImageUrl       string `json:"imageUrl"`
	UiOrderHint    int    `json:"uiOrderHint"`
}

// Target struct for the Metadata resource:
// /sc2/metadata/profile/:regionId/:realmId/:profileId
type Metadata struct {
	Name       string `json:"name"`
	ProfileUrl string `json:"profileUrl"`
	AvatarUrl  string `json:"avatarUrl"`
	ProfileId  string `json:"profileId"`
	RegionId   int    `json:"regionId"`
	RealmId    int    `json:"realmId"`
}

// Target struct for the Profile resource:
// /sc2/profile/:regionId/:realmId/:profileId
type Profile struct {
	Summary               ProfileSummary                 `json:"summary"`
	Snapshot              ProfileSnapshot                `json:"snapshot"`
	Career                ProfileCareer                  `json:"career"`
	SwarmLevels           ProfileSwarmLevels             `json:"swarmLevels"`
	Campaign              ProfileCampaign                `json:"campaign"`
	CategoryPointProgress []ProfileCategoryPointProgress `json:"categoryPointProgress"`
	AchievementShowcase   []string                       `json:"achievementShowcase"`
	EarnedRewards         []ProfileEarnedReward          `json:"earnedRewards"`
	EarnedAchievements    []ProfileEarnedAchievement     `json:"earnedAchievements"`
}

type ProfileSummary struct {
	Id                     string `json:"id"`
	Realm                  int    `json:"realm"`
	DisplayName            string `json:"displayName"`
	ClanName               string `json:"clanName"`
	ClanTag                string `json:"clanTag"`
	Portrait               string `json:"portrait"`
	DecalTerran            string `json:"decalTerran"`
	DecalZerg              string `json:"decalZerg"`
	DecalProtoss           string `json:"decalProtoss"`
	TotalSwarmLevel        int    `json:"totalSwarmLevel"`
	TotalAchievementPoints int    `json:"totalAchievementPoints"`
}

type ProfileSnapshot struct {
	SeasonSnapshot               ProfileSeasonSnapshot `json:"seasonSnapshot"`
	TotalRankedSeasonGamesPlayed int                   `json:"totalRankedSeasonGamesPlayed"`
}

type ProfileSeasonSnapshot struct {
	Solo   ProfileSeasonSnapshotLadder `json:"1v1"`
	Twos   ProfileSeasonSnapshotLadder `json:"2v2"`
	Threes ProfileSeasonSnapshotLadder `json:"3v3"`
	Fours  ProfileSeasonSnapshotLadder `json:"4v4"`
	Archon ProfileSeasonSnapshotLadder `json:"Archon"`
}

type ProfileSeasonSnapshotLadder struct {
	Rank       int    `json:"rank"`
	LeagueName string `json:"leagueName"`
	TotalGames int    `json:"totalGames"`
	TotalWins  int    `json:"totalWins"`
}

type ProfileCareer struct {
	TerranWins                int                 `json:"terranWins"`
	ZergWins                  int                 `json:"zergWins"`
	ProtossWins               int                 `json:"protossWins"`
	TotalCareerGames          int                 `json:"totalCareerGames"`
	TotalGamesThisSeason      int                 `json:"totalGamesThisSeason"`
	Current1v1LeagueName      string              `json:"current1v1LeagueName"`
	CurrentBestTeamLeagueName string              `json:"currentBestTeamLeagueName"`
	Best1v1Finish             ProfileBestFinished `json:"best1v1Finish"`
	BestTeamFinish            ProfileBestFinished `json:"bestTeamFinish"`
}

type ProfileBestFinished struct {
	LeagueName    string `json:"leagueName"`
	TimesAchieved int    `json:"timesAchieved"`
}

type ProfileSwarmLevels struct {
	Level   int                   `json:"level"`
	Terran  ProfileSwarmLevelRace `json:"terran"`
	Zerg    ProfileSwarmLevelRace `json:"zerg"`
	Protoss ProfileSwarmLevelRace `json:"protoss"`
}

type ProfileCampaign struct {
	DifficultyCompleted ProfileCampaignDifficultyCompleted `json:"difficultyCompleted"`
}

type ProfileCampaignDifficultyCompleted struct {
	WingsOfLiberty  string `json:"wings-of-liberty"`
	HeartOfTheSwarm string `json:"heart-of-the-swarm"`
	LegacyOfTheVoid string `json:"legacy-of-the-void"`
}

type ProfileSwarmLevelRace struct {
	Level              int `json:"level"`
	MaxLevelPoints     int `json:"maxLevelPoints"`
	CurrentLevelPoints int `json:"currentLevelPoints"`
}

type ProfileCategoryPointProgress struct {
	CategoryId   string `json:"categoryId"`
	PointsEarned int    `json:"pointsEarned"`
}

type ProfileEarnedReward struct {
	Category      string `json:"category"`
	RewardId      string `json:"rewardId"`
	AchievementId string `json:"achievementId"`
	Selected      bool   `json:"selected"`
}

type ProfileEarnedAchievement struct {
	AchievementId                    string                             `json:"achievementId"`
	CompletionDate                   float64                            `json:"completionDate"`
	NumCompletedAchievementsInSeries int                                `json:"numCompletedAchievementsInSeries"`
	TotalAchievementsInSeries        int                                `json:"totalAchievementsInSeries"`
	IsComplete                       bool                               `json:"isComplete"`
	InProgress                       bool                               `json:"inProgress"`
	Criteria                         []ProfileEarnedAchievementCriteria `json:"criteria"`
}

type ProfileEarnedAchievementCriteria struct {
	CriterionId string                                 `json:"criterionId"`
	Earned      ProfileEarnedAchievementCriteriaEarned `json:"earned"`
}

type ProfileEarnedAchievementCriteriaEarned struct {
	Quantity  int `json:"quantity"`
	StartTime int `json:"startTime"`
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
