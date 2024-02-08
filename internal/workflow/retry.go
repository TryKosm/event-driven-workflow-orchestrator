package workflow

func ShouldRetry(attempt int, max int) bool {
	return attempt < max
}
