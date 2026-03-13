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

func (s HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.
		Path("/tasks").
		Methods(http.MethodPost).
		HandlerFunc(s.httpServer.HandleCreateTask)

	router.
		Path("/tasks").
		Methods(http.MethodGet).
		HandlerFunc(s.httpServer.HandleGetTasks)

	router.
		Path("/tasks/{title}").
		Methods(http.MethodPatch).
		HandlerFunc(s.httpServer.HandleDoneTask)

	router.
		Path("/tasks/{title}").
		Methods(http.MethodDelete).
		HandlerFunc(s.httpServer.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
