package filterutils_test

import "github.com/hahnicity/filterutils"
import "sort"
import "testing"

var tests = []struct {
    name     string
    a        []string
    x        string
    expected bool
} {
    {"Has element", []string{"foo", "bar", "baz", "blah", "belp"}, "blah", true},
    {"Does not have element", []string{"tar", "var", "car", "a", "b", "d"}, "me", false}, 
}

func TestStringInSortedSlice(t *testing.T) {
    for _, test := range tests {
        sort.Strings(test.a)
        i := filterutils.StringInSortedSlice(test.a, test.x)
        if i != test.expected {
            t.Errorf("Test '%s' has failed. Expected: %s, Actual %s", test.name, test.expected, i)       
        }
    }
}
