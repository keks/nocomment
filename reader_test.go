package nocomment

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	type testcase struct {
		in  string
		out string
		n   int
		err error
	}

	mkTest := func(tc testcase) func(*testing.T) {
		return func(t *testing.T) {
			a := assert.New(t)

			out, err := ioutil.ReadAll(NewReader(strings.NewReader(tc.in)))
			n := len(out)
			a.Equal(len(tc.out), n, "length mismatch")

			if err != nil {
				a.Error(tc.err, err, "unexpected error")
			} else {
				a.NoError(err, "unexpected error")
			}

			a.Equal(tc.out, string(out), "unexpected output")
		}
	}

	tcs := []testcase{
		{
			in:  "test\n#foo",
			out: "test\n",
			err: nil,
		},
		{
			in:  "#foo\n\n",
			out: "\n\n",
			err: nil,
		},
		{
			in:  `Foobar!

bar foo foo bar bar # (barfoo, foobar bar foo)
bar foo bar barbar, foo bar bar foo, bar foofoofoo.
bar foo foo bar #foobar, bar foofoo foo.`,
			out:  `Foobar!

bar foo foo bar bar 
bar foo bar barbar, foo bar bar foo, bar foofoofoo.
bar foo foo bar `,
			err: nil,
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i), mkTest(tc))
	}
}
