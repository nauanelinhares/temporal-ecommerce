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

	workflowRun, err := s.temporalClient.ExecuteWorkflow(ctx, workflowOptions, workflows.CreateOrderWorkflow, order.UserID, order.ProductID, order.Quantity)
	if err != nil {
		return entities.Order{}, fmt.Errorf("failed to start CreateOrderWorkflow: %w", err)
	}

	err = workflowRun.Get(ctx, &order)
	if err != nil {
		return entities.Order{}, fmt.Errorf("CreateOrderWorkflow failed: %w", err)
	}

	log.Printf("Started workflow for order %s. WorkflowID: %s, RunID: %s", order.ID.String(), workflowRun.GetID(), workflowRun.GetRunID())

	return order, nil
}
