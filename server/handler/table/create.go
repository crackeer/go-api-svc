package table

import (
	"errors"
	"io/ioutil"
	"mime/multipart"

	ginhelper "github.com/crackeer/gopkg/gin"
	"github.com/gin-gonic/gin"
)

// Create ...
//  @param ctx
func Create(ctx *gin.Context) {
	tableObj, err := getTableObject(ctx)
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	params := ginhelper.AllParams(ctx)

	createData := map[string]interface{}{}

	for key, value := range params {
		if realValue, ok := value.([]*multipart.FileHeader); ok {
			data, err := readMutipartFile(realValue)
			if err != nil {
				ginhelper.Failure(ctx, -1, err.Error())
				return
			}
			createData[key] = data
			continue
		}
		createData[key] = value
	}

	createData, err = tableObj.Create(createData)
	tableObj.GetPrimaryKey()
	if err != nil {
		ginhelper.Failure(ctx, -1, err.Error())
		return
	}
	ginhelper.Success(ctx, createData)
}

func readMutipartFile(data []*multipart.FileHeader) (string, error) {

	if len(data) < 1 {
		return "", errors.New("mutipart file nil")
	}
	header := data[0]
	tmpFile, err := header.Open()

	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(tmpFile)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
