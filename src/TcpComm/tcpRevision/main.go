package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type crewMember struct {
	Id                int    `bson:"id"`
	Name              string `bson:"name"`
	SecurityClearance int    `bson:"SecurityClearance"`
	Position          string `bson:"Positiion"`
}

func main() {

	t := flag.String("type", "s", "Enter S or C for server or client")

	flag.Parse()
	switch strings.ToLower(*t) {
	case "s":
		ServerConnect()
	case "c":
		ClientConnect()
	default:
		fmt.Println("Please enter a valid value for type")
	}

	return
}

func ServerConnect() error {
	lConn, erL := net.Listen("tcp", ":9001")

	if erL != nil {
		log.Fatal(erL)
	}

	defer lConn.Close()

	for {
		sess, err := mgo.Dial("localhost")
		if err != nil {
			return fmt.Errorf("Not able to establish connection")
		}
		AConn, erlA := lConn.Accept()
		if erlA != nil {
			return fmt.Errorf("Could not accept connection to port 9000")
		}

		go handleConn(AConn, sess)

	}
}

func handleConn(c net.Conn, s *mgo.Session) error {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	// scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)

		if len(fields) == 0 {
			return fmt.Errorf("Please enter the type of operation to be performed")
		}

		switch fields[0] {
		case "GET":
			cm := new(crewMember)
			s := s.DB("hydra").C("hydra")
			v, _ := strconv.Atoi(fields[1])
			s.Find(bson.M{"id": v}).One(cm)
			fmt.Println(cm)
		case "SET":
			s := s.DB("hydra").C("hydra")
			sc, _ := strconv.Atoi(fields[3])
			id, _ := strconv.Atoi(fields[1])
			// fmt.Printf("%d, %d", sc, id)
			s.Insert(bson.M{"id": id, "name": fields[2], "SecurityClearance": sc, "Positiion": fields[4]})
			fmt.Println("Inserted item with id", fields[1])
		case "DEL":
			s := s.DB("hydra").C("hydra")
			s.RemoveId(bson.M{"id": fields[1]})
			fmt.Println("Removed item with id", fields[1])
		default:
			return fmt.Errorf("Enter valid selection value")
		}
	}
	return nil
}

func ClientConnect() {
	conn, err := net.Dial("tcp", ":9001")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write(append(scanner.Bytes(), '\n'))
	}
	return
}
