package dbtotablesonebranch

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/common/processors"
    "USERetl/etl/queries"
    "USERetl/etl/transformers"
    "USERetl/util"
)

func QueryToTransform() util.Query {
    fromDBEnvPrefix := buildenv.Forwarding()
    fromDatabase := util.BuildDatabase(fromDBEnvPrefix)

    selectUserData := queries.SelectFromTableByIDRange("user", 2000000, 2000002)
    userQuery := processors.NewSQLReader(fromDatabase, selectUserData)

    xformTables := transformers.NewXformTables()

    queryStage := common.NewPipelineStage(
      common.Do(userQuery).Outputs(xformTables),
    )

    staging := util.Query{
      PipelineStage: queryStage,
      XformedTables: xformTables,
    }

    return staging
}
