package coll

import (
	"log"
	"strconv"
)

type Comparable interface {
    CompareTo(other Comparable) int
}

type ComparableString string
type ComparableInt int

func (cs ComparableString) CompareTo(other Comparable) int {
    switch v := other.(type) {
    case ComparableString:
        if cs < v {
            return -1
        } else if cs > v {
            return 1
        }
        return 0
    case ComparableInt:
        csInt, err := strconv.Atoi(string(cs))
        if err != nil {
            log.Println("Error: Comparing different types")
            return 0
        }
        if csInt < int(v) {
            return -1
        } else if csInt > int(v) {
            return 1
        }
        return 0
    default:
        log.Println("Error: Comparing different types")
        return 0
    }
}

func (ci ComparableInt) CompareTo(other Comparable) int {
    switch v := other.(type) {
    case ComparableInt:
        if ci < v {
            return -1
        } else if ci > v {
            return 1
        }
        return 0
    case ComparableString:
        vInt, err := strconv.Atoi(string(v))
        if err != nil {
            log.Println("Error: Comparing different types")
            return 0
        }
        if int(ci) < vInt {
            return -1
        } else if int(ci) > vInt {
            return 1
        }
        return 0
    default:
        log.Println("Error: Comparing different types")
        return 0
    }
}