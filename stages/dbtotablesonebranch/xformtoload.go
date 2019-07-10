package dbtotablesonebranch

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/common/processors"
    "USERetl/util"
)

func TransformToLoad(query util.Query) util.LoadDB {
    toDBEnvPrefix := buildenv.Local()
    toDatabase := util.BuildDatabase(toDBEnvPrefix)

    loadAccts := processors.NewSQLWriter(toDatabase, "account")
    loadFlags := processors.NewSQLWriter(toDatabase, "flag")

    transformLoadStage := common.NewPipelineStage(
        common.Do(query.XformedTables).Outputs(loadAccts, loadFlags),
    )

    staging := util.LoadDB{
        PipelineStage: transformLoadStage,
        SQLWriteAccts: loadAccts,
        SQLWriteFlags: loadFlags,
    }

    return staging
}
