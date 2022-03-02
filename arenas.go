package als_golang

import (
	"fmt"
)

var BASE = "https://api.mozambiquehe.re/bridge?version=5"

type ArenasGeneral struct {
	arenas_wins   string
	arenas_kills  string
	arenas_damage string
}

/*
**Functionality**: Used to get the general statistics for a player from Arenas kills, damage done, wins.
**Possible platforms**: "PC", "PS4", "X1"
**Player Name**: You must use the Origin account name linked to the Steam account otherwise you may run into errors.
 */
func ArenasGeneralStats(APIToken string, platform string, player string) error {
	plat, platform_err := CheckPlatformArray(possible_platforms, platform)
	if platform_err != nil {
		return platform_err
	}

	res, err := GetHTTP(BASE + "&platform=" + plat.platform_type + "&player=" + player)
	if err != nil {
		return err
	}

	fmt.Printf("res: %v\n", string(res))
	return nil
}
