package sqlf

import (
	"fmt"
	"strings"
)

const sqlfPrefix = "-- name:"

// Load parses the input string and extracts SQL queries into a map.
// It expects the input to be a string containing SQL queries separated by a prefix.
// The queries are identified by lines prefixed with "---- sqlfName: $name".
// Comment lines starting with "--" are ignored.
//
// Parameters:
// - in: the input string containing the SQL queries
//
// Returns:
// - a map containing the extracted SQL queries, where the key is the name and the value is the query string
// The intended use is to call Load in an init script, and assign the output to a package level variable that then can be
// used then in the queries.
// //go:embed mySqlFile.sql
// var mySqlString string
// var mySqlQueries map[string]string
//
//	func init() {
//	 mySqlQueries = sqlf.Load(mySqlString)
//	}
//
//	func retrieve(....) .... {
//	 err = db.Get(&id, mySqlQueries["queryName"], arg)
//	}
//
// The idea behind it is to have sql queries separated from code to take advantage of ide tools, and also to make navigation
// in the project easier.
func Load(in string) map[string]string {
	queries := make(map[string]string)
	var currentName string
	for _, line := range strings.SplitAfter(in, "\n") {
		if strings.HasPrefix(line, sqlfPrefix) {
			nameFields := strings.Fields(strings.TrimSpace(strings.TrimPrefix(line, sqlfPrefix)))
			fmt.Println(nameFields, "hello")
			currentName = nameFields[0]
			continue
		}
		if strings.HasPrefix(line, "--") || strings.TrimSpace(line) == "" {
			continue
		}
		queries[currentName] += line
	}
	return queries
}
