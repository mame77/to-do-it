package score

import "database/sql"

type Repository interface {
	GetMotivationByUserID(userID string) (*Motivation, error)
	UpdateMotivation(motivation *Motivation) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetMotivationByUserID(userID string) (*Motivation, error) {
	query := `SELECT user_id, points, rank, level FROM motivation WHERE user_id = ?`

	motivation := &Motivation{}
	err := r.db.QueryRow(query, userID).Scan(&motivation.UserID, &motivation.Points, &motivation.Rank, &motivation.Level)

	if err == sql.ErrNoRows {
		// ユーザーのモチベーション情報が存在しない場合、初期値を作成
		motivation = &Motivation{
			UserID: userID,
			Points: 0,
			Rank:   "Bronze",
			Level:  1,
		}
		// DBに初期値を挿入
		insertQuery := `INSERT INTO motivation (user_id, points, rank, level) VALUES (?, ?, ?, ?)`
		_, err := r.db.Exec(insertQuery, motivation.UserID, motivation.Points, motivation.Rank, motivation.Level)
		if err != nil {
			return nil, err
		}
		return motivation, nil
	}

	if err != nil {
		return nil, err
	}

	return motivation, nil
}

func (r *postgresRepository) UpdateMotivation(motivation *Motivation) error {
	query := `UPDATE motivation SET points = ?, rank = ?, level = ? WHERE user_id = ?`
	_, err := r.db.Exec(query, motivation.Points, motivation.Rank, motivation.Level, motivation.UserID)
	return err
}
