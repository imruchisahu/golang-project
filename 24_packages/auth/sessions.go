package auth

func extractSession() string{ // if use small letter than it is private
	return "loggedin"
}
func GetSession() string {
	return extractSession()
}
