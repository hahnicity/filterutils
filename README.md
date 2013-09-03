# Filterutils
=============

Filter tools for Go!

## Usage
========
        func Search(n int, f, g func(int) bool) bool

Execute a Binary Search for a specific item in a list matching function g. If the item is
not found, continue the search based on specifications of function f. Return true/false
depending on whether an item is in our data set

        func StringInSortedSlice(a []string, x string) bool

Execute a Binary Search for a specifc string x in a slice a. Return true/false depending
on whether our string x is contained in slice a.

        func HasSuffix(a []string, suf string) []string

Search over a slice `a` for all elements with the suffix `suf`. Return a slice of strings

        func Filter(a []string, f func(int) bool) []string

Search over a slice a for all elements satisfying a specific function f. Return a slice
of all strings matching the criteria specified in f.
