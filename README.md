# Telbiz Client Libraries for Go

This repository contains Telbiz's Go client libraries. These libraries are officially maintained by Telbiz.
The libraries follow the Telbiz API closely and are designed to provide a native developer experience for Go developers.

## Documentation

The documentation for the Telbiz API can be found in [Docs](https://telbiz.la/pages/doc/user-guide).

## Installation

Use go get to retrieve the latest version of the client.

```bash
go get -u github.com/Codecamplao/Telbiz.SDK.Golang/v2
```

## Usage

```go
import telbiz "github.com/Codecamplao/Telbiz.SDK.Golang/v2"
```

Construct a new Telbiz client, then use the various services on the client to
access different parts of the Telbiz API. For example:

```go
ctx := context.Background()
client, err := telbiz.NewClient(ctx, "YOUR_API_KEY", "YOUR_SECRET_KEY")
if err != nil {
    panic(err)
}
```

Send SMS

```go
message := &telbiz.Message{
    To:    "2077805085",
    Title: telbiz.Info,
    Body:  "This is an open source Tizbiz Client Library for Go develop by jvonxay0308!",
}
sms, err := client.SendSMS(ctx, message)
if err != nil {
    panic(err)
}
```

## Go Versions Supported

This library supports the following Go implementations:

- Go 1.20.x
- Go 1.19.x
- Go 1.18.x

## Contributing

Contributions are welcome. Please open up an issue or create a pull request if you would like to help out.
