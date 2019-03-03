package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	log.SetFlags(0)

	var (
		passwordHash string
		passwd       []byte
	)

	flag.StringVar(&passwordHash, "h", "", "set the password hash to compare against (use this flag with single quotes like: -h='$HASH')")
	flag.Parse()
	passwd, err := readPassword("Enter password:")
	if err != nil {
		log.Fatal(err)
	}

	if passwordHash != "" {
		err := bcrypt.CompareHashAndPassword([]byte(passwordHash), passwd)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("OK")
	} else if len(passwd) > 0 {
		data, err := bcrypt.GenerateFromPassword(passwd, bcrypt.MinCost)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = bcrypt.CompareHashAndPassword(data, passwd)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("Hash:", string(data))
	} else {
		log.Fatal("Empty password given.")
	}

}

func readPassword(msg string) (passwd []byte, err error) {
	fmt.Println(msg)
	passwd, err = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return
}
