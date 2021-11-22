package repository

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ConvJson() (clients []Client, err error) {
	file, err := os.Open("D:\\Go\\src\\uaAlert\\config.txt")
	if err != nil {
		log.Fatalf("Can't open file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// var clients []Client
	for scanner.Scan() {
		s := scanner.Text()
		sp := strings.Split(s, ",")
		con := Client{
			LogFile:    sp[0],
			ClientName: sp[1],
		}
		// fmt.Println(sp[0])
		clients = append(clients, con)
	}

	// j, err := json.Marshal(clients)
	// if err != nil {
	// 	log.Fatalf("cannot marshal: %v", err)
	// }
	return clients, nil
}
