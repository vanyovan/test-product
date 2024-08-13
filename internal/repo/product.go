package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/vanyovan/test-product.git/internal/entity"
	"github.com/vanyovan/test-product.git/internal/helper"
	"github.com/vanyovan/test-product.git/internal/repo/wrapper"
)

type Repo struct {
	db *sql.DB
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error)
	UpdateStatusByUserId(ctx context.Context, status string, userId string) (updatedAt time.Time, err error)
	GetWalletByUserId(ctx context.Context, userId string) (result entity.Wallet, err error)
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateProduct(ctx context.Context, product entity.Product) (result entity.Product, err error) {
	tx, err := wrapper.FromContext(ctx)
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

func (r *Repo) GetWalletByUserId(ctx context.Context, userId string) (result entity.Wallet, err error) {
	query := "SELECT wallet_id, owned_by, status, enabled_at, disabled_at, balance FROM mst_wallet WHERE owned_by = ?"
	row := r.db.QueryRow(query, userId)
	result = entity.Wallet{}
	err = row.Scan(&result.WalletId, &result.OwnedBy, &result.Status, &result.EnabledAt, &result.DisabledAt, &result.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		} else {
			fmt.Println("Failed to retrieve row:", err)
		}
		return result, err
	}
	return result, nil
}

func (r *Repo) UpdateStatusByUserId(ctx context.Context, status string, userId string) (updatedAt time.Time, err error) {
	tx, err := wrapper.FromContext(ctx)
	if tx == nil || err != nil {
		tx, err = r.db.Begin()
		if err != nil {
			tx.Rollback()
			return time.Time{}, errors.New("failed to begin database transaction")
		}
	}

	timeNow := time.Now()
	if status == helper.ConstantEnabled {
		_, err = tx.Exec("UPDATE mst_wallet set status = ?, enabled_at = ? where owned_by = ?", status, timeNow, userId)
	} else {
		_, err = tx.Exec("UPDATE mst_wallet set status = ?, disabled_at = ? where owned_by = ?", status, timeNow, userId)
	}
	if err != nil {
		tx.Rollback()
		return time.Time{}, fmt.Errorf("failed to update wallet: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return time.Time{}, errors.New("failed to commit database transaction")
	}

	return timeNow, nil
}
