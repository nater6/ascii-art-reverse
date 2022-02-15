package asciiartreverse

func Createmap(slice []string) map[int][]string {
	mapslice := make(map[int][]string)
	ascii := 32
	count := 0
	split := 0
	empty := []string{}

	for split < len(slice)-9 {
		split = count * 9
		for i := split; i < split+9; i++ {
			empty = append(empty, slice[i])
		}
		mapslice[ascii] = empty
		empty = []string{}
		ascii++
		count++
	}

	return mapslice
}
