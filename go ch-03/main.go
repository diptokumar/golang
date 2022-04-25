package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	// var myString string;
	// // myString = "Hello, World!";
	// fmt.Scanln(&myString);
	// fmt.Println(myString);

	// reader:= bufio.NewReader(os.Stdin);

	// fmt.Println("Enter a Name: ");
	// name, _ := reader.ReadString('\n');
	// fmt.Printf(name);
	// reader:= bufio.NewReader(os.Stdin);
	// fmt.Println("Enter your reading ");
	// myReading, _ := reader.ReadString('\n');
	// fmt.Println(myReading);
	// fmt.Println("Enter your name ");
	// myName, _ := reader.ReadString('\n');
	// fmt.Println(myName)

	reader:= bufio.NewReader(os.Stdin);
	myRating, _ := reader.ReadString('\n');
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64);
	fmt.Println(myNumRating + 2);
} 