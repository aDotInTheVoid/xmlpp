package xmlpp

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestTree(t *testing.T) {
	blob := `<unit>
  <fn>
    <name>
      foo
    </name>
    <ret>
      int
    </ret>
    <args>
    </args>
    <body>
    </body>
  </fn>
  <fn>
    <name>
      bar
    </name>
    <ret>
      int
    </ret>
    <args>
    </args>
    <body>
    </body>
  </fn>
  <main>
  </main>
</unit>`
	tree, err := BuildTree(blob)
	if err != nil {
		t.Fatal(err)
	}
	eTree := XmlElement{
		"unit",
		[]XmlElement{
			{
				Name: "fn",
				Children: []XmlElement{
					{Name: "name", Text: "foo"},
					{Name: "ret", Text: "int"},
					{Name: "args"},
					{Name: "body"},
				},
			},
			{
				Name: "fn",
				Children: []XmlElement{
					{Name: "name", Text: "bar"},
					{Name: "ret", Text: "int"},
					{Name: "args"},
					{Name: "body"},
				},
			},
			{
				Name: "main",
			},
		},
		"",
	}
	if !TreeEq(&tree, &eTree) {
		t.Errorf("Expected %s, got %s", spew.Sdump(eTree), spew.Sdump(tree))
	}
}
