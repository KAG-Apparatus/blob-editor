package cfg

import (
	"bytes"
)

type Element interface {
	String() string
	Literal() []Token
}

type CFGFile struct {
	Elements []Element
}

func (p *CFGFile) String() string {
	var out bytes.Buffer
	for _, s := range p.Elements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Assignment struct {
	LeftHand  []Token
	RightHand []Token
}

func (a *Assignment) String() string {
	var out bytes.Buffer
	for _, t := range a.LeftHand {
		if t.Type == DOLLAR || t.Type == AT {
			out.WriteString(t.Literal)
			continue
		}
		out.WriteString(t.Literal + " ")
	}
	out.WriteString("= ")
	for _, t := range a.RightHand {
		if t.Type == NEWLINE {
			out.WriteString(t.Literal)
			continue
		}
		out.WriteString(t.Literal + " ")
	}
	return out.String()
}

func (as *Assignment) Literal() []Token {
	tokens := make([]Token, 0)
	tokens = append(tokens, as.LeftHand...)
	tokens = append(tokens, newToken(ASSIGN, '='))
	tokens = append(tokens, as.RightHand...)
	return tokens
}
