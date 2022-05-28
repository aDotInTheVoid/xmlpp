package xmlpp

import "strings"

// return the string, and if it is oneline
func Pp(tree *XmlElement) string {

	var buf strings.Builder
	buf.WriteRune('<')
	// Todo: namespace?
	buf.WriteString(tree.Name)
	if tree.Text == "" && len(tree.Children) == 0 {
		buf.WriteString("/>")
		return buf.String()
	} else if len(tree.Children) == 0 {
		buf.WriteRune('>')
		buf.WriteString(tree.Text)
		buf.WriteString("</")
		buf.WriteString(tree.Name)
		buf.WriteRune('>')
		return buf.String()
	} else if tree.Text == "" {
		buf.WriteRune('>')
		var cstrs []string
		for _, child := range tree.Children {
			cstrs = append(cstrs, Pp(&child))
		}
		if strslen(cstrs) > 70 {
			// SPlitline
			buf.WriteRune('\n')
			for _, cstr := range cstrs {
				buf.WriteString("  ")
				buf.WriteString(indent(cstr))
				buf.WriteRune('\n')
			}
			buf.WriteString("</")
			buf.WriteString(tree.Name)
			buf.WriteRune('>')
			return buf.String()
		} else {
			for _, cstr := range cstrs {
				buf.WriteString(cstr)
			}
		}
		buf.WriteString("</")
		buf.WriteString(tree.Name)
		buf.WriteRune('>')
		return buf.String()

	} else {
		panic("Unsupported")
	}
}

// AAAAAH
func strslen(strs []string) int {
	var l int
	for _, str := range strs {
		l += len(str)
	}
	return l
}

func indent(s string) string {
	return strings.Replace(s, "\n", "\n  ", -1)
}
