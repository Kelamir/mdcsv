package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Remove HTML tags using regular expressions
func stripHTMLTags(text string) string {
	re := regexp.MustCompile("<[^>]*>")
	return re.ReplaceAllString(text, "")
}

// Remove Markdown formatting
func stripMarkdown(text string) string {
	// Remove bold formatting (e.g., **bold**)
	text = strings.ReplaceAll(text, "**", "")

	// Remove italic formatting (e.g., *italic*)
	text = strings.ReplaceAll(text, "*", "")

	// Remove headers (e.g., ## Header)
	re := regexp.MustCompile(`(?m)^#+\s*`)
	text = re.ReplaceAllString(text, "")

	return text
}

func main() {
	// Open the input CSV file
	csvFile, err := os.Open("output.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create the output CSV file
	outputFile, err := os.Create("outputfinal.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)

	// Read and write CSV records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			return
		}

		// Strip HTML and Markdown formatting from the second column
		record[1] = stripHTMLTags(record[1])
		record[1] = stripMarkdown(record[1])

		// Write the modified record to the output CSV file
		err = writer.Write(record)
		if err != nil {
			fmt.Println("Error writing CSV record:", err)
			return
		}
	}

	// Flush any buffered content to the output file
	writer.Flush()

	fmt.Println("CSV processing completed.")
}
