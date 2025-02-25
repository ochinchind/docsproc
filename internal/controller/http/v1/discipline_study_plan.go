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

type disciplineStudyPlanRoutes struct {
	t  usecase.DisciplineStudyPlan
	l  logger.Interface
	rd *redis.Client
}

func newDisciplineStudyPlanRoutes(handler *gin.RouterGroup, t usecase.DisciplineStudyPlan, l logger.Interface, casbinEnforcer *casbin.Enforcer, rd *redis.Client) {
	r := &disciplineStudyPlanRoutes{t, l, rd}

	h := handler.Group("/disciplineStudyPlans").Use(auth.Auth(rd))
	{
		h.GET("", permission.Permission(casbinEnforcer, "disciplineStudyPlan", "read"), r.getDisciplineStudyPlans)
		h.GET("/:id", permission.Permission(casbinEnforcer, "disciplineStudyPlan", "read"), r.get)
		h.POST("", permission.Permission(casbinEnforcer, "disciplineStudyPlan", "write"), r.store)
		h.PATCH("/:id", permission.Permission(casbinEnforcer, "disciplineStudyPlan", "write"), r.update)
		h.DELETE("/:id", permission.Permission(casbinEnforcer, "disciplineStudyPlan", "write"), r.delete)
	}
}

type getDisciplineStudyPlansResponse struct {
	DisciplineStudyPlans []entity.DisciplineStudyPlan `json:"disciplineStudyPlans"`
	Total                int64                        `json:"total"`
}

// Get a list of disciplineStudyPlans.
//
// @Summary      Get DisciplineStudyPlans
// @Description  Fetch a list of disciplineStudyPlans.
// @ID           get_disciplineStudyPlans
// @Tags         DisciplineStudyPlans
// @Accept       json
// @Produce      json
// @Success      200 {object} getDisciplineStudyPlansResponse "Successful response with disciplineStudyPlan list"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineStudyPlans [get]
func (r disciplineStudyPlanRoutes) getDisciplineStudyPlans(context *gin.Context) {
	disciplineStudyPlans, total, err := r.t.Get(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, getDisciplineStudyPlansResponse{disciplineStudyPlans, total})
}

// Get a disciplineStudyPlan.
//
// @Summary      Get DisciplineStudyPlan
// @Description  Fetch a disciplineStudyPlan by ID.
// @ID           get_disciplineStudyPlan
// @Tags         DisciplineStudyPlans
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineStudyPlan ID"
// @Success      200 {object} entity.DisciplineStudyPlan "Successful response with disciplineStudyPlan"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineStudyPlans/{id} [get]
func (r disciplineStudyPlanRoutes) get(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	disciplineStudyPlan, err := r.t.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, disciplineStudyPlan)
}

// Update a disciplineStudyPlan.
//
// @Summary      Update DisciplineStudyPlan
// @Description  Update a disciplineStudyPlan.
// @ID           update_disciplineStudyPlan
// @Tags         DisciplineStudyPlans
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineStudyPlan ID"
// @Param        request body dto.UpdateDisciplineStudyPlanDTO true "DisciplineStudyPlan request body"
// @Success      200 {object} response "Successfully updated"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineStudyPlans [post]
func (r disciplineStudyPlanRoutes) update(context *gin.Context) {
	var updateDisciplineStudyPlanDTO = &dto.UpdateDisciplineStudyPlanDTO{}
	if err := context.BindJSON(updateDisciplineStudyPlanDTO); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(500, response{err.Error()})
		return
	}

	err = r.t.Update(id, updateDisciplineStudyPlanDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully updated"})
}

// Store a new disciplineStudyPlan.
//
// @Summary      Store DisciplineStudyPlan
// @Description  Store a new disciplineStudyPlan.
// @ID           store_disciplineStudyPlan
// @Tags         DisciplineStudyPlans
// @Accept       json
// @Produce      json
// @Param        request body dto.StoreDisciplineStudyPlanDTO true "DisciplineStudyPlan request body"
// @Success      200 {object} response "Successfully stored"
// @Failure      400 {object} response "Invalid request payload"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineStudyPlans [post]
func (r disciplineStudyPlanRoutes) store(context *gin.Context) {
	var disciplineStudyPlan = &dto.StoreDisciplineStudyPlanDTO{}
	if err := context.BindJSON(disciplineStudyPlan); err != nil {
		context.JSON(http.StatusBadRequest, response{err.Error()})
		return
	}

	err := r.t.Store(disciplineStudyPlan)
	if err != nil {
		context.JSON(http.StatusInternalServerError, response{err.Error()})
		return
	}

	context.JSON(http.StatusOK, response{Message: "Successfully created"})
}

// Delete a disciplineStudyPlan.
//
// @Summary      Delete DisciplineStudyPlan
// @Description  Delete a disciplineStudyPlan.
// @ID           delete_disciplineStudyPlan
// @Tags         DisciplineStudyPlans
// @Accept       json
// @Produce      json
// @Param        id path string true "DisciplineStudyPlan ID"
// @Success      200 {object} response "Successfully deleted"
// @Failure      500 {object} response "Internal server error"
// @Router       /v1/disciplineStudyPlans/{id} [delete]
func (r disciplineStudyPlanRoutes) delete(context *gin.Context) {
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
