package main

import (
	"Eagle/DbDR/internal/config"
	"Eagle/DbDR/internal/model"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
    var config config.ServiceConfig;

    configBytes, err := os.ReadFile("./configs/local.json")
    if err != nil {
        panic("Cannot read config.json")
    }
    
    err = json.Unmarshal(configBytes, &config)
    fmt.Println(config)
    if err != nil {
        panic("couldnt unmarshall config")
    }

    sa_pw, exists := os.LookupEnv("SA_PW")
    if !exists {
        panic("no SA password environment variable")
    }

    dsn := fmt.Sprintf("sqlserver://sa:%s@%s:%v?database=%s", sa_pw, config.Host, config.Port, config.Db)
    db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil{
        fmt.Println("Could not open database")
    }

    myModel := new(model.MyModel)
    db.AutoMigrate(&myModel)

    var record model.MyModel
    max_row := getMaxID(db)
    for i := range (config.Num_rows + 1) {
        max_row += 1
        record.SetIDTo(max_row)
        record.SetTimeTo(time.Now())
        record.SetUnitTo("C")
        record.SetValueTo(float64(i) * 1.313)
        db.Create(&record)
        time.Sleep(time.Duration(config.Ms_delay))
    }
}

func getMaxID (db *gorm.DB) (uint) {
    var max uint
    db.Raw("SELECT MAX(ID) FROM my_models").Scan(&max)
    return max
}
