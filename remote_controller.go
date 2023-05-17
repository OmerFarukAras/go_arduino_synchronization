package main

import (
	"bufio"
	"github.com/go-vgo/robotgo"
	"github.com/tarm/goserial"
	"io"
	"log"
)

func main() {
	robotgo.MouseSleep = 100
	c := &serial.Config{Name: "COM7", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		text := scanner.Text()
		log.Println(text)
		switch text {
		case "FFC23D":
			log.Println("Play/Pause COMMAND")

			robotgo.Scroll(0, -10)
			robotgo.Scroll(100, 0)

			robotgo.MilliSleep(100)
			robotgo.ScrollSmooth(-10, 6)

			robotgo.Move(10, 20)
			WriteSerial(s, "Play/Pause SUCCESS.")
		case "FF02FD":
			log.Println("Skip COMMAND")
			WriteSerial(s, "Skip SUCCESS.")
		case "FF22DD":
			log.Println("Previous COMMAND")
			WriteSerial(s, "Previous SUCCESS.")
		case "FFA857":
			log.Println("Volume UP COMMAND")
			WriteSerial(s, "Volume UP SUCCESS.")
		case "FFE01F":
			log.Println("Volume DOWN COMMAND")
			WriteSerial(s, "Volume DOWN SUCCESS.")
		}
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func WriteSerial(s io.ReadWriteCloser, text string) {
	_, err := s.Write([]byte(text))
	if err != nil {
		return
	}
}
