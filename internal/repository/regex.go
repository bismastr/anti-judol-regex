package repository

import (
	"context"
	"database/sql"
	"fmt"
)

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

const insertRegex = `
INSERT INTO dbo.regex (word, regex)
SELECT w.word, w.regex
FROM (VALUES (@p1, @p2)) AS w(word, regex)
WHERE NOT EXISTS (
    SELECT 1 FROM dbo.regex r 
    WHERE r.word = w.word
);`

func (q *Queries) InsertRegex(ctx context.Context, regex *Regex) error {
	_, err := q.db.ExecContext(ctx,
		insertRegex,
		sql.Named("p1", regex.Word),
		sql.Named("p2", regex.Regex),
	)

	if err != nil {
		return fmt.Errorf("failed to insert regexes: %w", err)
	}

	return nil
}

const wordExist = `SELECT CAST(CASE WHEN EXISTS(SELECT 1 FROM dbo.regex WHERE word = @word) THEN 1 ELSE 0 END AS BIT)`

func (q *Queries) WordExists(ctx context.Context, word string) (bool, error) {
	var exists bool
	err := q.db.QueryRowContext(ctx,
		wordExist,
		sql.Named("word", word),
	).Scan(&exists)

	return exists, err
}
