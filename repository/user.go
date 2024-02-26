package repository

import (
	"context"
	"github.com/JohnKucharsky/echo_pgx/models"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func CreateUser(pool *pgxpool.Pool, body serializer.UserBody) (
	*models.User,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx, `
        INSERT INTO users (first_name, last_name, created_at, updated_at)
        VALUES (@first_name, @last_name, @created_at, @updated_at)
        RETURNING id, first_name,last_name, created_at, updated_at`,
		pgx.NamedArgs{
			"first_name": body.FirstName,
			"last_name":  body.LastName,
			"created_at": time.Now().Local(),
			"updated_at": time.Now().Local(),
		},
	)
	if err != nil {
		return nil, err
	}

	user, errA := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[models.User],
	)
	if errA != nil {
		return nil, errA
	}

	return &user, nil
}

func GetUsers(pool *pgxpool.Pool) (
	[]*models.User,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx, `select * from users`,
	)
	if err != nil {
		return nil, err
	}

	userRes, errA := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[models.User],
	)

	if errA != nil {

		return nil, errA
	}

	return userRes, nil
}

func GetOneUser(pool *pgxpool.Pool, id int) (
	*models.User,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`select * from users where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	userRes, errA := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.User],
	)

	if errA != nil {

		return nil, errA
	}

	return userRes, nil
}

func UpdateUser(pool *pgxpool.Pool, body serializer.UserBody, id int) (
	*models.User,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`UPDATE users SET 
                updated_at = @updated_at,
                first_name = @first_name,
    			last_name = @last_name 
             WHERE id = @id returning id,created_at,updated_at,first_name,last_name`,
		pgx.NamedArgs{
			"id":         id,
			"updated_at": time.Now().Local(),
			"first_name": body.FirstName,
			"last_name":  body.LastName,
		},
	)
	if err != nil {
		return nil, err
	}

	userRes, errA := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.User],
	)

	if errA != nil {

		return nil, errA
	}

	return userRes, nil
}

func DeleteUser(pool *pgxpool.Pool, id int) (
	*models.User,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`delete from users where id = @id returning id,created_at,updated_at,first_name,last_name`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	userRes, errA := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.User],
	)

	if errA != nil {

		return nil, errA
	}

	return userRes, nil
}
