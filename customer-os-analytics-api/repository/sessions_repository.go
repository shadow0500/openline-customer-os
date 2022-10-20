package repository

import (
	"github.com.openline-ai.customer-os-analytics-api/repository/entity"
	"github.com.openline-ai.customer-os-analytics-api/repository/helper"
	"gorm.io/gorm"
)

type SessionsRepo struct {
	db *gorm.DB
}

type SessionsRepository interface {
	FindAllByApplication(appIdentifier entity.ApplicationUniqueIdentifier, page int, limit int) helper.QueryResult
}

func NewSessionsRepo(db *gorm.DB) *SessionsRepo {
	return &SessionsRepo{db: db}
}

func (r *SessionsRepo) FindAllByApplication(appIdentifier entity.ApplicationUniqueIdentifier, page int, limit int) helper.QueryResult {
	var sessions entity.SessionEntities

	pagination := helper.Pagination{
		Limit: limit,
		Page:  page,
	}

	find := r.db.
		Where(&entity.SessionEntity{Tenant: appIdentifier.Tenant, AppId: appIdentifier.AppId, TrackerName: appIdentifier.TrackerName})

	err := find.Scopes(helper.Paginate(sessions, &pagination, find)).
		Order("start_tstamp DESC").
		Find(&sessions).
		Error

	if err != nil {
		return helper.QueryResult{Error: err}
	}

	pagination.Rows = &sessions

	return helper.QueryResult{Result: &pagination}
}
