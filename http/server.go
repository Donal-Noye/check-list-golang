package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpServer *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpServer: httpHandlers,
	}
}

func (h HTTPHandlers) StartServer() error {
	router := mux.NewRouter()

	router.
		Path("/tasks").
		Methods(http.MethodGet).
		HandlerFunc(h.HandleCreateTask)

	router.
		Path("/tasks").
		Methods(http.MethodGet).
		HandlerFunc(h.HandleGetTasks)

	router.
		Path("/tasks/{title}").
		Methods(http.MethodPatch).
		HandlerFunc(h.HandleDoneTask)

	router.
		Path("/tasks/{title}").
		Methods(http.MethodDelete).
		HandlerFunc(h.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
