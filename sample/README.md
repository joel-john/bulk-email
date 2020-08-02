bulk-mailer test
===

## Test Environment Setup

Either you can use own SMTP Servers or get SMTP from services like Mailgun.

Free services like Gmail/Yahoo wont support sending bulk emails

### Using mailhog

A good option for testing emails is to use mailhog ([github.com/mailhog/MailHog](https://github.com/mailhog/MailHog)).

It will setup a test SMTP on localhost and display all emails that are sent from that 

#### MacOS
```bash
brew update && brew install mailhog
```

Then, start MailHog by running `mailhog` in the command line.

#### Debian / Ubuntu
```bash
sudo apt-get -y install golang-go
go get github.com/mailhog/MailHog
```

Then, start MailHog by running `/path/to/MailHog` in the command line.

E.g. the path to Go's bin files on Ubuntu is `~/go/bin/`, so to start the MailHog run:

```bash
~/go/bin/MailHog
```
#### Docker
[Run it from Docker Hub](https://registry.hub.docker.com/r/mailhog/mailhog/) 


## Files for Testing

By default, mailhog starts SMTP Server on port 1025 and HTTP Server on port 8025

So in ConfigFile.csv we can pass this as server information (username and password doesnt matter with mailhog since it catches all)

### Sample Files


Sample [template.html](./template.html) with sample HTML Template for email

Sample [ConfigFile.csv](./ConfigFile.csv) with 6 valid Server Records(for mailhog)

Sample [InvalidConfigFile.csv](./InvalidConfigFile.csv) with 4 valid and 2 invalid Server Records

Sample [RecipientList.csv](./RecipientList.csv) with 20 email records

Sample [LargeConfigFile.csv](./LargeConfigFile.csv) with 100 valid Server Records(for mailhog)

Sample [HugeRecipientList.csv](./HugeRecipientList.csv) with 1 million email records (contents of RecipientList is repeated multiple times)


## Test Run in Debian/ Ubuntu

Install using : 

```bash
go get github.com/joeljohn/bulk-mailer
```

Download all the Sample Files,

Open terminal in the Directory with sample files

Run MailHog using the above commands (with default mailhog configuration)

with default go bin path `/go/bin` in ubuntu run using:

```bash
~/go/bin/bulk-mailer -t template.html -r RecipientList.csv -c ConfigFile.csv -s "TEST SUBJECT" -d 50
```

### Viewing logs using mailhog 

on Webbrowser visit : 

`http://localhost:8025/`


## Sample Outputs

With delay of 100ms

```bash
~/go/bin/bulk-mailer -t template.html -r RecipientList.csv -c ConfigFile.csv -s "TEST SUBJECT" -d 100
```

Bash Output :

```
Successfully added config with username  server1@example.com
Successfully added config with username  server2@example.com
Successfully added config with username  server3@example.org
Successfully added config with username  server4@example.com
Successfully added config with username  server5@example.com
Successfully added config with username  server6@example.com
 Do you want to proceed Sending
Number of Email Records : 20
Number of Valid config  : 6

Email subject           : TEST SUBJECT
Delay between emails    : 100 ms
----------------------------------------------

Do you want to proceed Sending ? (Y/N)
y

Splitted Recipient List into 6 temporary files

Waiting To Finish Sending Emails
Email Sent to   user18@example.com
Email Sent to  jack@example.org
Email Sent to  johndoe@example.com
Email Sent to   user9@example.com
Email Sent to   user15@example.com
Email Sent to   user12@example.com
Email Sent to   user6@example.com
Email Sent to   user19@example.com
Email Sent to  jane@example.org
Email Sent to   user16@example.com
Email Sent to   user13@example.com
Email Sent to   user10@example.com
Email Sent to   user7@example.com
Email Sent to   user20@example.com
Email Sent to   user17@example.com
Email Sent to  richard@example.com
Email Sent to   user14@example.com
Email Sent to   user11@example.com
Email Sent to   user8@example.com
Email Sent to  dani@example.org

Terminating Program and removing temporary files
```
Mailhog Output:

![MailHog Output](./MailHog-Output.png?raw=true)

Note that with 20 recipients and 6 servers, it sent a total of 26 emails. 
The first 6 emails are for verifying smtp configuration


Email Content:

![Email Content](./Email-Content.png?raw=true)

