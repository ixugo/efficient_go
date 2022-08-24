package core

import (
	"context"
	"project/business/store/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Core struct {
	log  *zap.SugaredLogger
	user user.Store
}

func NewCore(log *zap.SugaredLogger, db *gorm.DB) Core {
	return Core{
		log:  log,
		user: user.NewStore(db),
	}
}

func (c Core) Create(ctx context.Context) error {
	c.user.Create()
	return nil
}
