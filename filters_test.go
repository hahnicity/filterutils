package filterutils_test

import "github.com/hahnicity/filterutils"
import "reflect"
import "sort"
import "testing"

var testStringInSortedSlice = []struct {
    name     string
    a        []string
    x        string
    expected bool
} {
    {"Has element", []string{"foo", "bar", "baz", "blah", "belp"}, "blah", true},
    {"Does not have element", []string{"tar", "var", "car", "a", "b", "d"}, "me", false}, 
}

func TestStringInSortedSlice(t *testing.T) {
    for _, test := range testStringInSortedSlice {
        sort.Strings(test.a)
        i := filterutils.StringInSortedSlice(test.a, test.x)
        if i != test.expected {
            // XXX Need to figure out the actual types for % to pass in for string formatting
            t.Errorf("Test '%s' has failed. Expected: %s, Actual %s", test.name, test.expected, i)  
        }
    }
}

func TestFilter(t *testing.T) {
    a := []string{"foo", "foo", "bar"}
    ex := []string{"foo", "foo"}
    i := filterutils.Filter(a, func(i int) bool { return a[i] == "foo"}) 
    if !reflect.DeepEqual(i, ex) {
        // XXX Arg... string formatting
        t.Errorf("Filter test has failed. Expected %d, Actual %d", ex, i)
    }
}
