# Caesar Cipher
This is a simple command line implementation of a Caesar cipher in Go. 

## Flags
- `-i` string: the path to the input file (text to encode or decode)
- `-o` string: the path to the desired output file (where to store encoded or decoded message)
- `-s` unsigned int: the number of characters to shift the message
- `-e` boolean: flag for encoding
- `-d` boolean: flag for decoding

Encoding is the default, with a default shift of 3. The default input file is cleartext.txt and the default output file is ciphertext.txt. If no shift is given when decoding, the program will print out all shifts between 1 and 26. Entering a newline will continue to the next shift, and entering any other character will stop the program. 

## Example
```
caesar -i cleartext.txt -s 5 -e -o ciphertext.txt
caesar -i ciphertext.txt -s 5 -d -o message.txt
```
