package resource

import (
	"context"
	"log"
	"net/http"
)

type Resource struct {
	ID   string
	Data string
}

type resources struct {
	rs map[string]*Resource
}

type Resources interface {
	List() []*Resource
	Get(string) *Resource
	Create(string, string) *Resource
	Update(string, string) *Resource
	Delete(string) bool
}

func (r *Resource) Update(data string) *Resource {
	r.Data = data
	return r
}

func (rs *resources) Get(id string) *Resource {
	r, ok := rs.rs[id]
	log.Println(r)
	if !ok {
		return new(Resource)
	}
	return r
}

func (rs *resources) Create(id string, data string) *Resource {
	r := new(Resource)
	r.ID = id
	r.Data = data
	rs.rs[r.ID] = r
	return r
}

func (rs *resources) Update(id string, data string) *Resource {
	r := rs.Get(id)
	r.Update(data)
	return r
}

func (rs *resources) Delete(id string) bool {
	_, ok := rs.rs[id]
	if ok {
		delete(rs.rs, id)
		return true
	}
	return false
}

func New() *resources {
	rs := new(resources)
	rs.rs = make(map[string]*Resource)
	return rs
}

func (rs *resources) List() []*Resource {
	var v []*Resource
	for _, value := range rs.rs {
		v = append(v, value)
	}
	return v
}

type contextKey string

const ContextKey = contextKey("resources")

func Middleware(next http.Handler) http.Handler {
	var rs = New()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKey, rs)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
