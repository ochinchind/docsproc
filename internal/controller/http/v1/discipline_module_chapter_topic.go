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

type disciplineModuleChapterTopicRoutes struct {
	t  usecase.DisciplineModuleChapterTopic
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineModuleChapterTopicRoutes(handler *gin.RouterGroup, t usecase.DisciplineModuleChapterTopic, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineModuleChapterTopicRoutes{t, l, rd}

	h := handler.Group("/disciplineModuleChapterTopics").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "disciplineModuleChapterTopic", "read"), r.getDisciplineModuleChapterTopics)
		h.GET("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapterTopic", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "disciplineModuleChapterTopic", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapterTopic", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapterTopic", "write"), r.delete)
	}
}

type getDisciplineModuleChapterTopicsResponse struct {
	DisciplineModuleChapterTopics []entity.DisciplineModuleChapterTopic `json:"disciplineModuleChapterTopics"`
	Total                         int64                                 `json:"total"`
}

// Get a list of disciplineModuleChapterTopics.
//
// @Summary      Get DisciplineModuleChapterTopics
// @Description  Fetch a list of disciplineModuleChapterTopics.
// @ID           get_disciplineModuleChapterTopics
// @Tags         DisciplineModuleChapterTopics
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplineModuleChapterTopicsResponse "Successful response with disciplineModuleChapterTopic list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapterTopics [get]
func (r disciplineModuleChapterTopicRoutes) getDisciplineModuleChapterTopics(context *gin.Context) {
	disciplineModuleChapterTopics, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplineModuleChapterTopicsResponse{disciplineModuleChapterTopics, total})
}

// Get a disciplineModuleChapterTopic.
//
// @Summary      Get DisciplineModuleChapterTopic
// @Description  Fetch a disciplineModuleChapterTopic by ID.
// @ID           get_disciplineModuleChapterTopic
// @Tags         DisciplineModuleChapterTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapterTopic ID"
// @Success      200 {object} entity.DisciplineModuleChapterTopic "Successful response with disciplineModuleChapterTopic"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapterTopics/{id} [get]
func (r disciplineModuleChapterTopicRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	disciplineModuleChapterTopic, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, disciplineModuleChapterTopic)
}

// Update a disciplineModuleChapterTopic.
//
// @Summary      Update DisciplineModuleChapterTopic
// @Description  Update a disciplineModuleChapterTopic.
// @ID           update_disciplineModuleChapterTopic
// @Tags         DisciplineModuleChapterTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapterTopic ID"
// @Param        request body dto.UpdateDisciplineModuleChapterTopicDTO true "DisciplineModuleChapterTopic request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapterTopics [post]
func (r disciplineModuleChapterTopicRoutes) update(context *gin.Context) {
	var updateDisciplineModuleChapterTopicDTO = &dto.UpdateDisciplineModuleChapterTopicDTO{}
	if err := context.BindJSON(updateDisciplineModuleChapterTopicDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineModuleChapterTopicDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new disciplineModuleChapterTopic.
//
// @Summary      Store DisciplineModuleChapterTopic
// @Description  Store a new disciplineModuleChapterTopic.
// @ID           store_disciplineModuleChapterTopic
// @Tags         DisciplineModuleChapterTopics
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineModuleChapterTopicDTO true "DisciplineModuleChapterTopic request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapterTopics [post]
func (r disciplineModuleChapterTopicRoutes) store(context *gin.Context) {
	var disciplineModuleChapterTopic = &dto.StoreDisciplineModuleChapterTopicDTO{}
	if err := context.BindJSON(disciplineModuleChapterTopic); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(disciplineModuleChapterTopic)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a disciplineModuleChapterTopic.
//
// @Summary      Delete DisciplineModuleChapterTopic
// @Description  Delete a disciplineModuleChapterTopic.
// @ID           delete_disciplineModuleChapterTopic
// @Tags         DisciplineModuleChapterTopics
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapterTopic ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapterTopics/{id} [delete]
func (r disciplineModuleChapterTopicRoutes) delete(context *gin.Context) {
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
