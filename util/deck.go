package util

import (
	"bufio"
	"os"
	"strings"
)

type deck []string

//where we save the deck
var DataDir = "/Users/xqy/dev/go/data/"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewDeck() deck {
	numlist := []string{"one", "two", "three", "four"}
	colorlist := []string{"Red", "Green", "Blue"}

	d := []string{}
	for _, color := range colorlist {
		for _, num := range numlist {
			d = append(d, num+" of "+color)

		}
	}

	//save deck into a file
	filename := DataDir + "mydeck.txt"

	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	s := strings.Join(d, ",")
	w := bufio.NewWriter(f)
	_, err2 := w.WriteString(s)
	check(err2)
	w.Flush()

	return d

}

func (d deck) Shuffle() {
	//open deck
	for i := range d { //only check index
		j := MyRandom(len(d))
		d[i], d[j] = d[j], d[i]
	}
}
