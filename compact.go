package piscine

func Compact(ptr *[]string) int {
	var compacted []string

	for _, str := range *ptr {
		if str != "" {
			compacted = append(compacted, str)
		}
	}

	*ptr = compacted

	return len(compacted)
}
