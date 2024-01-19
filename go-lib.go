/**
 * Author: Mitch Allen
 * File: go-lib.go
 */

package lib

import "strings"

// Returns the sum of two numbers
func Add(a int, b int) int {
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	return a - b
}

func Contains(lists string, fileEx string) bool {
	//  split string to array
	list := strings.Split(lists, ",")
	// check file extension
	for _, v := range list {
		if v == fileEx {
			return true
		}
	}
	return false

}
