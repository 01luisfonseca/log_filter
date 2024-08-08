package logic

import (
	"log"
	"regexp"
	"strings"
)

func EvaluateStringInLine(line *string, keywords *string) bool {
	return strings.Contains(*line, *keywords)
}

func EvaluateRegexInLine(line *string, regex *string) bool {
	evaluation, err := regexp.MatchString(*regex, *line)
	if err != nil {
		log.Println("Error evaluating regex")
		log.Printf("Error: %s", err)
		return false
	}
	return evaluation
}
