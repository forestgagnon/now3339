# now3339

Print the current time formatted as an RFC3339 timestamp.

## Installation

```
go install github.com/forestgagnon/now3339@latest
```

## Usage

```
now3339

now3339 plus 2h

now3339 minus 5m3s
```

Command substitution example (check out `jo` at https://github.com/jpmens/jo)

```
$ jo time=$(now3339 plus 24h) otherField="blah"

{"time":"2023-03-16T15:45:04-07:00","otherField":"blah"}
```
