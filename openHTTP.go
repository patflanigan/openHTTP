package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/toqueteos/webbrowser"
	"os"
	"time"
)

func main() {

	wordPort := flag.String("port", "80", "Enter a port number")
	wordHTTP := flag.String("web", "http", "options http or https")
	wordFile := flag.String("file", "ips.txt", "enter file name")

	flag.Parse()
	fmt.Println("\nopenHTTP v1.0                Pat Flanigan")
	fmt.Println("=========================================\n")
	fmt.Println("-web [options http or https]")
	fmt.Println("-port Enter a port number [default 80]")
    fmt.Println("-file <<file name with ip address>> [default ips.txt]")

	fmt.Println("Example: go run openHTTP.go -web http -port 80 -file ips.txt\n")
	port := *wordPort
	http := *wordHTTP
	ips := *wordFile

	// open the file filepath
	fileHandle, _ := os.Open(ips)
	defer fileHandle.Close()

	// create a scanner
	fs := bufio.NewScanner(fileHandle)

	counter := 1
	for fs.Scan() {
		txt := http + "://" + fs.Text() + ":" + port
		fmt.Println(txt)
		webbrowser.Open(txt)
		time.Sleep(time.Second * 1)

		counter++
		check := counter % 10

		if check == 0 {
			fmt.Print("Press 'Enter' to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

		}

	}
}
