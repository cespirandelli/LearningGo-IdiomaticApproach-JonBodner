package main

import "fmt"

func main() {
	// -----------------------
	// Integer Literals
	// -----------------------
	// Decimal (Base 10)
	decimal := 123

	// Binary (Base 2)
	binary := 0b1011 // 11 in decimal

	// Octal (Base 8)
	octal := 0o17 // 15 in decimal

	// Hexadecimal (Base 16)
	hexadecimal := 0x7F // 127 in decimal

	fmt.Println("---- Integer Literals ----")
	fmt.Println("Decimal:", decimal)
	fmt.Println("Binary:", binary)
	fmt.Println("Octal:", octal)
	fmt.Println("Hexadecimal:", hexadecimal)

	// -------------------------
	// Floating Point Literals
	// -------------------------
	// Decimal Floating Point
	decimalFloat := 3.14

	// Scientific notation (Exponent form)
	scientificNotation := 2.5e2 // 2.5 * 10^2 = 250

	// Hexadecimal Floating Point (with exponent part)
	hexFloat := 0x1.0p3 // 1.0 * 2^3 = 8

	fmt.Println("\n---- Floating Point Literals ----")
	fmt.Println("Decimal Floating Point:", decimalFloat)
	fmt.Println("Scientific Notation:", scientificNotation)
	fmt.Println("Hexadecimal Floating Point:", hexFloat)

	// -------------------------
	// Rune Literals
	// -------------------------
	// Rune as a character (Single Unicode character)
	runeChar := 'a'

	// Rune as an octal number
	runeOctal := '\141' // 'a' = 97

	// Rune as a hexadecimal number
	runeHex := '\x61' // 'a' = 97

	// Rune as a 16-bit hexadecimal number
	rune16Hex := '\u0061' // 'a' = 97

	// Rune as a 32-bit Unicode number
	rune32Hex := '\U00000061' // 'a' = 97

	fmt.Println("\n---- Rune Literals ----")
	fmt.Println("Rune as character:", runeChar)
	fmt.Println("Rune as octal:", runeOctal)
	fmt.Println("Rune as hexadecimal:", runeHex)
	fmt.Println("Rune as 16-bit hexadecimal:", rune16Hex)
	fmt.Println("Rune as 32-bit Unicode:", rune32Hex)

	// -------------------------
	// String Literals
	// -------------------------
	// Interpreted String Literal (uses double quotes)
	interpretedString := "Hello, World!"

	// Raw String Literal (uses backquotes)
	rawString := `Hello,
	World!`

	// String with escape sequences
	escapedString := "Hello,\n\"World!\""

	fmt.Println("\n---- String Literals ----")
	fmt.Println("Interpreted String:", interpretedString)
	fmt.Println("Raw String:", rawString)
	fmt.Println("String with Escape Sequences:", escapedString)
}
