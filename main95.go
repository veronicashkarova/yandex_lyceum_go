package main

func DeleteLongKeys(m map[string]int) map[string]int {
	new := make(map[string]int)

	for key, _ := range m {
		if len(key) > 5 {
			new[key] = m[key]
		}
	}

	return new
}
