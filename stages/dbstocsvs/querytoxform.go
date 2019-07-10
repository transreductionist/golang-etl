package dbstocsvs

import (
    "USERetl/buildenv"
    "USERetl/common"
    "USERetl/util"
    "USERetl/common/processors"
    "USERetl/etl/queries"
    "USERetl/etl/transformers"
)

func QueryToTransform() util.Query {
    userFwdQuery := getQuery(buildenv.Forwarding(), queries.SelectFromTableByIDRange(
        "user", 2000000, 2000002))
    userLocQuery := getQuery(buildenv.Local(), queries.SelectFromTableAll("flag"))

    xformAccounts := transformers.NewXformAccts()
    xformEmails := transformers.NewXformEmails()
    xformFlags := transformers.NewXformFlags()

    queryStage := common.NewPipelineStage(
        common.Do(userFwdQuery).Outputs(xformAccounts),
        common.Do(userFwdQuery).Outputs(xformEmails),
        common.Do(userLocQuery).Outputs(xformFlags),
    )

    staging := util.Query{
        PipelineStage: queryStage,
        XformedAccts: xformAccounts,
        XformedEmails: xformEmails,
        XformedFlags: xformFlags,
    }

    return staging
}

func getQuery(fromDBEnvPrefix string, query string) *processors.SQLReader {
    fromDatabase := util.BuildDatabase(fromDBEnvPrefix)
    return processors.NewSQLReader(fromDatabase, query)
}
