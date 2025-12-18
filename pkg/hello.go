package pkg

import "github.com/JusticeN/erstemodul/internal/auth"

func Hello() string {
	auth.Auth()
	//auth.auth()
	return "Version 0.0.1"
}
