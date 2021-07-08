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

func MultiSearchQuery(field []string, value string) map[string]interface{} {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  value,
				"fields": field,
			},
		},
	}

	return query
}

func UpdateQuery(key string, value string) map[string]interface{} {
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			key: value,
		},
	}
	return query
}
