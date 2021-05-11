package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func checkErr(err error, msg string) {

	if err != nil {
		exitMsg(msg)
	}
}

func exitMsg(m string) {

	fmt.Println(m)
	os.Exit(1)
}

func ConfigFlag() (CSVFileName *string) {

	CSVFileName = flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	return
}

func OpenCSVFile(fileName string) [][]string {

	file, err := os.Open(fileName)
	checkErr(err, "Failed to open the CSV file.")

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	checkErr(err, "Failed to parse the provided CSV file.")

	return lines
}
