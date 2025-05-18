package repository

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotExist     = errors.New("not exist")
)

func ErrCheckExistence(e error) error {
	return fmt.Errorf("failed to check existence: %w", e)
}

func ErrDatabase(e error) error {
	return fmt.Errorf("database error: %w", e)
}

func ErrBeginTransaction(e error) error {
	return fmt.Errorf("failed to begin transaction: %w", e)
}

func ErrCommitTransaction(e error) error {
	return fmt.Errorf("failed to commit transaction: %w", e)
}
