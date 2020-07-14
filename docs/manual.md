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
