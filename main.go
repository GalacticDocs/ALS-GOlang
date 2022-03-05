package als_golang

import (
	"fmt"
)

/*
**Functionality**: Used to get the general statistics for a player from Arenas kills, damage done, wins.
**Possible platforms**: "PC", "PS4", "X1"
**Player Name**: You must use the Origin account name linked to the Steam account otherwise you may run into errors.
 */
func ArenasGeneralStats(APIToken string, platform string, player string) error {
	
}

func main() {
	ArenasGeneralStats("", "PC", "MochiIsVibin")
}
