package usecase

import (
	"backend/internal/entities"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/lib/jwt"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB             *gorm.DB
	UserRepository *repositories.UserRepository
	Validate       *validator.Validate
}

func NewAuthUseCase(db *gorm.DB, UserRepository *repositories.UserRepository, Validate *validator.Validate) *AuthUsecase {
	return &AuthUsecase{
		DB:             db,
		UserRepository: UserRepository,
		Validate:       Validate,
	}
}

func (u *AuthUsecase) SignIn(data *models.SignInRequest) (*models.SignInResponse, error) {
	user := new(entities.User)
	err := u.UserRepository.First(user, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, errors.New(
			"invalid credentials",
		)
	}

	token, err := jwt.Sign(
		jwt.Claims{
			UserID:   user.ID,
			Username: user.Name,
		},
	)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &models.SignInResponse{
		Token: token,
		User: models.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil

}

func (u *AuthUsecase) SignUp(
	data *models.SignUpRequest,
) (*models.SignUpRequest, error) {

	user := u.UserRepository.First(&entities.User{
		Email: data.Email,
	}, map[string]interface{}{"email": data.Email})

	if user == nil {
		return data, errors.New("user already exists")
	}
	key, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return data, err
	}

	create := u.UserRepository.Create(&entities.User{
		ID:        uuid.NewString(),
		Email:     data.Email,
		Name:      data.Name,
		Password:  string(key),
		CreatedAt: time.Now(),
	})

	if create != nil {
		return data, create
	}

	return data, nil
}

func (u *AuthUsecase) GetProfile(id string) (*models.User, error) {
	user := new(entities.User)
	err := u.UserRepository.First(user, map[string]interface{}{"id": id})
	return &models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, err
}
