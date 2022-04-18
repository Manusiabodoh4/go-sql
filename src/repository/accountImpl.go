package repository

import (
	"context"
	"database/sql"
	"sync"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepoImpl struct {
	database *sql.DB
}

func NewAccountRepo(_database *sql.DB) AccountRepo {
	return &AccountRepoImpl{database: _database}
}

func (st *AccountRepoImpl) FindAll(ctx context.Context, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse) {

	defer group.Done()

	group.Add(1)

	script := "SELECT * FROM accounts"
	rows, err := st.database.QueryContext(ctx, script)

	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: nil}
		return
	}

	defer rows.Close()

	var res []entity.AccountEntity

	if !rows.Next() {
		channel <- entity.TemplateChannelResponse{Data: nil, Error: nil}
		return
	}

	for rows.Next() {
		account := entity.AccountEntity{}
		rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
		res = append(res, account)
	}

	channel <- entity.TemplateChannelResponse{Data: res, Error: nil}

}
func (st *AccountRepoImpl) FindByEmail(ctx context.Context, email string, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse) {

	defer group.Done()

	group.Add(1)

	script := "SELECT * FROM accounts WHERE email = $1 LIMIT 1"
	rows, err := st.database.QueryContext(ctx, script, email)

	var account entity.AccountEntity

	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: account}
		return
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&account.Nama, &account.Email, &account.Password, &account.Age)
		channel <- entity.TemplateChannelResponse{Error: nil, Data: account}
		return
	}

	channel <- entity.TemplateChannelResponse{Error: nil, Data: nil}

}
func (st *AccountRepoImpl) Insert(ctx context.Context, data entity.AccountEntity, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse) {

	defer group.Done()

	group.Add(1)

	script := "INSERT INTO accounts (nama, email, password, age) VALUES ($1,$2,$3,$4)"
	_, err := st.database.ExecContext(ctx, script, data.Nama, data.Email, data.Password, data.Age)

	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}

	channel <- entity.TemplateChannelResponse{Error: nil, Data: true}

}
func (st *AccountRepoImpl) Update(ctx context.Context, data entity.AccountEntity, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse) {

	defer group.Done()

	group.Add(1)

	script := "UPDATE accounts SET nama = $1 , password = $2, age = $3"
	_, err := st.database.ExecContext(ctx, script, data.Nama, data.Password, data.Age)

	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}

	channel <- entity.TemplateChannelResponse{Error: nil, Data: true}

}
func (st *AccountRepoImpl) DeleteById(ctx context.Context, email string, group *sync.WaitGroup, channel chan entity.TemplateChannelResponse) {

	defer group.Done()

	group.Add(1)

	script := "DELETE FROM accounts WHERE email = ?"
	_, err := st.database.ExecContext(ctx, script, email)

	if err != nil {
		channel <- entity.TemplateChannelResponse{Error: err, Data: false}
		return
	}

	channel <- entity.TemplateChannelResponse{Error: nil, Data: true}

}
