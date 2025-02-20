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

type qualificationRoutes struct {
	t  usecase.Qualification
	l  logger.Interface
	rd *redis.Client
}

func newQualificationRoutes(handler *gin.RouterGroup, t usecase.Qualification, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &qualificationRoutes{t, l, rd}

	h := handler.Group("/qualifications").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "qualification", "read"), r.getQualifications)
		h.GET("/:id", permission.Permission(casbinEnforcer, "qualification", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "qualification", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "qualification", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "qualification", "write"), r.delete)
	}
}

type getQualificationsResponse struct {
	Qualifications []entity.Qualification `json:"qualifications"`
	Total          int64                  `json:"total"`
}

// Get a list of qualifications.
//
// @Summary      Get Qualifications
// @Description  Fetch a list of qualifications.
// @ID           get_qualifications
// @Tags         Qualifications
// @Accept       json
// @Produce      json
// @Success      200 {object} getQualificationsResponse "Successful response with qualification list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/qualifications [get]
func (r qualificationRoutes) getQualifications(context *gin.Context) {
	qualifications, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getQualificationsResponse{qualifications, total})
}

// Get a qualification.
//
// @Summary      Get Qualification
// @Description  Fetch a qualification by ID.
// @ID           get_qualification
// @Tags         Qualifications
// @Accept       json
// @Produce      json
// @Param        id path string true "Qualification ID"
// @Success      200 {object} entity.Qualification "Successful response with qualification"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/qualifications/{id} [get]
func (r qualificationRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	qualification, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, qualification)
}

// Update a qualification.
//
// @Summary      Update Qualification
// @Description  Update a qualification.
// @ID           update_qualification
// @Tags         Qualifications
// @Accept       json
// @Produce      json
// @Param        id path string true "Qualification ID"
// @Param        request body dto.UpdateQualificationDTO true "Qualification request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/qualifications [post]
func (r qualificationRoutes) update(context *gin.Context) {
	var updateQualificationDTO = &dto.UpdateQualificationDTO{}
	if err := context.BindJSON(updateQualificationDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateQualificationDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new qualification.
//
// @Summary      Store Qualification
// @Description  Store a new qualification.
// @ID           store_qualification
// @Tags         Qualifications
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreQualificationDTO true "Qualification request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/qualifications [post]
func (r qualificationRoutes) store(context *gin.Context) {
	var qualification = &dto.StoreQualificationDTO{}
	if err := context.BindJSON(qualification); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(qualification)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a qualification.
//
// @Summary      Delete Qualification
// @Description  Delete a qualification.
// @ID           delete_qualification
// @Tags         Qualifications
// @Accept       json
// @Produce      json
// @Param        id path string true "Qualification ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/qualifications/{id} [delete]
func (r qualificationRoutes) delete(context *gin.Context) {
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
