package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BadgeType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const metadataTemplate = `{
    "description": "A symbol of allegiance to the Berachain ecosystem.",
    "external_url": "https://www.faucet.0xhoneyjar.xyz/",
    "image": "ipfs://QmU2CYcj82dajWhyR5pKh4VbkX9VAYoNVoLa9PDTF2zUEB/%d.png",
    "name": "%s",
    "attributes": []
}`

func main() {
	jsonData, err := readFile("badges.json")
	if err != nil {
		panic(err)
	}

	var badgeTypes []BadgeType
	err = json.Unmarshal(jsonData, &badgeTypes)
	if err != nil {
		panic(err)
	}

	outputDir := "generated_jsons"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		panic(err)
	}

	// generate a new json for each badge type
	for _, badgeType := range badgeTypes {

		finalJson := fmt.Sprintf(metadataTemplate, badgeType.ID, badgeType.Name)

		// Don't suffix filename with .json to make file names easier in contract
		outputFilename := fmt.Sprintf("%s/%d", outputDir, badgeType.ID)

		err = writeFile(outputFilename, []byte(finalJson))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Generated JSON files in", outputDir)
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return data, nil
}

func writeFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %w", filePath, err)
	}
	return nil
}
