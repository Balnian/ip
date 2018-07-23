package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	//DefaultListenPort default port number used
	DefaultListenPort = ":8080"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		//Get remote address, if theres a X-forwarded-for we take it over "req.RemoteAddr"
		remaddr, _, _ := net.SplitHostPort(req.RemoteAddr)
		if _, valid := req.Header["X-Forwarded-For"]; valid {
			remaddr = req.Header["X-Forwarded-For"][0]
		}

		w.Write([]byte(remaddr))
	})

	ListenPort := DefaultListenPort
	//Check for port on command line
	for i, v := range os.Args[1:] {
		if v == "-p" {
			if len(os.Args) > i+2 {
				ListenPort = os.Args[i+2]
			} else {
				log.Fatal("Missing argument after -p")
			}
		} else if v == "-h" || v == "--help" {
			printHelp()
			os.Exit(0)
		}
	}
	fmt.Println("Starting Server on port:", ListenPort)
	log.Fatal(http.ListenAndServe(ListenPort, nil))
}

func printHelp() {
	fmt.Println("USAGE")
	fmt.Println("	ip [OPTION]")
	fmt.Println("DESCRIPTION")
	fmt.Println("	Simple server returning the client Ip in plain text")
	fmt.Println("OPTION:")
	fmt.Println("	-p")
	fmt.Println("		Set the listen port. Default is :8080")
	fmt.Println("	-h, --help")
	fmt.Println("		Print this help page")
	fmt.Println("EXEMPLE")
	fmt.Println("	ip -p :2000 #Start the server on port 2000")
}
