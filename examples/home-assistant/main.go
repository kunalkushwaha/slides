package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Home Assistant System----\n\n")
	ip, err := externalIP()
	if err != nil {
		ip = "Unknown"
	}
	fmt.Fprintf(w, "Good Evening \"Kunal\" !\n")
	fmt.Fprintf(w, "How Can I help you?\n\n")

	fmt.Fprintf(w, "Server from IP : %s\n", ip)
}

func today(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Home Assistant System----\n\n")
	ip, err := externalIP()
	if err != nil {
		ip = "Unknown"
	}
	fmt.Fprintf(w, "Today's events : 8/12/2016\n")
	fmt.Fprintf(w, "Rancher Meetup @ Recruit, Nihombashi - 6:30 PM\n\n")
	fmt.Fprintf(w, "-----------\n")
	fmt.Fprintf(w, "[INFO] It will be Cold in evening, Get something to wear warm!\n\n")

	fmt.Fprintf(w, "Server from IP : %s\n", ip)
}

func tomorrow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Home Assistant System----\n\n")
	ip, err := externalIP()
	if err != nil {
		ip = "Unknown"
	}
	fmt.Fprintf(w, "Tomorrow's events : 9/12/2016\n")
	fmt.Fprintf(w, "[No events today] - Have a chill day!\n\n")
	fmt.Fprintf(w, "-----------\n")
	fmt.Fprintf(w, "[INFO] It will be Cold in evening, Get something to wear warm!\n\n")

	fmt.Fprintf(w, "Server from IP : %s\n", ip)
}

func todo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Home Assistant System----\n\n")
	ip, err := externalIP()
	if err != nil {
		ip = "Unknown"
	}
	fmt.Fprintf(w, "TODO List\n")
	fmt.Fprintf(w, "1. Fix Bingo-Bus toy.\n")
	fmt.Fprintf(w, "2. Get replaced water Filter\n")
	fmt.Fprintf(w, "3. Book ticket for holidays\n")
	fmt.Fprintf(w, "-----------\n\n")

	fmt.Fprintf(w, "Server from IP : %s\n", ip)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/today", today)
	http.HandleFunc("/tomorrow", tomorrow)
	http.HandleFunc("/todo", todo)
	fmt.Println("Server is listening at : 8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
