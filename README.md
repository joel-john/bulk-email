# Bulk-Mailer
Program to send emails to recipients from a huge list in a performant way.

A simple client program to send bulk emails. Written completely in Go. It will send email message (from given HTML Template file) to a list of recipients (from given csv file) using multiple smtp server relays (from given serverconfiguration csv file).

The arguments are given through command line interface ([CLI](https://github.com/urfave/cli))



## About

The main aim of this program is to send emails to a huge list of recipients (~1 Million). For sending emails, atleast one SMTP Server/ Relay is required.

But for ensuring fast mass delivery, multiple SMTP can be utilized. In doing so, the program will divide the list of recipients to number of SMTP relays available and send the emails concurrently  (implemented using goroutines)

    i.e, if number of recipients in list = 10000 and number of smtp relays are 4,
    then each of the smtp servers send email to 2500 recipients concurrently

In order to avoid emails marked as spam, a custom delay between each mail can be given (other wise default delay of 50ms is used)


## Usage

For Test Running Examples and Instructions visit ([example](/example))

### Commandline

```

./bmailer [--template TEMPLATEFILE] [--recipient RECIPIENTLISTFILE] [--config CONFIGFILE] [--subject "subject"] [--delay "delay"]

    --template FILE, -t FILE               Load HTML template from FILE                 [Required Argument]
    --recipient FILE, -r FILE              Load recipient list (csv) from FILE          [Required Argument]
    --config FILE, -c FILE                 Load SMTPConfig File (csv) from FILE         [Required Argument]
    --subject "subject", -s "subject"      Specify the "subject" for email (default: "Test Mail")
    --delay delay(in ms), -d delay(in ms)  Specify the delay(in ms) between each mail (default: 50)
    --help, -h                             show help

```

### File Structures

#### Template

An HTML File

The contents of email are provided by the HTML Template file. 
For mail-merging of Name, {{.Name}} Tag is used

Example :
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

Example :
```
John Doe,johndoe@example.com
Jane Doe,jane@example.org
Richard Wilson,richard@example.com
Daniel Wilson,dani@example.org

``` 

#### Config File

A CSV file with 4 columns - username,password,hostname,port

It contains the list of SMTP Servers. 
username column contains smtp login username 
password column contains smtp login password
hostname column contains smtp server host
port column contains smtp port

Invalid SMTP details are automatically removed
Atleast one valid SMTP Server should be there for successful program execution

Example :
```
hello@example.com,password,localhost,1025
welcome@example.com,password,localhost,1025
``` 

