package pqdsn

import (
	"bytes"
	"strconv"
	"strings"
)

type builder struct {
	bytes.Buffer
	escape bool
}

func (b *builder) addKey(k string) {
	if b.Len() > 0 {
		b.WriteRune(' ')
	}

	b.WriteString(k)
	b.WriteRune('=')
}

func (b *builder) addInt(k string, i int) {
	b.addKey(k)
	b.WriteString(strconv.Itoa(i))
}

func (b *builder) addString(k, s string) {
	b.addKey(k)

	if !b.escape {
		b.WriteString(s)
		return
	}

	if strings.ContainsRune(s, ' ') {
		b.WriteRune('\'')
		defer b.WriteRune('\'')
	}

	for _, c := range []rune(s) {
		if c == '\'' {
			b.WriteRune('\\')
		}
		b.WriteRune(c)
	}
}
