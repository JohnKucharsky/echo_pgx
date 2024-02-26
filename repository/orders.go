package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/JohnKucharsky/echo_pgx/models"
	"github.com/JohnKucharsky/echo_pgx/serializer"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func CreateOrder(pool *pgxpool.Pool, body serializer.OrderBody) (
	*models.Order,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx, `
        INSERT INTO orders (user_id, product_id, updated_at)
        VALUES (@user_id, @product_id, @updated_at)
        RETURNING id
        `,
		pgx.NamedArgs{
			"user_id":    body.UserId,
			"product_id": body.ProductId,
			"updated_at": time.Now().Local(),
		},
	)
	if err != nil {
		return nil, err
	}

	type returnedRow struct {
		ID uint `db:"id"`
	}

	row, errA := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[returnedRow],
	)
	if errA != nil {
		return nil, errA
	}

	order, err := GetOneOrder(pool, row.ID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func query() string {
	return `SELECT
			JSON_BUILD_OBJECT(
				'id', orders.id,
				'updated_at', orders.updated_at,
				'user', JSON_BUILD_OBJECT(
					'id', users.id,
				    'created_at', users.created_at,
				    'updated_at', users.updated_at,
				    'first_name', users.first_name,
				    'last_name', users.last_name
				),
				'product', JSON_BUILD_OBJECT(
					'id', products.id,
				    'created_at', products.created_at,
				    'updated_at', products.updated_at,
				    'name', products.name,
				    'serial_number', products.serial_number
				)
			)
		FROM orders
		    left join users on orders.user_id = users.id
		    left join products on orders.product_id = products.id`
}

func GetOrders(pool *pgxpool.Pool) (
	[]*models.Order,
	error,
) {
	ctx := context.Background()

	query := query()

	rows, err := pool.Query(
		ctx, query,
	)

	if err != nil {
		return nil, err
	}

	orders, errA := pgx.CollectRows(
		rows, pginit.JSONRowToAddrOfStruct[models.Order],
	)

	if errA != nil {
		return nil, errA
	}

	return orders, nil
}

func GetOneOrder(pool *pgxpool.Pool, id uint) (
	*models.Order,
	error,
) {
	ctx := context.Background()

	if id == 0 {
		return nil, errors.New("id is 0")
	}

	query := query()
	query += fmt.Sprintf(" where orders.id = %d", id)

	rows, err := pool.Query(
		ctx, query,
		pgx.NamedArgs{
			"id": id,
		},
	)

	if err != nil {
		return nil, err
	}

	order, errA := pgx.CollectExactlyOneRow(
		rows, pginit.JSONRowToAddrOfStruct[models.Order],
	)

	if errA != nil {
		return nil, errA
	}

	return order, nil
}

func UpdateOrder(pool *pgxpool.Pool, body serializer.OrderBody, id int) (
	*models.Order,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`UPDATE orders SET 
                updated_at = @updated_at,
                user_id = @user_id,
    			product_id = @product_id 
             WHERE id = @id returning id`,
		pgx.NamedArgs{
			"id":         id,
			"updated_at": time.Now().Local(),
			"user_id":    body.UserId,
			"product_id": body.ProductId,
		},
	)
	if err != nil {
		return nil, err
	}

	type returnedRow struct {
		ID uint `db:"id"`
	}

	row, errA := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[returnedRow],
	)
	if errA != nil {
		return nil, errA
	}

	order, err := GetOneOrder(pool, row.ID)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func DeleteOrder(pool *pgxpool.Pool, id int) (
	*uint,
	error,
) {
	ctx := context.Background()

	rows, err := pool.Query(
		ctx,
		`delete from orders where id = @id returning id`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	type returnedRow struct {
		ID uint `db:"id"`
	}

	row, errA := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[returnedRow],
	)
	if errA != nil {
		return nil, errA
	}

	return &row.ID, nil
}
