/*ASSIGNMENT 1
SET C
Q. 4 WAP in go language to check whether first string is
substring of another string or not*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	originalString := "This is our original string"
	sub_str := "Tour"
	if strings.Contains(originalString, sub_str) {
		fmt.Printf("%s is a substring of %s", sub_str, originalString)
	} else {
		fmt.Printf("%s is not a a substring of %s", sub_str, originalString)
	}
}
