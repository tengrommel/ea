package oone

func ThreeWords() string {
	threeWords := [3]string{
		"foo",
		"bar",
		"baz"}
	return threeWords[1]
}

func TenWords() string {
	tenWords := [10]string{"foo", "bar", "baz", "qux", "grault", "waldo", "plugh", "xyzzy", "thud", "spam"}
	return tenWords[6]
}
