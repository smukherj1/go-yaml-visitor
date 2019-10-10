package v1

import (
	"testing"
)

func TestYamlParse(t *testing.T) {
	y1 := []byte(`
- hi
- hello
- key1:
  - val1
  - innnerKey1: innerval1
    innerKey2: innerval2
  - val3
  key2:
  - val4
`)
	if err := walkYAML(y1); err != nil {
		t.Errorf("Failed to walk YAML: %v", err)
	}
}