package api

import (
	"encoding/json"
	"net/http"
	"test-go-server/logger"
	resource "test-go-server/pkg/resource"

	"github.com/go-chi/chi/v5"
)

type API struct {
	Resources resource.Resources
	Logger    logger.Logger
}

func (a *API) List(w http.ResponseWriter, r *http.Request) {
	rs := a.Resources.List()
	c, err := json.Marshal(rs)
	if err != nil {
		a.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(c))
	w.Header().Set("Content-Type", "application/json")
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	rs := a.Resources

	resource_id := chi.URLParam(r, "resource")
	data := r.URL.Query().Get("data")

	resource := rs.Create(resource_id, data)
	c, err := json.Marshal(resource)
	if err != nil {
		a.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// XXX: Needs to be done before writing body, else it will be ignored
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(c))
}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	rs := a.Resources

	resource_id := chi.URLParam(r, "resource")

	resource := rs.Get(resource_id)
	c, err := json.Marshal(resource)
	if err != nil {
		a.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(c))
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	rs := a.Resources

	resource_id := chi.URLParam(r, "resource")
	data := r.URL.Query().Get("data")

	resource := rs.Update(resource_id, data)
	c, err := json.Marshal(resource)
	if err != nil {
		a.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(c))
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	rs := a.Resources

	resource_id := chi.URLParam(r, "resource")

	rs.Delete(resource_id)

	w.Write([]byte(resource_id))
	w.Header().Set("Content-Type", "application/json")
}

func NewAPI(rs resource.Resources, l logger.Logger) API {
	return API{rs, l}
}
