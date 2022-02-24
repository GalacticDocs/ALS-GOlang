package als_golang

var possible_platforms = []string{
	"PC",
	"X1",
	"PS4",
}

type PlatformChecker struct {
	status        bool
	platform_type string
}

func CheckPlatformArray(array []string, value string) *PlatformChecker {
	var status []bool
	var platform_type []string

	for i := 0; i <= 2; i++ {
		if value != array[i] {
			append(&status)
		} else {
			checker = &PlatformChecker{
				status:        true,
				platform_type: array[i],
			}
		}
	}

	return checker
}
