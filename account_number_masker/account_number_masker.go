package account_number_masker

import (
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

var allowedDescriptionPattern = regexp.MustCompile(`(CHECK|DRAFT|SHARE DRAFT|DFT DEBIT)(\s*[#]*\s*\d{1,})`)

var sevenDigitPattern = regexp.MustCompile(`([-|\d]){7,}`)

type AccountNumberMasker struct {
	Parameters
}

type Parameters struct {
	ReplacementToken string
	ExposeLast       int
}

// NewAccountNumberMasker will return a AccountNumberMasker and ensure a ReplacementToken is set
func NewAccountNumberMasker(parameters Parameters) *AccountNumberMasker {
	if parameters.ReplacementToken == "" || len(parameters.ReplacementToken) > 1 {
		parameters.ReplacementToken = "X"
	}

	if parameters.ExposeLast < 0 {
		parameters.ExposeLast = 0
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

	matchedStrings := []string{}
	for _, match := range matches {
		matchedStrings = append(matchedStrings, desc[match[0]:match[1]])
	}

	for _, matchedNumber := range matchedStrings {
		if allowedDescriptionPattern.FindString(desc) != "" {
			continue
		}

		found := false
		for _, allowedPattern := range allowedNumericPatterns {
			if allowedPattern.FindString(matchedNumber) != "" {
				found = true
			}
		}

		if !found {
			desc = strings.ReplaceAll(desc, matchedNumber, maskAllButLastFour(matchedNumber, a.ReplacementToken, a.ExposeLast))
		}
	}

	return desc
}

var matchNumbers = regexp.MustCompile(`[\d]`)

func maskAllButLastFour(value string, replacementToken string, exposeLast int) string {
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
