package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Hardware Intrusion Monitor ----\n\n")
	ip, err := externalIP()
	if err != nil {
		ip = "Unknown"
	}
	fmt.Fprintf(w, "[INFO] No hardware intrusion found!\n\n")

	fmt.Fprintf(w, "Server from IP : %s\n", ip)

}

func main() {
	http.HandleFunc("/", handler)
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
