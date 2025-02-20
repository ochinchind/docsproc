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

type disciplineModuleRoutes struct {
	t  usecase.DisciplineModule
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineModuleRoutes(handler *gin.RouterGroup, t usecase.DisciplineModule, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineModuleRoutes{t, l, rd}

	h := handler.Group("/disciplineModules").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "disciplineModule", "read"), r.getDisciplineModules)
		h.GET("/:id", permission.Permission(casbinEnforcer, "disciplineModule", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "disciplineModule", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "disciplineModule", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "disciplineModule", "write"), r.delete)
	}
}

type getDisciplineModulesResponse struct {
	DisciplineModules []entity.DisciplineModule `json:"disciplineModules"`
	Total             int64                     `json:"total"`
}

// Get a list of disciplineModules.
//
// @Summary      Get DisciplineModules
// @Description  Fetch a list of disciplineModules.
// @ID           get_disciplineModules
// @Tags         DisciplineModules
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplineModulesResponse "Successful response with disciplineModule list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModules [get]
func (r disciplineModuleRoutes) getDisciplineModules(context *gin.Context) {
	disciplineModules, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplineModulesResponse{disciplineModules, total})
}

// Get a disciplineModule.
//
// @Summary      Get DisciplineModule
// @Description  Fetch a disciplineModule by ID.
// @ID           get_disciplineModule
// @Tags         DisciplineModules
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModule ID"
// @Success      200 {object} entity.DisciplineModule "Successful response with disciplineModule"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModules/{id} [get]
func (r disciplineModuleRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	disciplineModule, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, disciplineModule)
}

// Update a disciplineModule.
//
// @Summary      Update DisciplineModule
// @Description  Update a disciplineModule.
// @ID           update_disciplineModule
// @Tags         DisciplineModules
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModule ID"
// @Param        request body dto.UpdateDisciplineModuleDTO true "DisciplineModule request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModules [post]
func (r disciplineModuleRoutes) update(context *gin.Context) {
	var updateDisciplineModuleDTO = &dto.UpdateDisciplineModuleDTO{}
	if err := context.BindJSON(updateDisciplineModuleDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineModuleDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new disciplineModule.
//
// @Summary      Store DisciplineModule
// @Description  Store a new disciplineModule.
// @ID           store_disciplineModule
// @Tags         DisciplineModules
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineModuleDTO true "DisciplineModule request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModules [post]
func (r disciplineModuleRoutes) store(context *gin.Context) {
	var disciplineModule = &dto.StoreDisciplineModuleDTO{}
	if err := context.BindJSON(disciplineModule); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(disciplineModule)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a disciplineModule.
//
// @Summary      Delete DisciplineModule
// @Description  Delete a disciplineModule.
// @ID           delete_disciplineModule
// @Tags         DisciplineModules
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModule ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModules/{id} [delete]
func (r disciplineModuleRoutes) delete(context *gin.Context) {
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
