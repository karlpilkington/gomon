## monitor

### Introduction
`monitor` provides basic runtime monitoring for Go programs. It is
based on [monitor.py](https://github.com/kisom/pymods/blob/master/monitor.py),
a Python module originally written to ensure a Bitcoin broker continued
running even in the face of unanticipated exceptions. 

### Overview
The `monitor.Monitor` function takes a target function with the signature:

```go
func target() error {}
```

It is assumed the target function will use configuration files, `os.Args`,
environment variables, or some other mechanism to configure itself.

When the target function returns, `monitor` checks whether an error has
occurred and whether a panic has occurred, and sends a notification in
either case. If the function exits with no errors and hasn't panicked,
`monitor` assumes the program exits normally. This behaviour may be changed
by modifying the CanExit value in the package:

```go
monitor.CanExit = false     // disable exit on clean return
monitor.CanExit = true      // enable exit on clean return
```

### Example Usage

```go

func target() {
        // self-contained main program code
}

func main() {
        monitor.Monitor(target)
}
```

### Notifications
`monitor` supports [Pushover](https://www.pushover.net) and email 
notifications. To enable either, you can load the configuration from a
JSON file:

```json
{"mail": 
    {"port": "<smtp-port>",
     "pass": "<smtp-password>", 
     "user": "<smtp-user>", 
     "address": "<from-address>", 
     "server": "<smtp-server>",
     "to": ["<dev@yourdomain.tld>, <otherdev@theirdomain.tld>"]}, 
 "pushover": 
    {"token": "<api-key>", 
     "user": "<user-key>"}}
```

You may include only one of the two sections to enable that type of
notifications. By default, `monitor` looks for `monitor.json` in the
same directory as the code is being run from.

```go
        err := monitor.ConfigFromJson()
        if err != nil {
                fmt.Println("[!] error configuring monitor: ", err)
                os.Exit(1)
        }
        monitor.Monitor(target)
```

The configuration file can be set with the variable `ConfigFile`:

```go
monitor.ConfigFile = "secret.json"
```

Alternatively, you can load the configuration from the environment. The
environment variables support are:

* `MAIL_SERVER`
* `MAIL_USER`
* `MAIL_PASS`
* `MAIL_ADDRESS`
* `MAIL_PORT`
* `MAIL_TO` - this should be a comma separated list of address to email
when notifications go out. For example, "dev1@example.com,dev2@example.com".
* `PO_APIKEY`
* `PO_USER`

To load the configuration, call `ConfigFromEnv`:

```go
        monitor.ConfigFromEnv() 

        // check whether the notifications we expect to be enabled were enabled
        if monitor.PushoverEnabled() {
                log.Println("[!] pushover notifications not enabled!")
        }

        if monitor.EmailEnabled() {
                log.Println("[!] email notifications not enabled!")
        }
```
