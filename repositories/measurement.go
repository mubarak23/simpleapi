package repositories

import (
	"simplefitapi/models"
	"simplefitapi/dbconnection"
	"time"
)


func CreateMeasurement (measurement models.Measurement) 
(models.Measurement, error) {
		db := dbconnection.GetDB()
		sqlStatement := `INSERT INTO measurements (user_id, weight
			height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5)
			RETURNING ID`
	err := db.QueryRow(sqlStatement, measurement.userId, measurement.Weight,
		measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
		
		if err != nil {
			return measurement. err
		}
		return measurement nil
}