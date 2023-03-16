package main

import (
    "os"
    "fmt"
    "os/exec"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main(){

    // environment variables
    dbUsername := os.Getenv("MYSQL_USER")
    dbPassword := os.Getenv("MYSQL_PASS")
    dbIp       := os.Getenv("MYSQL_IP")
    dbPort     := os.Getenv("MYSQL_PORT")
    dbNameData := os.Getenv("MYSQL_DB_DATA")/*
    dbTableSensorOut   := os.Getenv("MYSQL_DB_DATA")
    dbTableActuatorOut := os.Getenv("MYSQL_DB_DATA")
    dbTableActuatorIn  := os.Getenv("MYSQL_DB_DATA")*/

    if(dbIp == ""){
        currentPath, _ := os.Getwd()
        configEnvPath := currentPath + "/env/config.sh"
        secretEnvPath := currentPath + "/env/secret.sh"
        err1 := exec.Command("source", configEnvPath).Run()
        err2 := exec.Command("source", secretEnvPath).Run()
        if(err1 != nil || err2 != nil){
            fmt.Println("Error: Cant apply environment variables\n - " + err1.Error() + "\n - " + err2.Error())
            os.Exit(1)
        }
    }

    // db connections
    dbDataConnString := dbUsername + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbNameData

    dbDataConn := sql.Open("mysql", dbDataConnString)
    defer dbDataConn.Close();

    
}
