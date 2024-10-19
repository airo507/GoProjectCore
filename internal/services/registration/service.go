package registration

import "context"

type RegService struct {
}

func NewRegistrationService() *RegService {
	return &RegService{}
}

func (s *RegService) Register(ctx context.Context, firstName string) (string, error) {
	return firstName, nil
}
