package main

import (
	"fmt"
	"net"
)

func main() {
	testCreate()
	testInsertBag()
	testSelectBag()
}

func testCreate() {
	statements := []string{
		"create bag person (name string, age int)",
		"create bag animals (name string)",
		"create table animals (name string)",
		"create table animals (name string)",
		"create function add (a int, b int) int",
		"create user developer",
		"create bag person (name string, age int)",
		"create user dba",
		"create user admin",
	}
	testStatements(statements)
}

func testInsertBag() {
	result := execStatement("insert into bag person (name, age) values (\"John\", 30)")
	fmt.Println(result)
}

func testSelectBag() {
	statements := []string{
		"select * from person where age = 20",
		"select * from person where age = 30",
		"select * from person where age = 40",
		"select * from role",
	}
	testStatements(statements)
}

func testStatements(statements []string) {
	c := make(chan string)
	for _, statement := range statements {
		go testStatement(statement, c)
	}
	for i := 0; i < len(statements); i++ {
		fmt.Println(<-c)
	}
}

func testStatement(statement string, c chan string) {
	result := execStatement(statement)
	c <- result
}

func execStatement(statement string) string {
	conn, err := net.Dial("tcp", "localhost:9000")
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
