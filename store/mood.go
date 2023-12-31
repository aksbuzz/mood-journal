package store

import (
	"context"
	"fmt"

	"github.com/aksbuzz/mood-journal/api"
)

func (s *Store) CreateMood(ctx context.Context, create *api.CreateMood) (*api.CreateMood, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO moods (
			mood,
			user_id,
			description,
			date
		)
		VALUES (?, ?, ?, ?)
		RETURNING id, mood, user_id, description, date
	`

	if err := tx.QueryRowContext(
		ctx, query, create.Mood, create.UserId, create.Description, create.Date,
	).Scan(
		&create.ID, &create.Mood, &create.UserId, &create.Description, &create.Date,
	); err != nil {
		return nil, FormatError(err)
	}
	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}
	response := create
	return response, nil
}

func (s *Store) ListMoods(ctx context.Context, find *api.FindMood) ([]*api.CreateMood, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	query := `SELECT * FROM moods`
	conditions := []string{}

	if v := find.UserId; v != nil {
		conditions = append(conditions, fmt.Sprintf("user_id = '%d'", *v))
	}
	if find.Mood != "" {
		conditions = append(conditions, fmt.Sprintf("mood = '%s'", find.Mood))
	}
	if find.TimeRange == api.Week {
		conditions = append(conditions, "date >= date('now', '-7 days')")
	} else if find.TimeRange == api.Month {
		conditions = append(conditions, "date >= date('now', '-30 days')")
	} else if find.TimeRange == api.Year {
		conditions = append(conditions, "date >= date('now', '-365 days')")
	}

	if len(conditions) > 0 {
		query += " WHERE " + joinStrings(conditions, " AND ")
	}

	query += " ORDER BY date DESC"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()
	moodsList := make([]*api.CreateMood, 0)
	for rows.Next() {
		var mood api.CreateMood
		if err := rows.Scan(
			&mood.ID, &mood.Mood, &mood.UserId, &mood.Description, &mood.Date,
		); err != nil {
			return nil, FormatError(err)
		}
		moodsList = append(moodsList, &mood)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return moodsList, nil
}

func joinStrings(slice []string, separator string) string {
	returnSlice := ""
	for i, s := range slice {
		if i > 0 {
			returnSlice += separator
		}
		returnSlice += s
	}
	return returnSlice
}
