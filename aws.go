package util

import "os"

// IsRunByAWS checks if the runner of this application is AWS by checking if AWS_ACCESS_KEY and AWS_SECRET_KEY environment variables exist
func IsRunByAWS() bool {
	_, ok := os.LookupEnv("BB_IS_AWS")
	if ok {
		return true
	}
	_, ok = os.LookupEnv("AWS_ACCESS_KEY")
	if !ok {
		return false
	}
	_, ok = os.LookupEnv("AWS_SECRET_KEY")
	return ok
}
