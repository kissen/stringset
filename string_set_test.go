package stringset

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringSet_General(t *testing.T) {
	set := New()

	words := strings.Split(
		"thy brother hecuba from dymas sprung a valiant warrior haughty bold and young",
		" ",
	)

	// add

	for _, word := range words {
		if added := set.Put(word); !added {
			t.Errorf("failed to add word=%v", word)
		}
	}

	// len

	if actual := set.Len(); actual != len(words) {
		t.Errorf("bad return of Len() expected=%v actual=%v", len(words), actual)
	}

	// get

	for _, word := range words {
		if !set.Contains(word) {
			t.Errorf("word=%v missing even though it should have been added", word)
		}
	}

	// remove existing

	if removed := set.Remove("valiant"); !removed {
		t.Errorf("refusing to remove added entry")
	}

	// try removing missing

	if removed := set.Remove("hector"); removed {
		t.Errorf("confirming removal of item that was not added")
	}
}

func TestStringSet_Uniques(t *testing.T) {
	uniques := []string{
		"what", "shall", "we", "do", "with", "the",
		"drunken", "sailor",
	}

	notinclude := []string{
		"dem", "morgenrot", "entgegen", "genossen",
	}

	s0 := New()
	s0.Put(uniques...)

	s1 := New()
	for _, u := range uniques {
		s1.Put(u)
	}

	s2 := NewWith(uniques...)

	sets := []StringSet{
		s0, s1, s2,
	}

	for _, s := range sets {
		if s.Len() != len(uniques) {
			t.Error("wrong length")
		}

		for _, u := range uniques {
			if !s.Contains(u) {
				t.Errorf("missing %v", u)
			}
		}

		for _, n := range notinclude {
			if s.Contains(n) {
				t.Errorf("found %v but was not added", n)
			}
		}
	}
}

func TestStringSet_Dups(t *testing.T) {
	notinclude := []string{
		"dem", "morgenrot", "entgegen", "genossen",
	}

	dups := []string{
		"what", "shall", "we", "do", "we", "with", "the",
		"with", "drunken", "sailor", "sailor",
	}

	s0 := New()
	s0.Put(dups...)

	s1 := New()
	for _, u := range dups {
		s1.Put(u)
	}

	s2 := NewWith(dups...)

	sets := []StringSet{
		s0, s1, s2,
	}

	for _, s := range sets {
		if s.Len() != 8 {
			t.Error("wrong length")
		}

		for _, d := range dups {
			if !s.Contains(d) {
				t.Errorf("missing %v", d)
			}
		}

		for _, n := range notinclude {
			if s.Contains(n) {
				t.Errorf("found %v but was not added", n)
			}
		}
	}
}

func TestStringSet_Remove(t *testing.T) {
	uniques := []string{
		"what", "shall", "we", "do", "with", "the",
		"drunken", "sailor",
	}

	notinclude := []string{
		"dem", "morgenrot", "entgegen", "genossen",
	}

	s0 := NewWith(uniques...)

	for _, u := range uniques {
		if !s0.Remove(u) {
			t.Error("reporting not deleted but should have been deleted")
		}
	}

	s1 := NewWith(uniques...)

	for _, n := range notinclude {
		if s1.Remove(n) {
			t.Errorf("reporting deleted but should not have been deleted")
		}
	}
}

func TestStringSet_PutReturn(t *testing.T) {
	uniques := []string{
		"what", "shall", "we", "do", "with", "the",
		"drunken", "sailor",
	}

	dups := []string{
		"what", "shall", "we", "do", "we", "with", "the",
		"with", "drunken", "sailor", "sailor",
	}

	s0 := New()

	for _, u := range uniques {
		if !s0.Put(u) {
			t.Errorf("reported not added when should have been added")
		}
	}

	s1 := New()

	if !s1.Put(dups...) {
		t.Error("reported not all added when all should have been added")
	}
}

func TestStringSet_PutALot(t *testing.T) {
	s := New()
	toPut := 100000

	for i := 0; i < toPut; i += 1 {
		v0 := fmt.Sprintf("v0=%v", i)

		v1 := []string{
			fmt.Sprintf("v1[0]=%v", i),
			fmt.Sprintf("v1[1]=%v", i),
		}

		v2 := []string{
			fmt.Sprintf("v2[0]=%v", i),
			fmt.Sprintf("v2[1]=%v", i),
			fmt.Sprintf("v2[3]=%v", i),
		}

		all := []string{
			v0, v1[0], v1[1], v2[0], v2[1], v2[2],
		}

		if !s.Put(v0) {
			t.Error("reported not added when should have been added")
		}

		if !s.Put(v1...) {
			t.Error("reported not added when should have been added")
		}

		if !s.Put(v2...) {
			t.Error("reported not added when should have been added")
		}

		for _, a := range all {
			if !s.Contains(a) {
				t.Error("reported not contained when should be contained")
			}

			if s.Put(a) {
				t.Error("reported added when should not have been added")
			}
		}
	}
}

func TestStringSet_Strings(t *testing.T) {
	dups := []string{
		"what", "shall", "we", "do", "we", "with", "the",
		"with", "drunken", "sailor", "sailor", "early",
		"so", "very", "early", "in", "the", "morning",
	}

	s := NewWith(dups...)
	slice := s.Strings()

	for _, d := range dups {
		found := false

		for _, sl := range slice {
			if d == sl {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("missing %v in slice", d)
		}
	}
}

func TestStringSet_Readme(t *testing.T) {
	ss := New()

	// Add some strings. Notice that "the" appears twice.
	ss.Put("the", "less", "I", "know", "the", "better")

	// Returns 5.
	if ss.Len() != 5 {
		t.Error("bad len")
	}

	// Returns true
	if !ss.Contains("better") {
		t.Error("missing string")
	}

	// Remove some strings.
	if !ss.Remove("the", "less") {
		t.Error("should have removed all of those")
	}

	// Now returns 3.
	if ss.Len() != 3 {
		t.Error("bad len after call to Remove")
	}

	// Returns {"I", "know", "better"}
	slice := ss.Strings()
	expected := []string{"I", "know", "better"}
	if len(slice) != 3 {
		t.Error("bad len of slice")
	}
	for _, e := range expected {
		found := false

		for _, s := range slice {
			if e == s {
				found = true
			}
		}

		if !found {
			t.Errorf("missing %v in slice", e)
		}
	}
}
