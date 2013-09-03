package filterutils

import "strings"

// Taken from the golang source code and modified 
//Do a binary search over a slice for an element.
//
//For large sets a bloom filter is more effective
func Search(n int, f func(int) bool, g func(int) bool) bool {
    // Define f(-1) == false and f(n) == true.
    // Invariant: f(i-1) == false, f(j) == true.
    i, j := 0, n
    for i < j {
        h := i + (j-i)/2 // avoid overflow when computing h
        if g(h) {
            return true
        } else if !f(h) { // i ≤ h < j
            i = h + 1 // preserves f(i-1) == false
        } else {
            j = h // preserves f(j) == true
        }
    }
    // i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
    return false
}

 
//Returns true if a string is in a slice, otherwise return false.
//
//Must be implemented with a sorted slice. Runtime is O(lg(n))
func StringInSortedSlice(a []string, x string) bool {
    return Search(
        len(a), 
        func(i int) bool { return a[i] > x}, 
        func(i int) bool { return a[i] == x},
    )
}

// Iterate over a slice of strings looking for all strings with a suffix suf
func HasSuffix(a []string, suf string) []string {
    return Filter(
        a, 
        func(i int) bool { return strings.HasSuffix(a[i], suf) },
    )
}
 
//Iterate over a slice and pick only elements which hold under function f
//    
//Runtime is O(n)
func Filter(a []string, f func(int) bool) []string {
    b := make([]string, 0)
    for i, _ := range a {
        if f(i) { 
            b = append(b, a[i])
        }
    }
    return b
}
