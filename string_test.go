package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	cases := map[string]string{
		"":                         "",
		"foobar":                   "foobar",
		"&FOO:BAR$BAZ":             "fooBarBaz",
		"fooBar":                   "fooBar",
		"FOObar":                   "foObar",
		"$foo%":                    "foo",
		"   $#$Foo   22    bar   ": "foo22Bar",
		"Foo-#1😄$_%^&*(1bar":       "foo11Bar",
	}

	for k, v := range cases {
		assert.Equal(v, CamelCase(k))
	}
}

func TestCapitalize(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	cases := map[string]string{
		"":        "",
		"Foo":     "Foo",
		"_foo":    "_foo",
		"foobar":  "Foobar",
		"fooBar":  "Foobar",
		"foo Bar": "Foo bar",
		"foo-bar": "Foo-bar",
		"$foo%":   "$foo%",
	}

	for k, v := range cases {
		assert.Equal(v, Capitalize(k))
	}
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "foo_bar",
		"--Foo---Bar-":             "foo_bar",
		"Foo Bar-":                 "foo_bar",
		"foo_Bar":                  "foo_bar",
		"fooBar":                   "foo_bar",
		"FOOBAR":                   "foobar",
		"FOO_BAR":                  "foo_bar",
		"__FOO_BAR__":              "foo_bar",
		"$foo@Bar":                 "foo_bar",
		"   $#$Foo   22    bar   ": "foo_22_bar",
		"Foo-#1😄$_%^&*(1bar":       "foo_1_1_bar",
	}

	for k, v := range cases {
		assert.Equal(v, SnakeCase(k))
	}
}

func TestUpperSnakeCase(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "FOO_BAR",
		"--Foo---Bar-":             "FOO_BAR",
		"Foo Bar-":                 "FOO_BAR",
		"foo_Bar":                  "FOO_BAR",
		"fooBar":                   "FOO_BAR",
		"FOOBAR":                   "FOOBAR",
		"FOO_BAR":                  "FOO_BAR",
		"__FOO_BAR__":              "FOO_BAR",
		"$foo@Bar":                 "FOO_BAR",
		"   $#$Foo   22    bar   ": "FOO_22_BAR",
		"Foo-#1😄$_%^&*(1bar":       "FOO_1_1_BAR",
	}

	for k, v := range cases {
		assert.Equal(v, UpperSnakeCase(k))
	}
}

func TestReverseString(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	assert.Equal("cba", ReverseString("abc"))
	assert.Equal("54321", ReverseString("12345"))
}
