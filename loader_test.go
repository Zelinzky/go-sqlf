package sqlf

import (
	"reflect"
	"testing"
)

const testFile = `-- name: GetAuthor
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor
INSERT INTO authors (
    name, bio
-- this is a commentary it should not be outputed to the final query
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: UpdateAuthor
UPDATE authors
SET name = $2,

-- additional lines to be ignored
    bio = $3


WHERE id = $1;

-- name: DeleteAuthor
DELETE FROM authors
WHERE id = $1;`

func TestLoad(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "basic test",
			args: args{in: testFile},
			want: map[string]string{
				"GetAuthor": `SELECT * FROM authors
WHERE id = $1 LIMIT 1;
`,
				"ListAuthors": `SELECT * FROM authors
ORDER BY name;
`,
				"CreateAuthor": `INSERT INTO authors (
    name, bio
) VALUES (
             $1, $2
         )
    RETURNING *;
`,
				"UpdateAuthor": `UPDATE authors
SET name = $2,
    bio = $3
WHERE id = $1;
`,
				"DeleteAuthor": `DELETE FROM authors
WHERE id = $1;`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Load(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RESULT = %v, \n EXPECTED = %v", got, tt.want)
			}
		})
	}
}
