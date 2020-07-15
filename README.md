gRPC SMTP Email Receiver
=======

### Simple SMTP connected using gRPC

Learning gRPC with Golang. Using go-smtp server to receive emails and then transfer to server using gRPC.

The following Golang Project uses:
- [gRPC](https://grpc.io/)
- [go-smtp](https://github.com/emersion/go-smtp)
- [viper](https://github.com/spf13/viper)

### Setup

Rename config-example.json to config.json

Change emailDomain and gRPCServer
- emailDomain - email domain for accepting emails
- gRPCServer - localhost:3001 is example in github/scorpionknifes/cmd/example-server

```json
{
  "SMTPPort": 25,
  "timeout": 10,
  "server": [
    {
      "emailDomain": "email.com",
      "gRPCServer": "localhost:3001"
    }
  ]
}
```

### How does it work

Using SMTP protocol we can receive emails using port 25. We can setup email domain to locate to SMTP server by using MX record for DNS. 
The SMTP receiver can only receive emails that are listed inside config.json, all other emails would be rejected by SMTP server.

When we receive an email through SMTP protocol we use the gRPC protocol to send Email content to gRPC Server.
There is an example receiver in github/scorpionknifes/cmd/example-server that is already correctly configured with config-example.json.

### How to run

Run gRPC SMTP receiver
- default port :3000

```
go run github.com/scorpionknifes/grpc-smtp-receiver
```

Run gRPC example server
- default port :3001

```
go run github.com/scorpionknifes/grpc-smtp-receiver/cmd/example-server
```

### How to update proto

Make sure to have Protocol Buffers v3 and protoc installed

Generate gRPC code using [protoc](https://developers.google.com/protocol-buffers/docs/downloads) tool

```
protoc --go_out=plugins=grpc:. *.proto
```


### Testing

Run gRPC SMTP receiver on port :3000 and gRPC example server on port :3001. 

Use Telnet to access and send email to SMTP server

```
telnet localhost 25
```

rcpt email must be setup correctly in "server" array in config.json

```
ehlo email.com
mail from: from@email.com
rcpt to: to@email.com
data
Test Data
.

```