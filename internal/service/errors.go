package service

import "fmt"

func ErrRequest(e error) error {
	return fmt.Errorf("request error: %w", e)
}

func ErrResponse(e error) error {
	return fmt.Errorf("failed to read response: %w", e)
}

func ErrParsing(e error) error {
	return fmt.Errorf("failed to parse json: %w", e)
}
