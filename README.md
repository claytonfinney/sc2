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
