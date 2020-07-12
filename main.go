package main

import (
	"bytes"
	"log"
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
