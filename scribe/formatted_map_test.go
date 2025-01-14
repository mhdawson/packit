package scribe_test

import (
	"testing"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testFormattedMap(t *testing.T, context spec.G, it spec.S) {
	var Expect = NewWithT(t).Expect

	context("String", func() {
		it("returns a formatted string representation of the map", func() {
			Expect(scribe.FormattedMap{
				"third":  3,
				"first":  1,
				"second": 2,
			}.String()).To(Equal("first  -> \"1\"\nsecond -> \"2\"\nthird  -> \"3\""))
		})
	})

	context("NewFormattedMapFromEnvironment", func() {
		context("for all packit env var operations", func() {
			it("prints the env in a well formatted map", func() {
				Expect(scribe.NewFormattedMapFromEnvironment(packit.Environment{
					"OVERRIDE.override": "some-value",
					"DEFAULT.default":   "some-value",
					"PREPEND.prepend":   "some-value",
					"PREPEND.delim":     ":",
					"APPEND.append":     "some-value",
					"APPEND.delim":      ":",
					"BOTH.append":       "appended-value",
					"BOTH.delim":        ":",
					"BOTH.prepend":      "prepended-value",
				})).To(Equal(scribe.FormattedMap{
					"OVERRIDE": "some-value",
					"DEFAULT":  "some-value",
					"PREPEND":  "some-value:$PREPEND",
					"APPEND":   "$APPEND:some-value",
					"BOTH":     "prepended-value:$BOTH:appended-value",
				}))
			})
		})
		context("for a standard string map", func() {
			it("prints the env in a well formatted map", func() {
				Expect(scribe.NewFormattedMapFromEnvironment(map[string]string{
					"SOME_ENV_VAR":       "some-value",
					"SOME_OTHER_ENV_VAR": "some-other-value",
				})).To(Equal(scribe.FormattedMap{
					"SOME_ENV_VAR":       "some-value",
					"SOME_OTHER_ENV_VAR": "some-other-value",
				}))
			})
		})
	})
}
