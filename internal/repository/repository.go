package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/nutochk/ef-test/internal/models"
)

type Repository interface {
	Create(p *models.PersonInfo) (*models.PersonInfo, error)
	Update(id int, i *models.Info) (*models.PersonInfo, error)
	Delete(id int) (bool, error)
	GetById(id int) (*models.PersonInfo, error)
	//TODO more get
}

type repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) *repo {
	return &repo{db: db}
}

func (r *repo) Create(p *models.PersonInfo) (*models.PersonInfo, error) {
	var exist bool
	err := r.db.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM people WHERE name = $1 AND surname = $2)`, p.Name, p.Surname).Scan(&exist)
	if err != nil {
		return nil, ErrCheckExistence(err)
	}
	if !exist {
		return nil, ErrAlreadyExist
	}

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return nil, ErrBeginTransaction(err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `INSERT INTO people (name,surname, patronymic) VALUES ($1, $2, $3)`, p.Name, p.Surname, p.Patronymic)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into people table: %w", err)
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO info (age, gender, nationality) VALUES ($1, $2, $3)`, p.Age, p.Gender, p.Nationality)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into info table: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, ErrCommitTransaction(err)
	}
	return p, nil
}

func (r *repo) Update(id int, i *models.Info) (*models.PersonInfo, error) {
	exist, err := checkExistence(r, id)
	if err != nil {
		return nil, ErrCheckExistence(err)
	}
	if !exist {
		return nil, ErrNotExist
	}

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return nil, ErrBeginTransaction(err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `UPDATE info SET age =$1, gender = $2, nationality = $3 WHERE id = $4`, i.Age, i.Gender, i.Nationality, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update info table: %w", err)
	}

	var p models.PersonInfo
	p.Age = i.Age
	p.Gender = i.Gender
	p.Nationality = i.Nationality
	err = tx.QueryRow(context.Background(), `SELECT name, surname, patronymic FROM people WHERE id = $1`, id).Scan(&p.Name, &p.Surname, &p.Patronymic)
	if err != nil {
		return nil, ErrDatabase(err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, ErrCommitTransaction(err)
	}
	return &p, nil
}

func (r *repo) Delete(id int) (bool, error) {
	exist, err := checkExistence(r, id)
	if err != nil {
		return false, ErrCheckExistence(err)
	}
	if !exist {
		return false, ErrNotExist
	}

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return false, ErrBeginTransaction(err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), `DELETE FROM people WHERE id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete from people table: %w", err)
	}

	_, err = tx.Exec(context.Background(), `DELETE FROM info WHERE person_id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete from info table: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return false, ErrCommitTransaction(err)
	}
	return true, nil
}

func (r *repo) GetById(id int) (*models.PersonInfo, error) {
	exist, err := checkExistence(r, id)
	if err != nil {
		return nil, ErrCheckExistence(err)
	}
	if !exist {
		return nil, ErrNotExist
	}

	var p models.PersonInfo
	query := `SELECT p.name, p.surname, p.patronymic, i.age, i.gender, i.nationality 
		FROM people p 
		JOIN info i ON p.id = i.person_id
		WHERE p.id = $1`
	err = r.db.QueryRow(context.Background(), query, id).Scan(&p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.Nationality)
	if err != nil {
		return nil, ErrDatabase(err)
	}
	return &p, nil
}

func checkExistence(r *repo, id int) (bool, error) {
	var exist bool
	err := r.db.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM people WHERE id = $1 )`, id).Scan(&exist)
	return exist, err
}
