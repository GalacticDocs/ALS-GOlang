package als

import (
	"strings"

	utils "github.com/GalacticDocs/ALS-GOlang/utils"
	"go.uber.org/zap"
)

var logConfig = zap.NewProductionConfig()

/* Checks for the main properties if they're defined/set in the config. */
func CheckDefaultConfig(config Config) bool {
	logConfig.Level.SetLevel(zap.DebugLevel)
	logger := utils.InjectableLogger(logConfig)

	if strings.TrimSpace(config.APIToken) == "" {
		logger.Error("The config requires a valid API Token but there was no token specified")
		return false
	}

	return true
}

func CheckBattleRoyaleConfig(config BattleRoyaleConfig) string {
	if config.GeneralOnly {
		return utils.GeneralStats
	} else if config.AdvancedOnly {
		return utils.AdvancedStats
	} else if config.Both {
		return utils.BothStats
	}

	return utils.GeneralStats
}

type Config struct {
	APIToken     string
	BattleRoyale *BattleRoyaleConfig
}

type BattleRoyaleConfig struct {
	GeneralOnly  bool
	AdvancedOnly bool
	Both         bool
}
