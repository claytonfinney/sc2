package sc2

import (
	"fmt"
)

// Target struct for the Legacy Profile resource:
// /sc2/legacy/profile/:regionId/:realmId/:profileId
type LegacyProfile struct {
	Id           string                    `json:"id"`
	Realm        int                       `json:"realm"`
	DisplayName  string                    `json:"displayName"`
	ClanName     string                    `json:"clanName"`
	ClanTag      string                    `json:"clanTag"`
	ProfilePath  string                    `json:"profilePath"`
	Portrait     LegacyIcon                `json:"portrait"`
	Career       LegacyProfileCareer       `json:"career"`
	SwarmLevels  LegacyProfileSwarmLevels  `json:"swarmLevels"`
	Campaign     LegacyProfileCampaign     `json:"campaign"`
	Season       LegacyProfileSeason       `json:"season"`
	Rewards      LegacyProfileRewards      `json:"rewards"`
	Achievements LegacyProfileAchievements `json:"achievements"`
}

type LegacyProfileSeason struct {
	SeasonId             int                  `json:"seasonId"`
	SeasonNumber         int                  `json:"seasonNumber"`
	SeasonYear           int                  `json:"seasonYear"`
	TotalGamesThisSeason int                  `json:"totalGamesThisSeason"`
	Stats                []LegacyProfileStats `json:"stats"`
}

type LegacyProfileCampaign struct {
	Wol  string `json:"wol"`
	Hots string `json:"hots"`
	Lotv string `json:"lotv"`
}

type LegacyProfileRewards struct {
	Selected []string `json:"selected"`
	Earned   []string `json:"earned"`
}

type LegacyProfileAchievements struct {
	Points        LegacyProfileAchievementsPoints         `json:"points"`
	Achievemenets []LegacyProfileAchievementsAchievements `json:"achievements"`
}

type LegacyProfileAchievementsPoints struct {
	TotalPoints    int            `json:"totalPoints"`
	CategoryPoints map[string]int `json:"categoryPoints"`
}

type LegacyProfileAchievementsAchievements struct {
	AchievementId  string `json:"achievementId"`
	CompletionDate int    `json:"completionDate"`
}

type LegacyProfileCareer struct {
	PrimaryRace      string `json:"primaryRace"`
	TerranWins       int    `json:"terranWins"`
	ProtossWins      int    `json:"protossWins"`
	ZergWins         int    `json:"zergWins"`
	Highest1v1Rank   string `json:"highest1v1Rank"`
	HighestTeamRank  string `json:"highestTeamRank"`
	SeasonTotalGames int    `json:"seasonTotalGames"`
	CareerTotalGames int    `json:"careerTotalGames"`
}

type LegacyProfileSwarmLevels struct {
	Level   int               `json:"level"`
	Terran  LegacyProfileRace `json:"terran"`
	Zerg    LegacyProfileRace `json:"zerg"`
	Protoss LegacyProfileRace `json:"protoss"`
}

type LegacyProfileStats struct {
	Type  string `json:"type"`
	Wins  int    `json:"wins"`
	Games int    `json:"games"`
}

type LegacyProfileRace struct {
	Level          int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

// Target struct for Legacy Rewards resource:
// /sc2/legacy/data/rewards/:regionId
type LegacyRewards struct {
	Portraits []LegacyReward `json:"portraits"`
	Skins     []LegacyReward `json:"skins"`
}

type LegacyReward struct {
	Title         string     `json:"title"`
	Id            string     `json:"id"`
	AchievementId string     `json:"achievementId"`
	Icon          LegacyIcon `json:"icon"`
}

// Target struct for Legacy Achievements resource:
// /sc2/legacy/data/achievements/:regionId
type LegacyAchievements struct {
	Achievements []LegacyAchievement         `json:"achievements"`
	Categories   []LegacyAchievementCategory `json:"categories"`
}

type LegacyAchievement struct {
	Title         string                    `json:"title"`
	Description   string                    `json:"description"`
	AchievementId string                    `json:"achievementId"`
	CategoryId    string                    `json:"categoryId"`
	Points        int                       `json:"points"`
	Icon          LegacyIcon                `json:"icon"`
	Category      LegacyAchievementCategory `json:"category"` //TODO: FIXME
}

type LegacyAchievementCategory struct {
	Title                 string                   `json:"title"`
	CategoryId            string                   `json:"categoryId"`
	FeaturedAchievementId string                   `json:"featuredAchievementId"`
	Children              []LegacyAchievementChild `json:"children"`
}

type LegacyAchievementChild struct {
	Title                 string `json:"title"`
	CategoryId            string `json:"categoryId"`
	FeaturedAchievementId string `json:"featuredAchievementId"`
}

// Target struct for Legacy Ladders resource:
// /sc2/legacy/:regionId/:realmId/:profileId/ladders
type LegacyLadders struct {
	CurrentSeason     []LegacySeason `json:"currentSeason"`
	PreviousSeason    []LegacySeason `json:"previousSeason"`
	ShowcasePlacement []LegacySeason `json:"showcasePlacement"`
}

type LegacySeason struct {
	Ladder     []LegacySeasonLadder `json:"ladder"`
	Characters []LegacyCharacter    `json:"characters"`
	NonRanked  []LegacyNonRanked    `json:"nonRanked"`
}

type LegacyNonRanked struct {
	Mmq         string `json:"mmq"`
	GamesPlayed int    `json:"gamesPlayed"`
}

type LegacySeasonLadder struct {
	LadderName       string `json:"ladderName"`
	LadderId         string `json:"ladderId"`
	Division         int    `json:"division"`
	Rank             int    `json:"rank"`
	League           string `json:"league"`
	MatchMakingQueue string `json:"matchMakingQueue"`
	Wins             int    `json:"wins"`
	Losses           int    `json:"losses"`
	Showcase         bool   `json:"showcase"`
}

// Target struct for the Legacy Ladder resource:
// /sc2/legacy/ladder/:regionId/:ladderId
type LegacyLadder struct {
	LadderMembers []LegacyLadderMember `json:"ladderMembers"`
}

type LegacyLadderMember struct {
	Character      LegacyCharacter `json:"character"`
	JoinTimestamp  int             `json:"joinTimestamp"`
	Points         int             `json:"points"`
	Wins           int             `json:"wins"`
	Losses         int             `json:"losses"`
	HighestRank    int             `json:"highestRank"`
	PreviousRank   int             `json:"previousRank"`
	FavoriteRaceP1 string          `json:"favoriteRaceP1"`
}

// Target struct for the Legacy Matches resource:
// /sc2/profile/:regionId/:realmId/:profileId/matches
type LegacyMatches struct {
	Matches []LegacyMatch `json:"matches"`
}

type LegacyMatch struct {
	Map      string `json:"map"`
	Type     string `json:"type"`
	Decision string `json:"decision"`
	Speed    string `json:"speed"`
	Date     int    `json:"date"`
}

// Some shared structs between a 2+ resources

type LegacyCharacter struct {
	Id          string `json:"id"`
	Realm       int    `json:"realm"`
	Region      int    `json:"region"`
	DisplayName string `json:"displayName"`
	ClanName    string `json:"clanName"`
	ClanTag     string `json:"clanTag"`
	ProfilePath string `json:"profilePath"`
}

type LegacyIcon struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
	Offset int    `json:"offset"`
	Url    string `json:"url"`
}

// API call functions

func (c *Client) GetLegacyProfile(region int, realm int, profile int) (*LegacyProfile, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/profile/%v/%v/%v", c.BaseUrl, region, realm, profile)
	p := &LegacyProfile{}
	err := c.Get(uri, p)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (c *Client) GetLegacyLadders(region int, realm int, profile int) (*LegacyLadders, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/%v/%v/%v/ladders", c.BaseUrl, region, realm, profile)
	l := &LegacyLadders{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetLegacyMatches(region int, realm int, profile int) (*LegacyMatches, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/%v/%v/%v/matches", c.BaseUrl, region, realm, profile)
	m := &LegacyMatches{}
	err := c.Get(uri, m)
	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *Client) GetLegacyAchievements(region int) (*LegacyAchievements, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/data/achievements/%v", c.BaseUrl, region)
	a := &LegacyAchievements{}
	err := c.Get(uri, a)
	if err != nil {
		return nil, err
	}

	return a, err
}

func (c *Client) GetLegacyLadder(region int, id int) (*LegacyLadder, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/ladder/%v/%v", c.BaseUrl, region, id)
	l := &LegacyLadder{}
	err := c.Get(uri, l)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (c *Client) GetLegacyRewards(region int) (*LegacyRewards, error) {
	uri := fmt.Sprintf("%s/sc2/legacy/data/rewards/%v", c.BaseUrl, region)
	r := &LegacyRewards{}
	err := c.Get(uri, r)
	if err != nil {
		return nil, err
	}

	return r, err
}
