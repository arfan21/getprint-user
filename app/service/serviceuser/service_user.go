package serviceuser

import (
	"fmt"
	"strings"

	"github.com/arfan21/getprint-user/app/model/modelresponse"
	"github.com/arfan21/getprint-user/app/model/modeluser"
	"github.com/arfan21/getprint-user/app/repository/mysql/mysqluser"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user modeluser.User) (*modelresponse.User, error)
	GetByID(id string) (*modelresponse.User, error)
	Update(user modeluser.User) (*modelresponse.User, error)
	Login(user modeluser.User) (*modelresponse.User, error)
	LoginUsingLine(dataLine *modelresponse.LineVerifyIdToken) (*modelresponse.User, error)
}

type services struct {
	userRepo mysqluser.UserRepository
}

func New(userRepo mysqluser.UserRepository) UserService {
	return &services{userRepo}
}

func (s *services) Create(user modeluser.User) (*modelresponse.User, error) {
	user.ID = uuid.NewV4()
	user.Identities.UserID = user.ID
	user.UserLog.UserID = user.ID
	if user.Identities.Provider == "" {
		user.Identities.Provider = "getprint"
		user.Identities.ProviderID = user.ID.String()
	}

	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	createdData, err := s.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	responseUser := new(modelresponse.User)
	responseUser.Set(*createdData)
	return responseUser, nil
}

func (s *services) GetByID(id string) (*modelresponse.User, error) {
	uuidFromString := uuid.FromStringOrNil(id)
	user, err := s.userRepo.GetByID(uuidFromString)
	if err != nil {
		return nil, err
	}

	return &modelresponse.User{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		Name:          user.Name,
		Picture:       user.Picture.String,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   user.PhoneNumber.String,
		Address:       user.Address.String,
		Role:          user.Role,
		Provider:      user.Identities.Provider,
		ProviderID:    user.Identities.ProviderID,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Update(user modeluser.User) (*modelresponse.User, error) {
	if user.Password.Valid {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		user.Password.Scan(string(hashedPassword))
	}

	err := s.userRepo.Update(&user)

	if err != nil {
		return nil, err
	}

	return &modelresponse.User{
		ID:            user.ID,
		CreatedAt:     user.CreatedAt,
		Name:          user.Name,
		Picture:       user.Picture.String,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		PhoneNumber:   user.PhoneNumber.String,
		Address:       user.Address.String,
		Role:          user.Role,
		Provider:      user.Identities.Provider,
		ProviderID:    user.Identities.ProviderID,
		LastLogin:     user.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) Login(user modeluser.User) (*modelresponse.User, error) {
	dataUser, err := s.userRepo.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password.String), []byte(user.Password.String))
	if err != nil {
		return nil, err
	}

	dataUser.UserLog.UserID = user.ID

	go func() {
		_ = s.userRepo.UpdateUserLog(&dataUser.UserLog)
	}()

	return &modelresponse.User{
		ID:            dataUser.ID,
		CreatedAt:     dataUser.CreatedAt,
		Name:          dataUser.Name,
		Picture:       dataUser.Picture.String,
		Email:         dataUser.Email,
		EmailVerified: dataUser.EmailVerified,
		PhoneNumber:   dataUser.PhoneNumber.String,
		Address:       dataUser.Address.String,
		Role:          dataUser.Role,
		Provider:      dataUser.Identities.Provider,
		LastLogin:     dataUser.UserLog.LastLogin.Time,
	}, nil
}

func (s *services) LoginUsingLine(dataLine *modelresponse.LineVerifyIdToken) (*modelresponse.User, error) {
	lineID := dataLine.Sub
	userData, err := s.userRepo.GetByProviderID(lineID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			userData := new(modeluser.User)

			userData.Identities.Provider = "line"
			userData.Identities.ProviderID = dataLine.Sub
			userData.Picture.Scan(dataLine.Picture)
			userData.Name = dataLine.Name
			userData.Email = fmt.Sprintf("%s@line.com", dataLine.Name)

			return s.Create(*userData)

		} else {
			return nil, err
		}

	}

	userData.UserLog.UserID = userData.ID

	go func() {
		_ = s.userRepo.UpdateUserLog(&userData.UserLog)
	}()

	return &modelresponse.User{
		ID:            userData.ID,
		CreatedAt:     userData.CreatedAt,
		Name:          userData.Name,
		Picture:       userData.Picture.String,
		Email:         userData.Email,
		EmailVerified: userData.EmailVerified,
		PhoneNumber:   userData.PhoneNumber.String,
		Address:       userData.Address.String,
		Role:          userData.Role,
		Provider:      userData.Identities.Provider,
		LastLogin:     userData.UserLog.LastLogin.Time,
	}, nil
}
