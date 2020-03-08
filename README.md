stringset
=========

Tiny Go package that provides `StringSet`, a set of `string`s.
No strings attached!

Usage
-----

	ss := stringset.New()

	// Add some strings. Notice that "the" appears twice.
	ss.Put("the", "less", "I", "know", "the", "better")

	// Returns 5.
	fmt.Println(ss.Len())

	// Returns true.
	ss.Contains("better")

	// Remove some strings.
	ss.Remove("the", "less")

	// Now returns 3.
	fmt.Println(ss.Len())

	// Returns {"I", "know", "better"}.
	fmt.Println(ss.Strings())

Or just read the `godoc`.
