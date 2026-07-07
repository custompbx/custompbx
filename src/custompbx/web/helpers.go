package web

import (
	"custompbx/mainStruct"
	"custompbx/webStruct"

	"github.com/custompbx/customorm"
)

const (
	defaultPageLimit = 250
	maxPageLimit     = 5000
	configPageLimit  = 25
	configMaxLimit   = 250
)

type paginatedResult struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

func errorResponse(event, message string) webStruct.UserResponse {
	return webStruct.UserResponse{Error: message, MessageType: event}
}

func dataResponse(event string, data interface{}) webStruct.UserResponse {
	return webStruct.UserResponse{Data: data, MessageType: event}
}

func normalizePagination(limit, page int) (int, int) {
	return normalizePaginationWithBounds(limit, page, defaultPageLimit, maxPageLimit)
}

func normalizeConfigPagination(limit, page int) (int, int) {
	return normalizePaginationWithBounds(limit, page, configPageLimit, configMaxLimit)
}

func normalizePaginationWithBounds(limit, page, defaultLimit, maxLimit int) (int, int) {
	if defaultLimit <= 0 {
		defaultLimit = defaultPageLimit
	}
	if maxLimit < defaultLimit {
		maxLimit = defaultLimit
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if page < 0 {
		page = 0
	}
	return limit, page * limit
}

func hasPagedRequest(request mainStruct.DBRequest) bool {
	return request.Limit != 0 || request.Filters != nil
}

func buildFilteredConfigRequest(base map[string]customorm.FilterFields, request mainStruct.DBRequest) customorm.Filters {
	limit, offset := normalizeConfigPagination(request.Limit, request.Offset)
	fields := cloneFilterFields(base)
	for _, filter := range request.Filters {
		fields[filter.Field] = customorm.FilterFields{
			Flag:     true,
			UseValue: true,
			Value:    filter.FieldValue,
			Operand:  filter.Operand,
		}
	}
	orderFields := append([]string(nil), request.Order.Fields...)
	return customorm.Filters{
		Fields: fields,
		Limit:  limit,
		Offset: offset,
		Order:  customorm.Order{Desc: request.Order.Desc, Fields: orderFields},
	}
}

func cloneFilterFields(fields map[string]customorm.FilterFields) map[string]customorm.FilterFields {
	cloned := make(map[string]customorm.FilterFields, len(fields))
	for key, value := range fields {
		cloned[key] = value
	}
	return cloned
}
