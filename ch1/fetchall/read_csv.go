package main

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func getRecords() []string {
	// Open the file
	csvfile, err := os.Open("/Users/leonardodalimasilva/go/src/book/src/gopl.io/ch1/fetchall/majestic_million.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	d := []string{"https://pudim.com.br"}
	// d := make([]string{"pudim.com.br"})

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if record[0] == "26629" {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		url := "https://" + record[2]
		d = append(d, url)
		// d = d + "," + record[2]

		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

	// fmt.Println(d)
	return d
}
