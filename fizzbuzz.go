// Package fizzbuzz has been thoroughly crafted as a tool for very
// intelligent programmers such as the creator of said package to
// have the ability to use a package that implements the algorithm
// which is required for appropriately running the FizzBuzz algorithm
// which was invented by Imran Ghory. The difficulties of porting this
// algorithm over to the Go language is not to be understood as an
// underachievment and as such should be looked upon with great jealousy.
// Please use this package for job interviews, iOS applications, and
// governmental secret cryptography NASA engineering.
package fizzbuzz

import (
	"errors"
	"fmt"
)

const (
	// Fizz is the effervescense of a soda or sparkling water.
	Fizz = "Fizz"
	// Buzz is the frenetic fluttering of a bumblebee's wings.
	Buzz = "Buzz"
)

var (
	// ErrNumberTooLow will be returned when the input number(s) are less than 1.
	ErrNumberTooLow = errors.New("fizzbuzz: input number(s) must be greater than 0")
	// ErrImproperRange will be returned when the end number input is larger
	// than then begin number input.
	ErrImproperRange = errors.New("fizzbuzz: end greater than beginning")
)

// Trigger is used as input to the functions RangeWithTriggers and
// UpToWithTriggers. You can define the trigger where you'd like the
// FizzBuzz algorithm to replace words.
type Trigger struct {
	Word    string
	Divisor int
}

func (t *Trigger) divisibleBy(numerator int) bool {
	return numerator%t.Divisor == 0
}

// RangeWithTriggers takes begin and end points along with a variatic number
// of Triggers to use in its algorithm. Any number within the range will
// be replaced if divisible by the Trigger. The order in which the Triggers
// are sent to the function affects outcome. If Buzz is sent before Fizz, the
// output will show BuzzFizz instead of FizzBuzz. Be aware of this.
func RangeWithTriggers(begin, end int, triggers ...Trigger) (string, error) {
	if end < begin {
		return "", ErrImproperRange
	}
	if begin < 1 {
		return "", ErrNumberTooLow
	}
	var output string
	for i := begin; i <= end; i++ {
		found := false
		for _, trigger := range triggers {
			if trigger.divisibleBy(i) {
				output += trigger.Word
				found = true
			}
		}
		if !found {
			output += fmt.Sprintf("%d", i)
		}
		output += "\n"
	}

	return output[:len(output)-1], nil

}

// Range takes a begin and end number. Range does not take in Triggers and
// instead returns the defualt triggers of Trigger{Fizz, 3}, Trigger{Buzz, 5}.
func Range(begin, end int) (string, error) {
	return RangeWithTriggers(begin, end, Trigger{Fizz, 3}, Trigger{Buzz, 5})
}

// UpTo functions similarly to Range with the default Triggers. UpTo takes a
// single end number input and starts counting from 1. This is the most basic
// FizzBuzz function.
func UpTo(end int) (string, error) {
	if end < 1 {
		return "", ErrNumberTooLow
	}
	return Range(1, end)
}

// UpToWithTriggers works functionally the same as RangeWithTriggers
// except instead of taking a range only takes an end number whilst starting
// with 1.
func UpToWithTriggers(end int, triggers ...Trigger) (string, error) {
	if end < 1 {
		return "", ErrNumberTooLow
	}
	return RangeWithTriggers(1, end, triggers...)
}
