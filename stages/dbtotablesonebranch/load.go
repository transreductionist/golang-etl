package dbtotablesonebranch

import (
    "USERetl/common"
    "USERetl/util"
)

func Load(load util.LoadDB) util.Write {

    writeStage := common.NewPipelineStage(
        common.Do(load.SQLWriteAccts),
        common.Do(load.SQLWriteFlags),
    )

    staging := util.Write{
        PipelineStage: writeStage,
    }

    return staging
}

