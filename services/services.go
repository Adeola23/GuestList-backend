package services

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/dto"
	
)


func AddGuestService(request dto.AddGuestRequest, response dto.GuestResponse) (int64, dto.GuestResponse){
	log.Print("nAME ", request.Name)
	if ok, err := dto.Validate(request.Name); !ok {
		fmt.Printf("[AddGuestService] Invalid input error for name: %v\n", err)
		return 400, response
	}

	if ok, err := dto.ValidateAccom(request.Accompanying_guests); !ok {
		fmt.Printf("[AddGuestService] Invalid input error for accompanying_guests: %v\n", err)
		return 400, response
	}
	isEnough, err := isCapacityEnough(request.Table, request.Accompanying_guests)

	if err !=nil {
		log.Printf("[AddGuestService] check-space db error: %v\n", err)
		return 500, response

	}

	if !isEnough {
		log.Printf("[AddGuestService] Not sufficient seats: %v\n", request)
		return 403, response
	}

	result, err := database.AddGuestToList(request.Table, request.Name, request.Accompanying_guests)

	if err != nil {
		log.Printf("[AddGuestService] Insert guest db error: %v\n", err)
		return 500, response
	}

	response.Name = result
	return 200, response

}

func AddTableService(request dto.AddTableRequest, response dto.AddTableResponse)(int64, dto.AddTableResponse){
	log.Print("service cap ", request.Capacity )
	_,id, err := database.AddTable(request.Capacity)
	

	if err != nil {
		log.Printf("[AddTableService] Insert table to tables db error: %v\n", err)
		return 500, response
	}

	response.Capacity = request.Capacity
	response.Id = id

	return 200, response

}
func GetGuestListService(response dto.GuestListResponse)(int64, dto.GuestListResponse){
	result, err := database.GetGuestList()
	if err != nil {
		log.Printf("[GetGuestListService] query db error: %v\n", err)
		return 500, response
	}
	response.Guests = result
	return 200, response


}

func GetArrivedGuestListService(response dto.GuestListResponse)(int64, dto.GuestListResponse){
	result, err := database.GetArrivedGuests()
	if err != nil {
		log.Printf("[GetArrivedGuestsService] query db error: %v\n", err)
		return 500, response
	}

	response.Guests = result
	return 200, response

}


func GetEmptySeatsService(response dto.EmptySeatsResponse)(int64, dto.EmptySeatsResponse){
	result, err := database.GetEmptySeats()
	if err != nil {
		log.Printf("[GetEmptySeatsService] query db error: %v\n", err)
		return 500, response
	}

	response.Seats_empty = result
	return 200, response

}

func CheckInGuestService(response dto.GuestResponse, request dto.UpdateGuestRequest)(int64, dto.GuestResponse){
	result, err := database.GuestArrive(request.Name, request.Accompanying_guests)
	if err != nil {
		log.Printf("[CheckInGuestService] query db error: %v\n", err)
		return 500, response
	}

	response.Name = result
	return 200, response

}


func DeleteGuestService(request dto.DeleteGuestRequest) (int64) {
	if ok, err := dto.Validate(request.Name); !ok {
		log.Printf("[DeleteGuestService] Invalid input error: %v\n", err)
		return 400
	}

	err := database.DeleteGuest(request.Name)
	if err != nil {
		log.Printf("[DeleteGuestService] delete guest db error: %v\n", err)
		return 500
	}
	return 204
}

func isCapacityEnough( table int, guests int) (bool, error) {
	capacity, err := database.GetTableCapacity(table)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[isCapacity] Not found table: %v\n", err)
			return false, err
		}
		log.Printf("[isCapacity] check max space process error: %v\n", err)
		return false, err
	}



	guestSelf := 1
	if ( guests + guestSelf) > capacity {
		log.Printf("[isSufficient] Not sufficient seats: %v\n%v", guests+guestSelf, capacity)
		return false, nil
	}
	return true, nil
}
