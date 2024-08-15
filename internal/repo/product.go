package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/vanyovan/test-product.git/internal/entity"
)

type Repo struct {
	db *sql.DB
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error)
	GetProducts(ctx context.Context) (result []entity.Product, err error)
	DeleteProductByProductID(ctx context.Context, id int64) (err error)
	UpdateProductByProductID(ctx context.Context, id int64, product entity.Product) (err error)
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error) {
	ctx = context.Background()
	tx, err := r.db.BeginTx(ctx, nil)
	if tx == nil || err != nil {
		tx, err = r.db.Begin()
		if err != nil {
			tx.Rollback()
			return result, errors.New("failed to begin database transaction")
		}
	}

	resultExec, err := tx.ExecContext(ctx, "INSERT INTO mst_product (name, description, price, variety, rating, stock) VALUES (?, ?, ?, ?, ?, ?)",
		product.ProductName, product.ProductDescription, product.ProductPrice, product.ProductVariety, product.ProductRating, product.ProductStock)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("failed to create wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return result, errors.New("failed to commit database transaction")
	}

	lastInsertID, _ := resultExec.LastInsertId()

	result = entity.Product{
		ProductID:          lastInsertID,
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		ProductPrice:       product.ProductPrice,
		ProductVariety:     product.ProductVariety,
		ProductRating:      product.ProductRating,
		ProductStock:       product.ProductStock,
	}
	return result, nil
}

func (r *Repo) GetProducts(ctx context.Context) (result []entity.Product, err error) {
	rows, err := r.db.Query("SELECT id, name, description, price, variety, rating, stock FROM mst_product")
	if err != nil {
		return result, err
	}

	for rows.Next() {
		product := &entity.Product{}
		err = rows.Scan(&product.ProductID, &product.ProductName, &product.ProductDescription, &product.ProductPrice, &product.ProductVariety, &product.ProductRating, &product.ProductStock)
		if err != nil {
			return result, err
		}
		result = append(result, *product)
	}

	return result, nil
}

func (r *Repo) DeleteProductByProductID(ctx context.Context, id int64) error {
	query := "DELETE FROM mst_product WHERE id = ?"

    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("could not delete product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("could not determine number of rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no product found with ID %d", id)
    }

    return nil
}

func (r *Repo) UpdateProductByProductID(ctx context.Context, id int64, product entity.Product) error {
	query := "UPDATE mst_product SET"
	var args []interface{}

	if &product.ProductName != nil {
        query += " name = ?,"
        args = append(args, product.ProductName)
    }
    if &product.ProductDescription != nil {
        query += " description = ?,"
        args = append(args, product.ProductDescription)
    }
    if &product.ProductPrice != nil {
        query += " price = ?,"
        args = append(args, product.ProductPrice)
    }
	if &product.ProductVariety != nil {
        query += " variety = ?,"
        args = append(args, product.ProductVariety)
    }
	if &product.ProductRating != nil {
        query += " rating = ?,"
        args = append(args, product.ProductRating)
    }
	if &product.ProductStock != nil {
        query += " stock = ?,"
        args = append(args, product.ProductStock)
    }

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = ?"
	args = append(args, id)


	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to begin database transaction")
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update transaction: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to commit database transaction")
	}

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("could not determine number of rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no product found with ID %d", id)
    }

	return nil
}