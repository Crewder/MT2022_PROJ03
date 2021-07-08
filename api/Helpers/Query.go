package Helpers

import "github.com/MT2022_PROJ03/Models"

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

func UpdateQuery(input Models.Book) map[string]interface{} {

	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"Abstract": input.Abstract,
			"Title":    input.Title,
			"Author":   input.Author,
		},
	}
	return query
}
