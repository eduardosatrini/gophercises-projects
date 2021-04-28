package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func ConfigFlag() (CSVFileName *string) {

	CSVFileName = flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	return
}

func ExitMsg(m string) {

	fmt.Println(m)
	os.Exit(1)
}

func OpenCSVFile(fileName string) [][]string {

	file, err := os.Open(fileName)
	if err != nil {
		ExitMsg("Failed to open the CSV file.")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		ExitMsg("Failed to parse the provided CSV file.")
	}

	return lines
}
