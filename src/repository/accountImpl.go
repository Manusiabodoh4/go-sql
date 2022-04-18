package repository

import (
	"context"
	"database/sql"

	"github.com/Manusiabodoh4/go-sql/src/entity"
)

type AccountRepoImpl struct {
	database *sql.DB
}

func NewAccountRepo(_database *sql.DB) AccountRepo {
	return &AccountRepoImpl{database: _database}
}

func (st *AccountRepoImpl) FindAll(ctx context.Context) ([]entity.AccountEntity, error) {

	script := "SELECT * FROM accounts"
	rows, err := st.database.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []entity.AccountEntity

	for rows.Next() {
		account := entity.AccountEntity{}
		rows.Scan(&account.Id, &account.Nama, &account.Email, &account.Password, &account.Age)
		res = append(res, account)
	}

	return res, nil

}
func (st *AccountRepoImpl) FindByEmail(ctx context.Context, email string) (entity.AccountEntity, error) {

	script := "SELECT * FROM accounts WHERE email = ? LIMIT 1"
	rows, err := st.database.QueryContext(ctx, script, email)

	var account entity.AccountEntity

	if err != nil {
		return account, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&account.Id, &account.Nama, &account.Email, &account.Password, &account.Age)
	}

	return account, nil

}
func (st *AccountRepoImpl) Insert(ctx context.Context, data entity.AccountEntity) (bool, error) {

	return true, nil
}
func (st *AccountRepoImpl) Update(ctx context.Context, data entity.AccountEntity) (bool, error) {
	return true, nil
}
func (st *AccountRepoImpl) DeleteById(ctx context.Context, id uint64) (bool, error) {
	return true, nil
}
