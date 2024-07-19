package db

import (
	"context"
	"medication-notifier/data"
	"medication-notifier/utils/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbMedicationDataService struct {
	conn *pgxpool.Pool
}

func NewDbMedicationDataService(conn *pgxpool.Pool) DbMedicationDataService {
	return DbMedicationDataService{
		conn: conn,
	}
}

func (s *DbMedicationDataService) Add(med data.Medication) error {
	sql := "insert into med.medications(user_id, name, day, time_of_day) values ($1, $2, $3, $4)"
	_, err := s.conn.Exec(
		context.Background(),
		sql,
		med.UserId,
		med.Name,
		med.Day,
		med.TimeOfDay,
	)
	return err
}

func (s *DbMedicationDataService) FindByUserId(userId string) []data.Medication {
	sql := `
	select id, user_id, name, day, time_of_day from med.medications
	where user_id=$1
	`
	result := []data.Medication{}
	rows, err := s.conn.Query(context.Background(), sql, userId)
	if err != nil {
		logger.Error("fetch medications by userId failed, err: %s", err)
		return result
	}

	for rows.Next() {
		var id string
		var name string
		var day string
		var timeOfDay string
		if err := rows.Scan(&id, &name, &day, &timeOfDay); err != nil {
			logger.Error("fetch medications by userId scan failed, err: %s", err)
			return []data.Medication{}
		}
		result = append(result, data.Medication{
			Id:        id,
			Name:      name,
			Day:       day,
			TimeOfDay: timeOfDay,
		})
	}

	return result
}
func (s *DbMedicationDataService) FindById(id string) (*data.Medication, error) {
	sql := `
	select id, user_id, name, day, time_of_day from med.medications
	where id=$1
	`
	row := s.conn.QueryRow(context.Background(), sql, id)

	var name string
	var day string
	var timeOfDay string
	if err := row.Scan(&id, &name, &day, &timeOfDay); err != nil {
		logger.Error("fetch medication by id scan failed, err: %s", err)
		return nil, err
	}

	return &data.Medication{
		Id:        id,
		Name:      name,
		Day:       day,
		TimeOfDay: timeOfDay,
	}, nil

}

func (s *DbMedicationDataService) RemoveById(id string) error {
	sql := "delete from med.medications where id=$1"
	_, err := s.conn.Exec(context.Background(), sql, id)

	return err
}

func (s *DbMedicationDataService) RemoveByUserId(userId string) error {
	sql := "delete from med.medications where user_id=$1"
	_, err := s.conn.Exec(context.Background(), sql, userId)

	return err
}
