package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	apps := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	for _, char := range log {
		if app, ok := apps[char]; ok {
			return app
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var r string
	for _, char := range log {
		if oldRune != char {
			r += string(char)

		} else {
			r += string(newRune)
		}
	}
	return r
}


// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := utf8.RuneCountInString(log)
	if length <= limit {
		return true
	}
	return false
}
