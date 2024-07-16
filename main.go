package main

import (
	"fmt"
	"net"
)

func main() {
	testCreateBag()
	testInsertBag()
	testSelectBag()
}

func testCreateBag() {
	result := testStatement("create bag person (name string, age int)")
	fmt.Println(result)
}

func testInsertBag() {
	result := testStatement("insert into bag person (name, age) values (\"John\", 30)")
	fmt.Println(result)
}

func testSelectBag() {
	statements := []string{
		"select * from person where age = 20",
		"select * from person where age = 30",
		"select * from person where age = 40",
	}
	c := make(chan string)
	for _, statement := range statements {
		go testSelect(statement, c)
	}
	for i := 0; i < len(statements); i++ {
		fmt.Println(<-c)
	}
}

func testSelect(statement string, c chan string) {
	result := testStatement(statement)
	c <- result
}

func testStatement(statement string) string {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return "Error:" + string(err.Error())
	}
	defer conn.Close()
	data := []byte(statement)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error:" + string(err.Error())
	}
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		return "Error:" + string(err.Error())
	}
	return string(response[:n])
}
