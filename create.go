package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var sm jwt.SigningMethod = jwt.SigningMethodHS256

func runner(path, name string, exp int64) int {
	if path == "" {
		log.Println("Please specify the secret key file.")
		return 1
	}
	privateKey, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err.Error())
		return 1
	}

	token := jwt.New(sm)
	claims := token.Claims.(jwt.MapClaims)
	// claims["admin"] = true
	// claims["sub"] = ""
	claims["name"] = name
	claims["exp"] = time.Now().Unix() + exp

	tokenString, _ := token.SignedString(privateKey)

	fmt.Println(tokenString)
	return 0
}

var (
	secret string
	name   string
	exp    int64
)

func init() {
	flag.StringVar(&secret, "s", "", "secret key file")
	flag.StringVar(&name, "n", "trrrrrys", "username")
	flag.Int64Var(&exp, "e", 3600, "expiration second")
	flag.Usage = func() {
		fmt.Printf("Usage: %v -s filename [-e second] [-n user_name]  \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	os.Exit(runner(secret, name, exp))
}
