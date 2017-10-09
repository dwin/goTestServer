package model

import (
	"encoding/json"
	"fmt"
	"time"

	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	_ "github.com/lib/pq"
)

type Request struct {
	ID          int64
	TokenID     string
	ReqMethod   string
	ClientIP    string
	ReqHeaders  json.RawMessage
	RespHeaders json.RawMessage
	CreatedAt   time.Time
}

func SaveRequest(tokenID string, c *gin.Context) {
	reqTime := time.Now()
	reqH := make(map[string]string)
	for k, v := range c.Request.Header {
		reqH[k] = fmt.Sprintf("%v", v)
	}
	reqHJSON, err := jsoniter.Marshal(&reqH)
	if err != nil {
		u.Log.Error().Err(err).Msg("Could not marshall request headers to JSON")
	}
	respH := make(map[string]string)
	respHeaders := c.Writer.Header()
	for k, v := range respHeaders {
		respH[k] = fmt.Sprintf("%v", v)
	}
	respHJSON, err := jsoniter.Marshal(&respH)
	if err != nil {
		u.Log.Error().Err(err).Msg("Could not marshall response headers to JSON")
	}
	req := Request{
		TokenID:     tokenID,
		ReqMethod:   c.Request.Method,
		ClientIP:    c.ClientIP(),
		ReqHeaders:  reqHJSON,
		RespHeaders: respHJSON,
		CreatedAt:   reqTime,
	}

	sStmt := "INSERT INTO request (token_id, req_method, client_ip, req_headers, resp_headers, created) values ($1, $2, $3, $4, $5, $6)"
	db := getDBconn()
	pStmt, err := db.Prepare(sStmt)
	defer pStmt.Close()
	defer db.Close()
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to prepare request statement to DB")
	}
	_, err = pStmt.Exec(req.TokenID, req.ReqMethod, req.ClientIP, req.ReqHeaders, req.RespHeaders, req.CreatedAt)
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to execute statement to insert new request to DB")
	}
}

func GetAllTokenRequests(tokenID string) (requests []Request, err error) {
	db := getDBconn()
	stmt, err := db.Prepare(`SELECT * FROM request WHERE token_id=$1`)
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to prepare statement for Get All Token Requests")
	}
	defer stmt.Close()
	defer db.Close()
	rows, err := stmt.Query(tokenID)
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to get query result for get all token requests")
		return requests, err
	}
	defer rows.Close()
	for rows.Next() {
		var req Request
		if err := rows.Scan(&req.ID, &req.TokenID, &req.ReqMethod, &req.ClientIP, &req.ReqHeaders, &req.RespHeaders, &req.CreatedAt); err != nil {
			u.Log.Error().Err(err).Msg("Error scanning query response rows")
			return requests, err
		}
		requests = append(requests, req)
	}
	err = rows.Err()
	if err != nil {
		u.Log.Error().Err(err)
		return requests, err
	}
	return requests, err
}
