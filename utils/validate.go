package utils

import (
	"fmt"
)

func validateParams(dest string, params map[string]string) bool {
	switch dest {
	case "Reddit":
		return validateReddit(params)
	case "Twitter":
		return validateTwitter(params)
	default:
		return false
	}
}

func validateReddit(params map[string]string) bool {

}

func validateTwitter(params map[string]string) bool {

}