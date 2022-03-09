package main

import (
	creds "go4audit/internal/credentials"
	remote "go4audit/internal/remoting"
)

func main() {
	var username, password, _ = creds.Credentials()
	var hosts = []string{"jira", "confluence", "graylog", "chat", "mail", "bitbucket", "repo"}

	remote.SSHing(username, password)
}
