// Code generated by goderive DO NOT EDIT.

package sort

import (
	"sort"
	"strings"
)

func deriveCompare(this, that []*MyStruct) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if len(this) != len(that) {
		if len(this) < len(that) {
			return -1
		}
		return 1
	}
	for i := 0; i < len(this); i++ {
		if c := deriveComparePtrToMyStruct(this[i], that[i]); c != 0 {
			return c
		}
	}
	return 0
}

func deriveSort(list []*MyStruct) []*MyStruct {
	sort.Slice(list, func(i, j int) bool { return deriveComparePtrToMyStruct(list[i], list[j]) < 0 })
	return list
}

func deriveComparePtrToMyStruct(this, that *MyStruct) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if c := deriveCompareint64(this.Int64, that.Int64); c != 0 {
		return c
	}
	if c := deriveComparePtrTostring(this.StringPtr, that.StringPtr); c != 0 {
		return c
	}
	return 0
}

func deriveCompareint64(this, that int64) int {
	if this != that {
		if this < that {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

func deriveComparePtrTostring(this, that *string) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	return deriveComparestring(*this, *that)
}

func deriveComparestring(this, that string) int {
	return strings.Compare(this, that)
}
