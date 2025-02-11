package service

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	commonService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
)

type Services struct {
	CommonServices *commonService.Services

	ContactService                  ContactService
	OrganizationService             OrganizationService
	CustomFieldService              CustomFieldService
	PhoneNumberService              PhoneNumberService
	EmailService                    EmailService
	UserService                     UserService
	FieldSetService                 FieldSetService
	EntityTemplateService           EntityTemplateService
	FieldSetTemplateService         FieldSetTemplateService
	CustomFieldTemplateService      CustomFieldTemplateService
	ConversationService             ConversationService
	TimelineEventService            TimelineEventService
	NoteService                     NoteService
	JobRoleService                  JobRoleService
	LocationService                 LocationService
	TagService                      TagService
	SearchService                   SearchService
	QueryService                    DashboardService
	DomainService                   DomainService
	IssueService                    IssueService
	InteractionSessionService       InteractionSessionService
	InteractionEventService         InteractionEventService
	PageViewService                 PageViewService
	AnalysisService                 AnalysisService
	AttachmentService               AttachmentService
	MeetingService                  MeetingService
	TenantService                   TenantService
	WorkspaceService                WorkspaceService
	SocialService                   SocialService
	PlayerService                   PlayerService
	OrganizationRelationshipService OrganizationRelationshipService
}

func InitServices(log logger.Logger, driver *neo4j.DriverWithContext, commonServices *commonService.Services, grpcClients *grpc_client.Clients) *Services {
	repositories := repository.InitRepos(driver)

	services := Services{
		CommonServices:                  commonServices,
		ContactService:                  NewContactService(log, repositories, grpcClients),
		OrganizationService:             NewOrganizationService(log, repositories, grpcClients),
		CustomFieldService:              NewCustomFieldService(log, repositories),
		PhoneNumberService:              NewPhoneNumberService(log, repositories, grpcClients),
		EmailService:                    NewEmailService(log, repositories),
		UserService:                     NewUserService(log, repositories, grpcClients),
		FieldSetService:                 NewFieldSetService(log, repositories),
		EntityTemplateService:           NewEntityTemplateService(log, repositories),
		FieldSetTemplateService:         NewFieldSetTemplateService(log, repositories),
		CustomFieldTemplateService:      NewCustomFieldTemplateService(log, repositories),
		ConversationService:             NewConversationService(log, repositories),
		JobRoleService:                  NewJobRoleService(log, repositories),
		LocationService:                 NewLocationService(log, repositories),
		TagService:                      NewTagService(log, repositories),
		DomainService:                   NewDomainService(log, repositories),
		IssueService:                    NewIssueService(log, repositories),
		PageViewService:                 NewPageViewService(log, repositories),
		AttachmentService:               NewAttachmentService(log, repositories),
		TenantService:                   NewTenantService(log, repositories),
		WorkspaceService:                NewWorkspaceService(log, repositories),
		SocialService:                   NewSocialService(log, repositories),
		OrganizationRelationshipService: NewOrganizationRelationshipService(log, repositories),
	}
	services.NoteService = NewNoteService(log, repositories, &services)
	services.TimelineEventService = NewTimelineEventService(log, repositories, &services)
	services.SearchService = NewSearchService(log, repositories, &services)
	services.QueryService = NewDashboardService(log, repositories, &services)
	services.InteractionEventService = NewInteractionEventService(log, repositories, &services)
	services.InteractionSessionService = NewInteractionSessionService(log, repositories, &services)
	services.AnalysisService = NewAnalysisService(log, repositories, &services)
	services.MeetingService = NewMeetingService(log, repositories, &services)
	services.PlayerService = NewPlayerService(repositories, &services)

	return &services
}
