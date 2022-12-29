package main

import (
	buffer "bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter URL:")

	url := buffer.NewReader(os.Stdin)
	line, _ := url.ReadString('\n')
	param := strings.Split(line, "?")[1]
	params := strings.Split(param, "&")

	print("\n")

	for _, param := range params {
		percentSplit := strings.Split(param, "%")

		var strParam string
		if len(percentSplit) > 1 {
			for i, j := range percentSplit {
				if i == 0 {
					strParam += j
				} else {
					bl, _ := hex.DecodeString(j[:2])
					string := string(bl)
					strParam += string
					strParam += j[2:]
				}
			}
		} else {
			strParam = param
		}
		fmt.Println(strParam)
	}
}
