package xmlpp

import (
	"testing"
)

func TestPp(t *testing.T) {
	blob := `
<unit>
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
</unit>
`

	tree, err := BuildTree(blob)
	if err != nil {
		t.Fatal(err)
	}

	ePrity := `<unit>
  <fn><name>foo</name><ret>int</ret><args/><body/></fn>
  <fn><name>bar</name><ret>int</ret><args/><body/></fn>
  <main/>
</unit>`

	pritty := Pp(&tree)
	if pritty != ePrity {
		t.Errorf("Expected %s, got %s", ePrity, pritty)
	}
}
