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

type competencyRoutes struct {
	t  usecase.Competency
	l  logger.Interface
	rd *redis.Client
}

func newCompetencyRoutes(handler *gin.RouterGroup, t usecase.Competency, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &competencyRoutes{t, l, rd}

	h := handler.Group("/competencies").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "competency", "read"), r.getCompetencies)
		h.GET("/:id", permission.Permission(casbinEnforcer, "competency", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "competency", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "competency", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "competency", "write"), r.delete)
	}
}

type getCompetenciesResponse struct {
	Competencies []entity.Competency `json:"competencies"`
	Total        int64               `json:"total"`
}

// Get a list of competencies.
//
// @Summary      Get Competencies
// @Description  Fetch a list of competencies.
// @ID           get_competencies
// @Tags         Competencies
// @Accept       json
// @Produce      json
// @Success      200 {object} getCompetenciesResponse "Successful response with competency list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/competencies [get]
func (r competencyRoutes) getCompetencies(context *gin.Context) {
	competencies, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getCompetenciesResponse{competencies, total})
}

// Get a competency.
//
// @Summary      Get Competency
// @Description  Fetch a competency by ID.
// @ID           get_competency
// @Tags         Competencies
// @Accept       json
// @Produce      json
// @Param        id path string true "Competency ID"
// @Success      200 {object} entity.Competency "Successful response with competency"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/competencies/{id} [get]
func (r competencyRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	competency, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, competency)
}

// Update a competency.
//
// @Summary      Update Competency
// @Description  Update a competency.
// @ID           update_competency
// @Tags         Competencies
// @Accept       json
// @Produce      json
// @Param        id path string true "Competency ID"
// @Param        request body dto.UpdateCompetencyDTO true "Competency request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/competencies [post]
func (r competencyRoutes) update(context *gin.Context) {
	var updateCompetencyDTO = &dto.UpdateCompetencyDTO{}
	if err := context.BindJSON(updateCompetencyDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateCompetencyDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new competency.
//
// @Summary      Store Competency
// @Description  Store a new competency.
// @ID           store_competency
// @Tags         Competencies
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreCompetencyDTO true "Competency request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/competencies [post]
func (r competencyRoutes) store(context *gin.Context) {
	var competency = &dto.StoreCompetencyDTO{}
	if err := context.BindJSON(competency); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(competency)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a competency.
//
// @Summary      Delete Competency
// @Description  Delete a competency.
// @ID           delete_competency
// @Tags         Competencies
// @Accept       json
// @Produce      json
// @Param        id path string true "Competency ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/competencies/{id} [delete]
func (r competencyRoutes) delete(context *gin.Context) {
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
