package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func main() {
	var duration int
	var secret string
	var issuer string
	claims := jwt.MapClaims{}

	flag.Func("kv", "key/value to add: ie foo=bar (can be used multiple times)", func(s string) error {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) != 2 {
			return errors.New("expected format foo=bar")
		}
		claims[kv[0]] = kv[1]
		return nil
	})
	flag.StringVar(&secret, "secret", "", "encoding secret string (mandatory)")
	flag.StringVar(&issuer, "issuer", "", "token issuer (mandatory)")
	flag.IntVar(&duration, "years", 10, "number of years before expiration")
	flag.Parse()

	if issuer == "" && secret == "" {
		fmt.Println("missing flags -issuer and -secret")
		flag.Usage()
		os.Exit(1)
	} else if issuer == "" {
		fmt.Println("missing flag -issuer")
		flag.Usage()
		os.Exit(1)
	} else if secret == "" {
		fmt.Println("missing flag -secret")
		flag.Usage()
		os.Exit(1)
	} else if len(claims) == 0 {
		fmt.Println("please provide at least one -kv flag")
		flag.Usage()
		os.Exit(1)
	}

	jwtKey := []byte(secret)

	now := time.Now()
	claims["iss"] = issuer
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour * 24 * 365 * time.Duration(duration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", tokenString)
}
