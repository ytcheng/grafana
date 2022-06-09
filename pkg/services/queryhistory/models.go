package queryhistory

import (
	"errors"

	"github.com/grafana/grafana/pkg/components/simplejson"
)

var (
	ErrQueryNotFound        = errors.New("query in query history not found")
	ErrStarredQueryNotFound = errors.New("starred query not found")
	ErrQueryAlreadyStarred  = errors.New("query was already starred")
)

// QueryHistory is the model for query history definitions
type QueryHistory struct {
	ID            int64  `xorm:"pk autoincr 'id'"`
	UID           string `xorm:"uid"`
	DatasourceUID string `xorm:"datasource_uid"`
	OrgID         int64  `xorm:"org_id"`
	CreatedBy     int64
	CreatedAt     int64
	Comment       string
	Queries       *simplejson.Json
}

// QueryHistory is the model for query history star definitions
type QueryHistoryStar struct {
	ID       int64  `xorm:"pk autoincr 'id'"`
	QueryUID string `xorm:"query_uid"`
	UserID   int64  `xorm:"user_id"`
}

type SearchInQueryHistoryQuery struct {
	DatasourceUIDs []string `json:"datasourceUids"`
	SearchString   string   `json:"searchString"`
	OnlyStarred    bool     `json:"onlyStarred"`
	Sort           string   `json:"sort"`
	Page           int      `json:"page"`
	Limit          int      `json:"limit"`
	From           int64    `json:"from"`
	To             int64    `json:"to"`
}

type QueryHistoryDTO struct {
	UID           string           `json:"uid" xorm:"uid"`
	DatasourceUID string           `json:"datasourceUid" xorm:"datasource_uid"`
	CreatedBy     int64            `json:"createdBy"`
	CreatedAt     int64            `json:"createdAt"`
	Comment       string           `json:"comment"`
	Queries       *simplejson.Json `json:"queries"`
	Starred       bool             `json:"starred"`
}

// QueryHistoryResponse is a response struct for QueryHistoryDTO
type QueryHistoryResponse struct {
	Result QueryHistoryDTO `json:"result"`
}

type QueryHistorySearchResult struct {
	TotalCount   int               `json:"totalCount"`
	QueryHistory []QueryHistoryDTO `json:"queryHistory"`
	Page         int               `json:"page"`
	PerPage      int               `json:"perPage"`
}

type QueryHistorySearchResponse struct {
	Result QueryHistorySearchResult `json:"result"`
}

// QueryHistoryDeleteQueryResponse is the response struct for deleting a query from query history
type QueryHistoryDeleteQueryResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type QueryToMigrate struct {
	DatasourceUID string           `json:"datasourceUid"`
	Queries       *simplejson.Json `json:"queries"`
	CreatedAt     int64            `json:"createdAt"`
	Comment       string           `json:"comment"`
	Starred       bool             `json:"starred"`
}

type QueryHistoryMigrationResponse struct {
	Message      string `json:"message"`
	TotalCount   int    `json:"totalCount"`
	StarredCount int    `json:"starredCount"`
}

// CreateQueryInQueryHistoryCommand is the command for adding query history
// swagger:model
type CreateQueryInQueryHistoryCommand struct {
	// UID of the data source for which are queries stored.
	// example: PE1C5CBDA0504A6A3
	DatasourceUID string `json:"datasourceUid"`
	// The JSON model of queries.
	// required: true
	// example: [ { "datasourceUid": "PE1C5CBDA0504A6A3", "queries": [ { "refId": "A", "key": "Q-87fed8e3-62ba-4eb2-8d2a-4129979bb4de-0", "scenarioId": "csv_content", "datasource": { "type": "testdata", "uid": "PD8C576611E62080A" } } ], "starred": false, "createdAt": 1643630762, "comment": "debugging" } ]
	Queries *simplejson.Json `json:"queries"`
}

// PatchQueryCommentInQueryHistoryCommand is the command for updating comment for query in query history
// swagger:model
type PatchQueryCommentInQueryHistoryCommand struct {
	// Updated comment
	Comment string `json:"comment"`
}

// MigrateQueriesToQueryHistoryCommand is the command used for migration of old queries into query history
// swagger:model
type MigrateQueriesToQueryHistoryCommand struct {
	// Array of queries to store in query history.
	// example: [ { "datasourceUid": "PE1C5CBDA0504A6A3", "queries": [ { "refId": "A", "key": "Q-87fed8e3-62ba-4eb2-8d2a-4129979bb4de-0", "scenarioId": "csv_content", "datasource": { "type": "testdata", "uid": "PD8C576611E62080A" } } ], "starred": false, "createdAt": 1643630762, "comment": "debugging" } ]
	Queries []QueryToMigrate `json:"queries"`
}
