# Bulk-Mailer
Program to send emails to recipients from a huge list in a performant way.

A simple client program to send bulk emails. Written completely in Go. It will send email message (from given HTML Template file) to a list of recipients (from given csv file) using multiple smtp server relays (from given serverconfiguration csv file).

The arguments are given through command line interface (CLI) [github.com/urfave/cli](https://github.com/urfave/cli))



## About

The main aim of this program is to send emails to a huge list of recipients (~1 Million). For sending emails, atleast one SMTP Server/ Relay is required.

But for ensuring fast mass delivery, multiple SMTP can be utilized. In doing so, the program will divide the list of recipients to number of SMTP relays available and send the emails concurrently  (implemented using goroutines)

    i.e, if number of recipients in list = 10000 and number of smtp relays are 4,
    then each of the smtp servers send email to 2500 recipients concurrently

In order to avoid emails marked as spam, a custom delay between each mail can be given (other wise default delay of 50ms is used)

CSV file validation, email format validation and SMTP server validations are built into the program


## Usage and Documentation

For Detailed Documentation visit ([./docs/manual.md](./docs/manual.md))

For Test Demo Examples and Instructions visit ([./example](./example))


## Prerequisites for building

cli package should be installed

```
$ GO111MODULE=on go get github.com/urfave/cli/v2
$ GO111MODULE=on go get github.com/urfave/cli
```

## Licence

MIT License, see [LICENSE](./LICENSE)
