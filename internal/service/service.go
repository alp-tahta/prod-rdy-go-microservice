package service

import "github.com/alp-tahta/prod-rdy-go-microservice/internal/repository"

type Service struct {
	r repository.RepositoryInterface
}

func NewService(r repository.RepositoryInterface) *Service {
	return &Service{r: r}
}

func (s *Service) CreateProduct() {
	s.r.CreateProduct()
}

func (s *Service) GetProduct() {
	s.r.GetProduct()
}

func (s *Service) UpdateProduct() {
	s.r.UpdateProduct()
}

func (s *Service) DeleteProduct() {
	s.r.DeleteProduct()
}
