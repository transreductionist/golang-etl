package dbstocsvs

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/common/processors"
    "USERetl/util"
    "os"
)

func TransformToLoad(query util.Query) util.LoadCSV {

    buildenv.FilePaths()

    filePointerAccount, filePointerEmail, filePointerFlag, filePointers := buildFiles()

    loadAccounts := processors.NewCSVWriter(filePointerAccount)
    loadEmails := processors.NewCSVWriter(filePointerEmail)
    loadFlags := processors.NewCSVWriter(filePointerFlag)

    transformLoadStage := common.NewPipelineStage(
        common.Do(query.XformedAccts).Outputs(loadAccounts),
        common.Do(query.XformedEmails).Outputs(loadEmails),
        common.Do(query.XformedFlags).Outputs(loadFlags),
    )

    staging := util.LoadCSV{
        PipelineStage: transformLoadStage,
        Accounts: loadAccounts,
        Emails: loadEmails,
        Flags: loadFlags,
        FilePointers: filePointers,
    }

    return staging

}

func buildFiles() (*os.File, *os.File, *os.File, util.FilePointers) {
    filePointers := util.FilePointers{}
    filePointerAccount := util.OpenFile("data","account.csv")
    filePointers = append(filePointers, filePointerAccount)
    filePointerEmail := util.OpenFile("data","email.csv")
    filePointers = append(filePointers, filePointerEmail)
    filePointerFlag := util.OpenFile("data","flag.csv")
    filePointers = append(filePointers, filePointerFlag)
    return filePointerAccount, filePointerEmail, filePointerFlag, filePointers
}
