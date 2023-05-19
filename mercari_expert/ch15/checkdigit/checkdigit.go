package checkdigit

import "fmt"

type luhn struct{}

func NewLuhn() Provider {
	return &luhn{}
}

func (l *luhn) Verify(code string) bool {
	if len(code) < 2 {
		return false
	}
	i, err := l.Generate(code[:len(code)-1])
	return err == nil && i == int(code[len(code)-1]-'0')
}

func (l *luhn) Generate(seed string) (int, error) {
	if seed == "" {
		return 0, fmt.Errorf("invalid argument: %q\n", seed)
	}

	sum, parity := 0, (len(seed)+1)%2
	for i, v := range seed {
		if !isNumber(v) {
			return 0, fmt.Errorf("invalid argument: %q\n", seed)
		}
		d := int(v - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	return sum * 9 % 10, nil
}

func isNumber(r rune) bool {
	digit := r - '0'
	if 0 <= int(digit) && int(digit) <= 9 {
		return true
	}
	return false
}

type Generator interface {
	Generate(seed string) (int, error)
}

type Verifier interface {
	Verify(code string) bool
}

type Provider interface {
	Generator
	Verifier
}
