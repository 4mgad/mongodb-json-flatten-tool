package main

import (
	"bufio"
	"fmt"
	"log"
	"main/util"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jsonStr := ""
	for scanner.Scan() {
		jsonStr += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(util.Flatten(jsonStr))
}
