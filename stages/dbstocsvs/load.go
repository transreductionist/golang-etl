package dbstocsvs

import (
    "USERetl/common"
    "USERetl/util"
)

func Load(load util.LoadCSV) util.Write {

    writeStage := common.NewPipelineStage(
        common.Do(load.Accounts),
        common.Do(load.Emails),
        common.Do(load.Flags),
    )

    staging := util.Write{
        PipelineStage: writeStage,
    }

    return staging
}
