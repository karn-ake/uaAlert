package repository

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ConvJson() (clients []Client, err error) {
	//Open config file location
	file, err := os.Open("D:\\Go\\src\\uaAlert\\config.txt")
	if err != nil {
		log.Fatalf("Can't open file %v", err)
	}
	defer file.Close()

	//Read config file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		sp := strings.Split(s, ",")
		con := Client{
			LogFile:    sp[0],
			ClientName: sp[1],
		}
		clients = append(clients, con)
	}
	return clients, nil
}
