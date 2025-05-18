package workflows

import (
	"time"

	"temporal-ecommerce/src/domain/entities"
	"temporal-ecommerce/src/temporal/order/activities"

	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"
)

type OrderWorkflowInput struct {
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int
}

// CreateOrderWorkflow orquestra a criação de um pedido.
func CreateOrderWorkflow(ctx workflow.Context, userID uuid.UUID, productID uuid.UUID, quantity int) (entities.Order, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var act activities.Activities
	var orderResult entities.Order

	initialOrder := entities.Order{

		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	err := workflow.ExecuteActivity(ctx, act.CreateOrderActivity, initialOrder).Get(ctx, &orderResult)
	if err != nil {
		return entities.Order{}, err
	}

	var product entities.Product
	err = workflow.ExecuteActivity(ctx, act.GetProductActivity, orderResult.ProductID.String()).Get(ctx, &product)
	if err != nil {
		orderResult.Status = entities.StatusCancelled
		workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, nil)
		return entities.Order{}, err
	}

	err = workflow.ExecuteActivity(ctx, act.ValidateStockActivity, product, orderResult).Get(ctx, &orderResult)
	if err != nil {
		orderResult.Status = entities.StatusCancelled
		workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, nil)
		return entities.Order{}, err
	}

	err = workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, &orderResult)
	if err != nil {
		return entities.Order{}, err
	}

	var user entities.User
	err = workflow.ExecuteActivity(ctx, act.GetUserActivity, orderResult.UserID.String()).Get(ctx, &user)
	if err != nil {
		orderResult.Status = entities.StatusCancelled
		workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, nil)
		return entities.Order{}, err
	}

	err = workflow.ExecuteActivity(ctx, act.ValidateUserBalanceActivity, user, orderResult).Get(ctx, &orderResult)
	if err != nil {
		orderResult.Status = entities.StatusPaymentFailed
		workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, nil)
		return entities.Order{}, err
	}

	err = workflow.ExecuteActivity(ctx, act.UpdateOrderActivity, orderResult).Get(ctx, &orderResult)
	if err != nil {
		return entities.Order{}, err
	}

	return orderResult, nil
}
