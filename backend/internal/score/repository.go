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
	// TODO: DBからユーザーのモチベーション情報を取得 [cite: 79]
	// SELECT points, rank, level FROM user_motivation WHERE user_id = $1
	motivation := &Motivation{UserID: userID}
	// ... (db.QueryRowContext...)
	return motivation, nil
}

func (r *postgresRepository) UpdateMotivation(motivation *Motivation) error {
	// TODO: DBのモチベーション情報を更新 [cite: 77-78]
	// UPDATE user_motivation SET points = $1, rank = $2, level = $3 WHERE user_id = $4
	return nil
}
