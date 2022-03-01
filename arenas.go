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
func ArenasGeneralStats(APIToken string, platform string, player string) (*ArenasGeneral, error) {
	plat, platform_err := CheckPlatformArray(possible_platforms, platform)
	if platform_err != nil {
		return nil, platform_err
	}

	res, http_err := client.Get(BASE + "&platform=" + plat.platform_type + "&player=" + player)
	if http_err != nil {
		return nil, http_err
	}

	fmt.Printf("res: %v\n", res)
	return nil, nil
}
