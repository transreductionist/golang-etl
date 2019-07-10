package dbtocsvs

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/common/processors"
    "USERetl/util"
    "os"
)

func TransformToLoad(query util.Query) util.LoadCSV {

    buildenv.FilePaths()

    filePointerAccount, filePointerEmail, filePointers := buildFiles()

    loadAccounts := processors.NewCSVWriter(filePointerAccount)
    loadEmails := processors.NewCSVWriter(filePointerEmail)

    transformLoadStage := common.NewPipelineStage(
        common.Do(query.XformedAccts).Outputs(loadAccounts),
        common.Do(query.XformedEmails).Outputs(loadEmails),
    )

    staging := util.LoadCSV{
        PipelineStage: transformLoadStage,
        Accounts: loadAccounts,
        Emails: loadEmails,
        FilePointers: filePointers,
    }

    return staging

}

func buildFiles() (*os.File, *os.File, util.FilePointers) {
    filePointers := util.FilePointers{}
    filePointerAccount := util.OpenFile("data","account.csv")
    filePointers = append(filePointers, filePointerAccount)
    filePointerEmail := util.OpenFile("data","email.csv")
    filePointers = append(filePointers, filePointerEmail)
    return filePointerAccount, filePointerEmail, filePointers
}
