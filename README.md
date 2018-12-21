# create-jwt
this is a cli-tools for create jwt.

# Run
```
$ go run ./
```

# Build
```
GOOS=[your OS] GOARCH=amd64 go build
```

# Usage
```
Usage: create-jwt -s filename [-e second] [-n user_name]
  -e int
        expiration second (default 3600)
  -n string
        username (default "trrrrrys")
  -s string
        secret key file
```