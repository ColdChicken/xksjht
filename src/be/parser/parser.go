package parser

import (
	xe "be/common/error"
	"be/common/log"
	"be/structs"
	"fmt"
	"regexp"
	"strings"
)

type Parser struct {

	// 是否处于引用
	inRef bool

	// 小节层次
	sectionLevelOne   int64
	sectionLevelTwo   int64
	sectionLevelThree int64

	// 当前处理中的content
	content *structs.ParserContent

	rawContent string
	lines      []string

	result *structs.ParserResult
}

func NewParser() *Parser {
	return &Parser{
		inRef:             false,
		sectionLevelOne:   0,
		sectionLevelTwo:   0,
		sectionLevelThree: 0,
		content:           &structs.ParserContent{Contents: []*structs.ParserContentBlock{}},
		result:            &structs.ParserResult{Contents: []*structs.ParserContent{}},
	}
}

func (p *Parser) Parser(rawContent string) error {
	p.rawContent = strings.TrimSpace(rawContent)
	p.lines = strings.Split(rawContent, "\n")
	err := p.parse()
	if err != nil {
		log.Errorf("解析内容失败, %s", err.Error())
		return err
	}
	if p.result.Title == "" {
		log.Errorln("标题信息不存在")
		return xe.New("标题信息不存在")
	}
	return nil
}

func (p *Parser) GetResult() *structs.ParserResult {
	return p.result
}

// 由于文章内容较短，这里直接使用字符串拼接的方式处理content，性能较差
func (p *Parser) parse() error {
	for _, line := range p.lines {
		log.Debugln(line)
		if strings.TrimSpace(line) == "" {
			continue
		}
		// 是否处于引用
		if p.inRef {
			if strings.HasPrefix(line, "```") {
				// 引用结束
				p.inRef = false
				p.content.Content = strings.Trim(p.content.Content, "\n")
				p.result.Contents = append(p.result.Contents, p.content)
				p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
			} else {
				// 引用的content
				p.content.Content += fmt.Sprintf("%s\n", line)
			}
			continue
		}

		if strings.HasPrefix(line, "#") {
			if err := p.parseSection(line); err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "```") {
			if err := p.parseRef(line); err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "*") {
			if err := p.parseList(line); err != nil {
				return err
			}
		} else if strings.HasPrefix(line, "$") {
			if err := p.parseOrderedList(line); err != nil {
				return err
			}
		} else {
			if err := p.parseBlock(line); err != nil {
				return err
			}
		}
	}
	if p.content.Type != "" {
		p.result.Contents = append(p.result.Contents, p.content)
	}
	return nil
}

// 处理 # 开头的内容
func (p *Parser) parseSection(line string) error {
	if p.content.Type != "" {
		p.result.Contents = append(p.result.Contents, p.content)
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
	}
	switch len(strings.Split(strings.Split(line, " ")[0], "#")) - 1 {
	case 1:
		p.result.Title = strings.Join(strings.Split(line, " ")[1:len(strings.Split(line, " "))], " ")
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
		return nil
	case 2:
		p.sectionLevelTwo = 0
		p.sectionLevelThree = 0
		p.sectionLevelOne++
		p.content.Type = "section"
		p.content.SectionID = fmt.Sprintf("%d", p.sectionLevelOne)
		p.content.Content = strings.Join(strings.Split(line, " ")[1:len(strings.Split(line, " "))], " ")
		return nil
	case 3:
		p.sectionLevelThree = 0
		p.sectionLevelTwo++
		p.content.Type = "section"
		p.content.SectionID = fmt.Sprintf("%d.%d", p.sectionLevelOne, p.sectionLevelTwo)
		p.content.Content = strings.Join(strings.Split(line, " ")[1:len(strings.Split(line, " "))], " ")
		return nil
	case 4:
		p.sectionLevelThree++
		p.content.Type = "section"
		p.content.SectionID = fmt.Sprintf("%d.%d.%d", p.sectionLevelOne, p.sectionLevelTwo, p.sectionLevelThree)
		p.content.Content = strings.Join(strings.Split(line, " ")[1:len(strings.Split(line, " "))], " ")
		return nil
	default:
		return xe.New("#开始的段落最多只能支持4个#")
	}
}

// 处理 ``` 开头的内容
func (p *Parser) parseRef(line string) error {
	if p.content.Type != "" {
		p.result.Contents = append(p.result.Contents, p.content)
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
	}
	p.inRef = true
	p.content.Type = "ref"
	sourcePos := strings.Index(line, "source")
	if sourcePos == -1 {
		if strings.Contains(line, "[") && strings.Contains(line, "]") {
			// link
			p.content.Source = "link"
			p.content.Value = line[3:len(line)]
		} else {
			// txt
			p.content.Source = "txt"
			p.content.Value = line[3:len(line)]
		}
	} else {
		source := strings.Split(strings.Split(line, ",")[0], "=")[1]
		value := strings.Split(line, "=")[len(strings.Split(line, "="))-1]
		if source != "link" && source != "txt" && source != "img" && source != "ref" {
			return xe.New("无效的ref类型")
		}
		p.content.Source = source
		p.content.Value = value
	}
	return nil
}

// 处理 * 开头的内容
func (p *Parser) parseList(line string) error {
	if p.content.Type != "" && p.content.Type != "list" {
		p.result.Contents = append(p.result.Contents, p.content)
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
	}
	p.content.Type = "list"
	p.content.Ordered = false
	data := line[1:len(line)]
	p.content.Contents = append(p.content.Contents, &structs.ParserContentBlock{Type: "block_txt", Content: data})
	return nil
}

// 处理 $ 开头的内容
func (p *Parser) parseOrderedList(line string) error {
	if p.content.Type != "" && p.content.Type != "list" {
		p.result.Contents = append(p.result.Contents, p.content)
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
	}
	p.content.Type = "list"
	p.content.Ordered = true
	data := line[1:len(line)]
	p.content.Contents = append(p.content.Contents, &structs.ParserContentBlock{Type: "block_txt", Content: data})
	return nil
}

// 处理段落
func (p *Parser) parseBlock(line string) error {
	if p.content.Type != "" {
		p.result.Contents = append(p.result.Contents, p.content)
		p.content = &structs.ParserContent{Contents: []*structs.ParserContentBlock{}}
	}

	p.content.Type = "block"
	if contents, err := p.parseBlockContent(line); err != nil {
		return err
	} else {
		p.content.Contents = contents
	}

	return nil
}

// 处理block
func (p *Parser) parseBlockContent(data string) ([]*structs.ParserContentBlock, error) {
	blocks := []*structs.ParserContentBlock{}
	linkFmt := `(\S*)\[\S*\]`
	linkRegexp := regexp.MustCompile(linkFmt)
	matcheIdxs := linkRegexp.FindAllStringIndex(data, -1)
	for idx, linkPos := range matcheIdxs {
		otherContent := ""
		if idx == 0 {
			otherContent = data[0:linkPos[0]]
		} else {
			otherContent = data[matcheIdxs[idx-1][1]:linkPos[0]]
		}
		// 分析txt和underline
		for idx, words := range strings.Split(otherContent, "**") {
			if idx%2 == 0 {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_txt", Content: strings.TrimSpace(words)})
			} else {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_underline", Content: strings.TrimSpace(words)})
			}
		}

		// 分析ref
		refData := data[linkPos[0]:linkPos[1]]
		refContent := strings.Split(refData, ")")[0][1:len(strings.Split(refData, ")")[0])]
		refLink := strings.Split(refData, "[")[1][0 : len(strings.Split(refData, "[")[1])-1]
		blocks = append(blocks, &structs.ParserContentBlock{Type: "block_ref", Content: refContent, Link: refLink})
	}
	// 处理剩余部分
	if len(matcheIdxs) == 0 {
		for idx, words := range strings.Split(data, "**") {
			if idx%2 == 0 {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_txt", Content: strings.TrimSpace(words)})
			} else {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_underline", Content: strings.TrimSpace(words)})
			}
		}
	} else {
		idx := len(matcheIdxs)
		otherContent := data[matcheIdxs[idx-1][1]:len(data)]
		for idx, words := range strings.Split(otherContent, "**") {
			if idx%2 == 0 {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_txt", Content: strings.TrimSpace(words)})
			} else {
				blocks = append(blocks, &structs.ParserContentBlock{Type: "block_underline", Content: strings.TrimSpace(words)})
			}
		}
	}
	return blocks, nil
}
