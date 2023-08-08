package repository

import (
	"database/sql"
	"suggester/db"
	"suggester/models"
)

type SuggestRepoInterface interface {
	GetOne(query string) (models.FullSuggest, error)
	Get(query string, limit int) ([]models.Suggest, error)
	Add(suggest models.Suggest) error
	IncrementAmount(id int) error
}

type SuggestRepo struct {
	db *sql.DB
}

func GetSuggestRepo() (SuggestRepo, error) {
	con, err := db.GetDBConnection()
	if err != nil {
		return SuggestRepo{con}, err
	} else {
		return SuggestRepo{con}, nil
	}
}

func (repo SuggestRepo) GetOne(query string) (models.FullSuggest, error) {
	var suggest models.FullSuggest

	row := repo.db.QueryRow(`SELECT "query_id", "query" FROM "query" WHERE query = $1`, query)

	if err := row.Scan(&suggest.Id, &suggest.Text); err != nil {
		return suggest, err
	}
	return suggest, nil
}

func (repo SuggestRepo) Get(query string, limit int) ([]models.Suggest, error) {
	suggesters := make([]models.Suggest, 0, limit)

	rows, err := repo.db.Query(
		`SELECT "query" FROM "query" 
               WHERE query LIKE $1 || '%' 
               ORDER BY "amount" DESC 
               LIMIT $2`, query, limit,
	)
	if err != nil {
		return suggesters, nil
	}

	for rows.Next() {
		var suggest models.Suggest
		if err := rows.Scan(&suggest.Text); err != nil {
			return suggesters, nil
		}
		suggesters = append(suggesters, suggest)
	}

	return suggesters, nil
}

func (repo SuggestRepo) Add(suggest models.Suggest) error {
	_, err := repo.db.Exec("INSERT INTO query(query, amount) VALUES ($1, 1)", suggest.Text)
	return err
}

func (repo SuggestRepo) IncrementAmount(id int) error {
	_, err := repo.db.Exec("UPDATE query SET amount = amount + 1 WHERE query_id = $1", id)
	return err
}
