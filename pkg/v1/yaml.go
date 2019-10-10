package v1

import (
	"gopkg.in/yaml.v2"
	"log"
	"fmt"
	//"reflect"
)

// Item represents an arbitrary YAML object.
type Item interface{}
// List represents a list of YAML objects.
type List []interface{}

// Map represents a YAML dictionary.
type Map map[interface{}]interface{}


func walkItem(i Item) {
	//log.Printf("walk item: %v, type(%v)", i, reflect.TypeOf(i))
	if s, ok := i.(string); ok {
		log.Printf("Encountered string: %q", s)
	}
	if l, ok := i.([]interface{}); ok {
		walkList(l)
	}
	if m, ok := i.(map[interface{}]interface{}); ok {
		walkMap(m)
	}
}

func walkList(l List) {
	//log.Printf("walk list: %v", l)
	for _, i := range l {
		walkItem(i)
	}
}

func walkMap(m Map) {
	//log.Printf("walk map: %v", m)
	for k, v := range m {
		walkItem(k)
		walkItem(v)
	}
}

func walkYAML(b []byte) error {
	var l List
	lErr := yaml.Unmarshal(b, &l)
	if lErr == nil {
		walkList(l)
		return nil
	}
	var m Map
	mErr := yaml.Unmarshal(b, &m)
	if mErr == nil {
		walkMap(m)
		return nil
	}

	return fmt.Errorf("unable to parse given blob as a YAML list or map: %v %v", lErr, mErr)
}