// namegen project namegen.go
package namegen

import (
	"math/rand"
	"strings"
	"time"
)

type Name struct {
	Parts map[string][]string
	Rule  []string
	Bad   []string
}

func (n Name) Generate(r *rand.Rand) string {
	var s string
	var rl []string

	for n.illegal(s) {
		rl = strings.Split(n.Rule[r.Intn(len(n.Rule))], "|")
		for _, v := range rl {
			s += n.randPart(r, v)
		}
	}
	return s
}

func (n Name) String() string {
	return n.Generate(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func (n Name) randPart(r *rand.Rand, i string) string {
	part, ok := n.Parts[i]
	if ok {
		return part[r.Intn(len(part))]
	}
	return ""
}

func (n Name) illegal(s string) bool {
	for _, b := range n.Bad {
		if strings.Contains(s, b) {
			return true
		}
	}
	return s == ""
}

type NameType interface {
	Male() Name
	Female() Name
}
