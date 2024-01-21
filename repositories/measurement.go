package repositories

import (
	"simplefitapi/models"
	"simplefitapi/dbconnection"
	"time"
)


func CreateUserMeasurement (measurement model.Measurements) (model.Measurements, error) {
		db := dbconnection.GetDB()
		sqlStatement := `INSERT INTO measurements (userid, weight, height, body_fat, created_at) 
		VALUES ($1, $2, $3, $4, $5)
			RETURNING ID`
			err := db.QueryRow(sqlStatement, measurement.UserID, measurement.Weight,
				measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
		
		
		if err != nil {
			return measurement, err
		}
		return measurement, nil
}

func UpdateMeasurement (measurement model.Measurements, id int) (model.Measurements, error) {
		db := dbconnection.GetDB()
		sqlStatement := `
			UPDATE measurements
			SET weight = $2, height = $3, body_fat = $4, created_at = $5
			WHERE id = $1
			RETURNING id
		`
		err := db.QueryRow(sqlStatement, id, measurement.Weight, measurement.Height,
		 measurement.BodyFat, time.Now()).Scan(&id)

		 if err != nil {
			return model.Measurements{}, err
		 }

		 measurement.Id = id
		 return measurement, nil
	} 