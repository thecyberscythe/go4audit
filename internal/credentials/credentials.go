package credentials
import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
	"syscall"
)

func Credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	Username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	} else {
		log.Println(err) // Log (optional)
	}

	Password := string(bytePassword)

	return strings.TrimSpace(Username), Password, err
}
