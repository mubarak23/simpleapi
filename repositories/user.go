package repositories

import (
	"time"
	 "simplefitapi/models"
	"simplefitapi/dbconnection"
)


func CreateUser (user model.User) (model.User, error) {
	db := dbconnection.GetDB()
	
	sqlStatement := `INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password ).Scan(&user.Id)

	if err != nil {
		return user, err
	}

	return user, nil

}

func UpdateUser (user model.User, id int) (model.User, error) {
	db := dbconnection.GetDB()
	sqlStatement := `
		UPDATE users
		SET name = $2, email = $3, password = $4, updated_at = $5
		WHERE id = $1
		RETURNING id 	
	`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, 
		time.Now()).Scan(&id)
	
	if err != nil {
		return model.User{}, err
	}
	user.Id = id
	return user, nil
	
}

// func AddNostrEvent (event model.Nostrevent) (model.Nostrevent, error) {
// 	db := dbconnection.GetDB()

// 	sqlStatement := `
// 		INSERT INTO nostrevents (pubkey, kind, ptags, etags, gtags, expiration, content)
// 	`
// 	err := db.QueryRow(sqlStatement, event.PubKey, event.Kind, event.Ptags, event.Etags, event.Gtags, event.Expiration, event.Content).Scan(&event.Id)

// 	if err != nil {
// 		return event, err
// 	}

// 	return event, nil

// }