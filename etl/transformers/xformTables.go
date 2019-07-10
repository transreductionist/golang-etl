package transformers

import (
    "USERetl/common/data"
    "USERetl/common/processors"
    "USERetl/common/util"
    "USERetl/etl/models"
    "strings"
    "time"
)

type MyXformTables struct{}

func NewXformTables() *MyXformTables {
    return &MyXformTables{}
}

func (t *MyXformTables) ProcessData(d data.JSON,
    outputChan chan data.JSON,
    killChan chan error) {
    var ultsysUsers []models.RcvdUltsysData
    err := data.ParseJSON(d, &ultsysUsers)
    util.KillPipelineIfErr(err, killChan)

    xformAccts, xformFlags := getXformTables(ultsysUsers)
    writeAccts := processors.SQLWriterData{TableName: "account", InsertData: xformAccts}
    writeFlags := processors.SQLWriterData{TableName: "flag", InsertData: xformFlags}
    writers := models.Transformed{SQLWriteAccts: writeAccts, SQLWriteFlags: writeFlags}

    dd, err := data.NewJSON(writers)
    util.KillPipelineIfErr(err, killChan)
    outputChan <- dd
}

func (t *MyXformTables) Finish(outputChan chan data.JSON,
    killChan chan error) {}

func getXformTables(ultsysUsers []models.RcvdUltsysData) ([]models.Account, []models.Flag) {
    account := models.Account{}
    createdAccounts := []models.Account{}
    flag := models.Flag{}
    createdFlags := []models.Flag{}
    now := time.Now()
    for _, ultsysUser := range ultsysUsers {
        id := ultsysUser.ID
        account.ID = id
        account.EmailID = 1
        account.UserName = ultsysUser.UserName
        account.Password = ultsysUser.Password
        account.Type = strings.ToLower(models.AccountTypes.Name(1))
        createdAccounts = append(createdAccounts, account)

        flag.ID = id
        flag.Name = "a-flag"
        flag.Type = strings.ToLower(models.FlagTypes.Name(1))
        flag.AgentID = 1
        flag.Date = now.Format("2006-01-02 15:04:05")
        createdFlags = append(createdFlags, flag)
    }

    return createdAccounts, createdFlags
}
