package maskerade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskWhenValueDoesNotContainACreditCardNumber(t *testing.T) {
	testCases := []string{
		"555555-555555-5555",
		"6449 000000 000000",
		"    1-800-555-5555    ",
		"1313131313131",
		"14141414141414",
		"starbucks store #555 1-800-555-5555",
		"business XXXX XXXX XXXX 1234 1234567890123456",
		"4111111111111112", // fails luhn test
		"4 1 3-11 93 3-1-2",
		"4        2      2",
		"677189---------3",
		"30888778502353", //wrong length; should be 16 not 14
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase)

		assert.Equal(t, testCase, got)
	}
}

func TestUTF8StringRebuilding(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "♡378282246310005♡", expected: "♡XXXXXXXXXXXXXXX♡"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got, "♡ should be present")
	}
}

func TestMaskWhenValueContainsAnAMEXCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "378282246310005", expected: "XXXXXXXXXXXXXXX"},
		{value: "371449635398431", expected: "XXXXXXXXXXXXXXX"},
		{value: "377154756108213", expected: "XXXXXXXXXXXXXXX"},
		{value: "378734493671000", expected: "XXXXXXXXXXXXXXX"},
		{value: "3771-547561-08213", expected: "XXXX-XXXXXX-XXXXX"},
		{value: "3771 547561-08213", expected: "XXXX XXXXXX-XXXXX"},
		{value: "3771 547561 08213", expected: "XXXX XXXXXX XXXXX"},
		{value: "3771 54756108213", expected: "XXXX XXXXXXXXXXX"},
		{value: "Starbucks Store #555 3771 547561 08213", expected: "Starbucks Store #555 XXXX XXXXXX XXXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsAnDiscoverCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "6011000990139424", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6011111111111117", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6011072668180642", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6011 0726 6818 0642", expected: "XXXX XXXX XXXX XXXX"},
		{value: "6011 0726 6818-0642", expected: "XXXX XXXX XXXX-XXXX"},
		{value: "6011-0726 6818 0642", expected: "XXXX-XXXX XXXX XXXX"},
		{value: "Starbucks Store #555 6011 0726 6818 0642", expected: "Starbucks Store #555 XXXX XXXX XXXX XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsADinersClubCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "30569309025904", expected: "XXXXXXXXXXXXXX"},
		{value: "38520000023237", expected: "XXXXXXXXXXXXXX"},
		{value: "5434731844539705", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5434 7318 4453 9705", expected: "XXXX XXXX XXXX XXXX"},
		{value: "5434-7318 4453-9705", expected: "XXXX-XXXX XXXX-XXXX"},
		{value: "5434 7318-4453 9705", expected: "XXXX XXXX-XXXX XXXX"},
		{value: "Starbucks Store #555 5434 7318 4453 9705", expected: "Starbucks Store #555 XXXX XXXX XXXX XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsAJCBCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "3566002020360505", expected: "XXXXXXXXXXXXXXXX"},
		{value: "3530111333300000", expected: "XXXXXXXXXXXXXXXX"},
		{value: "3088877850235318", expected: "XXXXXXXXXXXXXXXX"},
		{value: "3088 877850235318", expected: "XXXX XXXXXXXXXXXX"},
		{value: "3088-8778-5023-5318", expected: "XXXX-XXXX-XXXX-XXXX"},
		{value: "Starbucks Store #555 3088 8778 5023 5318", expected: "Starbucks Store #555 XXXX XXXX XXXX XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsAMasterCardCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "5555555555554444", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5157387747088202", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5105105105105100", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5105 1051 0510 5100", expected: "XXXX XXXX XXXX XXXX"},
		{value: "5105-1051-0510-5100", expected: "XXXX-XXXX-XXXX-XXXX"},
		{value: "5105 1051-0510 5100", expected: "XXXX XXXX-XXXX XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsAVisaCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "4222222222222220", expected: "XXXXXXXXXXXXXXXX"},
		{value: "4012888888881881", expected: "XXXXXXXXXXXXXXXX"},
		{value: "4111111111111111", expected: "XXXXXXXXXXXXXXXX"},
		{value: "4532026721946154", expected: "XXXXXXXXXXXXXXXX"},
		{value: "4532 0267 2194 6154", expected: "XXXX XXXX XXXX XXXX"},
		{value: "4532-0267-2194-6154", expected: "XXXX-XXXX-XXXX-XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsASwitchCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "5641826302446333", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6331-1005-9051-1593", expected: "XXXX-XXXX-XXXX-XXXX"},
		{value: "491192979470860192", expected: "XXXXXXXXXXXXXXXXXX"},   // 18-digit
		{value: "564182951770215274", expected: "XXXXXXXXXXXXXXXXXX"},   // 18-digit
		{value: "4903263854966737499", expected: "XXXXXXXXXXXXXXXXXXX"}, // 19-digit
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsASolohCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "6334850405102841", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6767 7753 0761 8444", expected: "XXXX XXXX XXXX XXXX"},
		{value: "6334-5957-4704-4538", expected: "XXXX-XXXX-XXXX-XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsADankortCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "5019600906451260", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5019 2174 5298 6407", expected: "XXXX XXXX XXXX XXXX"},
		{value: "5019-4930-7205-0938", expected: "XXXX-XXXX-XXXX-XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

// maestro CC # can be 12-19 digits
func TestMaskWhenValueContainsAMaestroCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "5094730275858403", expected: "XXXXXXXXXXXXXXXX"},
		{value: "5000-7429-6277", expected: "XXXX-XXXX-XXXX"},             // 12-digit
		{value: "6957-0438-4355-62", expected: "XXXX-XXXX-XXXX-XX"},       // 14-digit
		{value: "6848 4263 2113 25499", expected: "XXXX XXXX XXXX XXXXX"}, // 17-digit
		{value: "6878817132178558380", expected: "XXXXXXXXXXXXXXXXXXX"},   // 19-digit
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsAForbrugsforeningenCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "6007221289345722", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6007 2295 6187 0234", expected: "XXXX XXXX XXXX XXXX"},
		{value: "6007-2215-2343-5149", expected: "XXXX-XXXX-XXXX-XXXX"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsALaserCreditCardNumber(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "6709202628830898", expected: "XXXXXXXXXXXXXXXX"},
		{value: "6304-0626-5252-903482", expected: "XXXX-XXXX-XXXX-XXXXXX"},   // 18-digit
		{value: "6706 2312 0102 8896786", expected: "XXXX XXXX XXXX XXXXXXX"}, // 19-digit
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueHasAnUncommonLength(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "4222222222222", expected: "XXXXXXXXXXXXX"},                   // 13-digit visa (do these exist anymore?)
		{value: "4222-2222-2222-2", expected: "XXXX-XXXX-XXXX-X"},             // 13-digit visa
		{value: "4234567890123456782", expected: "XXXXXXXXXXXXXXXXXXX"},       // 19-digit visa (do these exist yet?)
		{value: "4234-5678-9012-3456782", expected: "XXXX-XXXX-XXXX-XXXXXXX"}, // 19-digit visa
		{value: "3607050000000000008", expected: "XXXXXXXXXXXXXXXXXXX"},       // 19-digit diners
		{value: "3607-050000-000000008", expected: "XXXX-XXXXXX-XXXXXXXXX"},   // 19-digit diners
		{value: "6544440044440046990", expected: "XXXXXXXXXXXXXXXXXXX"},       // 19-digit discover
		{value: "6544 4400 4444-0046990", expected: "XXXX XXXX XXXX-XXXXXXX"}, // 19-digit discover
		{value: "3569990010082211774", expected: "XXXXXXXXXXXXXXXXXXX"},       // 19-digit jcb
		{value: "3569 9900 1008 2211774", expected: "XXXX XXXX XXXX XXXXXXX"}, // 19-digit jcb
	}

	creditCardMasker := NewCreditCardMasker(Parameters{})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsMultipleCreditCardNumbers(t *testing.T) {
	testCases := []struct {
		value       string
		description string
		expected    string
	}{
		{
			value:    "4012888888881881,4012888888881881",
			expected: "XXXXXXXXXXXX1881,XXXXXXXXXXXX1881",
		},
		{
			value:    "4222222222222 5555555555554444",
			expected: "XXXXXXXXX2222 XXXXXXXXXXXX4444",
		},
		{
			value:    "3566002020360505\n6011000990139424\n30569309025904",
			expected: "XXXXXXXXXXXX0505\nXXXXXXXXXXXX9424\nXXXXXXXXXX5904",
		},
		{
			value:    "4532-0267-2194-6154 and 6011111111111117",
			expected: "XXXX-XXXX-XXXX-6154 and XXXXXXXXXXXX1117",
		},
		{
			value:    "4242 4242 4242 4242 4242 4242 4242 4242",
			expected: "XXXX XXXX XXXX 4242 XXXX XXXX XXXX 4242",
		},
		{
			value:       "4234-5678-9012-3456782 4234-5678-9012-3456790 4234-5678-9012-3456791",
			description: "valid 19 digit | valid 19 digit | failed luhn 19 digit",
			expected:    "XXXX-XXXX-XXXX-XXX6782 XXXX-XXXX-XXXX-XXX6790 4234-5678-9012-3456791",
		},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ExposeLast: 4,
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWhenValueContainsMultipleCreditCardNumbersWithReplacementText(t *testing.T) {
	testCases := []struct {
		value       string
		description string
		expected    string
	}{
		{
			value:    "4012888888881881,4012888888881881",
			expected: "[MASKED],[MASKED]",
		},
		{
			value:    "4222222222222 5555555555554444",
			expected: "[MASKED] [MASKED]",
		},
		{
			value:    "3566002020360505\n6011000990139424\n30569309025904",
			expected: "[MASKED]\n[MASKED]\n[MASKED]",
		},
		{
			value:    "4532-0267-2194-6154 and 6011111111111117",
			expected: "[MASKED] and [MASKED]",
		},
		{
			value:    "4532-0267-2194-6154 and 6011111111111117",
			expected: "[MASKED] and [MASKED]",
		},
		{
			value:    "4242 4242 4242 4242 4242 4242 4242 4242",
			expected: "[MASKED] [MASKED]",
		},
		{
			value:    "4242 4242 4242 4242 4111 1111 1111 1111",
			expected: "[MASKED] [MASKED]",
		},
		{
			value:       "4242 4242 4242 4242 4242 4242 4242 4241",
			description: "valid visa then invalid visa (failed luhn)",
			expected:    "[MASKED] 4242 4242 4242 4241",
		},
		{
			value:       "4234-5678-9012-3456782 4234-5678-9012-3456790 4234-5678-9012-3456791",
			description: "valid 19 digit | valid 19 digit | failed luhn 19 digit",
			expected:    "[MASKED] [MASKED] 4234-5678-9012-3456791",
		},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ReplacementText: "[MASKED]",
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got, testCase.description)
	}
}

func TestMansWithACustomReplacementCharacter(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "not a cc number", expected: "not a cc number"},
		{value: "378282246310005", expected: "***************"},
		{value: "5434 7318 4453 9705", expected: "**** **** **** ****"},
		{value: "6011072668180642", expected: "****************"},
		{value: "5157387747088202", expected: "****************"},
		{value: "4532-0267-2194-6154", expected: "****-****-****-****"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ReplacementToken: "*",
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskNonZeroExposeLast(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "not a cc number", expected: "not a cc number"},
		{value: "378282246310005", expected: "XXXXXXXXXXX0005"},
		{value: "5434 7318 4453 9705", expected: "XXXX XXXX XXXX 9705"},
		{value: "6011072668180642", expected: "XXXXXXXXXXXX0642"},
		{value: "5157387747088202", expected: "XXXXXXXXXXXX8202"},
		{value: "4532-0267-2194-6154", expected: "XXXX-XXXX-XXXX-6154"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ExposeLast: 4,
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWithCustomReplacementTokenAndNonZeroExposeLast(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{value: "not a cc number", expected: "not a cc number"},
		{value: "378282246310005", expected: "#############05"},
		{value: "5434 7318 4453 9705", expected: "#### #### #### ##05"},
		{value: "6011072668180642", expected: "##############42"},
		{value: "5157387747088202", expected: "##############02"},
		{value: "4532-0267-2194-6154", expected: "####-####-####-##54"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ReplacementToken: "#",
		ExposeLast:       2,
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

func TestMaskWithReplacementText(t *testing.T) {

	testCases := []struct {
		value    string
		expected string
	}{
		{value: "not a cc number", expected: "not a cc number"},
		{value: "pls 378282246310005", expected: "pls [MASKED]"},
		{value: "5434 7318 4453 9705", expected: "[MASKED]"},
		{value: "6011072668180642 ok", expected: "[MASKED] ok"},
		{value: "5157387747088202", expected: "[MASKED]"},
		{value: "4532-0267-2194-6154", expected: "[MASKED]"},
	}

	creditCardMasker := NewCreditCardMasker(Parameters{
		ReplacementText:  "[MASKED]",
		ReplacementToken: "%",
		ExposeLast:       3,
	})

	for _, testCase := range testCases {
		got := creditCardMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}

// func TestMask(t *testing.T) {

// 	t.Run("returns other text but not the fake visa", func(t *testing.T) {
// 		fakeVisa := "other text 4242 4242 4242 4242"

// 		parameters := Parameters{}

// 		got := Mask(fakeVisa, parameters)
// 		want := "other text XXXX XXXX XXXX XXXX"
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got [%s] want [%s]", got, want)
// 		}
// 	})

// 	t.Run("exposes the last 4", func(t *testing.T) {
// 		fakeVisa := "other text 4242 4242 4242 4242"

// 		parameters := Parameters{ExposeLast: 4}

// 		got := Mask(fakeVisa, parameters)
// 		want := "other text XXXX XXXX XXXX 4242"
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got [%s] want [%s]", got, want)
// 		}
// 	})

// 	t.Run("returns the incorrect luhn visa", func(t *testing.T) {
// 		fakeVisa := "other text 4242 4242 4242 4241"

// 		parameters := Parameters{}

// 		got := Mask(fakeVisa, parameters)
// 		want := "other text 4242 4242 4242 4241"
// 		if got != want {
// 			t.Errorf("got [%s] want [%s]", got, want)
// 		}
// 	})

// }

// func TestMaskOne(t *testing.T) {
// 	t.Run("masks the value", func(t *testing.T) {
// 		fakeVisa := "4242 4242 4242 4242"

// 		got := maskOne(fakeVisa)
// 		want := fakeVisa

// 		if got != want {
// 			t.Errorf("got [%s] want [%s]", got, want)
// 		}
// 	})
// }
