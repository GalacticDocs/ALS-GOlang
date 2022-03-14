package als

import (
	"fmt"

	http "github.com/GalacticDocs/ALS-GOlang/http_als"
	utils "github.com/GalacticDocs/ALS-GOlang/utils"
)

var Platforms []string = []string{
	utils.Origin,
	utils.XBox,
	utils.Playstation,
}

func GetArenasStats(platform string, player_name string) (*utils.GeneralArenas, error) {
	plat, platform_err := utils.CheckPlatformArray(Platforms, platform)
	if platform_err != nil {
		return nil, platform_err
	}

	res, err := http.GetHTTP(utils.BASE + "&platform=" + plat.platform_type + "&player=" + player)
	if err != nil {
		return nil, err
	}

	fmt.Printf("res: %v\n", string(res))
	return nil, nil
}
