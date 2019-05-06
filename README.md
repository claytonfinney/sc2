### sc2: A Golang client for the StarCraft II Web API

This package is intended for developers wanting to create applications using Blizzard's StarCraft II web API. All currently available endpoints are listed on the [Blizzard Developer Documentation](https://develop.battle.net/documentation/api-reference/starcraft-2-community-api).

All payloads are unmarshalled into structs with an almost entirely uniform naming scheme. In most cases, you can expect that a field's name is the same as Blizzard's JSON payload, just with a capital first letter. Some legacy endpoints have snake\_case field names in their payloads, which have been converted to CamelCase in their corresponding struct. 

All functions in this package that call Blizzard APIs follow the convention of \<Method\>\<EndpointName\>. The only exception to this convention is that there is an endpoint named `getLeagueData`, which does not have a corresponding function called `GetgetLeagueData`, just `GetLeagueData`.

#### Installation
````Go
go get github.com/claytonfinney/sc2
````

#### Example Usage
````Go
import (
	"fmt"
	"log"
	"github.com/claytonfinney/sc2"
)

func main() {
	conf := &sc2.ClientConfig{
		Region: "us",
		Client: "YOUR-CLIENT-HERE",
		Secret: "YOUR-SECRET-HERE",
		Test: false, // Not a necessary parameter in normal use
	}
	client := sc2.NewClient(conf)

	szn, err := client.GetSeason(1)
	if err != nil {
		log.Fatal(err)
	}
	curr := szn.SeasonId // corresponding Blizzard JSON field is seasonId
	// seasonId, queueId = 201 for LotV, teamType = 0 for 1v1, leagueId = 6 for Grandmaster)
	d, err := client.GetLeagueData(curr, 201, 0, 6)
	if err != nil {
		log.Fatal(err)
	}
	id := d.Tier[0].Division[0].LadderId
	grandmasters, err := client.GetLegacyLadder(1, id) // Currently the way to fetch ladder data without a unique player ID, may change soon
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grandmasters.LadderMembers[0].Character.DisplayName) // Gets the first place Grandmaster!
}

````

#### All available endpoints

For convenience:

**ladder.go**
`func (c *Client) GetLadderSummary(region int, realm int, profile int) (*LadderSummary, error)`
`func (c *Client) GetLadder(region int, realm int, profile int, ladder int) (*Ladder, error)`
`func (c *Client) GetGrandmaster(region int) (*Ladder, error)`
`func (c *Client) GetSeason(region int) (*Season, error)`

**league.go**
`func (c *Client) GetLeagueData(season int, queue int, team int, league int) (*LeagueData, error)`

**player.go**
Endpoint currently returns `[]`.

**legacy.go**
`func (c *Client) GetLegacyProfile(region int, realm int, profile int) (*LegacyProfile, error)`
`func (c *Client) GetLegacyLadders(region int, realm int, profile int) (*LegacyLadders, error)`
`func (c *Client) GetLegacyMatches(region int, realm int, profile int) (*LegacyMatches, error)`
`func (c *Client) GetLegacyAchievements(region int) (*LegacyAchievements, error)`
`func (c *Client) GetLegacyLadder(region int, id int) (*LegacyLadder, error)`
`func (c *Client) GetLegacyRewards(region int) (*LegacyRewards, error)`

**static.go**
`func (c *Client) GetStatic(region int) (*Static, error)`
`func (c *Client) GetMetadata(region int, realm int, profile int) (*Metadata, error)`
`func (c *Client) GetProfile(region int, realm int, profile int) (*Profile, error)`

#### Things I'm planning to add
Right now the actual client is pretty bare-bones, so some things I'd like to add are:
* A built-in rate limiter
* Refactor the client's `Get()` to allow for auto-retries

#### Contributing
If you'd like to contribute, whether it be something small like spelling or helping contribute the things listed above, please do open up a pull request! If your changes break the tests, don't waste your time trying to update the fixtures as the tests are currently pretty fickle (unless you want to!).
