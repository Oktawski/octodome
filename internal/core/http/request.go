package corehttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"octodome/internal/core"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParseJSON(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

func GetPathParam[T any](r *http.Request, key string) (T, error) {
	var zero T
	val := chi.URLParam(r, key)
	if val == "" {
		return zero, fmt.Errorf("missing path param: %s", key)
	}

	var result any
	var err error

	switch any(zero).(type) {
	case int:
		result, err = strconv.Atoi(val)
	case int64:
		result, err = strconv.ParseInt(val, 10, 64)
	case string:
		result = val
	default:
		return zero, fmt.Errorf("unsupported type for path param: %T", zero)
	}

	if err != nil {
		return zero, fmt.Errorf("invalid value for %s: %v", key, err)
	}

	return result.(T), nil
}

func GetQueryParam[T any](r *http.Request, key string) (T, error) {
	var zero T
	val := r.URL.Query().Get(key)
	if val == "" {
		return zero, fmt.Errorf("missing query param: %s", key)
	}

	var result any
	var err error

	switch any(zero).(type) {
	case int:
		result, err = strconv.Atoi(val)
	case int64:
		result, err = strconv.ParseInt(val, 10, 64)
	case bool:
		result, err = strconv.ParseBool(val)
	case string:
		result = val
	default:
		return zero, fmt.Errorf("unsupported type for query param: %T", zero)
	}

	if err != nil {
		return zero, fmt.Errorf("invalid value for %s: %v", key, err)
	}

	return result.(T), nil
}

func GetQueryParamOrDefault[T any](
	r *http.Request,
	key string,
	defaultVal T) T {

	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}

	var result any
	switch any(defaultVal).(type) {
	case int:
		if i, err := strconv.Atoi(val); err == nil {
			result = i
		} else {
			result = defaultVal
		}
	case bool:
		if b, err := strconv.ParseBool(val); err == nil {
			result = b
		} else {
			result = defaultVal
		}
	case string:
		result = val
	default:
		result = defaultVal
	}

	return result.(T)
}

func GetPagination(r *http.Request) core.Pagination {
	page := GetQueryParamOrDefault(r, "page", 1)
	pageSize := GetQueryParamOrDefault(r, "pageSize", 100)

	return core.Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

func GetID(r *http.Request) (uint, error) {
	id, err := GetPathParam[int](r, "id")
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
