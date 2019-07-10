package buildenv

import "os"

func Forwarding() string {
    _ = os.Setenv("FORWARDING_SERVER_HOST", "127.0.0.1")
    _ = os.Setenv("FORWARDING_SERVER_PORT", "3307")
    _ = os.Setenv("FORWARDING_SERVER_USER", "datastudy")
    _ = os.Setenv("FORWARDING_SERVER_PASSWORD", "bACjfbaQf0t9tbd0zpSYIxxwT4ZG3yFq")
    _ = os.Setenv("FORWARDING_SERVER_SCHEMA", "ultsys")
    return "FORWARDING_SERVER"
}