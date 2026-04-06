package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/miekg/dns"
)

var timezones = map[string]string{
	"london":     "Europe/London",
	"paris":      "Europe/Paris",
	"tokyo":      "Asia/Tokyo",
	"india":      "Asia/Kolkata",
	"california": "America/Los_Angeles",
	"new york":   "America/New_York",
	"dubai":      "Asia/Dubai",
	"sydney":     "Australia/Sydney",
	"singapore":  "Asia/Singapore",
}

func getTime(location string) string {

	if tz, ok := timezones[location]; ok {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return "error: bad timezone"
		}
		return time.Now().In(loc).Format("15:04:05 MST, Mon 02 Jan 2006")
	}
	return "error: location not supported"

}

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
	dns.HandleFunc(".", handleQuery)

	server := &dns.Server{
		Addr: ":8053",
		Net:  "udp",
	}

	fmt.Println("DNS server running on :8053")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error:", err)
	}
}
