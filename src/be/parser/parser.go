package parser

import (
	"be/structs"
	"fmt"
	"strings"
)

type Parser struct {

	// 是否处于引用
	inRef bool

	// 当前处理中的content
	content *structs.ParserContent

	rawContent string
	lines      []string

	result *structs.ParserResult
}

func NewParser() *Parser {
	return &Parser{
		inRef:   false,
		content: &structs.ParserContent{Contents: []*structs.ParserContentBlock{}},
		result:  &structs.ParserResult{Contents: []*structs.ParserContent{}},
	}
}

func (p *Parser) Parser(rawContent string) {
	p.rawContent = strings.TrimSpace(rawContent)
	p.lines = strings.Split(rawContent, "\n")
	p.parse()
}

func (p *Parser) GetResult() *structs.ParserResult {
	return p.result
}

// 由于文章内容较短，这里直接使用字符串拼接的方式处理content，性能较差
func (p *Parser) parse() {
	for _, line := range p.lines {
		line = strings.TrimSpace(line)
		// 是否处于引用
		if p.inRef {
			if strings.HasPrefix(line, "```") {
				// 引用结束
				p.inRef = false
				p.content.Content = strings.Trim(p.content.Content, "\n")
				p.result.Contents = append(p.result.Contents, p.content)
				p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
				continue
			} else {
				// 引用的content
				p.content.Content += fmt.Sprintf("%s\n", line)
			}
		}
	}
}
