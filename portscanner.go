package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

func worker(ports, results chan int) {
	for p := range ports {
		connec, error := net.Dial("tcp", os.Args[3] + ":" + strconv.Itoa(p))
		if error != nil {
			results <- 0
			continue
		}
		connec.Close()
		results <- p
	}

}

func main() {
	if(len(os.Args) != 4) {
		fmt.Println("Please give all argument")
		os.Exit(1)
	}
	// In CLI $go run port.go numberofclient protrange website(without -> https://) 
	// example go run portscanner.go 120 2:250 google.com 
	numberofclients, _ := strconv.Atoi(os.Args[1])
	portrange := strings.Split(os.Args[2], ":")
	startport, _ := strconv.Atoi(portrange[0])
	endport, _ := strconv.Atoi(portrange[1])

	ports := make(chan int, numberofclients)
	results := make(chan int)

	var openport []int

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func () {
		for i := startport; i <= endport; i++ {
			ports <- i 
		}
	}()

	for i := startport; i <= endport; i++ {
		port := <- results
		if port != 0 {
			fmt.Printf("%d port is open\n", port)
			openport = append(openport, port)
		}

		
	}

	close(ports)
	close(results)

	sort.Ints(openport)
	fmt.Printf("List of openport -> ")
	fmt.Println(openport)
}