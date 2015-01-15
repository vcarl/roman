package main

import (
	"bytes"
	"flag"
	"fmt"
	"math"
	"strconv"
)

type Roman struct {
	runes    *[]rune
	numerals *map[rune]int
	buf      bytes.Buffer
}

// ToRoman accepts an int and returns a string representing the roman
// numeral equivalent.
func (r *Roman) ToRoman(x int) string {
	r.buf = *new(bytes.Buffer)
	// If the list of runes is nil, initialize it.
	if r.runes == nil {
		// List all possible roman numerals, plus 2 bogus ones to simplify code.
		r.runes = &[]rune{'I', 'V', 'X', 'L', 'C', 'D', 'M', 'E', 'E'}
	}
	// If the number is greater than 3000, return an empty string.
	if x > 3000 {
		return ""
	}

	for i := 3; i >= 0; i-- {
		// Break up the number into its component digits (e.g. 1066 to 1, 0, 6, 6), and
		// pass that digit and what numerals represent 1, 5, and 10 for that order of
		// magnitude (e.g. 10, 50, 100; 100, 500, 1000).
		r.buf.WriteString(r.romanHelper((x/int(math.Pow10(i)))%10, (*r.runes)[i*2], (*r.runes)[i*2+1], (*r.runes)[i*2+2], new(bytes.Buffer)))
	}

	return r.buf.String()
}

// romanHelper is a helper function called from within ToRoman.
func (r *Roman) romanHelper(x int, one, five, ten rune, buf *bytes.Buffer) string {
	if x >= 10 || x <= 0 {
		return ""
	}
	// Write the appropriate runes for each number.
	switch {
	case x == 4:
		buf.WriteRune(one)
		buf.WriteRune(five)
	case x == 5:
		buf.WriteRune(five)
	case x == 9:
		buf.WriteRune(one)
		buf.WriteRune(ten)
	case x > 5:
		diff := x - 5
		buf.WriteRune(five)
		for i := 0; i < diff; i++ {
			buf.WriteRune(one)
		}
	case x < 5:
		for i := 0; i < x; i++ {
			buf.WriteRune(one)
		}
	}
	return buf.String()
}

// FromRoman converts a Roman numeral into a number.
// No validation is currently performed, garbage in garbage out.
// An integer of the determined value is returned.
func (r *Roman) FromRoman(x string) int {
	// TODO: validate that the passed string is a valid Roman numeral.

	// If there isn't a list of numerals, initialize it.
	if r.numerals == nil {
		r.numerals = &map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}
	}

	var next, cur, total int
	// Grab the length of the string given.
	// TODO: not 100% sure how this handles improper Unicode characters.
	length := len([]rune(x))

	// Loop through the string and interpret the numerals.
	// Track the current and next numeral so we can check if the current numeral
	// should be subtracted from the next.
	for i := 0; i < length; i++ {
		// Get the numeric representation of the current and next numerals.
		// If we're at the last rune, just give it -1 for next.
		cur = (*r.numerals)[rune(x[i])]
		if i+1 < length {
			next = (*r.numerals)[rune(x[i+1])]
		} else {
			next = -1
		}
		// If the current numeral is smaller than the next, subtract the current from
		// the next. Since we've now handled two runes at once, increment the loop
		// counter.
		if cur < next {
			cur = (next - cur)
			i++
		}
		// Add the current value to the total.
		total += cur
	}

	return total
}

func main() {
	var from = flag.Bool("from", false, "testing!")
	var to = flag.Bool("to", false, "testing!")
	flag.Parse()

	roman := new(Roman)

	if len(flag.Args()) != 0 {
		var arg = flag.Args()[0]

		if *from {
			fmt.Println(roman.FromRoman(arg))
		} else if *to {
			var val, _ = strconv.Atoi(arg)
			fmt.Println(roman.ToRoman(val))
		}
	}
}
