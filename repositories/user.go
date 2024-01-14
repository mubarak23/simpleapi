package repositories

import (
	"simplefitapi/models"
	"simplefitapi/dbconnection"
)


func CreateUser (user models.User) (models.User, error) {
	db := dbconnection.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password) 
		VALUE ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password ).Scan(&user.Id)

	if err != nil {
		return user, err
	}

	return user, nil

}