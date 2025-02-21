package amqp_rpc

//
//import (
//	"context"
//	"fmt"
//	"github.com/ochinchind/docsproc/internal/entity"
//	"github.com/ochinchind/docsproc/internal/usecase"
//	"github.com/ochinchind/docsproc/pkg/rabbitmq/rmq_rpc/server"
//	"github.com/streadway/amqp"
//)
//
//type disciplineRoutes struct {
//	disciplineUseCase usecase.Discipline
//}
//
//func newDisciplineRoutes(routes map[string]server.CallHandler, t *usecase.Services) {
//	r := &disciplineRoutes{t.Discipline}
//	{
//		routes["getDiscipline"] = r.getDiscipline()
//	}
//}
//
//type getDisciplinesResponse struct {
//	Disciplines []entity.Discipline `json:"disciplines"`
//	Total       int64               `json:"total"`
//}
//
//func (r *disciplineRoutes) getDiscipline() server.CallHandler {
//	return func(d *amqp.Delivery) (interface{}, error) {
//		disciplines, total, err := r.disciplineUseCase.Get(context.Background())
//		if err != nil {
//			return nil, fmt.Errorf("amqp_rpc - disciplineRoutes - getHistory - r.disciplineUseCase.History: %w", err)
//		}
//
//		response := getDisciplinesResponse{disciplines, total}
//
//		return response, nil
//	}
//}
