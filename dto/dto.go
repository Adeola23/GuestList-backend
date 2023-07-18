package dto


type AddGuestRequest struct {
	Name                string `param:"name" `
	Table               int    `json:"table" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
}

type AddTableRequest struct {
	Capacity int    `json:"capacity" binding:"required"`
}


type DeleteGuestRequest struct {
	Name string `param:"name" binding:"required"`
}


type UpdateGuestRequest struct {
	Name                string `param:"name" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
}

type AddTableResponse struct {
	Id int64   `json:"id" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`

}

type GuestResponse struct {
	Name string `json:"name" binding:"required"`
}

type EmptySeatsResponse struct {
	Seats_empty int `json:"seats_empty" binding:"required" `
}

type Guest struct {
	Name                string `json:"name" binding:"required"`
	Table               int    `json:"table,omitempty" `
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
	Time_arrived        string `json:"time_arrived,omitempty"`
}

type GuestListResponse struct {
	Guests []Guest `json:"guests"`
}

type ArrivedGuest struct {
	Name                string `json:"name" binding:"required"`
	Accompanying_guests int    `json:"accompanying_guests" binding:"required"`
	Time_arrived        string `json:"time_arrived" binding:"required"`
}

type HandleError struct {
    message string
}

func (e HandleError) Error() string {
    return e.message
}

func Validate(s string) (bool, error) {
	if s == "" {
		return false, HandleError{"Empty string"}
	}
	return true, nil
}

func ValidateAccom(n int) (bool, error) {
	if n < 0 {
		return false, HandleError{"accompanying guest is less than zero "}
	}
	return true, nil
}