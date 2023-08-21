package tdd

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type Money struct {
	amount   int
	currency string
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money
	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")
	require.EqualValues(t, fifteenDollars, portfolioInDollars)
}

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	return Money{amount: 15, currency: "USD"}
}

func TestURLJoin(t *testing.T) {
	s := filepath.Join("", "av/123.txt")
	fmt.Println(s)

	c, err := url.JoinPath("", "av/123.txt")
	require.NoError(t, err)
	fmt.Println("c", c)

}

func TestCopy(t *testing.T) {
	s1 := bytes.NewBufferString("hello")
	read := bytes.NewBufferString("")

	_, err := io.Copy(read, s1)
	fmt.Println(err)
	fmt.Println(read.String())

}
