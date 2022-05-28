package main

import (
	"encoding/xml"
	"strings"

	"github.com/adotinthevoid/xmlpp"
	"github.com/davecgh/go-spew/spew"
)

func main() {
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

	d := xml.NewDecoder(strings.NewReader(blob))

	t, err := xmlpp.BuildTree(d)
	if err != nil {
		panic(err)
	}
	spew.Dump(t)
}
