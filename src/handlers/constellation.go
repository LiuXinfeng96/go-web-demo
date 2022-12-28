package handlers

import (
	"go-web-demo/src/db"
	"go-web-demo/src/models"
	"go-web-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlAddConstellation(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddConstellationReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.ConstellationId, req.ConstellationName,
			req.SatelliteLinkState)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.ConstellationId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		satelliteLinkState, ok := db.StateValue[req.SatelliteLinkState]
		if !ok {
			ParamsValueJSONResp("satellite link state type not as expected", c)
			return
		}

		constellation := &db.Constellation{
			ConstellationId:    req.ConstellationId,
			ConstellationName:  req.ConstellationName,
			SatelliteLinkState: satelliteLinkState,
			SatelliteTotalNum:  req.SatelliteTotalNum,
			SatelliteUpNum:     req.SatelliteUpNum,
			SatelliteDownNum:   req.SatelliteDownNum,
		}

		err = s.InsertOneObjertToDB(constellation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetConstellationList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Constellation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "constellation_id")
			params.SearchIndex = append(params.SearchIndex, "constellation_name")
		}

		sqlRows, total, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.ConstellationInfo, 0)

		for sqlRows.Next() {
			var constellation db.Constellation
			err := s.ScanRows(sqlRows, &constellation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ConstellationInfo{
				ConstellationId:    constellation.ConstellationId,
				ConstellationName:  constellation.ConstellationName,
				SatelliteLinkState: db.StateName[constellation.SatelliteLinkState],
				SatelliteTotalNum:  constellation.SatelliteTotalNum,
				SatelliteUpNum:     constellation.SatelliteUpNum,
				SatelliteDownNum:   constellation.SatelliteDownNum,
				BaseRespInfo: models.BaseRespInfo{
					Id:       constellation.Id,
					LastTime: constellation.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func TraceGetConstellationList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Constellation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "constellation_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "constellation_id")
			params.SearchIndex = append(params.SearchIndex, "constellation_name")
		}

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.ConstellationInfo, 0)

		for sqlRows.Next() {
			var constellation db.Constellation
			err := s.ScanRows(sqlRows, &constellation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ConstellationInfo{
				ConstellationId:    constellation.ConstellationId,
				ConstellationName:  constellation.ConstellationName,
				SatelliteLinkState: db.StateName[constellation.SatelliteLinkState],
				SatelliteTotalNum:  constellation.SatelliteTotalNum,
				SatelliteUpNum:     constellation.SatelliteUpNum,
				SatelliteDownNum:   constellation.SatelliteDownNum,
				BaseRespInfo: models.BaseRespInfo{
					Id:       constellation.Id,
					LastTime: constellation.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
