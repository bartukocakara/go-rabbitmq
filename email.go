package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)
func main() {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", os.Getenv("FROM_USERNAME"), os.Getenv("FROM_PASSWORD"), "smtp.mailtrap.io")
	// Here we do it all: connect to our server, set up a message and send it
	to := []string{os.Getenv("TO_EMAİL")}
	msg := []byte("To:"+ os.Getenv("TO_EMAİL") +"\r\n" +
		"Subject: Sales reports are here please check!\r\n" +
		"\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, os.Getenv("TO_EMAİL"), to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success")
}	