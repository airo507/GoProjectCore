package auth

import (
	"context"
	"encoding/json"
	"fmt"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"os"
	"strings"
)

type UserRepository interface {
	Register(ctx context.Context, userId string, userData userEntity.UserData) (userEntity.User, error)
	Login(ctx context.Context, login string, password string) (string, error)
}

type Service struct {
	repo UserRepository
}

func NewRegistrationService(userRepo UserRepository) *Service {
	return &Service{
		repo: userRepo,
	}
}

func (s *Service) Register(ctx context.Context, userId string, userInfo userEntity.UserData) error {
	fileName := "users.json"

	userCreated, err := s.repo.Register(ctx, userId, userInfo)

	err = s.CheckUser(fileName, userCreated)
	if err != nil {
		fmt.Println("Can't register user!")
	}
	return err
}

func (s *Service) CheckUser(fileName string, userCreated userEntity.User) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("File not exist")
		if os.IsNotExist(err) {
			err = s.WriteToFile(fileName, userCreated)
		}
	}
	defer file.Close()

	usersFromFiles := s.ReadFile(fileName)

	for _, user := range usersFromFiles {
		if user.UserData.Login == userCreated.UserData.Login {
			usersFromFiles = append(usersFromFiles, userCreated)
			break
		}
	}

	file.Close()
	return nil
}

func (s *Service) WriteToFile(fileName string, userCreated userEntity.User) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(userCreated)
	return nil
}

func (s *Service) ReadFile(fileName string) []userEntity.User {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	strFromFile := string(data)
	reader := strings.NewReader(strFromFile)
	decoder := json.NewDecoder(reader)

	var usersFromFile []userEntity.User
	err = decoder.Decode(&usersFromFile)
	if err != nil {
		panic(err)
	}

	return usersFromFile
}

func (s *Service) Login(ctx context.Context, login string, password string) error {

	return nil
}
