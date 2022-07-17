package cmd

func SearchByTags(tag string) ([]string, error) {
	tpTable, err := RecoverTagToPathTable(Tp)
	if err != nil {
		return []string{}, err
	}

	_, ok := tpTable.Table[tag]
	if ok {
		paths := []string{}
		for path := range tpTable.Table[tag] {
			paths = append(paths, path)
		}
		return paths, nil
	}

	return []string{}, nil
}
