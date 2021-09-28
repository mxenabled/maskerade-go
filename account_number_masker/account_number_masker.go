package account_number_masker

import (
	"fmt"
	"github.com/mxenabled/maskerade-go/helpers"
	"regexp"
	"strings"
)

var allowedNumericPatterns = []*regexp.Regexp{
	regexp.MustCompile(`\b[[:digit:]]-[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}\b`),
	regexp.MustCompile(`\b[[:digit:]]{4}-[[:digit:]]{3}-[[:digit:]]{4}\b`),
	regexp.MustCompile(`\b[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}\b`),
	regexp.MustCompile(`\b[[:digit:]]{4}-[[:digit:]]{7}\b`),
	regexp.MustCompile(`\b[[:digit:]]{3}-[[:digit:]]{7}\b`),
	regexp.MustCompile(`\b[[:digit:]]{3}-[[:digit:]]{4}\b`),
}

// var allowedDescriptionPattern = regexp.MustCompile(`(?i)(CHECK|DRAFT|SHARE DRAFT|DFT DEBIT)`)
// var allowedDescriptionPattern = regexp.MustCompile(`(?i)(CHECK|DRAFT|SHARE DRAFT|DFT DEBIT)[[[:space:]]]*[#]*[[[:space:]]]*[[[:digit:]]]{1,}`)
var allowedDescriptionPattern = regexp.MustCompile(`(CHECK|DRAFT|SHARE DRAFT|DFT DEBIT)(\s*[#]*\s*\d{1,})`)

var sevenDigitPattern = regexp.MustCompile(`([-|\d]){7,}`)

// var replaceAllDigitsPattern = regexp.MustCompile(`[[:digit:]](?=[-[[:digit:]]]{4})`)
var replaceAllDigitsPattern = regexp.MustCompile(`([-|\d]){7,}`)

type AccountNumberMasker struct {
	Parameters
}

type Parameters struct {
	ReplacementText  string
	ReplacementToken string
	ExposeLast       int
}

// NewAccountNumberMasker will return a AccountNumberMasker and ensure a ReplacementToken is set
func NewAccountNumberMasker(parameters Parameters) *AccountNumberMasker {
	if parameters.ReplacementToken == "" {
		parameters.ReplacementToken = "X"
	}

	return &AccountNumberMasker{
		parameters,
	}
}

func (a *AccountNumberMasker) Mask(desc string) string {
	if desc == "" {
		return ""
	}

	matches := sevenDigitPattern.FindAllStringIndex(desc, -1)
	fmt.Println("matches count", matches)

	matchedStrings := []string{}
	for _, match := range matches {
		matchedStrings = append(matchedStrings, desc[match[0]:match[1]])
	}

	for _, matchedNumber := range matchedStrings {
		fmt.Println("matches", matchedNumber, allowedDescriptionPattern.FindString(desc))

		if allowedDescriptionPattern.FindString(desc) != "" {
			continue
		}

		found := false
		for _, allowedPattern := range allowedNumericPatterns {
			if allowedPattern.FindString(matchedNumber) != "" {
				found = true
				fmt.Println("found in allowed strings")
			}
		}

		if !found {
			desc = strings.ReplaceAll(desc, matchedNumber, maskAllButLastFour(matchedNumber, a.ReplacementToken))

			fmt.Println("desc", desc, "not found in allowedNumericPatterns")
		}
	}

	return desc
}

var matchNumbers = regexp.MustCompile(`[\d]`)

func maskAllButLastFour(value string, replacementToken string) string {
	exposeLast := 4

	allNumberIndices := matchNumbers.FindAllStringIndex(value, -1)
	numbersToKeep := len(allNumberIndices) - exposeLast

	indicesToMask := []int{}

	for _, r := range allNumberIndices {
		if len(indicesToMask) < numbersToKeep {
			indicesToMask = append(indicesToMask, r[0])
		}
	}

	return helpers.ReplaceCharacterAtIndexInString(value, replacementToken, indicesToMask)
}
