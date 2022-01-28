package compare

import "github.com/sergi/go-diff/diffmatchpatch"

var DefaultTextRange = 0.3

type Text interface {
	Text() string
	Length() int
}

func word(txt1, txt2 Text, checkSize bool, limit float64) (int, bool) {
	if checkSize {
		if l1, l2 := float64(txt1.Length()), float64(txt2.Length()); l1 > l2*(1+limit) || l1 < l2*(1-limit) {
			return 0, false
		}
	}

	diff := diffmatchpatch.New()
	diffs := diff.DiffMain(txt1.Text(), txt2.Text(), false)

	return diff.DiffLevenshtein(diffs), true
}

func Word(txt1, txt2 Text, checkSize bool) (int, bool) {
	return word(txt1, txt2, checkSize, DefaultTextRange)
}

func WordWithRange(txt1, txt2 Text, limit float64) (int, bool) {
	return word(txt1, txt2, true, limit)
}
