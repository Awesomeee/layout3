package layout

import (
	"github.com/gin-gonic/gin"
	
	//"errors"
	"fmt"
	"net/http"
	
	db "github.com/hpt/go-client/app/db/sqlc"
	"github.com/hpt/go-client/app/util"
	"log"
)

type Body struct {
	// json tag to de-serialize json body
	AdmID string `json:"admID"`
}

func checkDuplicateLayoutID(ctx *gin.Context) {

	/*ctx.JSON(200, gin.H{
		"message": "uploadProcess " + " successfully",
	})*/
	
	/*admID := ctx.Param("admID")
	fmt.Println("admID " + admID)
	if admID == "" {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(errors.New("admID param is required")))
		return
	}*/
	
	body:=Body{}
	// using BindJson method to serialize body with struct
	if err:=ctx.BindJSON(&body);err!=nil{
		log.Fatal(err)
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	}
	fmt.Println(body.AdmID)
	
	res, err := getLayoutsService(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	dbBpmAdmLayout := res
	var result string
	result = "" //1: trùng, 0: không trùng
	for i := 0; i < len(*dbBpmAdmLayout); i++ {
		fmt.Println("i " + fmt.Sprint(i) + "- dbBpmAdmLayoutIDElement " + fmt.Sprint((*dbBpmAdmLayout)[i].ID))
		if fmt.Sprint((*dbBpmAdmLayout)[i].ID) == body.AdmID {
			result = "1"
			break
		}
		if i == len(*dbBpmAdmLayout)-1 {
			result = "0"
		}
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Success", result))

}

func getLayoutsService(ctx *gin.Context) (*[]db.BpmAdmLayout, error) {
	getLayoutParam := db.GetLayoutsParams{Limit:1, Offset:0, Code:"11", IDAsc:true, IDDesc:false, CodeAsc:true, CodeDesc:false}
	layouts, err := db.StoreDB.GetLayouts(ctx,getLayoutParam)
	if err != nil {
		return nil, err
	}
	return &layouts, nil
}