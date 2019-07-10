package transformers

import (
    "USERetl/common/data"
    "USERetl/common/processors"
    "USERetl/common/util"
    "USERetl/etl/models"
    "strings"
    "time"
)

type MyXformFlags struct{}

func NewXformFlags() *MyXformFlags {
   return &MyXformFlags{}
}

func (t *MyXformFlags) ProcessData(d data.JSON,
   outputChan chan data.JSON,
   killChan chan error) {
   var ultsysUsers []models.RcvdUltsysData
   err := data.ParseJSON(d, &ultsysUsers)
   util.KillPipelineIfErr(err, killChan)

   xformFlags := getXformFlags(ultsysUsers)
   writeData := processors.SQLWriterData{"flag", xformFlags}

   dd, err := data.NewJSON(writeData)
   util.KillPipelineIfErr(err, killChan)
   outputChan <- dd
}

func (t *MyXformFlags) Finish(outputChan chan data.JSON,
   killChan chan error) {}

func getXformFlags(ultsysUsers []models.RcvdUltsysData) ([]models.Flag) {
   flag := models.Flag{}
   createdFlags := []models.Flag{}
   now := time.Now()
    for _, ultsysUser := range ultsysUsers {
        flag.ID = ultsysUser.ID
        flag.Name = "a-flag"
        flag.Type = strings.ToLower(models.FlagTypes.Name(1))
        flag.AgentID = 1
        flag.Date = now.Format("2006-01-02 15:04:05")
        createdFlags = append(createdFlags, flag)
   }
   return createdFlags
}
