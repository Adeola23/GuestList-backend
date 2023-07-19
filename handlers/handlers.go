package handlers

import (
	"net/http"
	"log"

	"github.com/getground/tech-tasks/backend/dto"
	"github.com/getground/tech-tasks/backend/services"
	"github.com/gin-gonic/gin"
)

// @Summary Add table
// @Description Add table to sql db
// @Tags Add
// @Accept json
// @Produce json
// @Param body body dto.AddTableRequest true "body"
// @Success 200 {object} dto.GuestResponse "id" "capacity"
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 403 {object} errors.ErrorInfo "Forbiden to add"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /tables [POST]
func AddTable(c *gin.Context) {
	response := dto.AddTableResponse{}
	request := dto.AddTableRequest{}
	c.BindJSON(&request)
	statusCode, result := services.AddTableService(request,response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, "Invalid input")
	case 403:
		c.JSON(http.StatusForbidden, "Forbiden to update")
	case 500:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	default:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

// @Summary Add guest
// @Description Add guest to sql db
// @Tags Add
// @Accept json
// @Produce json
// @Param body body dto.AddGuestRequest true "body"
// @Success 200 {object} dto.GuestResponse "name"
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 403 {object} errors.ErrorInfo "Forbiden to add"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guest_list/:name [POST]
func AddGuest(c *gin.Context) {
	request := dto.AddGuestRequest{}
	response := dto.GuestResponse{}
	request.Name = c.Param("name")
	c.BindJSON(&request)
	log.Printf("[AddGuest Handler] request=%+v\n", request)
	statusCode, result := services.AddGuestService(request, response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, nil)
	case 403:
		c.JSON(http.StatusForbidden, nil)
	case 500:
		c.JSON(http.StatusInternalServerError, nil)
	default:
		c.JSON(http.StatusInternalServerError, nil)
	}

}

// @Summary Delete guest
// @Description Delete guest from db
// @Tags Delete
// @Success 204 Delete successful
// @Failure 400 {object} errors.ErrorInfo "Invalid input"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guests/:name [DELETE]
func DeleteGuest(c *gin.Context) {
	request := dto.DeleteGuestRequest{}
	request.Name = c.Param("name")
	log.Printf("[DeleteGuest Handler] request=%+v\n", request)
	statusCode := services.DeleteGuestService(request)
	switch statusCode {
	case 204:
		c.JSON(http.StatusNoContent, nil)
	case 400:
		c.JSON(http.StatusBadRequest, nil)
	case 500:
		c.JSON(http.StatusInternalServerError, nil)
	default:
		c.JSON(http.StatusInternalServerError, nil)
	}
}

// @Summary Update guest
// @Description Update guest's Accompanying guest number
// @Tags UPDATE
// @Accept json
// @Produce json
// @Param body body dto.UpdateGuestRequest true "body"
// @Success 200 {object} dto.GuestResponse "name"
// @Failure 400 {object} errors: "Invalid input"
// @Failure 403 {object} errors: "Forbiden to update"
// @Failure 500 {object} errors: "Internal server error"
// @Router /guests/:name [PUT]
func CheckInGuest(c *gin.Context) {
	request := dto.UpdateGuestRequest{}
	response := dto.GuestResponse{}
	request.Name = c.Param("name")
	c.BindJSON(&request)
	log.Printf("[CheckInGuest Handler] request=%+v\n", request)
	statusCode, result := services.CheckInGuestService (response, request)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 400:
		c.JSON(http.StatusBadRequest, "Invalid input")
	case 403:
		c.JSON(http.StatusForbidden, "Forbiden to update")
	case 500:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	default:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

// @Summary Get empty seats
// @Description get all empty seats number
// @Tags GET
// @Produce json
// @Success 200 {object} dto.EmptySeatsResponse "seats_empty"
// @Failure 500 {object} errors: "Internal server error"
// @Router /seats_empty [GET]
func GetEmptySeats(c *gin.Context) {
	response := dto.EmptySeatsResponse{}
	statusCode, result:= services.GetEmptySeatsService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, "Internal server error")
	default:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

// @Summary Get guest lists
// @Description get all guest list
// @Tags GET
// @Produce json
// @Success 200 {object} dto.GuestListResponse "guests[]"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guest_list [GET]
func GetGuestLists(c *gin.Context) {
	response := dto.GuestListResponse{}
	statusCode, result := services.GetGuestListService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	default:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}

// @Summary Get guest lists with arrived time
// @Description get all guest list with arrived time
// @Tags GET
// @Produce json
// @Success 200 {object} dto.GuestListResponse "guests[]"
// @Failure 500 {object} errors.ErrorInfo "Internal server error"
// @Router /guests [GET]
func GetArrivedGuestLists(c *gin.Context) {
	response := dto.GuestListResponse{}
	statusCode, result := services.GetArrivedGuestListService(response)
	switch statusCode {
	case 200:
		c.JSON(http.StatusOK, result)
	case 500:
		c.JSON(http.StatusInternalServerError, "Internal server error")
	default:
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}


// @Summary health checker API
// @Success 200 {string} string "ok"
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}