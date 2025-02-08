package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/airo507/GoProjectCore/internal/app/user"
	userEntity "github.com/airo507/GoProjectCore/internal/entity/user"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

type UserRepository interface {
	Register(ctx context.Context, userId string, userData userEntity.User) (userEntity.User, error)
	Get(userId string) error
}

type Service struct {
	repo UserRepository
}

//var SecretKey = []byte("secretkey1")

func NewRegistrationService(userRepo UserRepository) *Service {
	return &Service{
		repo: userRepo,
	}
}

func (s *Service) Register(ctx context.Context, userId string, userInfo user.ResponseUser) error {
	fileName := "users.json"

	userData := userEntity.User{
		Id:        userInfo.UserId,
		Login:     userInfo.Login,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Email:     userInfo.Email,
		Password:  userInfo.Password,
	}

	userCreated, err := s.repo.Register(ctx, userId, userData)

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
		if user.Login == userCreated.Login && user.Password == userCreated.Password {
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
		return fmt.Errorf("%s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(userCreated)
	return nil
}

func (s *Service) ReadFile(fileName string) []userEntity.User {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("%s", err)
		return nil
	}

	strFromFile := string(data)
	reader := strings.NewReader(strFromFile)
	decoder := json.NewDecoder(reader)

	var usersFromFile []userEntity.User
	err = decoder.Decode(&usersFromFile)
	if err != nil {
		fmt.Errorf("%s", err)
		return nil
	}

	return usersFromFile
}

func (s *Service) checkUsersInFile(fileName string, userData user.InputUser) error {
	usersFromFiles := s.ReadFile(fileName)

	for _, user := range usersFromFiles {
		if user.Login == userData.Login && user.Password == userData.Password {
			return fmt.Errorf("%s", "User is exist")
		}
	}
	return fmt.Errorf("%s", "User is not exist")
}

func (s *Service) Login(ctx context.Context, userData user.InputUser) error {
	fileName := "users.json"
	err := s.checkUsersInFile(fileName, userData)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GenerateToken(login string) (string, error) {
	claims := jwt.MapClaims{
		"username": login,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
