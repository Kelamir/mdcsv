package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Open CSV writer
	csvFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(csvFile)

	defer csvWriter.Flush()
	defer csvFile.Close()

	scanner := bufio.NewScanner(file)
	var question string
	var answer strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// check if the line contains "###"
		if strings.Contains(line, "###") {
			if question != "" {
				// Write the previous question and answer to the csv file
				csvWriter.Write([]string{question, answer.String()})
				answer.Reset()
			}
			// Get the question from the line, without the "###"
			question = strings.TrimSpace(strings.Split(line, "###")[1])
		} else {
			// Add the line to the current answer
			answer.WriteString(line + "\n")
		}
	}

	// Check for any scanner error
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write the last question and answer to the csv file
	if question != "" {
		csvWriter.Write([]string{question, answer.String()})
	}

	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}
}
