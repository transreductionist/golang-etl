package transformers

import (
    "USERetl/common/data"
    "USERetl/common/processors"
    "USERetl/common/util"
    "USERetl/etl/models"
    "strings"
)

type MyXformAccts struct{}

func NewXformAccts() *MyXformAccts {
    return &MyXformAccts{}
}

func (t *MyXformAccts) ProcessData(d data.JSON,
    outputChan chan data.JSON,
    killChan chan error) {
    var ultsysUsers []models.RcvdUltsysData
    err := data.ParseJSON(d, &ultsysUsers)
    util.KillPipelineIfErr(err, killChan)

    xformAccts := getXformAccts(ultsysUsers)
    writeData := processors.SQLWriterData{"account", xformAccts}

    dd, err := data.NewJSON(writeData)
    util.KillPipelineIfErr(err, killChan)
    outputChan <- dd
}

func (t *MyXformAccts) Finish(outputChan chan data.JSON,
    killChan chan error) {}

func getXformAccts(ultsysUsers []models.RcvdUltsysData) ([]models.Account) {
    account := models.Account{}
    createdAccounts := []models.Account{}
    for _, ultsysUser := range ultsysUsers {
        account.ID = ultsysUser.ID
        account.EmailID = 1
        account.UserName = ultsysUser.UserName
        account.Password = ultsysUser.Password
        account.Type = strings.ToLower(models.AccountTypes.Name(1))
        createdAccounts = append(createdAccounts, account)
    }
    return createdAccounts
}
