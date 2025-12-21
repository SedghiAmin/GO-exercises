package main

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
