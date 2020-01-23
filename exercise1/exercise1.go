package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fpath := flag.String("csv", "questions.csv", "The path to the CVS file")
	limit := flag.Int("limit", 30, "Time for the quiz in seconds (default 30)")
	flag.Parse()
	content, err := ioutil.ReadFile(*fpath)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(content)))
	i, score := 1, 0

	go func(limit int) {
		timer := time.NewTimer(time.Duration(limit) * time.Second)
		<-timer.C
		fmt.Println("\n\nTime's up!\n\nScore:", score)
		os.Exit(0)
	}(*limit)

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("\nScore:", score)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Problem#%d: "+record[0]+" = ? ", i)
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")
		if err == io.EOF {
			break
		}
		if answer != record[1] {
			fmt.Println("\nScore:", score)
			break
		}
		score++
		i++
	}
}
