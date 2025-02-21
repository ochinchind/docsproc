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

type disciplineRoutes struct {
	t  usecase.Discipline
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineRoutes(handler *gin.RouterGroup, t usecase.Discipline, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineRoutes{t, l, rd}

	h := handler.Group("/disciplines").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "discipline", "read"), r.getDisciplines)
		h.GET("/:id", permission.Permission(casbinEnforcer, "discipline", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "discipline", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "discipline", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "discipline", "write"), r.delete)
	}
}

type getDisciplinesResponse struct {
	Disciplines []entity.Discipline `json:"disciplines"`
	Total       int64               `json:"total"`
}

// Get a list of disciplines.
//
// @Summary      Get Disciplines
// @Description  Fetch a list of disciplines.
// @ID           get_disciplines
// @Tags         Disciplines
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplinesResponse "Successful response with discipline list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplines [get]
func (r disciplineRoutes) getDisciplines(context *gin.Context) {
	disciplines, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplinesResponse{disciplines, total})
}

// Get a discipline.
//
// @Summary      Get Discipline
// @Description  Fetch a discipline by ID.
// @ID           get_discipline
// @Tags         Disciplines
// @Accept       json
// @Produce      json
// @Param        id path string true "Discipline ID"
// @Success      200 {object} entity.Discipline "Successful response with discipline"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplines/{id} [get]
func (r disciplineRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	discipline, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, discipline)
}

// Update a discipline.
//
// @Summary      Update Discipline
// @Description  Update a discipline.
// @ID           update_discipline
// @Tags         Disciplines
// @Accept       json
// @Produce      json
// @Param        id path string true "Discipline ID"
// @Param        request body dto.UpdateDisciplineDTO true "Discipline request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplines [post]
func (r disciplineRoutes) update(context *gin.Context) {
	var updateDisciplineDTO = &dto.UpdateDisciplineDTO{}
	if err := context.BindJSON(updateDisciplineDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new discipline.
//
// @Summary      Store Discipline
// @Description  Store a new discipline.
// @ID           store_discipline
// @Tags         Disciplines
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineDTO true "Discipline request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplines [post]
func (r disciplineRoutes) store(context *gin.Context) {
	var discipline = &dto.StoreDisciplineDTO{}
	if err := context.BindJSON(discipline); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	userId := context.MustGet("auth_user_id").(uint)

	err := r.t.Store(discipline, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a discipline.
//
// @Summary      Delete Discipline
// @Description  Delete a discipline.
// @ID           delete_discipline
// @Tags         Disciplines
// @Accept       json
// @Produce      json
// @Param        id path string true "Discipline ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplines/{id} [delete]
func (r disciplineRoutes) delete(context *gin.Context) {
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
