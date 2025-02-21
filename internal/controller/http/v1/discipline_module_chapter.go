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

type disciplineModuleChapterRoutes struct {
	t  usecase.DisciplineModuleChapter
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineModuleChapterRoutes(handler *gin.RouterGroup, t usecase.DisciplineModuleChapter, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineModuleChapterRoutes{t, l, rd}

	h := handler.Group("/disciplineModuleChapters").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "disciplineModuleChapter", "read"), r.getDisciplineModuleChapters)
		h.GET("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapter", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "disciplineModuleChapter", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapter", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "disciplineModuleChapter", "write"), r.delete)
	}
}

type getDisciplineModuleChaptersResponse struct {
	DisciplineModuleChapters []entity.DisciplineModuleChapter `json:"disciplineModuleChapters"`
	Total                    int64                            `json:"total"`
}

// Get a list of disciplineModuleChapters.
//
// @Summary      Get DisciplineModuleChapters
// @Description  Fetch a list of disciplineModuleChapters.
// @ID           get_disciplineModuleChapters
// @Tags         DisciplineModuleChapters
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplineModuleChaptersResponse "Successful response with disciplineModuleChapter list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapters [get]
func (r disciplineModuleChapterRoutes) getDisciplineModuleChapters(context *gin.Context) {
	disciplineModuleChapters, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplineModuleChaptersResponse{disciplineModuleChapters, total})
}

// Get a disciplineModuleChapter.
//
// @Summary      Get DisciplineModuleChapter
// @Description  Fetch a disciplineModuleChapter by ID.
// @ID           get_disciplineModuleChapter
// @Tags         DisciplineModuleChapters
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapter ID"
// @Success      200 {object} entity.DisciplineModuleChapter "Successful response with disciplineModuleChapter"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapters/{id} [get]
func (r disciplineModuleChapterRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	disciplineModuleChapter, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, disciplineModuleChapter)
}

// Update a disciplineModuleChapter.
//
// @Summary      Update DisciplineModuleChapter
// @Description  Update a disciplineModuleChapter.
// @ID           update_disciplineModuleChapter
// @Tags         DisciplineModuleChapters
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapter ID"
// @Param        request body dto.UpdateDisciplineModuleChapterDTO true "DisciplineModuleChapter request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapters [post]
func (r disciplineModuleChapterRoutes) update(context *gin.Context) {
	var updateDisciplineModuleChapterDTO = &dto.UpdateDisciplineModuleChapterDTO{}
	if err := context.BindJSON(updateDisciplineModuleChapterDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineModuleChapterDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new disciplineModuleChapter.
//
// @Summary      Store DisciplineModuleChapter
// @Description  Store a new disciplineModuleChapter.
// @ID           store_disciplineModuleChapter
// @Tags         DisciplineModuleChapters
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineModuleChapterDTO true "DisciplineModuleChapter request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapters [post]
func (r disciplineModuleChapterRoutes) store(context *gin.Context) {
	var disciplineModuleChapter = &dto.StoreDisciplineModuleChapterDTO{}
	if err := context.BindJSON(disciplineModuleChapter); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(disciplineModuleChapter)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a disciplineModuleChapter.
//
// @Summary      Delete DisciplineModuleChapter
// @Description  Delete a disciplineModuleChapter.
// @ID           delete_disciplineModuleChapter
// @Tags         DisciplineModuleChapters
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineModuleChapter ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineModuleChapters/{id} [delete]
func (r disciplineModuleChapterRoutes) delete(context *gin.Context) {
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
