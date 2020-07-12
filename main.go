package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// app := &cli.App{
	// 	Name:  "Bulk Email",
	// 	Usage: "Send Bulk Emails",
	// 	Action: func(c *cli.Context) error {
	// 		fmt.Println("For help run mail --help")
	// 		return nil
	// 	},
	// }

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

//ServerSetup For setting up server
func ServerSetup() {

}

//Read reads list of recipients from csv file
func Read() {

	csvFile, _ := os.Open("recipients.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

}

//Send for sending email
func Send() {

	fmt.Println("Email Sent!")

}
