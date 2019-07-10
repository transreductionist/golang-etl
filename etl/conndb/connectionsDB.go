package conndb

import (
    "USERetl/common/util"
    "fmt"
    "os"
    "strconv"
)

type Resources interface{
    GetDatabase() util.Database
}

func GetDatabase(envPrefix string) util.Database {
    envPort := fmt.Sprintf("%s_%s", envPrefix, "PORT")
    portAsString := os.Getenv(envPort)
    portAsInt, errAtoi := strconv.Atoi(portAsString)
    if errAtoi != nil {
        fmt.Println(errAtoi)
        os.Exit(2)
    }

    db := util.Database{
        Host:      os.Getenv(fmt.Sprintf("%s_%s", envPrefix, "HOST")),
        Port:      portAsInt,
        UserID:    os.Getenv(fmt.Sprintf("%s_%s", envPrefix, "USER")),
        PassWord:  os.Getenv(fmt.Sprintf("%s_%s", envPrefix, "PASSWORD")),
        Database:  os.Getenv(fmt.Sprintf("%s_%s", envPrefix, "SCHEMA")),
        ParseTime: true}
    return db
}
