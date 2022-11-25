package hello

import "rsc.io/quote/v3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}

func Proverb2() string {
	return quote.Concurrency()
}
