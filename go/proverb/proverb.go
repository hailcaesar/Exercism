// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	ans := make([]string, len(rhyme))

	if len(ans) >= 2 {
		for idx := 0; idx < len(ans)-1; idx++ {
			ans[idx] = "For want of a " + rhyme[idx] + " the " + rhyme[idx+1] + " was lost."
		}
	}
	if len(rhyme) > 0 {
		ans[len(ans)-1] = "And all for the want of a " + rhyme[0] + "."
	}
	return ans
}
