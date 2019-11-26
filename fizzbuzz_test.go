package fizzbuzz

import (
	"errors"
	"testing"
)

func TestRangeWithTriggers(t *testing.T) {
	got, err := RangeWithTriggers(5, 15, Trigger{"Zazz", 5}, Trigger{"Zuzz", 3})
	if err != nil {
		t.Error(err)
	}
	want := `Zazz
Zuzz
7
8
Zuzz
Zazz
11
Zuzz
13
14
ZazzZuzz`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	got, err = RangeWithTriggers(5, 15, Trigger{"Zazz", 5}, Trigger{"Zuzz", 3}, Trigger{"Zarz", 7})
	if err != nil {
		t.Error(err)
	}
	want = `Zazz
Zuzz
Zarz
8
Zuzz
Zazz
11
Zuzz
13
Zarz
ZazzZuzz`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	got, err = RangeWithTriggers(1, 10)
	if err != nil {
		t.Error(err)
	}
	want = `1
2
3
4
5
6
7
8
9
10`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRange(t *testing.T) {
	got, err := Range(5, 15)
	if err != nil {
		t.Error(err)
	}
	want := `Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	got, err = Range(5, 5)
	if err != nil {
		t.Error(err)
	}
	want = "Buzz"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	_, err = Range(5, 4)
	if !errors.Is(err, ErrImproperRange) {
		t.Errorf("got %v, expected %v", err, ErrNumberTooLow)
	}

	_, err = Range(-2, -1)
	if !errors.Is(err, ErrNumberTooLow) {
		t.Errorf("got %v, expected %v", err, ErrNumberTooLow)
	}
}

func TestUpTo(t *testing.T) {
	got, err := UpTo(100)
	if err != nil {
		t.Error(err)
	}
	want := `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
Buzz
Fizz
52
53
Fizz
Buzz
56
Fizz
58
59
FizzBuzz
61
62
Fizz
64
Buzz
Fizz
67
68
Fizz
Buzz
71
Fizz
73
74
FizzBuzz
76
77
Fizz
79
Buzz
Fizz
82
83
Fizz
Buzz
86
Fizz
88
89
FizzBuzz
91
92
Fizz
94
Buzz
Fizz
97
98
Fizz
Buzz`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	_, err = UpTo(0)
	if !errors.Is(err, ErrNumberTooLow) {
		t.Errorf("got %v, expected %v", err, ErrNumberTooLow)
	}
}

func TestDivisibleBy(t *testing.T) {
	trigger := Trigger{Fizz, 3}
	got := trigger.divisibleBy(3)
	want := true
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
	got = trigger.divisibleBy(2)
	want = false
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
	trigger = Trigger{Buzz, 5}
	got = trigger.divisibleBy(0)
	want = true
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
