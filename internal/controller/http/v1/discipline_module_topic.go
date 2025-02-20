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

type disciplineModuleTopicRoutes struct {
	t  usecase.DisciplineModuleTopic
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineModuleTopicRoutes(handler *gin.RouterGroup, t usecase.DisciplineModuleTopic, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineModuleTopicRoutes{t, l, rd}

	h := handler.Group("/disciplineModuleTopics").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "disciplineModuleTopic", "read"), r.getDisciplineModuleTopics)
		h.GET("/:id", permission.Permission(casbinEnforcer, "disciplineModuleTopic", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "disciplineModuleTopic", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "disciplineModuleTopic", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "disciplineModuleTopic", "write"), r.delete)
	}
}

type getDisciplineModuleTopicsResponse struct {
	DisciplineModuleTopics []entity.DisciplineModuleTopic `json:"disciplineModuleTopics"`
	Total                  int64                          `json:"total"`
}

// Get a list of disciplineModuleTopics.
//
// @Summary      Get DisciplineModuleTopics
// @Description  Fetch a list of disciplineModuleTopics.
// @ID           get_disciplineModuleTopics
// @Tags         DisciplineModuleTopics
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplineModuleTopicsResponse "Successful response with disciplineModuleTopic list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleTopics [get]
func (r disciplineModuleTopicRoutes) getDisciplineModuleTopics(context *gin.Context) {
	disciplineModuleTopics, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplineModuleTopicsResponse{disciplineModuleTopics, total})
}

// Get a disciplineModuleTopic.
//
// @Summary      Get DisciplineModuleTopic
// @Description  Fetch a disciplineModuleTopic by ID.
// @ID           get_disciplineModuleTopic
// @Tags         DisciplineModuleTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleTopic ID"
// @Success      200 {object} entity.DisciplineModuleTopic "Successful response with disciplineModuleTopic"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleTopics/{id} [get]
func (r disciplineModuleTopicRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	disciplineModuleTopic, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, disciplineModuleTopic)
}

// Update a disciplineModuleTopic.
//
// @Summary      Update DisciplineModuleTopic
// @Description  Update a disciplineModuleTopic.
// @ID           update_disciplineModuleTopic
// @Tags         DisciplineModuleTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleTopic ID"
// @Param        request body dto.UpdateDisciplineModuleTopicDTO true "DisciplineModuleTopic request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleTopics [post]
func (r disciplineModuleTopicRoutes) update(context *gin.Context) {
	var updateDisciplineModuleTopicDTO = &dto.UpdateDisciplineModuleTopicDTO{}
	if err := context.BindJSON(updateDisciplineModuleTopicDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineModuleTopicDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new disciplineModuleTopic.
//
// @Summary      Store DisciplineModuleTopic
// @Description  Store a new disciplineModuleTopic.
// @ID           store_disciplineModuleTopic
// @Tags         DisciplineModuleTopics
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineModuleTopicDTO true "DisciplineModuleTopic request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleTopics [post]
func (r disciplineModuleTopicRoutes) store(context *gin.Context) {
	var disciplineModuleTopic = &dto.StoreDisciplineModuleTopicDTO{}
	if err := context.BindJSON(disciplineModuleTopic); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(disciplineModuleTopic)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a disciplineModuleTopic.
//
// @Summary      Delete DisciplineModuleTopic
// @Description  Delete a disciplineModuleTopic.
// @ID           delete_disciplineModuleTopic
// @Tags         DisciplineModuleTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleTopic ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleTopics/{id} [delete]
func (r disciplineModuleTopicRoutes) delete(context *gin.Context) {
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
