package operations

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filePath string) ([][]string){
	// filePath := `dataSeeding/data/movies.csv`
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	fmt.Println("............CSV read successfully............")

	return records
}
