package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var infile string
	var outfile string
	var s int
	var to_encrypt bool
	var to_decrypt bool
	flag.StringVar(&infile, "i", "cleartext.txt", "Specify input file. Default is cleartext.txt.")
	flag.StringVar(&outfile, "o", "ciphertext.txt", "Specify output file. Default is ciphertext.txt.")
	flag.IntVar(&s, "s", -1, "Specify shift. Should be 8 bit unsigned integer.")
	flag.BoolVar(&to_encrypt, "e", false, "Specify encryption. Default is true.")
	flag.BoolVar(&to_decrypt, "d", false, "Specify decryption. Default is false.")

	flag.Parse()
	if !to_encrypt && !to_decrypt {
		to_encrypt = true
	}
	var shift uint8
	if s != -1 {
		shift = uint8(s)
	} else {
		shift = 0
	}

	if to_encrypt {
		cleartext, err := os.ReadFile(infile)
		fileReadError(err, infile)
		err = os.WriteFile(outfile, []byte(Encrypt(string(cleartext), shift)), 0644)
		fileReadError(err, outfile)
	} else {
		ciphertext, err := os.ReadFile(infile)
		fileReadError(err, infile)
		if s != -1 {
			err = os.WriteFile(outfile, []byte(Decrypt(string(ciphertext), shift)), 0644)
			fileReadError(err, outfile)
		} else {
			user_break := false
			input_reader := bufio.NewReader(os.Stdin)
			for !user_break && shift < 26 {
				fmt.Println("Press enter to continue. Enter 's' to stop.")
				shift++
				fmt.Println(Decrypt(string(ciphertext), shift))
				input, err := input_reader.ReadString('\n')
				if err != nil {
					fmt.Println(err)
					fmt.Println("Error reading your input")
					os.Exit(1)
				}
				user_break = input != "\n"
			}
		}
	}
}

func fileReadError(e error, file string) {
	if e != nil {
		fmt.Println(e)
		fmt.Printf("The file %s could not be used.\n", file)
		os.Exit(1)
	}
}

func Encrypt(cleartext string, shift uint8) string {
	text := allCaps(removePunctuation(cleartext))
	b := []byte(text)
	for i, x := range b {
		b[i] = ((x-65)+shift)%26 + 65
	}
	return string(b)
}

func Decrypt(ciphertext string, shift uint8) string {
	b := []byte(ciphertext)
	shift = 26 - (shift % 26)
	for i, x := range b {
		b[i] = ((x-65)+shift)%26 + 65
	}
	return string(b)
}

func removePunctuation(text string) string {
	b := []byte(text)
	out := make([]byte, 0)
	for _, x := range b {
		if (x >= 97 && x <= 122) || (x >= 65 && x <= 90) {
			out = append(out, x)
		}
	}
	return string(out)
}

func allCaps(text string) string {
	b := []byte(text)
	out := make([]byte, len(text))
	for i, x := range b {
		if x >= 97 && x <= 122 {
			out[i] = x - 32
		} else {
			out[i] = x
		}
	}
	return string(out)
}
