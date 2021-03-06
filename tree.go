package xmlpp

import (
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

type XmlElement struct {
	Name     string
	Children []XmlElement
	Text     string
}

var ErrUnbalancedXml = errors.New("unbalanced xml")
var ErrUnsupportedXml = errors.New("unsupported xml")

func BuildTree(s string) (XmlElement, error) {

	d := xml.NewDecoder(strings.NewReader(strings.TrimSpace(s)))

	var stack []XmlElement
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return XmlElement{}, err
		}
		switch v := tok.(type) {
		case xml.StartElement:
			elem := XmlElement{
				Name: v.Name.Local,
			}
			stack = append(stack, elem)
		case xml.EndElement:
			if len(stack) > 1 {
				stack[len(stack)-2].Children = append(stack[len(stack)-2].Children, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		case xml.CharData:
			stack[len(stack)-1].Text += strings.TrimSpace(string(v))
		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		default:
			return XmlElement{}, ErrUnsupportedXml
		}
	}
	if len(stack) != 1 {
		return XmlElement{}, ErrUnbalancedXml
	}
	return stack[0], nil
}

func TreeEq(a, b *XmlElement) bool {
	return a.Name == b.Name &&
		a.Text == b.Text &&
		treesEq(a.Children, b.Children)
}

func treesEq(as, bs []XmlElement) bool {
	if len(as) != len(bs) {
		return false
	}
	for i, a := range as {
		if !TreeEq(&a, &bs[i]) {
			return false
		}
	}
	return true
}
