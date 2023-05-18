package location_validation

import (
	"bytes"
	"context"
	"encoding/json"
	common_module "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	utils_common "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/location/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/location/commands"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/location/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/location/models"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type LocationEventHandler struct {
	repositories     *repository.Repositories
	locationCommands *commands.LocationCommands
	log              logger.Logger
	cfg              *config.Config
}

type LocationValidateRequest struct {
	Address string `json:"address" validate:"required"`
}

type LocationValidationResponseV1 struct {
	Address *ValidatedAddress `json:"address"`
	Valid   bool              `json:"valid"`
	Error   *string           `json:"error"`
}
type ValidatedAddress struct {
	Country      string   `json:"country"`
	Region       string   `json:"region"`
	District     string   `json:"district"`
	Locality     string   `json:"locality"`
	Street       string   `json:"street"`
	Zip          string   `json:"zip"`
	PostalCode   string   `json:"postalCode"`
	AddressLine1 string   `json:"addressLine1"`
	AddressLine2 string   `json:"addressLine2"`
	AddressType  string   `json:"addressType"`
	HouseNumber  string   `json:"houseNumber"`
	PlusFour     string   `json:"plusFour"`
	Commercial   bool     `json:"commercial"`
	Predirection string   `json:"predirection"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
	TimeZone     string   `json:"timeZone"`
	UtcOffset    int      `json:"utcOffset"`
}

func (h *LocationEventHandler) OnLocationCreate(ctx context.Context, evt eventstore.Event) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LocationEventHandler.OnLocationCreate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.LocationCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	tenant := eventData.Tenant
	locationId := aggregate.GetLocationObjectID(evt.AggregateID, tenant)

	if eventData.RawAddress == "" && eventData.LocationAddress.Address1 == "" && (eventData.LocationAddress.Street == "" || eventData.LocationAddress.HouseNumber == "") {
		h.locationCommands.SkipLocationValidation.Handle(ctx, commands.NewSkippedLocationValidationCommand(locationId, tenant, "", "Missing raw Address"))
	} else {
		rawAddress := strings.TrimSpace(eventData.RawAddress)
		if rawAddress == "" {
			rawAddress = constructRawAddressForValidationFromLocationAddressFields(eventData)
		}

		locationValidateRequest := LocationValidateRequest{
			Address: rawAddress,
		}

		preValidationErr := validator.GetValidator().Struct(locationValidateRequest)
		if preValidationErr != nil {
			tracing.TraceErr(span, preValidationErr)
			h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, preValidationErr.Error())
			return nil
		} else {
			evJSON, err := json.Marshal(locationValidateRequest)
			if err != nil {
				tracing.TraceErr(span, err)
				h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, err.Error())
				return nil
			}
			requestBody := []byte(string(evJSON))
			req, err := http.NewRequest("POST", h.cfg.Services.ValidationApi+"/validateAddress", bytes.NewBuffer(requestBody))
			if err != nil {
				tracing.TraceErr(span, err)
				h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, err.Error())
				return nil
			}
			// Set the request headers
			req.Header.Set(common_module.ApiKeyHeader, h.cfg.Services.ValidationApiKey)
			req.Header.Set(common_module.TenantHeader, tenant)

			// Make the HTTP request
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				tracing.TraceErr(span, err)
				h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, err.Error())
				return nil
			}
			defer response.Body.Close()
			var result LocationValidationResponseV1
			err = json.NewDecoder(response.Body).Decode(&result)
			if err != nil {
				tracing.TraceErr(span, err)
				h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, err.Error())
				return nil
			}
			if !result.Valid {
				h.sendLocationFailedValidationEvent(ctx, tenant, locationId, rawAddress, *result.Error)
				return nil
			}

			locationAddressFields := models.LocationAddressFields{
				Country:      result.Address.Country,
				Region:       result.Address.Region,
				District:     result.Address.District,
				Locality:     result.Address.Locality,
				Street:       result.Address.Street,
				Address1:     result.Address.AddressLine1,
				Address2:     result.Address.AddressLine2,
				Zip:          result.Address.Zip,
				AddressType:  result.Address.AddressType,
				HouseNumber:  result.Address.HouseNumber,
				PostalCode:   result.Address.PostalCode,
				PlusFour:     result.Address.PlusFour,
				Commercial:   result.Address.Commercial,
				Predirection: result.Address.Predirection,
				Latitude:     result.Address.Latitude,
				Longitude:    result.Address.Longitude,
				TimeZone:     result.Address.TimeZone,
				UtcOffset:    result.Address.UtcOffset,
			}

			h.locationCommands.LocationValidated.Handle(ctx, commands.NewLocationValidatedCommand(locationId, tenant, rawAddress, locationAddressFields))
		}
	}

	return nil
}

func constructRawAddressForValidationFromLocationAddressFields(eventData events.LocationCreatedEvent) string {
	rawAddress :=
		eventData.LocationAddress.HouseNumber + " " +
			eventData.LocationAddress.Street + " " +
			eventData.LocationAddress.Address1 + " " +
			eventData.LocationAddress.Address2 + " " +
			utils_common.StringFirstNonEmpty(eventData.LocationAddress.Zip, eventData.LocationAddress.PostalCode) + ", " +
			eventData.LocationAddress.Locality
	if eventData.LocationAddress.Locality != "" {
		rawAddress += ","
	}
	rawAddress += " " + eventData.LocationAddress.District + " " +
		eventData.LocationAddress.Region + " " +
		eventData.LocationAddress.Country
	return rawAddress
}

func (h *LocationEventHandler) sendLocationFailedValidationEvent(ctx context.Context, tenant, locationId, rawAddress, error string) {
	h.log.Errorf("Failed to validate location %s for tenant %s: %s", locationId, tenant, error)
	h.locationCommands.FailedLocationValidation.Handle(ctx, commands.NewFailedLocationValidationCommand(locationId, tenant, rawAddress, error))
}
