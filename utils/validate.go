package util

func validateParams(dest string, params map[string]string) bool {
	switch dest {
	case "Reddit":
		return validateReddit(params)
	case "Twitter":
		return validateTwitter(params)
	case "Scheduler":
		return validateScheduler(params)
	default:
		return false
	}
}

func validateReddit(params map[string]string) bool {
	_, valid1 := params["after"]
	_, valid2 := params["before"]
	_, valid3 := params["subreddit"]
	return valid1 && valid2 && valid3
}

func validateTwitter(params map[string]string) bool {
	_, valid := params["sinceId"]
	return valid
}

func validateScheduler(params map[string]string) bool {
	// add once scheduler is completed
	return true
}