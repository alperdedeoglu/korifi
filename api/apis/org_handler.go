package apis

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	controllerruntime "sigs.k8s.io/controller-runtime"

	"code.cloudfoundry.org/cf-k8s-controllers/api/authorization"
	"code.cloudfoundry.org/cf-k8s-controllers/api/payloads"
	"code.cloudfoundry.org/cf-k8s-controllers/api/presenter"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
	"code.cloudfoundry.org/cf-k8s-controllers/controllers/webhooks/workloads"
)

const (
	OrgsEndpoint = "/v3/organizations"
)

//counterfeiter:generate -o fake -fake-name OrgRepositoryProvider . OrgRepositoryProvider

type OrgRepositoryProvider interface {
	OrgRepoForRequest(request *http.Request) (repositories.CFOrgRepository, error)
}

type OrgHandler struct {
	logger          logr.Logger
	apiBaseURL      url.URL
	orgRepoProvider OrgRepositoryProvider
}

func NewOrgHandler(apiBaseURL url.URL, orgRepoProvider OrgRepositoryProvider) *OrgHandler {
	return &OrgHandler{
		logger:          controllerruntime.Log.WithName("Org Handler"),
		apiBaseURL:      apiBaseURL,
		orgRepoProvider: orgRepoProvider,
	}
}

func (h *OrgHandler) orgCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload payloads.OrgCreate
	rme := decodeAndValidateJSONPayload(r, &payload)
	if rme != nil {
		writeRequestMalformedErrorResponse(w, rme)

		return
	}

	org := payload.ToMessage()
	org.GUID = uuid.NewString()

	orgRepo, err := h.orgRepoProvider.OrgRepoForRequest(r)
	if err != nil {
		if authorization.IsInvalidAuth(err) {
			h.logger.Error(err, "unauthorized to create org")
			writeInvalidAuthErrorResponse(w)

			return
		}

		if authorization.IsNotAuthenticated(err) {
			h.logger.Error(err, "unauthorized to create org")
			writeNotAuthenticatedErrorResponse(w)

			return
		}

		h.logger.Error(err, "failed to create org repo for the authorization header")
		writeUnknownErrorResponse(w)

		return
	}

	record, err := orgRepo.CreateOrg(r.Context(), org)
	if err != nil {
		if workloads.HasErrorCode(err, workloads.DuplicateOrgNameError) {
			errorDetail := fmt.Sprintf("Organization '%s' already exists.", org.Name)
			h.logger.Info(errorDetail)
			writeUnprocessableEntityError(w, errorDetail)
			return
		}
		h.logger.Error(err, "Failed to create org", "Org Name", payload.Name)
		writeUnknownErrorResponse(w)
		return
	}

	err = writeJsonResponse(w, presenter.ForCreateOrg(record, h.apiBaseURL), http.StatusCreated)
	if err != nil {
		h.logger.Error(err, "Failed to render response")
		writeUnknownErrorResponse(w)
	}
}

func (h *OrgHandler) orgListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	var names []string
	namesList := r.URL.Query().Get("names")
	if len(namesList) > 0 {
		names = strings.Split(namesList, ",")
	}

	orgRepo, err := h.orgRepoProvider.OrgRepoForRequest(r)
	if err != nil {
		if authorization.IsNotAuthenticated(err) {
			h.logger.Error(err, "unauthorized to list orgs")
			writeNotAuthenticatedErrorResponse(w)

			return
		}

		if authorization.IsInvalidAuth(err) {
			h.logger.Error(err, "unauthorized to list orgs")
			writeInvalidAuthErrorResponse(w)

			return
		}

		h.logger.Error(err, "failed to create org repo for the authorization header")
		writeUnknownErrorResponse(w)

		return
	}

	orgs, err := orgRepo.ListOrgs(ctx, names)
	if err != nil {
		h.logger.Error(err, "failed to fetch orgs")
		writeUnknownErrorResponse(w)

		return
	}

	err = writeJsonResponse(w, presenter.ForOrgList(orgs, h.apiBaseURL, *r.URL), http.StatusOK)
	if err != nil {
		h.logger.Error(err, "Failed to render response")
		writeUnknownErrorResponse(w)
	}
}

func (h *OrgHandler) RegisterRoutes(router *mux.Router) {
	router.Path(OrgsEndpoint).Methods("GET").HandlerFunc(h.orgListHandler)
	router.Path(OrgsEndpoint).Methods("POST").HandlerFunc(h.orgCreateHandler)
}
