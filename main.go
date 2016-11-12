package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	digits  = "0123456789"
	lowers  = "abcdefghijklmnopqrstuvwxyz"
	uppers  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols = "!\"#$%&'()*+,-./0123456789:;<=>?@[\\]^_`{|}~"
)

func Generate(charset []byte, length int) (string, error) {
	max := big.NewInt(int64(len(charset)))
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}
	return string(password), nil
}

var length = 8
var charset = [](byte)(lowers)

func main() {

	var disableLower, disableUpper, disableDigit, disableSymbol bool
	var length int
	flag.BoolVar(&disableDigit, "disable-digit", false, "Disable digits in password")
	flag.BoolVar(&disableLower, "disable-lower", false, "Disable lower cacses in password")
	flag.BoolVar(&disableUpper, "disable-upper", false, "Disable upper cacses in password")
	flag.BoolVar(&disableSymbol, "disable-symbol", false, "Disable symbols in password")
	flag.IntVar(&length, "length", 8, "Length of password (default is 8)")
	flag.Parse()

	var charset []byte

	if !disableDigit {
		charset = append(charset, []byte(digits)...)
	}
	if !disableLower {
		charset = append(charset, []byte(lowers)...)
	}
	if !disableUpper {
		charset = append(charset, []byte(uppers)...)
	}
	if !disableSymbol {
		charset = append(charset, []byte(symbols)...)
	}

	if len(charset) == 0 {
		fmt.Fprintln(os.Stderr, "The character set of the password is empty")
		os.Exit(1)
	}

	password, err := Generate(charset, length)
	if err != nil {
		panic(err)
	}
	fmt.Println(password)
}
