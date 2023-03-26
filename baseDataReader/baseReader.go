package main

import (
    //"os"
    "fmt"
    //"time"
    //"os/exec"
    //"database/sql"

    "github.com/jacobsa/go-serial/serial"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
/*
    // environment variables
    var dbUsername         = os.Getenv("MYSQL_USER")
    var dbPassword         = os.Getenv("MYSQL_PASS")
    var dbIp               = os.Getenv("MYSQL_IP")
    var dbPort             = os.Getenv("MYSQL_PORT")
    var dbNameData         = os.Getenv("MYSQL_DB_DATA")
    var dbTableActuatorOut = os.Getenv("MYSQL_TB_ACT_OUT")
    var dbTableSensorOut   = os.Getenv("MYSQL_TB_SEN_OUT")/*
    var dbTableActuatorIn  = os.Getenv("MYSQL_TB_ACT_IN")

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

    dbDataConn, err := sql.Open("mysql", dbDataConnString)
    if(err != nil){
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }
    defer dbDataConn.Close();
*/
    // db querys
/*    sensorFailures := "";
    temperature := "";
    humidity := "";
    accelerometer := "";
    barometer := "";
    gps := "";
    gyroscope := "";
    status := "";
    warning := "";
    battery1 := "";
    battery2 := "";
    airPump := "";
    parachuteServo := "";
    gasValve := "";*/

    // usb reader
    options := serial.OpenOptions{
        PortName:        "/dev/cu.usbmodem21101",
        BaudRate:        9600,
        DataBits:        8,
        StopBits:        1,
        MinimumReadSize: 4,
    }
    port, err := serial.Open(options)
    if err != nil {
        fmt.Println("Error: Cant open serial com\n - " + err.Error())
    }
    buf := make([]byte, 100)
    n, err := port.Read(buf)
    if err != nil {
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }
    fmt.Println("Readen", n)
    fmt.Println(string(buf))
/*
    ticket := time.NewTicker(0.2 * time.Second)
    for range ticket.C {
        //string := getDataFromUsb();



        dbSensorDataQuery := "INSERT INTO " + dbTableSensorOut + " (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES (" + sensorFailures + "," + temperature + "," + humidity + "," + accelerometer + "," + barometer + "," + gps + "," + gyroscope + ");"
        dbSensorDataQuery := "INSERT INTO " + dbTableActuatorOut + " (status, warning, battery1, battery2, airPump, parachuteServo, gasValve) VALUES (" + status + "," + warning + "," + battery1 + "," + battery2 + "," + airPump + "," + parachuteServo + "," + gasValve + ");"

    }*/
}
