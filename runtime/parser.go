package runtime

import (
	"github.com/Mrs4s/MiraiGo/message"
	"regexp"
	"strings"
)

type Parser struct {
	Command string
	Args    []string
	Extra   []message.IMessageElement
}

func (p *Parser) Parse(e []message.IMessageElement) {
	var foundCommand = false
	for _, element := range e {
		if te, ok := element.(*message.TextElement); ok {
			if foundCommand {
				continue
			}
			text := strings.TrimSpace(te.Content)
			if text == "" {
				continue
			}
			splitStr := argSplit(text)
			if len(splitStr) >= 1 {
				p.Command = strings.TrimSpace(splitStr[0])
				for _, s := range splitStr[1:] {
					p.Args = append(p.Args, strings.TrimSpace(s))
				}
				foundCommand = true
			}
		} else {
			p.Extra = append(p.Extra, element)
		}
	}
}

func (p *Parser) FilterExtra(etype message.ElementType) []message.IMessageElement {
	var result []message.IMessageElement
	for _, e := range p.Extra {
		if e.Type() == etype {
			result = append(result, e)
		}
	}
	return result
}

func (p *Parser) FilterExtraFunc(f func(e message.IMessageElement) bool) []message.IMessageElement {
	var result []message.IMessageElement
	for _, e := range p.Extra {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

func NewParser() *Parser {
	return new(Parser)
}

func NewParserFromMessage(e []message.IMessageElement) *Parser {
	p := new(Parser)
	p.Parse(e)
	return p
}

// goroutine safe for FindAllString
var ruleRex = regexp.MustCompile(`[^\s"]+|"([^"]*)"`)

func argSplit(str string) (result []string) {
	match := ruleRex.FindAllString(str, -1)
	for _, s := range match {
		result = append(result, strings.Trim(strings.TrimSpace(s), `" `))
	}
	return
}
