package entity

import (
	"fmt"
	"time"
)

type InteractionEventEntity struct {
	Id            string
	CreatedAt     time.Time
	Channel       string
	Content       string
	ContentType   string
	Source        DataSource
	SourceOfTruth DataSource
	AppSource     string

	DataloaderKey string
}

func (interactionEventEntity InteractionEventEntity) ToString() string {
	return fmt.Sprintf("id: %s", interactionEventEntity.Id)
}

type InteractionEventEntities []InteractionEventEntity

func (InteractionEventEntity) Labels(tenant string) []string {
	return []string{"InteractionEvent", "InteractionEvent_" + tenant}
}
