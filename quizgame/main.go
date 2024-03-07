package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func playIt(csvFileName string, goodA, badA *int) {
	scanner := bufio.NewScanner(os.Stdin)
	fileContent, err := os.ReadFile(csvFileName)

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(bytes.NewReader(fileContent))

	for i := 1; ; i++ {

		quiz, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Problem #%v: %v = ", i, quiz[0])
		scanner.Scan()
		text := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		if text == quiz[1] {
			*goodA++
		} else {
			*badA++
		}
	}
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "Set csv file")
	settedTime := flag.Int("time", 60*60*60, "Set time to finish (in seconds)")

	flag.Parse()

	finishTime := time.After(time.Duration(*settedTime) * time.Second)

	goodA, badA := 0, 0
	go playIt(*csvFileName, &goodA, &badA)

	<-finishTime

	fmt.Printf("Good Answers: %v \n Bad Answers: %v \n", goodA, badA)

}
