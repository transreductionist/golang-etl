package queries

import (
    "fmt"
)

func SelectFromTableAll(tableName string) string {
    return fmt.Sprintf(`SELECT * FROM %v`, tableName)
}

func SelectFromTableByID(tableName string, ID int) string {
    return fmt.Sprintf(`SELECT * FROM %v WHERE ID = %v`, tableName, ID)
}

func SelectFromTableByIDRange(tableName string, ID0 int, ID1 int) string {
    return fmt.Sprintf(`SELECT * FROM %v WHERE ID >= %v AND ID <= %v`, tableName, ID0, ID1)
}
