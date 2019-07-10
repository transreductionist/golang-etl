package main

import (
    "USERetl/common"
    "USERetl/stages/dbtotablesonebranch"
    "USERetl/util"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    query := dbtotablesonebranch.QueryToTransform()
    load := dbtotablesonebranch.TransformToLoad(query)
    write := dbtotablesonebranch.Load(load)

    layout, _ := common.NewPipelineLayout(
        query.PipelineStage,
        load.PipelineStage,
        write.PipelineStage,)

    pipeline := common.NewBranchingPipeline(layout)
    pipeline.Name = "DBtoDBMultipleTables."

    util.RunPipeline(pipeline)

    stats := pipeline.Stats()
    fmt.Printf("\n ***** Stats:\n\n %s", stats)
}
