package model

import "os"
import "github.com/m3ng9i/qreader/global"


// Get size of sqlite3 database file.
func DBSize() (size int64, err error) {

    info, err := os.Stat(global.PathDB)
    if err != nil {
        return
    }

    size = info.Size()
    return
}
