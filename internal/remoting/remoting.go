package remoting

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

func SSHing(username string, password string) {
	//TODO WIP. Move ssh routine here

	var results = make(chan string, 10)
	var timeout = time.After(120 * time.Second)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	/////////////////////////////////////////
	//Connect to All Hosts and Run commands//
	/////////////////////////////////////////

	for _, hostname := range hosts {
		go func(hostname string) {
			executeCmd(wget, hostname, config)
			results <- executeCmd(cmd, hostname, config)
		}(hostname)
	}
	/////////////////////////////////////////
	//Return Results//
	/////////////////////////////////////////
	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			fmt.Print(res)
		case <-timeout:
			fmt.Println("Timed out!")
			return
		}
	}
}

func executeCmd(cmd, hostname string, config *ssh.ClientConfig) string {

	client, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	if err := session.Run(cmd); err != nil {
		log.Print("Failed to run: " + err.Error())
	}

	return hostname + ": " + stdoutBuf.String()

}
}
