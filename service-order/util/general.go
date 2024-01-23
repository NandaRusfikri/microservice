package util



func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func RemoveDuplicateInt(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//func SplitOrderQuery(order string) (string, string) {
//	orderA := strings.Split(order, "|")
//	if len(orderA) > 1 {
//		if orderA[0] != "" && orderA[1] != "" {
//			return orderA[0], orderA[1]
//		}
//	}
//
//	return orderA[0], "ASC"
//}
