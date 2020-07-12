package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
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

//Message struct
type Message struct {
	to      string
	from    string
	subject string
	body    string
}

//ServerSetup For setting up server
func ServerSetup() {

}

//ParseTemplate parses the template
func ParseTemplate(templateFileName string, data interface{}) string {

	// Open the file
	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Fatalln("Couldn't open the template", err)
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, data)
	return buf.String()
}

//ReadRecipient reads list of recipients from csv file
func ReadRecipient(recipientListFileName, templateFileName, from, subject string) {

	// Open the file
	csvFile, err := os.Open(recipientListFileName)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data := struct {
			Name string
		}{
			Name: record[1],
		}
		body := ParseTemplate(templateFileName, data)
		m := Message{
			to:      record[3],
			subject: subject,
			body:    body,
		}
	}

}
