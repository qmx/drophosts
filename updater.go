package main

import "strings"

const MARKER = "## drophosts ##"

func UpdateHosts(original string, newContent string) string {
	parts := strings.Split(original, MARKER)
	if len(parts) == 3 {
		return parts[0] + newContent
	}
	return original + "\n" + newContent
}
