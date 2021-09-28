package account_number_masker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMask(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{

		{value: "ROCKY MOUNTAIN P BILL P *******8001 REF # 111111111113225", expected: "ROCKY MOUNTAIN P BILL P *******8001 REF # XXXXXXXXXXX3225"},
		{value: "ATM WITHDRAWAL - 6756 COVINGTON HWY LITHONIA GA 4039 0000214", expected: "ATM WITHDRAWAL - 6756 COVINGTON HWY LITHONIA GA 4039 XXX0214"},
		{value: "ATM Point of Sale; SNAP RUSH CITY MN-CK; 320-358-0091, MN; - Unknown Code", expected: "ATM Point of Sale; SNAP RUSH CITY MN-CK; 320-358-0091, MN; - Unknown Code"},
		{value: "ATM/ACH DEBIT 245-SPANKY'S TOBACCOTELECHK 697-9263 PURCHASE   POP:168      OI", expected: "ATM/ACH DEBIT 245-SPANKY'S TOBACCOTELECHK 697-9263 PURCHASE   POP:168      OI"},
		{value: "56374628546", expected: "XXXXXXX8546"},
		{value: "555-555-5555", expected: "555-555-5555"},
		{value: "555-5555", expected: "555-5555"},
		// {value: "--55---", expected: "--X5---"}, // not working test case at the moment
		{value: "5555-5555", expected: "XXXX-5555"},
		{value: "5555-5555-5555-5555", expected: "XXXX-XXXX-XXXX-5555"},
		{value: "555-5555 5555-5555-5555-5555", expected: "555-5555 XXXX-XXXX-XXXX-5555"},
		{value: "5555-5555-5555-5555 555-5555", expected: "XXXX-XXXX-XXXX-5555 555-5555"},
		{value: "46352846478 Subway 5736285638", expected: "XXXXXXX6478 Subway XXXXXX5638"},
		{value: "Check #55555", expected: "Check #55555"},
		{value: "LA FITNESS 555-555-5555555555", expected: "LA FITNESS XXX-XXX-XXXXXX5555"},
		{value: "LA FITNESS 555-555-5555", expected: "LA FITNESS 555-555-5555"},
		{value: "CHECK CRD PURCHASE 03/23 EDDIE BAUER 0835 OREM UT 486831XXXXXX6803 112233445566778 ?MCC=5655", expected: "CHECK CRD PURCHASE 03/23 EDDIE BAUER 0835 OREM UT 486831XXXXXX6803 XXXXXXXXXXX6778 ?MCC=5655"},
		{value: "CHECK 01069, 5544332211", expected: "CHECK 01069, 5544332211"},
		{value: "Share Draft # 555555", expected: "Share Draft # 555555"},
	}

	accountNumberMasker := NewAccountNumberMasker(Parameters{})

	for _, testCase := range testCases {
		got := accountNumberMasker.Mask(testCase.value)

		assert.Equal(t, testCase.expected, got)
	}
}
