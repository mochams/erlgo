# erlgo

A lightweight Go library that implements the Erlang port protocol for communication between Go and Erlang/OTP applications.

## Features

- Implements standard Erlang port protocol (4-byte length prefix)
- Clean, simple API for reading and writing messages
- Zero dependencies beyond Go standard library
- Designed for efficiency with buffered reading

## Installation

```bash
go get github.com/mochams/erlgo@latest
```

## Usage

```go
package main

import (
    "log"
    "github.com/mochams/erlgo"
)

func main() {
    // Read a message from Erlang
    messageBytes, err := erlgo.Receive()
    if err != nil {
        log.Fatal(err)
    }
    
    // Process your message bytes here...
    
    // Write a response back to Erlang
    response := []byte("your response data")
    if err := erlgo.Send(response); err != nil {
        log.Fatal(err)
    }
}
```

## API

### Receive

```go
func Receive() ([]byte, error)
```

Reads a length-prefixed message from Erlang through stdin. Returns the message bytes and any error encountered.

### Send

```go
func Send(messageBytes []byte) error
```

Writes a length-prefixed message to Erlang through stdout. Takes the message bytes and returns any error encountered.

## Erlang Port Protocol

The library implements the standard Erlang port protocol where each message is prefixed with a 4-byte length in big-endian order, followed by the actual message content.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

Inspired by the needs of the Erlang/OTP community for reliable Go port communication.
