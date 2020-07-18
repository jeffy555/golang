package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "Hello\n. Jefferson Immanuel"
	//New scanner has a pointer to a *scanner, it has method Text()
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		//Scan() advances the scanner to the next token, which will then be available through bytes or text method
		line := scanner.Text()
		fmt.Println(line)
	}
}
