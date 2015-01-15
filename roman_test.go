package main

import "testing"

func TestToRoman(t *testing.T) {
	numerals := map[int]string {
		1:    "I",
		4:    "IV",
		5:    "V",
		10:   "X",
		99:   "XCIX",
		149:  "CXLIX",
		499:  "CDXCIX",
		2999: "MMCMXCIX",
		3000: "MMM",
	}
	r := new(Roman)
	for number, numeral := range numerals {
		str := r.ToRoman(number)
		if str != numeral {
			t.Errorf("ToRoman: %d produced %v, expected %v", number, str, numeral)
		}
	}
}

func TestFromRoman(t *testing.T) {
	numerals := map[int]string {
		1:    "I",
		4:    "IV",
		5:    "V",
		10:   "X",
		99:   "XCIX",
		149:  "CXLIX",
		499:  "CDXCIX",
		2999: "MMCMXCIX",
		3000: "MMM",
	}
	r := new(Roman)
	for number, numeral := range numerals {
		res := r.FromRoman(numeral)
		if res != number {
			t.Errorf("ToRoman: %v produced %d, expected %d", numeral, res, number)
		}
	}
}

func BenchmarkToRoman(b *testing.B) {
	numerals := map[int]string {
		1:    "I",
		4:    "IV",
		5:    "V",
		10:   "X",
		99:   "XCIX",
		149:  "CXLIX",
		499:  "CDXCIX",
		2999: "MMCMXCIX",
		3000: "MMM",
	}
	r := new(Roman)
	for i := 0; i < b.N; i++ {
		for number := range numerals {
			r.ToRoman(number)
		}	
	}
}

func BenchmarkFromRoman(b *testing.B) {
	numerals := map[int]string {
		1:    "I",
		4:    "IV",
		5:    "V",
		10:   "X",
		99:   "XCIX",
		149:  "CXLIX",
		499:  "CDXCIX",
		2999: "MMCMXCIX",
		3000: "MMM",
	}
	r := new(Roman)
	for i := 0; i < b.N; i++ {
		for _, numeral := range numerals {
			r.FromRoman(numeral)
		}	
	}
}
