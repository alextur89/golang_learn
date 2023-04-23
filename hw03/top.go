package hw03frequencyanalysis

import "sort"
import "strings"

func Top10(s string) []string {
    counters := make(map[string]uint64) 
    for _, w := range(strings.Fields(s)) {
        counters[w] += 1
    }
    keys := make([]string, 0, len(counters))
    for key := range counters {
        keys = append(keys, key)
    }
    sort.SliceStable(keys, func(i, j int) bool{
        if counters[keys[i]] == counters[keys[j]] {
            return strings.Compare(keys[i],keys[j]) < 0
        }
        return counters[keys[i]] > counters[keys[j]]
    })
    result := make([]string, 0, 10)
    for i:=0; i < len(keys) && i < 10; i++ {
        if k := keys[i]; len(k) > 0 {
            result = append(result, keys[i])
        }
    }
	return result
}
