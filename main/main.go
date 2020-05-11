// https://www.hackerrank.com/challenges/contacts/problem

package main

import (
	"strings"
)

func contacts(queries [][]string) []int32 {
	var contactMap = make(map[rune][]string)
	var result = make([]int32, 0)

	for i := 0; i < len(queries); i++ {
		var command = queries[i][0]
		var param = queries[i][1]

		switch command {
		case "add":
			contacts, exist := contactMap[rune(param[0])]

			if exist {
				// Append the new contact to existing contact list based on the first character of the name
				contactMap[rune(param[0])] = append(contacts, param)
			} else {
				// Create a new entry of contact list based on the first character of the name
				contactMap[rune(param[0])] = []string{param}
			}

		case "find":
			// Find the contact list based on the first character of the name
			contacts, exist := contactMap[rune(param[0])]
			var count int32

			if exist {
				// if contact list found, only then search the list
				for j := 0; j < len(contacts); j++ {
					if strings.HasPrefix(contacts[j], param) {
						count++
					}
				}
			}

			result = append(result, count)
		}
	}

	return result
}

func main() {

}
