package breed

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/simply-alliv/tigris-go-explore/pkg/shared/params"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func GetAllBreeds(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 1
		}
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 20
		}
		paginate, err := strconv.ParseBool(r.URL.Query().Get("paginate"))
		if err != nil {
			paginate = true
		}
		creationType := r.URL.Query().Get("creationType")
		qp := params.PaginationQueryParams{
			Page:     page,
			Limit:    limit,
			Paginate: paginate,
		}
		var bqp params.BreedQueryParams
		if creationType != "" {
			bqp = params.BreedQueryParams{CreationType: &creationType}
		} else {
			bqp = params.BreedQueryParams{CreationType: nil}
		}
		fmt.Printf("GET /breeds - PaginationQueryParams: %+v - BreedQueryParams: %+v\n", qp, bqp)
		data, _, err := s.GetAllBreeds(r.Context(), qp, bqp)
		if err != nil {
			log.Fatalf("Unable to get all breeds: %+v\n", err)
		}

		// create a new Response struct
		response := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    data,
		}
		writeResponse(w, response)
	}
}

func GetSingleBreed(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			panic(ErrBadRouting)
		}

		data, err := s.GetSingleBreed(r.Context(), id)
		if err != nil {
			log.Fatalf("Unable to get single breed by id (%s): %+v\n", id, err)
		}

		fmt.Printf("GET /breed/%s\n", id)
		// create a new Response struct
		response := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    data,
		}
		writeResponse(w, response)
	}
}

func CreateSingleBreed(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto CreateBreed
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			panic(err)
		}
		// Set default timestamps
		dto.CreatedAt = time.Now().UTC()
		dto.UpdatedAt = time.Now().UTC()

		fmt.Printf("POST /breeds - CreateBreedDTO: %+v\n", dto)
		data, err := s.CreateSingleBreed(r.Context(), dto)
		if err != nil {
			log.Fatalf("Unable to create single breed %+v\n", err)
		}

		// create a new Response struct
		response := Response{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    data,
		}
		writeResponse(w, response)
	}
}

func UpdateSingleBreed(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			panic(ErrBadRouting)
		}

		var dto UpdateBreed
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			panic(err)
		}
		// Set default timestamps
		dto.UpdatedAt = time.Now().UTC()

		fmt.Printf("PATCH /breeds - UpdateBreedDTO: %+v\n", dto)
		data, err := s.UpdateSingleBreed(r.Context(), id, dto)
		if err != nil {
			log.Fatalf("Unable to create single breed by id (%s): %+v\n", id, err)
		}

		// create a new Response struct
		response := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    data,
		}
		writeResponse(w, response)
	}
}

func DeleteeSingleBreed(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			panic(ErrBadRouting)
		}

		fmt.Printf("DELETE /breeds/%s\n", id)
		err := s.DeleteSingleBreed(r.Context(), id)
		if err != nil {
			log.Fatalf("Unable to delete single breed by id (%s): %+v\n", id, err)
		}

		// create a new Response struct
		response := Response{
			Status:  http.StatusOK,
			Message: "success",
			Data:    nil,
		}
		writeResponse(w, response)
	}
}

func writeResponse(w http.ResponseWriter, response Response) {
	// set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// encode the response struct as JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		// handle the error
		http.Error(w, "error encoding JSON response", http.StatusInternalServerError)
		return
	}
}
