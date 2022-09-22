package services

import "hexagonal-architecture/internal/core/ports"

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) GetUser() error {
	err := s.userRepository.GetUser()
	if err != nil {
		return err
	}

	return nil
}
