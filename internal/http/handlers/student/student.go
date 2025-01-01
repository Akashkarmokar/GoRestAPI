package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Akashkarmokar/GoRestAPI/internal/http/types"
	"github.com/Akashkarmokar/GoRestAPI/internal/http/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Creating a student")

		response.WriteJson(w, http.StatusCreated, student)
	}
}
