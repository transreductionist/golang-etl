package transformers

import (
    "USERetl/common/data"
    "USERetl/common/util"
    "USERetl/etl/models"
)

type MyXformEmails struct{}

func NewXformEmails() *MyXformEmails {
    return &MyXformEmails{}
}

func (t *MyXformEmails) ProcessData(d data.JSON,
    outputChan chan data.JSON,
    killChan chan error) {

    var ultsysUsers []models.RcvdUltsysData
    var XformToEmails []models.Email

    err := data.ParseJSON(d, &ultsysUsers)
    util.KillPipelineIfErr(err, killChan)

    for _, ultsysUser := range ultsysUsers {
        XformToEmail := models.Email{}
        XformToEmail.ID = ultsysUser.ID
        XformToEmail.AccountID = ultsysUser.ID
        XformToEmail.Email = ultsysUser.UserName
        XformToEmail.LastOpenInUTC = ultsysUser.LastOpen.Format("2006-01-02 15:04:05")
        XformToEmail.LastClickInUTC = ultsysUser.LastClick.Format("2006-01-02 15:04:05")
        XformToEmail.LastBounceInUTC = ultsysUser.LastBounce.Format("2006-01-02 15:04:05")
        XformToEmail.TotalOpens = ultsysUser.OpenCount
        XformToEmail.TotalClicks = ultsysUser.ClickCount
        XformToEmail.TotalBounces = 0
        XformToEmail.CreatedInUTC = ultsysUser.DateJoined.Format("2006-01-02 15:04:05")
        XformToEmail.LastVerificationInUTC = ultsysUser.DateJoined.Format("2006-01-02 15:04:05")
        XformToEmails = append(XformToEmails, XformToEmail)
    }

    if len(XformToEmails) > 0 {
        dd, err := data.NewJSON(XformToEmails)
        util.KillPipelineIfErr(err, killChan)
        outputChan <- dd
    }
}

func (t *MyXformEmails) Finish(outputChan chan data.JSON,
    killChan chan error) {}
