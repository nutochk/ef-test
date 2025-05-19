package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/nutochk/ef-test/internal/models"
)

type Repository interface {
	Create(p *models.PersonInfo) (int, error)
	Update(id int, i *models.Person) (*models.PersonInfo, error)
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

func (r *repo) Create(p *models.PersonInfo) (int, error) {
	//var exist bool
	//err := r.db.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM people WHERE name = $1 AND surname = $2)`, p.Name, p.Surname).Scan(&exist)
	//if err != nil {
	//	return nil, ErrCheckExistence(err)
	//}
	//if !exist {
	//	return nil, ErrAlreadyExist
	//}

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return 0, ErrBeginTransaction(err)
	}
	defer tx.Rollback(context.Background())

	var id int
	err = tx.QueryRow(context.Background(), `INSERT INTO people (name,surname, patronymic) VALUES ($1, $2, $3) RETURNING id`, p.Name, p.Surname, p.Patronymic).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into people table: %w", err)
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO info (person_id, age, gender, gender_probability) VALUES ($1, $2, $3, $4)`, id, p.Age, p.Gender, p.GenderProbability)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into info table: %w", err)
	}

	for _, n := range p.Nationality {
		_, err = tx.Exec(context.Background(), `INSERT INTO countries (person_id, nationality, probability) VALUES ($1, $2, $3)`, id, n.CountryId, n.Probability)
		if err != nil {
			return 0, fmt.Errorf("failed to insert into info table: %w", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return 0, ErrCommitTransaction(err)
	}
	return id, nil
}

func (r *repo) Update(id int, p *models.Person) (*models.PersonInfo, error) {
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

	_, err = tx.Exec(context.Background(), `UPDATE people SET name =$1, surname = $2, patronymic = $3 WHERE id = $4`, p.Name, p.Surname, p.Patronymic, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update info table: %w", err)
	}

	var pi models.PersonInfo
	pi.Name = p.Name
	pi.Surname = p.Surname
	pi.Patronymic = p.Patronymic
	err = tx.QueryRow(context.Background(), `SELECT age, gender, gender_probability FROM info WHERE person_id = $1`, id).Scan(&pi.Age, &pi.Gender, &pi.GenderProbability)
	if err != nil {
		return nil, ErrDatabase(err)
	}

	rows, err := tx.Query(context.Background(), `SELECT nationality, probability FROM countries WHERE person_id = $1`, id)
	if err != nil {
		return nil, ErrDatabase(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Country
		err = rows.Scan(&c.CountryId, &c.Probability)
		if err != nil {
			return nil, ErrDatabase(err)
		}
		pi.Nationality = append(pi.Nationality, c)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, ErrCommitTransaction(err)
	}
	return &pi, nil
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

	_, err = tx.Exec(context.Background(), `DELETE FROM countries WHERE person_id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete from info table: %w", err)
	}

	_, err = tx.Exec(context.Background(), `DELETE FROM info WHERE person_id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete from info table: %w", err)
	}

	_, err = tx.Exec(context.Background(), `DELETE FROM people WHERE id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete from people table: %w", err)
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
	query := `SELECT p.name, p.surname, p.patronymic, i.age, i.gender, i.gender_probability 
		FROM people p 
		JOIN info i ON p.id = i.person_id
		WHERE p.id = $1`
	err = r.db.QueryRow(context.Background(), query, id).Scan(&p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &p.GenderProbability)
	if err != nil {
		return nil, ErrDatabase(err)
	}

	rows, err := r.db.Query(context.Background(), `SELECT nationality, probability FROM countries WHERE person_id = $1`, id)
	if err != nil {
		return nil, ErrDatabase(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Country
		err = rows.Scan(&c.CountryId, &c.Probability)
		if err != nil {
			return nil, ErrDatabase(err)
		}
		p.Nationality = append(p.Nationality, c)
	}

	return &p, nil
}

func checkExistence(r *repo, id int) (bool, error) {
	var exist bool
	err := r.db.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM people WHERE id = $1 )`, id).Scan(&exist)
	return exist, err
}
