package als_golang

import "strings"

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
func ArenasGeneralStats(APIToken string, platform string, player string) ArenasGeneral {
	plat := CheckPlatformArray(possible_platforms,platform)

	res, err := client.Get(BASE + "&platform=" + platform + "&player=" + player)
}