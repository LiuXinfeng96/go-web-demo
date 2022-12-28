package handlers

import (
	"go-web-demo/src/db"
	"go-web-demo/src/models"
	"go-web-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExecAddCommState(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddCommStateReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.SatelliteId, req.SatelliteName, req.LinkLoad,
			req.OrbitId, req.CommState, req.CommPort, req.CommDelay, req.CommBandwidth)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.SatelliteId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		state, ok := db.StateValue[req.CommState]
		if !ok {
			ParamsValueJSONResp("comm state type not as expected", c)
			return
		}

		commState := &db.CommState{
			SatelliteId:   req.SatelliteId,
			SatelliteName: req.SatelliteName,
			OrbitId:       req.OrbitId,
			CommState:     state,
			CommDelay:     req.CommDelay,
			CommPort:      req.CommPort,
			CommBandwidth: req.CommBandwidth,
			LinkLoad:      req.LinkLoad,
		}

		err = s.InsertOneObjertToDB(commState)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetCommStateList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
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
			ModelStruct: new(db.CommState),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, total, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.CommStateInfo, 0)

		for sqlRows.Next() {
			var commState db.CommState
			err := s.ScanRows(sqlRows, &commState)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.CommStateInfo{
				SatelliteId:   commState.SatelliteId,
				SatelliteName: commState.SatelliteName,
				OrbitId:       commState.OrbitId,
				CommState:     db.StateName[commState.CommState],
				CommPort:      commState.CommPort,
				CommDelay:     commState.CommDelay,
				CommBandwidth: commState.CommBandwidth,
				LinkLoad:      commState.LinkLoad,
				BaseRespInfo: models.BaseRespInfo{
					Id:       commState.Id,
					LastTime: commState.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func TraceGetCommStateList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.CommState),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "satellite_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.CommStateInfo, 0)

		for sqlRows.Next() {
			var commState db.CommState
			err := s.ScanRows(sqlRows, &commState)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.CommStateInfo{
				SatelliteId:   commState.SatelliteId,
				SatelliteName: commState.SatelliteName,
				OrbitId:       commState.OrbitId,
				CommState:     db.StateName[commState.CommState],
				CommPort:      commState.CommPort,
				CommDelay:     commState.CommDelay,
				CommBandwidth: commState.CommBandwidth,
				LinkLoad:      commState.LinkLoad,
				BaseRespInfo: models.BaseRespInfo{
					Id:       commState.Id,
					LastTime: commState.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
