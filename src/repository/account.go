package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepo struct {
	Database *sql.DB
}

func NewAccountRepo(_db *sql.DB) Repository {
	return &AccountRepo{Database: _db}
}

func (st *AccountRepo) Find(ctx context.Context, channel chan entity.TemplateChannelResponse) {
	query := "SELECT * FROM accounts"
	rows, err := st.Database.QueryContext(ctx, query)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: nil}
		return
	}
	defer rows.Close()
	var response []entity.AccountEntity
	var account entity.AccountEntity
	if !rows.Next() {
		channel <- entity.TemplateChannelResponse{Error: nil, Data: nil}
		return
	}
	rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
	response = append(response, account)
	for rows.Next() {
		rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
		response = append(response, account)
	}
	channel <- entity.TemplateChannelResponse{Error: nil, Data: response}
}
func (st *AccountRepo) FindWithParam(ctx context.Context, channel chan entity.TemplateChannelResponse, param string, args ...interface{}) {
	query := "SELECT * FROM accounts WHERE " + param
	rows, err := st.Database.QueryContext(ctx, query, args...)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: nil}
		return
	}
	defer rows.Close()
	var response []entity.AccountEntity
	var account entity.AccountEntity
	if !rows.Next() {
		channel <- entity.TemplateChannelResponse{Error: nil, Data: nil}
		return
	}
	rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
	response = append(response, account)
	for rows.Next() {
		rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
		response = append(response, account)
	}
	channel <- entity.TemplateChannelResponse{Error: nil, Data: response}
}
func (st *AccountRepo) InsertOne(ctx context.Context, channel chan entity.TemplateChannelResponse, args ...interface{}) {
	query := "INSERT INTO accounts (nama, email, password, age) VALUES ($1, $2, $3, $4)"
	_, err := st.Database.ExecContext(ctx, query, args...)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}
	channel <- entity.TemplateChannelResponse{Error: err, Data: true}
}
func (st *AccountRepo) InsertMany(ctx context.Context, channel chan entity.TemplateChannelResponse, value []map[string]interface{}) {
	query := "INSERT INTO accounts (nama, email, password, age) VALUES "
	var data []interface{}
	for _, row := range value {
		query += "(?, ?, ?, ?),"
		data = append(data, row["Nama"], row["Email"], row["Password"], row["Age"])
	}
	query = strings.TrimSuffix(query, ",")
	stmt, err := st.Database.PrepareContext(ctx, query)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}
	_, err = stmt.ExecContext(ctx, data...)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}
	channel <- entity.TemplateChannelResponse{Error: nil, Data: true}
}
func (st *AccountRepo) Update(ctx context.Context, channel chan entity.TemplateChannelResponse, set string, param string, args ...interface{}) {
	query := "UPDATE accounts SET " + set + " WHERE " + param
	_, err := st.Database.ExecContext(ctx, query, args...)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}
	channel <- entity.TemplateChannelResponse{Error: err, Data: true}
}
func (st *AccountRepo) Delete(ctx context.Context, channel chan entity.TemplateChannelResponse, param string, args ...interface{}) {
	query := "DELETE FROM accounts WHERE " + param
	_, err := st.Database.ExecContext(ctx, query, args...)
	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}
	channel <- entity.TemplateChannelResponse{Error: err, Data: true}
}
