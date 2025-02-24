package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func printHelp() {
	fmt.Println("Usage: static-server [options]")
	fmt.Println("Options:")
	fmt.Println("  --ip string      IP address to bind the server (default \"0.0.0.0\")")
	fmt.Println("  --port string    Port to bind the server (default \"8080\")")
	fmt.Println("  --directory string  Directory to serve files from (default \"./\")")
	fmt.Println("  --help          Show this help message")
}

func main() {
	ip := flag.String("ip", "0.0.0.0", "IP static server")
	port := flag.String("port", "8080", "Port static server")
	directory := flag.String("directory", "./", "Directory static server")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	fmt.Printf("Starting server at %s:%s, serving files from %s\n", *ip, *port, *directory)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*directory))))
	err := http.ListenAndServe(*ip+":"+*port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}