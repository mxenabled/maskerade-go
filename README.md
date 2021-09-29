![maskerade.png](maskerade.png)

# Maskerade-Go

## Account Number Masker Usage

```go

masker := NewAccountNumberMasker(Parameters{
    ReplacementToken: "X",
    ExposeLast: 2,
})
masker.Mask("5555-5555-5555-5555") // "XXXX-XXXX-XXXX-XX55"
```

## Credit Card Number Masker Usage

```go

masker := NewCreditCardMasker(Parameters{
    ReplacementToken: "X",
    ExposeLast: 2,
})
masker.Mask("5555-5555-5555-5555") // "XXXX-XXXX-XXXX-XX55"
```

```go

masker := NewCreditCardMasker(Parameters{
    ReplacementText: "[MASKED]"
})
masker.Mask("my credit card number is 4242 4242 4242 4242") // "my credit card number is[MASKED]"
```

## Notes

Golang Port of https://github.com/mxenabled/maskerade
* Works for multiple credit card numbers in the same string.
* Works for account numbers over 7 digits & hyphens long
