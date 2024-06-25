package name

import (
	"fmt"
	"strings"
)

// Upper returns the uppercase of the given string argument.
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower returns the lowercase of the given string argument.
func Lower(s string) string {
	return strings.ToLower(s)
}

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func myStringIntBoolPrint[printThis string | int | bool](what printThis) {
	fmt.Println(what)
}

func myGenericPrint[printThis any](what printThis) {
	fmt.Println(what)
}
