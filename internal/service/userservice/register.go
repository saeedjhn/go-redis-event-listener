package userservice

import (
	"context"
	"errors"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-redis-event-listener/internal/domain/entity"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/kind"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-redis-event-listener/internal/infrastructure/security/bcrypt"
	"github.com/saeedjhn/go-redis-event-listener/pkg/message"
)

func (u *UserInteractor) Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	const op = message.OpUserUsecaseRegister

	isUnique, err := u.repository.IsMobileUnique(req.Mobile)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	if !isUnique {
		return userdto.RegisterResponse{},
			richerror.New(op).
				WithErr(errors.New(message.ErrorMsgMobileIsNotUnique)). //nolint:err113
				WithMessage(message.ErrorMsgInvalidInput).
				WithKind(kind.KindStatusBadRequest)
	}

	user := entity.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptedPass, _ := bcrypt.Generate(req.Password, bcrypt.DefaultCost) // Check err
	user.Password = encryptedPass

	createdUser, err := u.repository.Create(user)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	// Fire event => 'user.event.created'
	u.event.Created(context.Background(), createdUser) // Check err

	return userdto.RegisterResponse{
		User: userdto.UserInfo{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Mobile:    createdUser.Mobile,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		}, // Or
		//User: createdUser.ToUserInfoDTO(),
	}, nil
}
