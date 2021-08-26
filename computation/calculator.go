package computation

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

type Calculator struct {
	number  int
	meaning string
}

func sha1Converter(s string, factor int) string {
	sh := sha1.New()
	sh.Write([]byte(s))
	if factor == 0 {
		return hex.EncodeToString(sh.Sum(nil))
	} else {
		return hex.EncodeToString(sh.Sum(make([]byte, factor)))
	}

}

func New(question string) *Calculator {
	answer := rand.Intn(len(question))
	guide_for_humans := sha1Converter(question, 0)
	return &Calculator{
		answer,
		guide_for_humans,
	}
}

func (c *Calculator) AddComplexity(a int) {
	c.number = rand.Intn(c.number + a)
	c.meaning = sha1Converter(c.meaning, a)
}

func (c *Calculator) GetWiser() (int, string) {
	return c.number, c.meaning
}
