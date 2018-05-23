package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hex, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println("error decoding")
	}
	fmt.Println(hex)

	output := base64.StdEncoding.EncodeToString(hex)
	fmt.Println(output)
}
