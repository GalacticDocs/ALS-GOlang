package main

import (
	http "github.com/GalacticDocs/ALS-GOlang/http_als"
	utils "github.com/GalacticDocs/ALS-GOlang/utils"
)

func New(conf Config) IArenas {
	if conf.APIToken == nil {}
}

func GetGeneralArenas() /*(*utils.PlatformChecker,*/ error {
	plat, platform_err := utils.CheckPlatformArray(utils.possible_platforms, platform)
	if platform_err != nil {
		return platform_err
	}

	res, err := http.GetHTTP(utils.BASE + "&platform=" + plat.platform_type + "&player=" + player)
	if err != nil {
		return err
	}

	fmt.Printf("res: %v\n", string(res))
	return nil
}


type IArenas struct {

}

type ArenasProps interface {

}

type Config struct {
	APIToken string
}