package simpleid

import "testing"

func TestNew(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", id)
	if len(id) != 8 {
		t.Fatal("length error")
	}
}

func TestNewWithLength(t *testing.T) {
	id, err := NewWithLength(20)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", id)
	if len(id) != 20 {
		t.Fatal("length error")
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New()
	}
}
