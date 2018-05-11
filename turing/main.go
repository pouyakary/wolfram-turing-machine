package main

// ────────────────────────────────────────────────────────────────────────────────

import (
	"fmt"
	"strings"
)

// ────────────────────────────────────────────────────────────────────────────────

type rule struct {
	left  string
	right string
}

// ────────────────────────────────────────────────────────────────────────────────

func getRules() []rule {
	return []rule{
		{"-7", "62"},
		{"17", "72"},
		{"27", "82"},
		{"-8", "61"},
		{"18", "71"},
		{"28", "81"},
		{"6-", "23"},
		{"61", "24"},
		{"62", "25"},
		{"4-", "-6"},
		{"41", "-7"},
		{"42", "-8"},
		{"5-", "13"},
		{"51", "14"},
		{"52", "15"},
		{"-3", "61"},
		{"13", "71"},
		{"23", "81"},
	}
}

// ────────────────────────────────────────────────────────────────────────────────

func evaluate(code string) string {
	rules := getRules()
	for true {
		newCode := ""
		fmt.Println("   " + code)

		for _, rule := range rules {
			if newCode == "" {
				newCode = strings.Replace(code, rule.left, rule.right, -1)
			} else {
				newCode = strings.Replace(newCode, rule.left, rule.right, -1)
			}
		}
		if newCode == code {
			return code
		}

		code = newCode
	}
	return ""
}

// ────────────────────────────────────────────────────────────────────────────────

func main() {

	fmt.Println("\n    ✣ The most simple Turing Machine ✣\n")

	evaluate("------------------6-----------------")
}

// ────────────────────────────────────────────────────────────────────────────────
