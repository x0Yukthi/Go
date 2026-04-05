package main

import (
	"fmt"
	"strings"
)

func parseQuery(domain string) (command, location string) {
	domain = strings.TrimSuffix(domain, ".")
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return "", ""
	}
	command = parts[0]
	strings.Join(parts[1:], " ")
	location = strings.Join(parts[1:], " ")

	if command != "time" && command != "weather" {
		return "", ""
	}

	if location == "" {
		return "", ""
	}

	return command, location
}

func main() {
	fmt.Println(parseQuery("time.london"))         // time london
	fmt.Println(parseQuery("weather.tokyo"))       // weather tokyo
	fmt.Println(parseQuery("time.new.york"))       // time new york
	fmt.Println(parseQuery("weather.los.angeles")) // weather los angeles
	fmt.Println(parseQuery("hello.world"))         // "" ""
	fmt.Println(parseQuery("time."))               // "" "" (no location)
}
