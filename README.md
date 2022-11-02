# Card Validator
Credit Card Validator - Takes in a credit/debit card number from a common card vendor (Visa, MasterCard, MIR) and validates [(using the Luna algorithm)](https://github.com/faringet/algorithms_and_sorts#algorithms--sorts) it to make sure that it is a valid number.

___
 Visa Features
- Have 16 digits.
- Can be validated using the Luhn algorithm.
- Are actively issued.

Issuer identification number (IIN)
The IIN identifies the card issuing institution that issued the card to the card holder. Visa credit cards begin with the following prefixes: 4539, 4556, 4916, 4532, 4929, 40240071, 4485, 4716, 4.

regex pattern:
```javascript 
^4[0-9]{12}(?:[0-9]{3})?$
```

___
Mastercard Features
- Have 16 digits.
- Can be validated using the Luhn algorithm.
- Are actively issued.

Issuer identification number (IIN)
The IIN identifies the card issuing institution that issued the card to the card holder. Mastercard credit cards begin with the following prefixes: 51, 52, 53, 54, 55, 2221, 2720.

regex pattern:
```javascript 
^5[1-5][0-9]{14}$
```
___
MIR Features
- Have 16 digits.
- Can be validated using the Luhn algorithm.
- Are actively issued.

Issuer identification number (IIN)
The IIN identifies the card issuing institution that issued the card to the card holder. Mir credit cards begin with the following prefixes: 2200, 2201, 2202, 2203, 2204.

regex pattern:
```javascript 
^(?:220[0-4])\d+$
```
