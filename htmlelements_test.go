package htmlelements

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var testDoc = `
<html>
<head>
<title>Test Document</title>
</head>
<body>
<div id="content">
<p class="para">
Paragraph 1
</p>
<p class="para">
Paragraph 2
</p>
</div>
</body>
</html>
`

func TestElements(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(testDoc))
	if err != nil {
		t.Fatalf("failed to parse test doc: %v", err)
	}
	n := GetElementByID(doc, "content")
	if n == nil {
		t.Errorf("failed to find content node")
	}
	t.Logf("content node %#v", n)
	parasClass := GetElementsByClassName(n, "para")
	if len(parasClass) != 2 {
		t.Errorf("did not find two paragraphs by class")
	}
	parasTag := GetElementsByTagName(n, "p")
	if len(parasTag) != 2 {
		t.Errorf("did not find two paragraphs by tag")
	}
}
