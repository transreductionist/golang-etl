package util

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/common/logger"
    "USERetl/common/processors"
    "USERetl/common/util"
    "USERetl/etl/conndb"
    "USERetl/etl/transformers"
    "database/sql"
    "log"
    "os"
)

type Query struct {
    PipelineStage *common.PipelineStage
    XformedTables  *transformers.MyXformTables
    XformedAccts  *transformers.MyXformAccts
    XformedEmails  *transformers.MyXformEmails
    XformedFlags  *transformers.MyXformFlags
}

type LoadDB struct {
    PipelineStage *common.PipelineStage
    SQLWriteTables *processors.SQLWriter
    SQLWriteAccts *processors.SQLWriter
    SQLWriteEmails *processors.SQLWriter
    SQLWriteFlags *processors.SQLWriter
}

type Write struct {
    PipelineStage *common.PipelineStage
}

type FilePointers []*os.File

type LoadCSV struct {
    PipelineStage *common.PipelineStage
    Accounts *processors.CSVWriter
    Emails *processors.CSVWriter
    Flags *processors.CSVWriter
    FilePointers FilePointers
}

func RunPipeline( pipeline *common.Pipeline){
    errPipeline := <-pipeline.Run()
    if errPipeline != nil {
        logger.ErrorWithoutTrace(pipeline.Name, ":", errPipeline)
        logger.ErrorWithoutTrace(pipeline.Stats())
    } else {
        logger.Info(pipeline.Name, ": Completed successfully.")
    }
}

func BuildDatabase(dbEnvPrefix string) *sql.DB{
    connection := conndb.GetDatabase(dbEnvPrefix)
    dev0ConnectionDB := util.ConnectionStringDB(connection)
    return util.DBConnect(dev0ConnectionDB)
}

func OpenFile(category string, fileName string) (filePointer *os.File) {
    filePath := buildenv.GetFilePath(category,fileName)
    filePointer, errFileAccount := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
    if errFileAccount != nil {
        log.Fatal(errFileAccount)
    }
    return filePointer
}

func CloseFile(filePointers []*os.File){
    for _, filePointer := range filePointers {
        if errFileClose := filePointer.Close(); errFileClose != nil{
            log.Fatal(errFileClose)}
    }
}


