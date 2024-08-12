package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Todos struct {
	ID          int
	Description string
	Completed   bool
}

type TodosModel struct {
	DB *pgxpool.Conn
}

func (t *TodosModel) GetAll() ([]Todos, error) {
	query := `SELECT id, description, completed FROM todos ORDER by id`

	rows, err := t.DB.Query(context.Background(), query)
	if err != nil {
		return []Todos{}, err
	}

	defer rows.Close()

	var todos []Todos

	for rows.Next() {
		var t Todos

		err = rows.Scan(&t.ID, &t.Description, &t.Completed)
		if err != nil {
			return []Todos{}, err
		}

		todos = append(todos, t)
	}

	return todos, nil

}

func (t *TodosModel) GetID(id int) (Todos, error) {
	if id < 1 {
		return Todos{}, errors.New("record not found")
	}

	query := "SELECT id, description, completed FROM todos WHERE id = $1"

	row := t.DB.QueryRow(context.Background(), query, id)

	var todo Todos

	err := row.Scan(&todo.ID, &todo.Description, &todo.Completed)
	if err != nil {
		return Todos{}, err
	}

	return todo, nil
}

func (t *TodosModel) MarkComplete(todo Todos) error {
	stmt := `
						UPDATE todos SET completed = $1 
						WHERE id = $2
						RETURNING id, description, completed
						`
	args := []any{
		todo.Completed,
		todo.ID,
	}

	err := t.DB.QueryRow(context.Background(), stmt, args...).Scan(&todo.ID, &todo.Description, &todo.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodosModel) UpdateDescription(todo Todos) error {
	if todo.ID < 1 {
		return errors.New("record not found")
	}

	stmt := `
						UPDATE todos SET description = $1
						WHERE id = $2
						RETURNING id, description, completed
					`

	args := []any{
		todo.Description,
		todo.ID,
	}

	err := t.DB.QueryRow(context.Background(), stmt, args...).Scan(&todo.ID, &todo.Description, &todo.Completed)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodosModel) Insert(todo Todos) error {
	stmt := `
						INSERT INTO todos (description, completed)
						VALUES ($1, $2)
						RETURNING id, description, completed
					`
	args := []any{
		todo.Description,
		todo.Completed,
	}

	err := t.DB.QueryRow(context.Background(), stmt, args...).Scan(&todo.ID, &todo.Description, &todo.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (t *TodosModel) Delete(todo Todos) error {
	if todo.ID < 1 {
		return errors.New("record not found")
	}

	query := `
						DELETE FROM todos WHERE id = $1
					`

	res, err := t.DB.Exec(context.Background(), query, todo.ID)
	if err != nil {
		return err
	}

	row := res.RowsAffected()
	if row == 0 {
		return errors.New("record not found")
	}

	return nil

}
