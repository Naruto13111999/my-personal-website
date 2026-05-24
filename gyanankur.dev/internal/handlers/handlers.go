package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"gyanankur.dev/internal/data"
)

type Handler struct {
	templates *template.Template
	profile   data.Profile
}

func New() (*Handler, error) {
	funcMap := template.FuncMap{
		"techBadge": data.TechBadgeFor,
		"techBadges": data.TechBadges,
		"initial": func(s string) string {
			if s == "" {
				return "?"
			}
			return string([]rune(s)[0])
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("web/templates/*.html")
	if err != nil {
		return nil, err
	}
	return &Handler{
		templates: tmpl,
		profile:   data.GetProfile(),
	}, nil
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.RenderIndex(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) RenderIndex(w io.Writer) error {
	return h.templates.ExecuteTemplate(w, "index.html", h.profile)
}

func (h *Handler) APIProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
