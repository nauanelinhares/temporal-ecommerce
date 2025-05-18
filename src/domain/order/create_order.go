package order

import (
	"context"
	"fmt"
	"log"
	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/temporal/order/workflows"

	// Assuming your workflow input struct is here or defined directly

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func (s *OrderService) CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error) {

	workflowOptions := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("order-%s", uuid.New().String()),
		TaskQueue: "ORDER_TASK_QUEUE",
	}

	we, err := s.temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.CreateOrderWorkflow, order.UserID, order.ProductID, order.Quantity)

	if err != nil {
		return entities.Order{}, err
	}

	log.Printf("Started workflow for order %s. WorkflowID: %s, RunID: %s", order.ID.String(), we.GetID(), we.GetRunID())

	return order, nil
}
