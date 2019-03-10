package sc2

import (
	"errors"
)

func ConvertLeagueToID(league string) (int, error) {
	switch league {
	case "bronze":
		return 0, nil
	case "silver":
		return 1, nil
	case "gold":
		return 2, nil
	case "platinum":
		return 3, nil
	case "diamond":
		return 4, nil
	case "master":
		return 5, nil
	case "grandmaster":
		return 6, nil
	default:
		return 0, errors.New("Error in call to ConvertLeagueToID: invalid string passed as argument")
	}
}
