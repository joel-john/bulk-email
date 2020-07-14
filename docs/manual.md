bmailer manual
===
<!-- toc -->

- [Commands](#commands)
    - [Files](#files)
        - [File Structure](#file-structures)
            - [Template](#template)
            - [RecipientList](#RecipientList)
            - [ConfigFile](#ConfigFile)
    - [Subject](#subject)
    - [Delay](#delay)
- [Functions](#functions)
    - [main()](#main)
    - [SplitRecipients()](#SplitRecipients)
    - [VerifyCSV()](#VerifyCSV)
    - [ParseTemplate()](#ParseTemplate)
    - [ParseServerConfig()](#ParseServerConfig)
    - [ParseRecipient()](#ParseRecipient)
    - [SendEmail()](#SendEmail)
    - [ValidateFormat()](#ValidateFormat)



<!-- tocstop -->

## Commands

```

./bmailer [--template TEMPLATEFILE] [--recipient RECIPIENTLISTFILE] [--config CONFIGFILE] [--subject "subject"] [--delay "delay"]

    --template FILE, -t FILE               Load HTML template from FILE                 [Required Argument]
    --recipient FILE, -r FILE              Load recipient list (csv) from FILE          [Required Argument]
    --config FILE, -c FILE                 Load SMTPConfig File (csv) from FILE         [Required Argument]
    --subject "subject", -s "subject"      Specify the "subject" for email (default: "Test Mail")
    --delay delay(in ms), -d delay(in ms)  Specify the delay(in ms) between each mail (default: 50)
    --help, -h                             show help

```

### Files

The files are passed as CLI flags, all File flags are mandatory

Usage:

    --template TemplateFile, -t TemplateFile                Load HTML template from TemplateFile
    --recipient RecipientListFile, -r RecipientListFile     Load recipient list (csv) from RecipientListFile
    --config ConfigFile, -c ConfigFile                      Load SMTPConfig File (csv) from ConfigFile

#### Files Structures

#### Template

An HTML File

The contents of email are provided by the HTML Template file. 
For mail-merging of Name, {{.Name}} Tag is used

Example Template File :
```
<html>
   <head>
      <title></title>
      <meta content="">
      <style></style>
   </head>
   <body>
      <p>Hello {{.Name}}</p>
      <p><br></p>
      <p>This is a Test Email</p>
      <p><br></p>
      <p>Thanks,</p>
      <p>John Doe</p>
   </body>
</html>
```

#### RecipientList

A CSV file with two columns - Name, Email

It contains the list of email recipients. 
Name column contains Name of Recipients (Which is used for mail-merge)
Email column contains email address of recipients (Used for sending emails)
The program uses a basic regex validation for validating email address formats.

Example RecipientList File:
```
John Doe,johndoe@example.com
Jane Doe,jane@example.org
Richard Wilson,richard@example.com
Daniel Wilson,dani@example.org

``` 

#### ConfigFile

A CSV file with 4 columns - username,password,hostname,port

It contains the list of SMTP Servers. 
username column contains smtp login username 
password column contains smtp login password
hostname column contains smtp server host
port column contains smtp port

Invalid SMTP details are automatically removed
Atleast one valid SMTP Server should be there for successful program execution

Example Config File:
```
hello@example.com,password,localhost,1025
welcome@example.com,password,localhost,1025
``` 
### Subject

The subject of email is passed as a CLI argument. 
If no subject is given, default subject "Test Mail" is used

Usage:

    --subject "Hello World!", -s "Hello World!"      sets 'Hello World!' as the subject for email

### Delay

If custom delays between emails are needed, this flag should be set,
If no arguments are given, default delay of 50ms is used

    --delay 500, -d 500             sets a delay of 500 ms between each email


## Functions

### main()

main controls overall flow of the program. The cli.App is implemented which takes arguments through command line interface.

* The input from command line flags are stored
* Calls [VerifyCSV](#verifycsv) on recipientListFile, configFile and stores the computed record legth
* The info for server configuration and server counf is obtained by calling [ParseServerConfig](#ParseServerConfig) on ConfigFile
* If there are any valid server configurations, It displays the necessary informations to the user and waits for user input before proceeding (Y/N)
* If the user proceeds, [SplitRecipients](#SplitRecipients) is called, which creates temporary files(according to number of valid servers)
* Each temporary recipientFiles and individual server configurations are passed to [ParseRecipient](#ParseRecipient) along with delay
* The ReadRecipient is called using goroutines for ensuring concurrent execution
* At the end of program, temporary recipientFiles are removed


### SplitRecipients()

`func SplitRecipients(recipientListFileName string, serverCount, recordLength int)`

SplitRecipients splits the recipientListFileName into different files (According to the number of server count)

i.e,If SplitRecipient is called with the following RecipientListFile.csv

```
John Doe,johndoe@example.com
Jane Doe,jane@example.org
Richard Wilson,richard@example.com
Daniel Wilson,dani@example.org
Jack,jack@example.org

``` 

If there are 2 valid servers, then the RecipientList will be split into two files

BMail_recipientList1.csv

```
John Doe,johndoe@example.com
Jane Doe,jane@example.org
Richard Wilson,richard@example.com
```

BMail_recipientList2.csv

```
Daniel Wilson,dani@example.org
Jack,jack@example.org
```

This process will make it easier for sending emails to each list from respective SMTPs
    


### VerifyCSV()

`func VerifyCSV(recipientListFileName, configFileName string) int `

VerifyCSV verifies the CSV Files (recipientListFile and configFile), If error occurs during parsing, it is thrown

It also returns the number of records in reipientList

Example RecipientList File:
```
John Doe,johndoe@example.com
Jane Doe,jane@example.org
Richard Wilson,richard@example.com
Daniel Wilson,dani@example.org

``` 

Calling VerifyCSV on this will return 5 since it contains 5 records

### ParseTemplate()

`func ParseTemplate(templateFileName string, data interface{}) string`

ParseTemplate mail-merges the data into templateFile

example :
template.html
```
<html>
   <head>
      <title></title>
      <meta content="">
      <style></style>
   </head>
   <body>
      <p>Hello {{.Name}}</p>
      <p><br></p>
      <p>This is a Test Email</p>
      <p><br></p>
      <p>Thanks,</p>
      <p>John Doe</p>
   </body>
</html>
```
After using parseTemplate with Data as Name:`John Doe`

The output file will be 

```
<html>
   <head>
      <title></title>
      <meta content="">
      <style></style>
   </head>
   <body>
      <p>Hello John Doe</p>
      <p><br></p>
      <p>This is a Test Email</p>
      <p><br></p>
      <p>Thanks,</p>
      <p>John Doe</p>
   </body>
</html>

```

`John Doe` is inserted into `{{.Name}}`

### ParseServerConfig()

`func ParseServerConfig(configFileName string) (int, []string, []string, []string, []string)`

ParseServerConfig parses the configFile

* Checks whether the SMTP configs are valid by sending a test mail to the `username`
* If valid each information (username, password, hostname, password) are returned as slices
* Also returns the number of valid SMTP configs


### ParseRecipient()

`func ParseRecipient(recipientListFileName, templateFileName, configFileName, subject, from string, serverstruct ServerConfig, delay int, wg *sync.WaitGroup)`

ParseRecipient reads the list of recipients from recipientListFile and perform the following for each recipient :

* calls ValidateFormat to validate email address structure
* if valid, mail-merges template using the name from recipients to create body of email
* Creates message structure according to parameters read from file and passed as arguments
* Call the [sendEmail](#sendemail) with arguments
* Sleeps for the set delay

### SendEmail()

`func (m *Message) SendEmail(username, password, hostname, port string)`

SendEmail sends the email with ServerInfo and Message details

If an error occurs during sending email, it will wait for some time and tries to send email again. If it happens for more than 10 times, It will print a console output and skips sending

### ValidateFormat()

`func ValidateFormat(email string) error`

ValidateFormats checks the format of an email address using regex
If validation fails, an error is thrown