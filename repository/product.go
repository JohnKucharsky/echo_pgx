package repository

import (
	"context"
	"github.com/JohnKucharsky/echo_pgx/models"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func CreateProduct(pool *pgxpool.Pool, body serializer.ProductBody) (
	*models.Product,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx, `
        INSERT INTO products (name, serial_number, created_at, updated_at)
        VALUES (@name, @serial_number, @created_at, @updated_at)
        RETURNING id, name,serial_number, created_at, updated_at`,
		pgx.NamedArgs{
			"name":          body.Name,
			"serial_number": body.SerialNumber,
			"created_at":    time.Now().Local(),
			"updated_at":    time.Now().Local(),
		},
	)
	if err != nil {
		return nil, err
	}

	product, errA := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[models.Product],
	)
	if errA != nil {
		return nil, errA
	}

	return &product, nil
}

func GetProduct(pool *pgxpool.Pool) (
	[]*models.Product,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx, `select * from products`,
	)
	if err != nil {
		return nil, err
	}

	product, errA := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[models.Product],
	)

	if errA != nil {

		return nil, errA
	}

	return product, nil
}

func GetOneProduct(pool *pgxpool.Pool, id int) (
	*models.Product,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`select * from products where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	product, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.Product],
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProduct(pool *pgxpool.Pool, body serializer.ProductBody, id int) (
	*models.Product,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`UPDATE products SET 
                updated_at = @updated_at,
                name = @name,
    			serial_number = @serial_number 
             WHERE id = @id returning id,created_at,updated_at,name,serial_number`,
		pgx.NamedArgs{
			"id":            id,
			"updated_at":    time.Now().Local(),
			"name":          body.Name,
			"serial_number": body.SerialNumber,
		},
	)
	if err != nil {
		return nil, err
	}

	product, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.Product],
	)

	if err != nil {

		return nil, err
	}

	return product, nil
}

func DeleteProduct(pool *pgxpool.Pool, id int) (
	*models.Product,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`delete from products where id = @id returning id,created_at,updated_at,name,serial_number`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	product, errA := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[models.Product],
	)

	if errA != nil {

		return nil, errA
	}

	return product, nil
}
