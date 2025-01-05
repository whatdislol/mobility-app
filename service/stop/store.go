package stop

import (
	"database/sql"

	"github.com/whatdislol/mobility-app/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateStop(stop types.Stop) error {
	_, err := s.db.Exec("INSERT INTO stops (latitude, longitude) VALUES ($1, $2)", stop.Latitude, stop.Longitude)
	if err != nil {
		return err
	}

	return nil
}