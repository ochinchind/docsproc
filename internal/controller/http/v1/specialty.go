package v1

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ochinchind/docsproc/internal/dto"
	"github.com/ochinchind/docsproc/internal/entity"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/auth"
	"github.com/ochinchind/docsproc/internal/http-server/middleware/permission"
	"github.com/ochinchind/docsproc/internal/usecase"
	"github.com/ochinchind/docsproc/pkg/logger"
	"net/http"
	"strconv"
)

type specialtyRoutes struct {
	t  usecase.Specialty
	l  logger.Interface
	rd *redis.Client
}

func newSpecialtyRoutes(handler *gin.RouterGroup, t usecase.Specialty, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &specialtyRoutes{t, l, rd}

	h := handler.Group("/specialties").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "specialty", "read"), r.getSpecialties)
		h.GET("/:id", permission.Permission(casbinEnforcer, "specialty", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "specialty", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "specialty", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "specialty", "write"), r.delete)
	}
}

type getSpecialtiesResponse struct {
	Specialties []entity.Specialty `json:"specialties"`
	Total       int64              `json:"total"`
}

// Get a list of specialties.
//
// @Summary      Get Specialties
// @Description  Fetch a list of specialties.
// @ID           get_specialties
// @Tags         Specialties
// @Accept       json
// @Produce      json
// @Success      200 {object} getSpecialtiesResponse "Successful response with specialty list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/specialties [get]
func (r specialtyRoutes) getSpecialties(context *gin.Context) {
	specialties, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getSpecialtiesResponse{specialties, total})
}

// Get a specialty.
//
// @Summary      Get Specialty
// @Description  Fetch a specialty by ID.
// @ID           get_specialty
// @Tags         Specialties
// @Accept       json
// @Produce      json
// @Param        id path string true "Specialty ID"
// @Success      200 {object} entity.Specialty "Successful response with specialty"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/specialties/{id} [get]
func (r specialtyRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	specialty, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, specialty)
}

// Update a specialty.
//
// @Summary      Update Specialty
// @Description  Update a specialty.
// @ID           update_specialty
// @Tags         Specialties
// @Accept       json
// @Produce      json
// @Param        id path string true "Specialty ID"
// @Param        request body dto.UpdateSpecialtyDTO true "Specialty request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/specialties [post]
func (r specialtyRoutes) update(context *gin.Context) {
	var updateSpecialtyDTO = &dto.UpdateSpecialtyDTO{}
	if err := context.BindJSON(updateSpecialtyDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateSpecialtyDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new specialty.
//
// @Summary      Store Specialty
// @Description  Store a new specialty.
// @ID           store_specialty
// @Tags         Specialties
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreSpecialtyDTO true "Specialty request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/specialties [post]
func (r specialtyRoutes) store(context *gin.Context) {
	var specialty = &dto.StoreSpecialtyDTO{}
	if err := context.BindJSON(specialty); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(specialty)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully stored"})
}

// Delete a specialty.
//
// @Summary      Delete Specialty
// @Description  Delete a specialty.
// @ID           delete_specialty
// @Tags         Specialties
// @Accept       json
// @Produce      json
// @Param        id path string true "Specialty ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/specialties/{id} [delete]
func (r specialtyRoutes) delete(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Delete(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully deleted"})
}
