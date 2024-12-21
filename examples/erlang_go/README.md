# erlang_go

An OTP library that communicates to go via ports.

## Prerequisites

- Go (1.23)
- Erlang (27.1)
- Rebar3

## Build

### Build Go

```bash
make build
```

### Build Erlang

```bash
rebar3 compile
```

## Run

```bash
# Open Shell
rebar3 shell

# Call Go 
erlang_go:handle_call().
```
