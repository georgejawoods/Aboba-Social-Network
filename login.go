package aboba

import "context"

type LoginInput struct {
	Email    string
	Username *string
}

func (svc *Service) Login(ctx context.Context, in LoginInput) error {
	return nil
}
