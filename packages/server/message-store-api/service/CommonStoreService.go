package service

import (
	msProto "github.com/openline-ai/openline-customer-os/packages/server/message-store-api/proto/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/repository/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type commonStoreService struct {
}

type CommonStoreService interface {
	ConvertMSTypeToEntityType(channel msProto.MessageType) entity.EventType
	ConvertEntityTypeToMSType(eventType entity.EventType) msProto.MessageType

	ConvertMSSubtypeToEntitySubtype(channel msProto.MessageSubtype) entity.EventSubtype
	ConvertEntitySubtypeToMSSubtype(eventType entity.EventSubtype) msProto.MessageSubtype

	ConvertMSSenderTypeToEntitySenderType(direction msProto.SenderType) entity.SenderType
	ConvertEntitySenderTypeToMSSenderType(direction entity.SenderType) msProto.SenderType

	ConvertEntityDirectionToMSDirection(direction entity.Direction) msProto.MessageDirection
	ConvertMSDirectionToEntityDirection(direction msProto.MessageDirection) entity.Direction

	EncodeConversationToMS(conversation Conversation) *msProto.FeedItem
	EncodeMessageIdToMS(conversation Conversation) *msProto.MessageId
	EncodeConversationEventToMS(conversationEvent entity.ConversationEvent) *msProto.Message
}

func (s *commonStoreService) EncodeConversationEventToMS(conversationEvent entity.ConversationEvent) *msProto.Message {
	return &msProto.Message{
		MessageId:         s.EncodeMessageIdToMs(conversationEvent),
		InitiatorUsername: s.EncodeUsernameToMs(conversationEvent.InitiatorUsername),
		Type:              s.ConvertEntityTypeToMSType(conversationEvent.Type),
		Subtype:           s.ConvertEntitySubtypeToMSSubtype(conversationEvent.Subtype),
		Content:           conversationEvent.Content,
		Direction:         s.ConvertEntityDirectionToMSDirection(conversationEvent.Direction),
		Time:              timestamppb.New(conversationEvent.CreateDate),
		SenderId:          conversationEvent.SenderId,
		SenderType:        s.ConvertEntitySenderTypeToMSSenderType(conversationEvent.SenderType),
		SenderUsername:    s.EncodeUsernameToMs(conversationEvent.SenderUsername),
	}
}

func (s *commonStoreService) EncodeMessageIdToMs(conversationEvent entity.ConversationEvent) *msProto.MessageId {
	return &msProto.MessageId{ConversationEventId: conversationEvent.ID, ConversationId: conversationEvent.ConversationId}
}

func (s *commonStoreService) EncodeUsernameToMs(username string) *msProto.ParticipantId {
	if strings.HasPrefix(username, "mailto:") {
		return &msProto.ParticipantId{Type: msProto.ParticipantIdType_MAILTO, Identifier: strings.TrimPrefix(username, "mailto:")}
	} else if strings.HasPrefix(username, "tel:") {
		return &msProto.ParticipantId{Type: msProto.ParticipantIdType_TEL, Identifier: strings.TrimPrefix(username, "tel:")}
	} else {
		// fail back to assuming email id
		return &msProto.ParticipantId{Type: msProto.ParticipantIdType_MAILTO, Identifier: username}
	}
}

func (s *commonStoreService) ConvertMSParticipantIdToUsername(participantId *msProto.ParticipantId) string {
	if participantId == nil {
		return ""
	}
	if participantId.Type == msProto.ParticipantIdType_MAILTO {
		return "mailto:" + participantId.Identifier
	} else if participantId.Type == msProto.ParticipantIdType_TEL {
		return "tel:" + participantId.Identifier
	} else {
		// shouldn't reach here
		return participantId.Identifier
	}
}

func (s *commonStoreService) EncodeConversationToMS(conversation Conversation) *msProto.FeedItem {
	return &msProto.FeedItem{
		Id:                  conversation.Id,
		InitiatorFirstName:  conversation.InitiatorFirstName,
		InitiatorLastName:   conversation.InitiatorLastName,
		InitiatorUsername:   s.EncodeUsernameToMs(conversation.InitiatorUsername),
		InitiatorType:       conversation.InitiatorType,
		LastSenderFirstName: conversation.LastSenderFirstName,
		LastSenderLastName:  conversation.LastSenderLastName,
		LastContentPreview:  conversation.LastContentPreview,
		LastTimestamp:       timestamppb.New(conversation.UpdatedAt),
	}
}

func (s *commonStoreService) ConvertMSTypeToEntityType(channel msProto.MessageType) entity.EventType {
	switch channel {
	case msProto.MessageType_WEB_CHAT:
		return entity.WEB_CHAT
	case msProto.MessageType_EMAIL:
		return entity.EMAIL
	case msProto.MessageType_VOICE:
		return entity.VOICE
	default:
		return entity.WEB_CHAT
	}
}

func (s *commonStoreService) ConvertEntityTypeToMSType(eventType entity.EventType) msProto.MessageType {
	switch eventType {
	case entity.WEB_CHAT:
		return msProto.MessageType_WEB_CHAT
	case entity.EMAIL:
		return msProto.MessageType_EMAIL
	case entity.VOICE:
		return msProto.MessageType_VOICE
	default:
		return msProto.MessageType_WEB_CHAT
	}
}

func (s *commonStoreService) ConvertMSSubtypeToEntitySubtype(subtype msProto.MessageSubtype) entity.EventSubtype {
	switch subtype {
	case msProto.MessageSubtype_MESSAGE:
		return entity.TEXT
	case msProto.MessageSubtype_FILE:
		return entity.FILE
	default:
		return entity.TEXT
	}
}

func (s *commonStoreService) ConvertEntitySubtypeToMSSubtype(eventType entity.EventSubtype) msProto.MessageSubtype {
	switch eventType {
	case entity.TEXT:
		return msProto.MessageSubtype_MESSAGE
	case entity.FILE:
		return msProto.MessageSubtype_FILE
	default:
		return msProto.MessageSubtype_MESSAGE
	}
}

func (s *commonStoreService) ConvertMSSenderTypeToEntitySenderType(direction msProto.SenderType) entity.SenderType {
	switch direction {
	case msProto.SenderType_CONTACT:
		return entity.CONTACT
	case msProto.SenderType_USER:
		return entity.USER
	default:
		return entity.CONTACT
	}
}

func (s *commonStoreService) ConvertEntitySenderTypeToMSSenderType(direction entity.SenderType) msProto.SenderType {
	switch direction {
	case entity.CONTACT:
		return msProto.SenderType_CONTACT
	case entity.USER:
		return msProto.SenderType_USER
	default:
		return msProto.SenderType_CONTACT
	}
}

func (s *commonStoreService) ConvertEntityDirectionToMSDirection(direction entity.Direction) msProto.MessageDirection {
	switch direction {
	case entity.INBOUND:
		return msProto.MessageDirection_INBOUND
	case entity.OUTBOUND:
		return msProto.MessageDirection_OUTBOUND
	default:
		return msProto.MessageDirection_INBOUND
	}
}

func (s *commonStoreService) ConvertMSDirectionToEntityDirection(direction msProto.MessageDirection) entity.Direction {
	switch direction {
	case msProto.MessageDirection_INBOUND:
		return entity.INBOUND
	case msProto.MessageDirection_OUTBOUND:
		return entity.OUTBOUND
	default:
		return entity.INBOUND
	}
}

func NewCommonStoreService() *commonStoreService {
	commonStoreService := new(commonStoreService)
	return commonStoreService
}
