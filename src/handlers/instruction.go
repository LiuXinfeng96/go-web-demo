package handlers

import (
	"go-web-demo/src/db"
	"go-web-demo/src/models"
	"go-web-demo/src/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ControlAddInstruction(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddInstructionReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.InstructionId,
			req.InstructionContent, req.DebrisId, req.DebrisName,
			req.SatelliteId, req.SatelliteName)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.InstructionId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get token failed", c)
			return
		}

		genInstructionTime := time.Now().Unix()
		noExecI := &db.Instruction{
			InstructionId:      req.InstructionId,
			InstructionSource:  claims.Name,
			Type:               db.OPERATION,
			ExecState:          db.NOTEXEC,
			InstructionContent: req.InstructionContent,
			DebrisId:           req.DebrisId,
			DebrisName:         req.DebrisName,
			SatelliteId:        req.SatelliteId,
			SatelliteName:      req.SatelliteName,
			GenInstructionTime: genInstructionTime,
		}
		// 未执行指令信息入库
		err = s.InsertOneObjertToDB(noExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		operation := &db.Operation{
			Operator:        claims.Name,
			OperatorIp:      c.ClientIP(),
			OperationTime:   genInstructionTime,
			SatelliteId:     req.SatelliteId,
			SatelliteName:   req.SatelliteName,
			OperationRecord: "编辑指令：" + req.InstructionId,
		}

		// 编辑操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		execInstructionTime := time.Now().Unix()
		inExecI := &db.Instruction{
			InstructionId:       req.InstructionId,
			InstructionSource:   claims.Name,
			Type:                db.OPERATION,
			ExecState:           db.INEXEC,
			InstructionContent:  req.InstructionContent,
			DebrisId:            req.DebrisId,
			DebrisName:          req.DebrisName,
			SatelliteId:         req.SatelliteId,
			SatelliteName:       req.SatelliteName,
			GenInstructionTime:  genInstructionTime,
			ExecInstructionTime: execInstructionTime,
		}
		// 执行指令入库
		err = s.InsertOneObjertToDB(inExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		//假设执行中-------------------------------
		time.Sleep(time.Millisecond * 500)
		//-----------------------------------------

		endExecI := &db.Instruction{
			InstructionId:       req.InstructionId,
			InstructionSource:   claims.Name,
			Type:                db.OPERATION,
			ExecState:           db.EXECSUCCESS,
			InstructionContent:  req.InstructionContent,
			DebrisId:            req.DebrisId,
			DebrisName:          req.DebrisName,
			SatelliteId:         req.SatelliteId,
			SatelliteName:       req.SatelliteName,
			GenInstructionTime:  genInstructionTime,
			ExecInstructionTime: execInstructionTime,
		}

		// 执行结果入库
		err = s.InsertOneObjertToDB(endExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		operation = &db.Operation{
			Operator:        claims.Name,
			OperationTime:   execInstructionTime,
			OperatorIp:      c.ClientIP(),
			SatelliteId:     req.SatelliteId,
			SatelliteName:   req.SatelliteName,
			OperationRecord: "执行指令：" + req.InstructionId,
		}

		// 执行操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetInstructionList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			QueryMap:    make(map[string]string),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		params.QueryMap["exec_state"] = strconv.Itoa(int(db.NOTEXEC))

		sqlRows, total, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:       instruction.Id,
					LastTime: instruction.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func ExecGetExecResultList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		sqlRows, total, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionDetails, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionDetails{
				InstructionInfo: models.InstructionInfo{
					InstructionId:       instruction.InstructionId,
					InstructionSource:   instruction.InstructionSource,
					InstructionContent:  instruction.InstructionContent,
					InstructionType:     db.InstructionTypeName[instruction.Type],
					ExecInstructionTime: instruction.ExecInstructionTime,
					GenInstructionTime:  instruction.GenInstructionTime,
					DebrisId:            instruction.DebrisId,
					DebrisName:          instruction.DebrisName,
					SatelliteId:         instruction.SatelliteId,
					SatelliteName:       instruction.SatelliteName,
					BaseRespInfo: models.BaseRespInfo{
						Id:       instruction.Id,
						LastTime: instruction.LastTime,
					},
				},
				ExecState: db.ExecStateName[instruction.ExecState],
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func TraceGetInstructionList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "instruction_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:       instruction.Id,
					LastTime: instruction.LastTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
