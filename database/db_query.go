package database

import (
	"database/sql"
	
	"log"

	"github.com/getground/tech-tasks/backend/dto"
)





func AddGuestToList(table int, name string, partner int) (string, error) {
	db := DbConn()
	_, err := db.Exec("INSERT INTO guest_list(table_id, guest_name, accompanying_guests) VALUES(?, ?,?)", table, name, partner)
	if err != nil {
		log.Printf("[AddGuestToList] insert guest  to list db error: %v\n", err)
		return "", err
	}
	return name, nil

}

func AddTable(capacity int)(int, int64, error, ){

	db := DbConn()
	result, err := db.Exec("INSERT INTO tables(capacity) VALUES(?)", capacity)
	if err != nil {
		log.Printf("[AddTable] insert table  to tables db error: %v\n", err)
		return 0, 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Printf("[AddTable] retrieve last insert ID error: %v\n", err)
		return 0, 0, err
	}

	log.Print("capacity ", capacity)

	return capacity, id,nil
}

func DeleteGuest(name string) error {
	db := DbConn()
	_, err := db.Exec("DELETE FROM guest_list WHERE guest_name=?", name)
	if err != nil {
		log.Printf("[DeleteGuest] delete guest db error: %v\n", err)
		return err
	}
	log.Printf("Guest %s: successfully deleted from the guest list", name)
	return nil
}

func UpdateGuest(name string, partner int) (bool, error) {
	db := DbConn()
	_, err := db.Exec("UPDATE guest_list SET accompanying_guests=? WHERE guest_name=?", partner, name)
	if err != nil {
		log.Printf("[UpdateGuest] update guest data error: %v\n", err)
		return false, err
	}
	return true, nil
}

func QueryTableByName(name string) (int, error) {
	db := DbConn()
	tableId := 0
	var n sql.NullInt32
	err := db.QueryRow("SELECT table_id FROM guest_list WHERE name = ?", name).Scan(&n)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[QueryTableByName] Name not found error: %v\n", err)
			return tableId, err
		}
		log.Printf("[QueryTableByName] query Table number db error: %v\n", err)
		return -1, err
	}
	if n.Valid {
		tableId = int(n.Int32)
	}
	return tableId, nil

}
func GetTableCapacity(tableId int) (int, error) {
	db := DbConn()
	var availableCapacity int
	// Select all available seats
	rows, err := db.Query("SELECT capacity from tables WHERE table_id=?", tableId)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		// Scan rows into variable
		if err := rows.Scan(&availableCapacity); err != nil {
			log.Println(err)
			return 0, err
		}
	}
	return availableCapacity, nil
}

func GuestArrive(name string, partner int) (string, error) {
	// Let the guest in and update the status and actual arrived guests. Arrival time will get updated automatically.
	
	db := DbConn()
	// Execute query
	_, err := db.Exec("UPDATE guest_list set status=?, accompanying_guests=? WHERE guest_name=?", "ARRIVED", partner, name)
	if err != nil {
		log.Println(err)
		return name, err
	}
	log.Printf("Guest %s: successfully updated from the guest list", name)
	return name, err
}



func GetArrivedGuests()([]dto.Guest, error){

	db := DbConn()
	rows, err := db.Query("SELECT guest_name, accompanying_guests,time_arrived FROM guest_list")
	if err != nil {
		log.Printf("[GetArrivedGuests] query arrived guest error: %v\n", err)
		return []dto.Guest{}, err
	}

	guest := dto.Guest{} 
	guest_list := []dto.Guest{}

	for rows.Next() {
		err = rows.Scan(&guest.Name, &guest.Accompanying_guests, &guest.Time_arrived)
		if err != nil {
			log.Printf("[QueryArrivedGuests] Query next guest error: %v\n", err)
			return []dto.Guest{}, err
		}
		guest_list = append(guest_list, guest)
	}

	defer db.Close()
	return guest_list, nil
}

func GetArrivedGuestsByName(name string)(int , error){

	db := DbConn()
	rows, err := db.Query("SELECT actual_accompanying_guests FROM guest_list WHERE guest_name=?  VALUES(?)", name)
	if err != nil {
		log.Printf("[GetArrivedGuestsByName] Query guest error: %v\n", err)
		return 0, err
	}
	var availableCapacity int

	defer rows.Close()
	for rows.Next() {
		// Scan rows into variable
		if err := rows.Scan(&availableCapacity); err != nil {
			log.Println(err)
			return 0, err
		}
	}
	return availableCapacity, nil
}

func GetGuestList()([]dto.Guest, error){
	db := DbConn()
	rows, err := db.Query("SELECT guest_name,table_id, accompanying_guests FROM guest_list")
	if err != nil {
		log.Printf("[QueryArrivedGuests] query guest list error: %v\n", err)
		return []dto.Guest{}, err
	}

	guest := dto.Guest{} 
	guest_list := []dto.Guest{}

	for rows.Next() {
		err = rows.Scan(&guest.Name, &guest.Table, &guest.Accompanying_guests)
		if err != nil {
			log.Printf("[GetGuestsList] query next guest error: %v\n", err)
			return []dto.Guest{}, err
		}
		guest_list = append(guest_list, guest)
	}

	defer db.Close()
	return guest_list, nil


}



func GetEmptySeats() (int, error) {
	db := DbConn()
	rows, err := db.Query("SELECT SUM(accompanying_guests + 1)"+
		"FROM guest_list WHERE status=?", "ARRIVED")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer rows.Close()

	var totalArrivedGuests int
	for rows.Next() {
		// Scan rows into variable
		if err := rows.Scan(&totalArrivedGuests); err != nil {
			log.Println(err)
			return 0, err
		}
	}

	rows, err = db.Query("SELECT SUM(capacity) FROM tables")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer rows.Close()

	var totalSeats int
	for rows.Next() {
		// Scan rows into variable
		if err := rows.Scan(&totalSeats); err != nil {
			log.Println(err)
			return 0, err
		}
	}

	return totalSeats - totalArrivedGuests, nil


}


