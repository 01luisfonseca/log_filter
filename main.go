package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"log_filter/logic"
	"os"
	"path/filepath"
)

type Flags struct {
	inputFile        string
	outputFileName   string
	filterType       string
	filterDefinition string
}

func main() {
	log.Println("Starting log filter")
	flags, err := flagCheck()
	if err != nil {
		log.Println("Error parsing flags")
		log.Fatal(err)
		return
	}

	file, err := os.Open(flags.inputFile)
	if err != nil {
		log.Println("Error opening file")
		log.Fatal(err)
		return
	}
	defer file.Close()

	inputFilePath, err := filepath.Abs(flags.inputFile)
	if err != nil {
		log.Println("Error getting absolute path")
		log.Fatal(err)
		return
	}

	inputDir := filepath.Dir(inputFilePath)
	outputFilePath := filepath.Join(inputDir, flags.outputFileName)

	if _, err := os.Stat(outputFilePath); err == nil {
		err = os.Remove(outputFilePath)
		if err != nil {
			log.Println("Error removing output file")
			log.Fatal(err)
			return
		}
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("Error creating output file")
		log.Fatal(err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	linesFound := 0
	linesProcessed := 0
	for scanner.Scan() {
		line := scanner.Text()
		linesProcessed++
		found := lineCheck(&line, &flags.filterDefinition, &flags.filterType)
		if found {
			linesFound++
			lineSave(&line, outputFile)
		}
	}
	log.Println("Finished filtering. Lines found:", linesFound, "Lines processed:", linesProcessed)
}

func flagCheck() (Flags, error) {
	flags := Flags{}
	flag.StringVar(&flags.inputFile, "inputFile", "", "Input file to filter")
	flag.StringVar(&flags.outputFileName, "outputFileName", "log_filtered.log", "Output file name")
	flag.StringVar(&flags.filterType, "filterType", "string", "Filter type (string or regex)")
	flag.StringVar(&flags.filterDefinition, "filterDefinition", "Apollo", "Filter definition")
	flag.Parse()
	if flags.inputFile == "" {
		return flags, errors.New("input file was not provided in flag parameters")
	}
	return flags, nil
}

func lineCheck(line *string, filter *string, filter_type *string) bool {
	switch *filter_type {
	case "string":
		return logic.EvaluateStringInLine(line, filter)
	case "regex":
		return logic.EvaluateRegexInLine(line, filter)
	default:
		return false
	}
}

func lineSave(line *string, outputFile *os.File) bool {
	_, err := outputFile.WriteString(*line + "\n")
	if err != nil {
		log.Println("Error writing to output file")
		log.Printf("Error: %s", err)
		return false
	}
	return true
}
