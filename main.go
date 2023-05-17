package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/goserial"
	"log"
	"os"
	"strings"
)

func main() {
	c := &serial.Config{Name: "COM8", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		data := ""
		if strings.HasPrefix(text, "Card UID:") {
			data = strings.ReplaceAll(text, "Card UID: ", "")
			fmt.Println("UID DATA : " + data)
		} else if strings.HasPrefix(text, "PICC type:") {
			data = strings.ReplaceAll(text, "Card UID: ", "")
			fmt.Println("PICC DATA : " + data)
		} else if strings.HasPrefix(text, "#WAITING APPROVE") {
			//FUNCSIN HERE
			reader := bufio.NewReader(os.Stdin)
			_, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				return
			}
			_, err = s.Write([]byte("Transfer approved."))
			if err != nil {
				return
			}
		}
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}
