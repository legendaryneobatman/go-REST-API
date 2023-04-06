package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "go-server-copy/models"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(userId int, list todo.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	var user_id int
	var title string
	var description string
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING *", listsTable)
	row := tx.QueryRow(createListQuery, userId, list.Title, list.Description)
	if err := row.Scan(&id, &user_id, &title, &description); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ListPostgres) GetAll(userId int) ([]todo.List, error) {
	var lists []todo.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", listsTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var list todo.List
		if err := rows.Scan(&list.Id, &list.UserId, &list.Title, &list.Description); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ListPostgres) GetById(listId string) (todo.List, error) {
	var list todo.List
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", listsTable)
	row := r.db.QueryRow(query, listId)
	if err := row.Scan(&list.Id, &list.UserId, &list.Title, &list.Description); err != nil {
		return todo.List{}, err
	}
	return list, nil
}

func (r *ListPostgres) Update(listId string, input todo.List) error {
	targetList, err := r.GetById(listId)
	if err != nil {
		return err
	}
	if input.Title != "" {
		targetList.Title = input.Title
	}
	if input.Description != "" {
		targetList.Description = input.Description
	}
	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2 WHERE id = $3", listsTable)
	row := r.db.QueryRow(query, targetList.Title, targetList.Description, listId)

	return row.Scan()
}
