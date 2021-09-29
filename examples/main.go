package main

import (
	"fmt"
	accountMasker "github.com/mxenabled/maskerade-go/account_number_masker"
	creditMasker "github.com/mxenabled/maskerade-go/credit_card_masker"
)

func main() {
	amasker := accountMasker.NewAccountNumberMasker(accountMasker.Parameters{
		ReplacementToken: "X",
		ExposeLast:       2,
	})
	fmt.Println(amasker.Mask("5555-5555-5555-5555"))
	// XXXX-XXXX-XXXX-XX55

	cmasker := creditMasker.NewCreditCardMasker(creditMasker.Parameters{
		ReplacementToken: "X",
		ExposeLast:       2,
	})
	fmt.Println(cmasker.Mask("my credit card is 4242 4242 4242 4242"))
	// my credit card is XXXX XXXX XXXX XX42

	cmasker = creditMasker.NewCreditCardMasker(creditMasker.Parameters{
		ReplacementText: "[MASKED]",
	})
	fmt.Println(cmasker.Mask("my credit card is 4242 4242 4242 4242"))
	// my credit card is [MASKED]
}
