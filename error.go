package aoc

import "fmt"

// NetworkError occurs when something went wrong with the HTTP request.
type NetworkError struct {
	Original error
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("Network error: %s", e.Original.Error())
}

// DownloadError occurs when the AoC server returns with a status code
// other than 200.
type DownloadError struct {
	StatusCode int
}

func (e DownloadError) Error() string {
	return fmt.Sprintf("Download error: %d", e.StatusCode)
}

// SubmitError occurs when the AoC server returns with a status code
// other than 200.
type SubmitError struct {
	StatusCode int
}

func (e SubmitError) Error() string {
	return fmt.Sprintf("Submit error: %d", e.StatusCode)
}

// IncorrectAnswer occurs when the provided answer is not correct.
type IncorrectAnswer struct {
	Hint string
	Wait string
}

func (e IncorrectAnswer) Error() string {
	return fmt.Sprintf("Error: %s\n%s", e.Hint, e.Wait)
}
