package recommender

import (
	"strings"
)

type replace struct {
	new string
	old string
}

func getWords(body string) []string {
	var result []string

	body = replacer(body)
	result = strings.Split(body, " ")
	return result
}

func replacer(body string) string {
	var replaces = []replace{
		replace{
			new: "",
			old: ",",
		},
		replace{
			new: "",
			old: ".",
		},
		replace{
			old: "“",
			new: "",
		},
		replace{
			old: ")",
			new: "",
		},
		replace{
			old: "(",
			new: "",
		},
		replace{
			old: "!",
			new: "",
		},
		replace{
			old: " — ",
			new: " ",
		},
		replace{
			old: "”",
			new: "",
		},
		replace{
			old: "?",
			new: "",
		},
	}

	for _, r := range replaces {
		body = strings.Replace(body, r.old, r.new, -1)
	}
	return body
}
