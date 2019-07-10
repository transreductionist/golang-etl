package dbtocsvs

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/util"
    "USERetl/common/processors"
    "USERetl/etl/queries"
    "USERetl/etl/transformers"
)

func QueryToTransform() util.Query {
    fromDBEnvPrefix := buildenv.Forwarding()
    fromDatabase := util.BuildDatabase(fromDBEnvPrefix)

    selectUserData := queries.SelectFromTableByIDRange("user", 2000000, 2000002)
    userQuery := processors.NewSQLReader(fromDatabase, selectUserData)

    xformAccts := transformers.NewXformAccts()
    xformEmails := transformers.NewXformEmails()

    queryStage := common.NewPipelineStage(
        common.Do(userQuery).Outputs(xformAccts),
        common.Do(userQuery).Outputs(xformEmails),
    )

    staging := util.Query{
        PipelineStage: queryStage,
        XformedAccts: xformAccts,
        XformedEmails: xformEmails,
    }

    return staging
}
