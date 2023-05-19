package checkdigit

import "fmt"

type gtin struct {
	digit   int
	posCorr bool
}

func NewJAN13() Provider {
	return &gtin{
		digit:   13,
		posCorr: true,
	}
}

func (g *gtin) Verify(code string) bool {
	if len(code) != g.digit {
		return false
	}
	i, err := g.Generate(code[:len(code)-1])
	return err == nil && i == int(code[len(code)-1]-'0')
}

func (g *gtin) Generate(seed string) (int, error) {
	if len(seed) != g.digit-1 {
		return 0, fmt.Errorf("invalid argument: %q\n", seed)
	}

	var oddSum, evenSum int
	for i, n := range seed {
		if !isNumber(n) {
			return 0, fmt.Errorf("invalid argument: %q\n", seed)
		}
		if g.posCorr {
			i++
		}
		if i%2 == 0 {
			evenSum += int(n - '0')
		} else {
			oddSum += int(n - '0')
		}
	}
	d := 10 - (evenSum*3+oddSum)%10
	if d == 10 {
		d = 0
	}

	return d, nil
}
