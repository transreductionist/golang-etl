package buildenv

import "os"

func Remote() string {
    _ = os.Setenv("REMOTE_SERVER_HOST", "199.167.77.248")
    _ = os.Setenv("REMOTE_SERVER_PORT", "3306")
    _ = os.Setenv("REMOTE_SERVER_USER", "datastudy")
    _ = os.Setenv("REMOTE_SERVER_PASSWORD", "bACjfbaQf0t9tbd0zpSYIxxwT4ZG3yFq")
    _ = os.Setenv("REMOTE_SERVER_SCHEMA", "ultsys")
    return "REMOTE_SERVER"
}
