package als_utils

import (
	"errors"
)

var (
	possible_platforms = []string{"PC", "X1", "PS4"}
)

type PlatformChecker struct {
	status        bool
	platform_type string
}

func CheckPlatformArray(array []string, value string) (*PlatformChecker, error) {
	var (
		checker *PlatformChecker = &PlatformChecker{
			status:        false,
			platform_type: "unknown",
		}
		InvPlatfErr = errors.New("The value \"" + value + "\" is not a valid platform, the platform must be \"PC\", \"X1\" or \"PS4\".")
	)

	for i := 0; i <= len(array); i++ {
		if array[i] == value {
			checker = &PlatformChecker{
				status:        true,
				platform_type: value,
			}
		}
	}

	if !checker.status || checker.platform_type == "unknown" {
		return nil, InvPlatfErr
	}

	return checker, nil
}
