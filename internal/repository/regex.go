package repository

import "context"

const getRegexList = `SELECT id, word, regex, created_at
FROM dbo.regex
ORDER BY created_at DESC;`

func (q *Queries) GetRegexList(ctx context.Context) ([]Regex, error) {
	rows, err := q.db.QueryContext(ctx, getRegexList)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []Regex
	for rows.Next() {
		var r Regex
		if err := rows.Scan(
			&r.Id,
			&r.Word,
			&r.Regex,
			&r.CreatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

const insertRegex = `INSERT INTO dbo.regex (word, regex) VALUES (@word, @regex)`

func (q *Queries) InsertRegex(ctx context.Context, word string, regex string) {

}
