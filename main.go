package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func stripHTMLTags(text string) string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(text, "")
}

func stripMarkdown(text string) string {
	text = strings.ReplaceAll(text, "**", "")
	text = strings.ReplaceAll(text, "*", "")
	re := regexp.MustCompile(`(?m)^#+\s*`)
	text = re.ReplaceAllString(text, "")

	return text
}

func main() {
	csvFile, err := os.Open("output.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	outputFile, err := os.Create("outputfinal.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	reader := csv.NewReader(csvFile)

	writer := csv.NewWriter(outputFile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			return
		}

		record[1] = stripHTMLTags(record[1])
		record[1] = stripMarkdown(record[1])

		err = writer.Write(record)
		if err != nil {
			fmt.Println("Error writing CSV record:", err)
			return
		}
	}

	writer.Flush()

	fmt.Println("CSV processing completed.")
}
