package buildenv

import "os"

func Local() string {
    _ = os.Setenv("LOCAL_SERVER_HOST", "127.0.0.1")
    _ = os.Setenv("LOCAL_SERVER_PORT", "3306")
    _ = os.Setenv("LOCAL_SERVER_USER", "root")
    _ = os.Setenv("LOCAL_SERVER_PASSWORD", "Password")
    _ = os.Setenv("LOCAL_SERVER_SCHEMA", "nusa_user")
    return "LOCAL_SERVER"
}
