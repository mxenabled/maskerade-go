![maskerade.png](maskerade.png)

# Maskerade-Go

## Account Number Masker Usage

```go
import (
    accountMasker "github.com/mxenabled/maskerade-go/account_number_masker"
)

masker := accountMasker.NewAccountNumberMasker(accountMasker.Parameters{
    ReplacementToken: "X",
    ExposeLast: 2,
})
masker.Mask("5555-5555-5555-5555") // "XXXX-XXXX-XXXX-XX55"
```

## Credit Card Number Masker Usage

```go
import (
    creditMasker "github.com/mxenabled/maskerade-go/credit_card_masker"
)

masker := creditMasker.NewCreditCardMasker(creditMasker.Parameters{
    ReplacementToken: "X",
    ExposeLast: 2,
})
masker.Mask("5555-5555-5555-5555") // "XXXX-XXXX-XXXX-XX55"
```

```go
import (
    creditMasker "github.com/mxenabled/maskerade-go/credit_card_masker"
)

masker := creditMasker.NewCreditCardMasker(creditMasker.Parameters{
    ReplacementText: "[MASKED]",
})
masker.Mask("my credit card number is 4242 4242 4242 4242") // "my credit card number is[MASKED]"
```

## Notes

Golang Port of https://github.com/mxenabled/maskerade
* Works for multiple credit card numbers in the same string.
* Works for account numbers over 7 digits
