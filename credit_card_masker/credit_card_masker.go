package maskerade

import (
	"github.com/joeljunstrom/go-luhn"
	"github.com/mxenabled/maskerade-go/ranges"
	"github.com/mxenabled/maskerade-go/helpers"
	"regexp"
	"strings"
)

var creditCardPatterns = []*regexp.Regexp{
	/* maestro */ regexp.MustCompile(`\b(?:5[06-8]|6\d)\d{2}[ -]?\d{4}[ -]?\d{4}[ -]?\d{0,7}\b`),
	/* visa */ regexp.MustCompile(`\b4\d{3}[ -]?\d{4}[ -]?\d{4}[ -]?(?:\d|\d{4}|\d{7})\b`),
	/* amex */ regexp.MustCompile(`\b3[47]\d{2}[ -]?\d{6}[ -]?\d{5}\b`),
	/* diners_club */ regexp.MustCompile(`\b3(?:0[0-5]|[68]\d)\d[ -]?\d{6}[ -]?(?:\d{4}|\d{9})\b`),
	/* mastercard */ regexp.MustCompile(`\b(?:5[1-5]\d{2}[ -]?\d{2}|6771[ -]?89|222[1-9][ -]?\d{2}|22[3-9]\d[ -]?\d{2}|2[3-6]\d{2}[ -]?\d{2}|27[01]\d[ -]?\d{2}|2720[ -]?\d{2})\d{2}[ -]?\d{4}[ -]?\d{4}\b`),
	/* discover */ regexp.MustCompile(`\b(?:6011|65\d{2}|64[4-9]\d)[ -]?\d{4}[ -]?\d{4}[ -]?(?:\d{4}|\d{7})\b`),
	/* union_pay */ regexp.MustCompile(`\b62\d{2}[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}\b`),
	/* jcb */ regexp.MustCompile(`\b(?:308[89]|309[0-4]|309[6-9]|310[0-2]|311[2-9]|3120|315[8-9]|333[7-9]|334[0-9]|352[89]|35[3-8]\d)[ -]?\d{4}[ -]?\d{4}[ -]?(?:\d{4}|\d{7})\b`),
	/* switch_short_iin */ regexp.MustCompile(`\b(?:4903|4905|4911|4936|6333|6759)[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}(?:\d{2,3})?\b`),
	/* switch_long_iin */ regexp.MustCompile(`\b(?:5641[ -]?82|6331[ -]?10)\d{2}[ -]?\d{4}[ -]?\d{4}(?:\d{2,3})?\b`),
	/* solo */ regexp.MustCompile(`\b(?:6334|6767)[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}(?:\d{2,3})?\b`),
	/* dankort */ regexp.MustCompile(`\b5019[ -]?\d{4}[ -]?\d{4}[ -]?\d{4}\b`),
	/* forbrugsforeningen */ regexp.MustCompile(`\b6007[ -]?22\d{2}[ -]?\d{4}[ -]?\d{4}\b`),
	/* laser */ regexp.MustCompile(`\b(?:6304|6706|6709|6771[^89])[ -]?\d{4}[ -]?\d{4}[ -]?(?:\d{4}|\d{6,7})?\b`),
}

type CreditCardMasker struct {
	Parameters
}

type Parameters struct {
	ReplacementText  string
	ReplacementToken string
	ExposeLast       int
}

// NewCreditCardMasker will return a CreditCardMasker and ensure a ReplacementToken is set
func NewCreditCardMasker(parameters Parameters) *CreditCardMasker {
	if parameters.ReplacementToken == "" {
		parameters.ReplacementToken = "X"
	}

	return &CreditCardMasker{
		parameters,
	}
}

func (c *CreditCardMasker) Mask(value string) string {
	matchedPatterns := tryAllPatterns(value, creditCardPatterns)

	originalValue := value
	for _, matchedPattern := range matchedPatterns {
		oldString := originalValue[matchedPattern.Start:matchedPattern.End]
		maskedString := maskOne(oldString, c)
		value = strings.ReplaceAll(value, oldString, maskedString)
	}
	return value
}

func tryAllPatterns(value string, creditCardPatterns []*regexp.Regexp) []ranges.MatchRange {
	allMatches := []ranges.MatchRange{}
	for _, pattern := range creditCardPatterns {
		matches := pattern.FindAllStringSubmatchIndex(value, -1)
		if len(matches) != 0 {
			for _, match := range matches {
				allMatches = append(allMatches, ranges.MatchRange{Start: match[0], End: match[1]})
			}
		}
	}

	contingousMatches := ranges.FindLargest(allMatches)
	if len(contingousMatches) == 0 {
		return nil
	} else {
		return contingousMatches
	}
}

var matchNonNumbers = regexp.MustCompile(`[^0-9]`)
var matchNumbers = regexp.MustCompile(`[0-9]`)

func maskOne(value string, options *CreditCardMasker) string {
	cleaned := matchNonNumbers.ReplaceAllString(value, "")

	if !luhn.Valid(cleaned) {
		return value
	}

	if options.ReplacementText != "" {
		return options.ReplacementText
	}

	allNumberIndices := matchNumbers.FindAllStringIndex(value, -1)
	numbersToKeep := len(allNumberIndices) - options.ExposeLast
	indicesToMask := []int{}

	for _, r := range allNumberIndices {
		if len(indicesToMask) < numbersToKeep {
			indicesToMask = append(indicesToMask, r[0])
		}
	}

	return replaceCharacterAtIndexInString(value, options.ReplacementToken, indicesToMask)
}
