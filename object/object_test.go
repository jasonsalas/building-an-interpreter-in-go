package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "hello world"}
	hello2 := &String{Value: "hello world"}
	diff1 := &String{Value: "my name is jas"}
	diff2 := &String{Value: "my name is jas"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
}
