package auth

import "fmt"

type Session struct {
	Hostname string
}

func StartSession(hostname string) (Session, error) {
	var (
		result Session
		args   []string
	)

	args = append(args, hostname)

	fmt.Println(args, hostname)

	return result, nil
}
