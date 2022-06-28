package v1

// import (
//   "fmt"
//   "math"
//   "net/http"
//   "strconv"

//   "github.com/360EntSecGroup-Skylar/excelize"
//   "github.com/gin-gonic/gin"

//   "gitlab.udevs.io/safia/safia_planner_service/api/models"
//   "gitlab.udevs.io/safia/safia_planner_service/config"
//   mainModels "gitlab.udevs.io/safia/safia_planner_service/models"
//   "gitlab.udevs.io/safia/safia_planner_service/pkg/logger"
//   "gitlab.udevs.io/safia/safia_planner_service/pkg/util"
//   "gitlab.udevs.io/safia/safia_planner_service/storage"
// )

// const (
//   BucketName        = "safia"
//   ReportsBucketName = "reports"
// )

// type handlerV1 struct {
//   log     logger.Logger
//   cfg     *config.Config
//   storage storage.StorageI
// }

// type HandlerV1Options struct {
//   Log     logger.Logger
//   Cfg     *config.Config
//   Storage storage.StorageI
// }

// func New(options *HandlerV1Options) *handlerV1 {
//   return &handlerV1{
//     log:     options.Log,
//     cfg:     options.Cfg,
//     storage: options.Storage,
//   }
// }

// func ParseQueryParam(c *gin.Context, key string, defaultValue string) (int32, error) {
//   valueStr := c.DefaultQuery(key, defaultValue)

//   value, err := strconv.Atoi(valueStr)

//   if err != nil {
//     c.JSON(http.StatusBadRequest, gin.H{
//       "error": err.Error(),
//     })
//     // c.Abort()
//     return 0, err
//   }

//   return int32(value), nil
// }

// func (h *handlerV1) HandleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
//   h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
//   c.JSON(code, models.ErrorModel{
//     Code:    code,
//     Message: message,
//     Error:   err,
//   })
// }
// func (h *handlerV1) HandleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
//   c.JSON(code, models.SuccessModel{
//     Code:    code,
//     Message: message,
//     Data:    data,
//   })
// }

// func ASCIToChar(number int32) string {
//   return string(number)
// }

// // "sort"
// func WriteToExcel(f *excelize.File, data []interface{}, row int) {
//   var (
//     cellNum = 65
//     cell    string
//   )
//   for i := 0; i < len(data); i++ {
//     if cellNum > 90 {
//       cell = ASCIToChar(int32(65)) + ASCIToChar(int32(cellNum-26)) + strconv.Itoa(row)
//     } else {
//       cell = ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//     }
//     f.SetCellValue("Sheet1", cell, data[i])
//     cellNum = cellNum + 1
//   }

// }

// func WriteComeNGoExcelHeader(f *excelize.File, data []interface{}, row int) {
//   var (
//     cellNum = 65
//     cell    string
//   )
//   style, _ := f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"}, 
//                     "font":{"bold":true,"italic":false}}`)

//   for i := 0; i < len(data); i++ {
//     if cellNum > 90 && cellNum <= 116 {
//       cell = ASCIToChar(int32(65)) + ASCIToChar(int32(cellNum-26)) + strconv.Itoa(row)
//       extended := ASCIToChar(int32(65)) + ASCIToChar(int32(cellNum-25)) + strconv.Itoa(row)
//       f.MergeCell("Sheet1", cell, extended)
//     } else if cellNum > 116 {
//       cell = ASCIToChar(int32(66)) + ASCIToChar(int32(cellNum-52)) + strconv.Itoa(row)
//       extended := ASCIToChar(int32(66)) + ASCIToChar(int32(cellNum-51)) + strconv.Itoa(row)
//       f.MergeCell("Sheet1", cell, extended)
//     } else {
//       cell = ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//       extended := ASCIToChar(int32(cellNum+1)) + strconv.Itoa(row)
//       f.MergeCell("Sheet1", cell, extended)
//     }

//     f.SetCellValue("Sheet1", cell, data[i])
//     f.SetColWidth("Sheet1", ASCIToChar(int32(cellNum)), ASCIToChar(int32(cellNum)), 30)
//     f.SetCellStyle("Sheet1", cell, cell, style)
//     cellNum = cellNum + 2
//   }
// }

// func WriteTimeTableByWeekHeader(f *excelize.File, data []interface{}, row int) {
//   cellNum := 65
//   columnWidth := 0
//   for i := 0; i < len(data); i++ {
//     if i == 0 {
//       columnWidth = 30
//     } else if i == 1 {
//       columnWidth = 40
//     } else {
//       columnWidth = 60
//     }
//     cell := ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//     f.SetCellValue("Sheet1", cell, data[i])
//     f.SetColWidth("Sheet1", ASCIToChar(int32(cellNum)), ASCIToChar(int32(cellNum)), float64(columnWidth))
//     style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}, 
//         "font":{"bold":true,"italic":false}}`)
//     f.SetCellStyle("Sheet1", cell, cell, style)
//     cellNum = cellNum + 1
//   }

// }

// func WriteHeaderFunc(f *excelize.File, data []DaysWithShifts) {
//   cellNum := 65
//   subCellNum := 66
//   for i := 0; i < len(data); i++ {
//     if i == 0 {
//       cell := ASCIToChar(int32(cellNum)) + "1"
//       f.SetCellValue("Sheet1", cell, data[i].Day)
//       f.MergeCell("Sheet1", cell, "A2")
//       cellNum += 1
//     } else {
//       if len(data[i].Shifts) > 0 {
//         currentCell := ASCIToChar(int32(cellNum)) + "1"
//         extendedCell := ASCIToChar(int32(cellNum)+int32(len(data[i].Shifts)-1)) + "1"
//         f.SetCellValue("Sheet1", currentCell, data[i].Day)
//         f.MergeCell("Sheet1", currentCell, extendedCell)
//         for k := 0; k < len(data[i].Shifts); k++ {
//           currentCell := ASCIToChar(int32(subCellNum)) + "2"
//           f.SetCellValue("Sheet1", currentCell, data[i].Shifts[k])
//           subCellNum += 1
//         }
//         cellNum = cellNum + len(data[i].Shifts)
//       } else {
//         currentCell := ASCIToChar(int32(cellNum)) + "1"
//         extendedCell := ASCIToChar(int32(cellNum)) + "2"
//         f.SetCellValue("Sheet1", currentCell, data[i].Day)
//         f.MergeCell("Sheet1", currentCell, extendedCell)
//         cellNum = cellNum + 1
//         subCellNum += 1
//       }

//     }

//   }
// }

// func WriteEmployeeHeaderFunc(f *excelize.File, data []DaysWithShifts) {
//   cellNum := 65
//   subCellNum := 67
//   for i := 0; i < len(data); i++ {
//     if i == 0 {
//       cell := ASCIToChar(int32(cellNum)) + "1"
//       f.SetCellValue("Sheet1", cell, data[i].Day)
//       f.MergeCell("Sheet1", cell, "A2")
//       cellNum = cellNum + 1
//     } else if i == 1 {
//       cell := ASCIToChar(int32(cellNum)) + "1"
//       f.SetCellValue("Sheet1", cell, data[i].Day)
//       f.MergeCell("Sheet1", cell, "B2")
//       cellNum = cellNum + 1
//     } else {

//       if len(data[i].Shifts) > 0 {
//         currentCell := ASCIToChar(int32(cellNum)) + "1"
//         extendedCell := ASCIToChar(int32(cellNum)+int32(len(data[i].Shifts)-1)) + "1"
//         f.SetCellValue("Sheet1", currentCell, data[i].Day)
//         f.MergeCell("Sheet1", currentCell, extendedCell)
//         for k := 0; k < len(data[i].Shifts); k++ {
//           currentCell := ASCIToChar(int32(subCellNum)) + "2"
//           f.SetCellValue("Sheet1", currentCell, data[i].Shifts[k])
//           subCellNum += 1
//         }
//         cellNum = cellNum + len(data[i].Shifts)
//       } else {
//         currentCell := ASCIToChar(int32(cellNum)) + "1"
//         extendedCell := ASCIToChar(int32(cellNum)) + "2"
//         f.SetCellValue("Sheet1", currentCell, data[i].Day)
//         f.MergeCell("Sheet1", currentCell, extendedCell)
//         cellNum = cellNum + 1
//         subCellNum = subCellNum + 1
//       }

//     }

//   }
// }

// func WriteBody(f *excelize.File, data []mainModels.TimetableBody) {
//   cellNum := 65
//   row := 3
//   for i := 0; i < len(data); i++ {
//     cellNum = 65
//     cell := ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//     f.SetCellValue("Sheet1", cell, data[i].PositionName)
//     for _, day := range data[i].Days {
//       for _, shift := range day.Shifts {
//         cellNum += 1
//         cell = ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//         f.SetCellValue("Sheet1", cell, shift.Count)
//       }
//       if len(day.Shifts) == 0 {
//         cellNum += 1
//       }
//     }
//     row += 1
//   }
// }

// func WiriteEmployeeBody(f *excelize.File, data []mainModels.TimetableEmployeesBody) {
//   positionCellNum := 65
//   emplCellNum := 66
//   cellNum := 67
//   rowPosiotion := 3
//   row := 3
//   emplRow := 3
//   for i := 0; i < len(data); i++ {
//     cellPosition := ASCIToChar(int32(positionCellNum)) + strconv.Itoa(rowPosiotion)
//     f.SetCellValue("Sheet1", cellPosition, data[i].PositionName)
//     rowPosiotion = rowPosiotion + len(data[i].Employees) - 1
//     f.MergeCell("Sheet1", cellPosition, "A"+strconv.Itoa(rowPosiotion))
//     rowPosiotion = rowPosiotion + 1
//     if len(data[i].Employees) == 0 {
//       rowPosiotion = rowPosiotion + 1
//     }

//     for _, employee := range data[i].Employees {
//       cellNum = 67
//       cellEmpl := ASCIToChar(int32(emplCellNum)) + strconv.Itoa(emplRow)
//       f.SetCellValue("Sheet1", cellEmpl, employee.FullName)
//       emplRow = emplRow + 1

//       for _, day := range employee.Shifts {
//         if len(day.Shifts) == 0 {
//           cellNum = cellNum + 1
//         }
//         for _, shift := range day.Shifts {
//           shiftCell := ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
//           if shift.Assigned {
//             f.SetCellValue("Sheet1", shiftCell, "+")
//             cellNum = cellNum + 1
//           } else {
//             f.SetCellValue("Sheet1", shiftCell, "-")
//             cellNum = cellNum + 1
//           }

//         }
//       }
//       row += 1
//     }
//     if len(data[i].Employees) == 0 {
//       emplRow = emplRow + 1
//     }

//   }
// }

// func containsShifts(s []mainModels.ExcelShift, searchterm string) bool {

//   for _, shift := range s {
//     if shift.ShiftName == searchterm {
//       return true
//     }
//   }

//   return false
// }

// func FillDays(input []mainModels.ExcelEmployee, date string) (resp []mainModels.ExcelEmployee) {
//   weekdays := util.CurrentWeekDates(date)
//   for _, day := range weekdays {
//     var exist bool
//     for _, existDate := range input {
//       if day == existDate.Day {
//         resp = append(resp, existDate)
//         exist = true
//       }
//     }
//     if !exist {
//       var temp mainModels.ExcelEmployee
//       temp.Day = day
//       resp = append(resp, temp)
//     }

//   }
//   return
// }

// func containsPostions(positons []mainModels.TimetableByWeekExcelBody, positionName string) bool {

//   for _, positon := range positons {
//     if positon.PositionName == positionName {
//       return true
//     }
//   }

//   return false
// }

// func ConvertDataToExcelBody(data []mainModels.GetAllByWeek, date string) (resp []mainModels.TimetableByWeekExcelBody) {
//   for i := 0; i < len(data); i++ {
//     for _, position := range data[i].DayResult {
//       var positionTemp mainModels.TimetableByWeekExcelBody
//       positionTemp.PositionName = position.PositionName
//       if !containsPostions(resp, position.PositionName) {
//         resp = append(resp, positionTemp)
//       }

//     }
//   }

//   for _, day := range data {
//     for _, position := range day.DayResult {
//       for _, shift := range position.PositionShifts {
//         for i, existingPosition := range resp {
//           if position.PositionName == existingPosition.PositionName {
//             if !containsShifts(existingPosition.Shifts, shift.ShiftName) {
//               var shiftExcel mainModels.ExcelShift
//               shiftExcel.ShiftName = shift.ShiftName
//               resp[i].Shifts = append(resp[i].Shifts, shiftExcel)
//             }
//           }
//         }
//       }
//     }
//   }
//   for n, excelPosition := range resp {
//     for _, day := range data {
//       for _, dbPosition := range day.DayResult {
//         if excelPosition.PositionName == dbPosition.PositionName {
//           for i, excelShift := range excelPosition.Shifts {
//             for _, dbShift := range dbPosition.PositionShifts {
//               var tempEmployees mainModels.ExcelEmployee
//               if dbShift.ShiftName == excelShift.ShiftName {
//                 tempEmployees.Day = day.Date
//                 tempEmployees.Employees = dbShift.ShiftEmployees
//                 resp[n].Shifts[i].EmployeesByDay = append(resp[n].Shifts[i].EmployeesByDay, tempEmployees)
//               }
//             }
//           }
//         }
//       }
//     }
//   }

//   for i, position := range resp {
//     for n, shift := range position.Shifts {
//       filledDays := FillDays(shift.EmployeesByDay, date)
//       resp[i].Shifts[n].EmployeesByDay = filledDays
//     }
//   }
//   return
// }

// func maxEmployeeColumn(daily []int32) int32 {
//   var maxNum int32 = 0
//   for _, maxSize := range daily {

//     if maxSize > maxNum {
//       maxNum = maxSize
//     }
//   }

//   return maxNum
// }

// func WriteTimeTableByWeekBody(f *excelize.File, data []mainModels.GetAllByWeek, date string) {
//   body := ConvertDataToExcelBody(data, date)

//   var (
//     positionCellNum = 65
//     rowPosition     = 2
//     shiftCellNum    = 66
//     rowShift        = 2
//   )

//   for _, position := range body {
//     var (
//       maxEmployees    = 0
//       employeeCellNum = 67
//       rowEmployee     = 2
//     )
//     cellPosition := ASCIToChar(int32(positionCellNum)) + strconv.Itoa(rowPosition)
//     style, _ := f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"}, 
//                     "font":{"bold":true,"italic":false}}`)
//     f.SetCellValue("Sheet1", cellPosition, position.PositionName)
//     f.SetColWidth("Sheet1", ASCIToChar(int32(positionCellNum)), ASCIToChar(int32(positionCellNum)), 30)
//     f.SetCellStyle("Sheet1", cellPosition, cellPosition, style)
//     for _, shift := range position.Shifts {
//       cellShift := ASCIToChar(int32(shiftCellNum)) + strconv.Itoa(rowShift)
//       style, _ := f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"}, 
//                     "font":{"bold":true,"italic":false}}`)
//       f.SetCellValue("Sheet1", cellShift, shift.ShiftName)
//       f.SetColWidth("Sheet1", ASCIToChar(int32(shiftCellNum)), ASCIToChar(int32(shiftCellNum)), 40)
//       f.SetCellStyle("Sheet1", cellShift, cellShift, style)
//       var maxNum []int32 = []int32{}
//       employeeCellNum = 67

//       if shift.EmployeesByDay == nil {
//         employeeCellNum += 1
//       }
//       for _, employee := range shift.EmployeesByDay {
//         var (
//           count = 0
//         )
//         rowDailyEmployee := 0
//         rowEmployee = rowShift
//         for index, employeeInfo := range employee.Employees {

//           if index == 0 {
//             rowDailyEmployee = 0
//           } else {
//             rowDailyEmployee = 1
//           }
//           rowEmployee += rowDailyEmployee
//           cellEmployee := ASCIToChar(int32(employeeCellNum)) + strconv.Itoa(rowEmployee)
//           count++
//           var name string
//           style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}, 
//                     "font":{"bold":false,"italic":false}}`)
//           if employeeInfo.EmployeeName == nil {
//             name = "Не Прикреплен"
//             f.SetCellValue("Sheet1", cellEmployee, name)
//             f.SetColWidth("Sheet1", ASCIToChar(int32(employeeCellNum)), ASCIToChar(int32(employeeCellNum)), 60)
//             f.SetCellStyle("Sheet1", cellEmployee, cellEmployee, style)
//           } else {
//             if *employeeInfo.AttendanceStatus == "absent" {
//               name = *employeeInfo.EmployeeName + " (Не пришел)"
//               f.SetCellValue("Sheet1", cellEmployee, name)
//               f.SetColWidth("Sheet1", ASCIToChar(int32(employeeCellNum)), ASCIToChar(int32(employeeCellNum)), 60)
//               f.SetCellStyle("Sheet1", cellEmployee, cellEmployee, style)
//             } else {
//               name = *employeeInfo.EmployeeName + " (Пришел)"
//               f.SetCellValue("Sheet1", cellEmployee, name)
//               f.SetColWidth("Sheet1", ASCIToChar(int32(employeeCellNum)), ASCIToChar(int32(employeeCellNum)), 60)
//               f.SetCellStyle("Sheet1", cellEmployee, cellEmployee, style)
//             }
//           }
//         }

//         if count > maxEmployees {
//           maxEmployees = count
//         }
//         maxNum = append(maxNum, int32(maxEmployees))
//         employeeCellNum += 1
//       }
//       maxShift := maxEmployeeColumn(maxNum)
//       rowShift += int(maxShift)
//       rowEmployee = rowShift
//       f.MergeCell("Sheet1", cellShift, "B"+strconv.Itoa(rowEmployee-1))
//     }
//     rowPosition = rowShift
//     f.MergeCell("Sheet1", cellPosition, "A"+strconv.Itoa(rowPosition-1))
//   }

// }

// func removerDuplicateName(data []mainModels.ComeNgoReportResponse) []string {
//   check := make(map[string]int)
//   res := make([]string, 0)

//   for _, val := range data {
//     check[val.FullName] = 1
//   }

//   for name, _ := range check {
//     res = append(res, name)
//   }

//   return res
// }

// func contains(name, date string, data []mainModels.ComeNgoReportResponse) bool {

//   for _, val := range data {
//     if val.FullName == name && val.Day == date {
//       return true
//     }
//   }

//   return false
// }

// func containsFullName(name string, data []mainModels.ComeNGoExcelBody) bool {

//   for _, val := range data {

//     if val.FullName == name {
//       return true
//     }
//   }

//   return false
// }

// func ConvertDataToAttendanceTemplate(data []mainModels.ComeNgoReportResponse) []mainModels.ComeNGoExcelBody {
//   var (
//     comeNGoExcelBodies []mainModels.ComeNGoExcelBody
//   )

//   for _, value := range data {

//     if !containsFullName(value.FullName, comeNGoExcelBodies) {
//       var (
//         comeNGoExcelBody mainModels.ComeNGoExcelBody
//         date             mainModels.DateRange
//       )
//       comeNGoExcelBody.FullName = value.FullName
//       comeNGoExcelBody.Branch = value.CompanyBranchName
//       comeNGoExcelBody.Position = value.PositionName
//       date.Date = value.Day
//       date.FactInTime = value.FactInTime
//       date.FactOutTime = value.FactOutTime
//       fact, _ := strconv.ParseFloat(value.FactHours, 8)
//       date.FactHours = float32(fact)
//       plan, _ := strconv.ParseFloat(value.GraphHours, 8)
//       date.GraphicHours = float32(plan)
//       overallHours, _ := strconv.ParseFloat(value.ResultingHours, 8)
//       date.OverallHours = float32(overallHours)
//       comeNGoExcelBody.DataRanges = append(comeNGoExcelBody.DataRanges, date)
//       comeNGoExcelBodies = append(comeNGoExcelBodies, comeNGoExcelBody)
//       elaborationTimeInFloat32, _ := strconv.ParseFloat(value.RevisionHours, 8)
//       comeNGoExcelBody.ElaborationTime = float32(elaborationTimeInFloat32)
//     } else {
//       for ind, val := range comeNGoExcelBodies {
//         var (
//           date mainModels.DateRange
//         )

//         date.Date = value.Day
//         date.FactInTime = value.FactInTime
//         date.FactOutTime = value.FactOutTime
//         fact, _ := strconv.ParseFloat(value.FactHours, 8)
//         date.FactHours = float32(fact)
//         plan, _ := strconv.ParseFloat(value.GraphHours, 8)
//         date.GraphicHours = float32(plan)
//         overallHours, _ := strconv.ParseFloat(value.ResultingHours, 8)
//         date.OverallHours = float32(overallHours)
//         elaborationTime, _ := strconv.ParseFloat(value.RevisionHours, 8)
//         date.ElaborationTime = float32(elaborationTime)

//         if value.FullName == val.FullName {
//           comeNGoExcelBodies[ind].DataRanges = append(comeNGoExcelBodies[ind].DataRanges, date)
//         }
//       }
//     }

//   }

//   return comeNGoExcelBodies

// }

// func convertTime(data string) string {

//   result := data[6:10] + "-" + data[3:5] + "-" + data[0:2]
//   return result
// }

// func ConvertToTime(text float64) (resp string) {
//   var hour, minute, second float64
//   var hourString, minuteString, secondString, hourRes, minuteRes, secondRes string

//   hour = math.Floor(text)
//   minute = math.Floor((text - float64(hour)) * 60)
//   second = math.Floor(((text-float64(hour))*60 - minute) * 60)
//   h := int(hour)
//   m := int(minute)
//   s := int(second)
//   hourString = strconv.Itoa(h)
//   minuteString = strconv.Itoa(m)
//   secondString = strconv.Itoa(s)

//   if len(hourString) == 1 {
//     hourRes = "0" + strconv.Itoa(int(hour))
//   } else {
//     hourRes = strconv.Itoa(int(hour))
//   }
//   if len(minuteString) == 1 {
//     minuteRes = "0" + strconv.Itoa(int(minute))
//   } else {
//     minuteRes = strconv.Itoa(int(minute))
//   }
//   if len(secondString) == 1 {
//     secondRes = "0" + strconv.Itoa(int(second))
//   } else {
//     secondRes = strconv.Itoa(int(second))
//   }

//   resp = hourRes + ":" + minuteRes + ":" + secondRes
//   return resp
// }
// func WriteAttendaceComeNGoBody(f *excelize.File, data []mainModels.ComeNgoReportResponse, dates []string) {

//   responses := ConvertDataToAttendanceTemplate(data)

//   var (
//     fullNameCellNum        = 65
//     branchCellNum          = 67
//     positionCellNum        = 69
//     elaborationTimeCellNum = 71
//     onScheduleCellNum      = 73
//     overAllDaysCellNum     = 75
//     overAllHoursCellNum    = 77

//     rows = 2
//   )

//   for _, response := range responses {
//     var (
//       overAllDays       int32
//       resultDateCellNum = 79
//       cellResultDate    string
//     )
//     style, _ := f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"}, 
//     "font":{"bold":false,"italic":false}}`)

//     cellFullName := ASCIToChar(int32(fullNameCellNum)) + strconv.Itoa(rows)
//     extendedFullName := ASCIToChar(int32(fullNameCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellFullName, extendedFullName)
//     cellBranch := ASCIToChar(int32(branchCellNum)) + strconv.Itoa(rows)
//     extendedBranch := ASCIToChar(int32(branchCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellBranch, extendedBranch)
//     cellPosition := ASCIToChar(int32(positionCellNum)) + strconv.Itoa(rows)
//     extendedPosition := ASCIToChar(int32(positionCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellPosition, extendedPosition)
//     cellElaboration := ASCIToChar(int32(elaborationTimeCellNum)) + strconv.Itoa(rows)
//     extendedElaboration := ASCIToChar(int32(elaborationTimeCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellElaboration, extendedElaboration)
//     cellSchedule := ASCIToChar(int32(onScheduleCellNum)) + strconv.Itoa(rows)
//     extendedSchedule := ASCIToChar(int32(onScheduleCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellSchedule, extendedSchedule)
//     cellOverAllDays := ASCIToChar(int32(overAllDaysCellNum)) + strconv.Itoa(rows)
//     extendedOverAllDays := ASCIToChar(int32(overAllDaysCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellOverAllDays, extendedOverAllDays)
//     cellOverAllHours := ASCIToChar(int32(overAllHoursCellNum)) + strconv.Itoa(rows)
//     extendedOverAllHours := ASCIToChar(int32(overAllHoursCellNum+1)) + strconv.Itoa(rows)
//     f.MergeCell("Sheet1", cellOverAllHours, extendedOverAllHours)

//     f.SetCellValue("Sheet1", cellFullName, response.FullName)
//     f.SetColWidth("Sheet1", ASCIToChar(int32(fullNameCellNum)), ASCIToChar(int32(fullNameCellNum)), 30)
//     f.SetCellStyle("Sheet1", cellFullName, cellFullName, style)

//     f.SetCellValue("Sheet1", cellBranch, response.Branch)
//     f.SetColWidth("Sheet1", ASCIToChar(int32(branchCellNum)), ASCIToChar(int32(branchCellNum)), 30)
//     f.SetCellStyle("Sheet1", cellBranch, cellBranch, style)

//     f.SetCellValue("Sheet1", cellPosition, response.Position)
//     f.SetColWidth("Sheet1", ASCIToChar(int32(positionCellNum)), ASCIToChar(int32(positionCellNum)), 30)
//     f.SetCellStyle("Sheet1", cellPosition, cellPosition, style)

//     for _, day := range dates {
//       for _, date := range response.DataRanges {
//         if resultDateCellNum > 116 && resultDateCellNum <= 141 {
//           cellResultDate = ASCIToChar(int32(66)) + ASCIToChar(int32(resultDateCellNum-52)) + strconv.Itoa(rows)
//           extendedDate := ASCIToChar(int32(66)) + ASCIToChar(int32(resultDateCellNum-51)) + strconv.Itoa(rows)
//           if date.Date == day[:10] {
//             for k := 0; k < 2; k++ {
//               if k == 0 {
//                 f.SetCellValue("Sheet1", cellResultDate, date.FactInTime)
//               }
//               if k == 1 {
//                 f.SetCellValue("Sheet1", extendedDate, date.FactOutTime)
//               }
//             }
//             response.ElaborationTime += date.ElaborationTime
//             response.GraphicTime += date.GraphicHours
//             response.OverAll += date.OverallHours
//             overAllDays += 1
//           } //else {
//           //   f.MergeCell("Sheet1", cellResultDate, extendedDate)
//           //   f.SetCellValue("Sheet1", cellResultDate, "выходной")
//           // }
//         } else if resultDateCellNum > 90 && resultDateCellNum <= 116 {
//           cellResultDate = ASCIToChar(int32(65)) + ASCIToChar(int32(resultDateCellNum-26)) + strconv.Itoa(rows)
//           extendedDate := ASCIToChar(int32(65)) + ASCIToChar(int32(resultDateCellNum-25)) + strconv.Itoa(rows)
//           if date.Date == day[0:10] {
//             for k := 0; k < 2; k++ {
//               if k == 0 {
//                 f.SetCellValue("Sheet1", cellResultDate, date.FactInTime)
//               }
//               if k == 1 {
//                 f.SetCellValue("Sheet1", extendedDate, date.FactOutTime)
//               }
//             }
//             response.ElaborationTime += date.ElaborationTime
//             response.GraphicTime += date.GraphicHours
//             response.OverAll += date.OverallHours
//             overAllDays += 1
//           } //else {
//           // f.MergeCell("Sheet1", cellResultDate, extendedDate)
//           // f.SetCellValue("Sheet1", cellResultDate, "выходной")
//           // }
//         } else if resultDateCellNum > 66 && resultDateCellNum <= 90 {
//           cellResultDate = ASCIToChar(int32(resultDateCellNum)) + strconv.Itoa(rows)
//           extendedDate := ASCIToChar(int32(resultDateCellNum+1)) + strconv.Itoa(rows)
//           fmt.Println("data.Data : ", date.Date)
//           if date.Date == day[0:10] {
//             fmt.Println("Cell : ", resultDateCellNum)
//             for k := 0; k < 2; k++ {
//               if k == 0 {
//                 f.SetCellValue("Sheet1", cellResultDate, date.FactInTime)
//               }
//               if k == 1 {
//                 f.SetCellValue("Sheet1", extendedDate, date.FactOutTime)
//               }
//             }
//             response.ElaborationTime += date.ElaborationTime
//             response.GraphicTime += date.GraphicHours
//             response.OverAll += date.OverallHours
//             overAllDays += 1
//           } //else {
//           // f.MergeCell("Sheet1", cellResultDate, extendedDate)
//           // f.SetCellValue("Sheet1", cellResultDate, "выходной")
//           // }

//         }
//         resultDateCellNum += 2
//       }

//       f.SetCellValue("Sheet1", cellOverAllHours, ConvertToTime(float64(response.OverAll)))
//       f.SetColWidth("Sheet1", ASCIToChar(int32(overAllHoursCellNum)), ASCIToChar(int32(overAllHoursCellNum)), 30)
//       f.SetCellStyle("Sheet1", cellOverAllHours, cellOverAllHours, style)

//       f.SetCellValue("Sheet1", cellElaboration, ConvertToTime(float64(response.ElaborationTime)))
//       f.SetColWidth("Sheet1", ASCIToChar(int32(elaborationTimeCellNum)), ASCIToChar(int32(elaborationTimeCellNum)), 30)
//       f.SetCellStyle("Sheet1", cellElaboration, cellElaboration, style)

//       f.SetCellValue("Sheet1", cellSchedule, ConvertToTime(float64(response.GraphicTime)))
//       f.SetColWidth("Sheet1", ASCIToChar(int32(onScheduleCellNum)), ASCIToChar(int32(onScheduleCellNum)), 30)
//       f.SetCellStyle("Sheet1", cellSchedule, cellSchedule, style)
//       f.SetCellValue("Sheet1", cellOverAllDays, int(overAllDays))
//       f.SetColWidth("Sheet1", ASCIToChar(int32(overAllDaysCellNum)), ASCIToChar(int32(overAllDaysCellNum)), 30)
//       f.SetCellStyle("Sheet1", cellOverAllDays, cellOverAllDays, style)

//       //  f.SetCellStyle("Sheet1", cell, cell, style)

//     }
//     rows += 1
//   }
// }