package piscine

func SortWordArr(a []string) {
	for i := range a {
		for j := i+1; j<len(a); j++ {
			if a[i] > a[j] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
}

// or
// func SortWordArr(a []string) {
// 	for i := range a {
// 		for j := range a[i:] {
// 			if a[i] > a[j+i] {
// 				a[i],a[j+i] = a[j+i],a[i]
// 			}
// 		}
// 	}
// }
