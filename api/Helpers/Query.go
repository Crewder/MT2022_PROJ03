package Helpers

func SearchQuery(field string, value string) map[string]interface{} {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				field: value,
			},
		},
	}

	return query
}

func MultiSearchQuery(value string) map[string]interface{} {

	match := []interface{}{
		map[string]interface{}{"match": map[string]string{"abstract": value}},
		map[string]interface{}{"match": map[string]string{"title": value}},
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": match,
			},
		},
	}

	return query
}
