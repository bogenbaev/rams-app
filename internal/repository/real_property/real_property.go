package real_property

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"rams/pkg/models"
	"time"
)

type RealPropertyRepo struct {
	db *sqlx.DB
}

func NewRealPropertyRepository(db *sqlx.DB) *RealPropertyRepo {
	return &RealPropertyRepo{
		db: db,
	}
}

func (r *RealPropertyRepo) Create(ctx context.Context, realPro models.RealProperty) error {
	location, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return err
	}

	// Get the current time in Almaty
	now := time.Now().In(location)
	query := fmt.Sprintf(`
		INSERT INTO %s
					(
					 	property_type_id,
					 	property_type,
					 	address,
					 	price,
					 	rooms,
					 	area,
					 	description,
					 	created_at
					)
			values
			    	(
			    		$1, $2, $3, $4, $5, $6, $7, $8
			    	)
		returning id
	`, models.RealPropertyTable)

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err = tx.QueryRow(query,
		realPro.PropertyTypeID,
		realPro.PropertyType,
		realPro.Address,
		realPro.Price,
		realPro.Rooms,
		realPro.Area,
		realPro.Description,
		now,
	).Scan(&realPro.ID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RealPropertyRepo) GetList(ctx context.Context) (realPros []models.RealProperty, err error) {
	query := fmt.Sprintf(`
		SELECT * from %s 
	`, models.RealPropertyTable)

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return realPros, err
	}
	defer tx.Rollback()

	if err = tx.Select(&realPros, query); err != nil {
		return realPros, err
	}

	return realPros, tx.Commit()
}

func (r *RealPropertyRepo) GetByID(ctx context.Context, realPro models.RealProperty) (models.RealProperty, error) {
	query := fmt.Sprintf(`
		SELECT * from %s
			where id = $1
	`, models.RealPropertyTable)

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	if err != nil {
		return realPro, err
	}

	if err = r.db.Get(&realPro, query, realPro.ID); err != nil {
		return realPro, err
	}

	return realPro, tx.Commit()
}
