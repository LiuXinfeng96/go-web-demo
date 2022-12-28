package handlers

import (
	"go-web-demo/src/db"
	"go-web-demo/src/models"
	"go-web-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlGetOperationList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")
		sortTypeStr := c.Query("sortType")
		searchInput := c.Query("searchConditions")

		err := isStringRequiredParamsEmpty(pageSizeStr, pageStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		sortType, ok := services.SortTypeValue[sortTypeStr]
		if !ok {
			sortType = services.SORTTYPE_TIME
		}

		params := &services.QueryObjectsParams{
			ModelStruct: new(db.Operation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "operator")
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, total, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.OperationInfo, 0)

		for sqlRows.Next() {
			var operation db.Operation
			err := s.ScanRows(sqlRows, &operation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OperationInfo{
				Operator:        operation.Operator,
				OperationTime:   operation.OperationTime,
				OperationIp:     operation.OperatorIp,
				OperationRecord: operation.OperationRecord,
				SatelliteId:     operation.SatelliteId,
				SatelliteName:   operation.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:       operation.Id,
					LastTime: operation.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func TraceGetOperationList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")
		sortTypeStr := c.Query("sortType")
		searchInput := c.Query("searchConditions")

		err := isStringRequiredParamsEmpty(pageSizeStr, pageStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		sortType, ok := services.SortTypeValue[sortTypeStr]
		if !ok {
			sortType = services.SORTTYPE_TIME
		}

		params := &services.QueryLatestObjectsParams{
			ModelStruct: new(db.Operation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "satellite_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "operator")
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.OperationInfo, 0)

		for sqlRows.Next() {
			var operation db.Operation
			err := s.ScanRows(sqlRows, &operation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OperationInfo{
				Operator:        operation.Operator,
				OperationTime:   operation.OperationTime,
				OperationIp:     operation.OperatorIp,
				OperationRecord: operation.OperationRecord,
				SatelliteId:     operation.SatelliteId,
				SatelliteName:   operation.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:       operation.Id,
					LastTime: operation.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
