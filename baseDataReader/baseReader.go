package main

import (
    "os"
    "fmt"
    "time"
    "strings"
    "os/exec"
    "database/sql"

    "github.com/jacobsa/go-serial/serial"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    // environment variables
    var dbUsername         = os.Getenv("MYSQL_USER")
    var dbPassword         = os.Getenv("MYSQL_PASS")
    var dbIp               = os.Getenv("MYSQL_IP")
    var dbPort             = os.Getenv("MYSQL_PORT")
    var dbNameData         = os.Getenv("MYSQL_DB_DATA")/*
    var dbTableSensorOut   = os.Getenv("MYSQL_TB_SEN_OUT")
    var dbTableActuatorOut = os.Getenv("MYSQL_TB_ACT_OUT")
    var dbTableActuatorIn  = os.Getenv("MYSQL_TB_ACT_IN")*/

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

    // db querys
/*  sensorFailures := "";
    temperature := "";
    humidity := "";
    accelerometer := "";
    barometer := "";
    gps := "";
    gyroscope := "";
    status := ""
    warning := ""
    battery1 := "";
    battery2 := "";
    airPump := "";
    parachuteServo := "";
    gasValve := "";*/

    // usb reader

    ticket := time.NewTicker(200 * time.Millisecond)
    for range ticket.C {
        sensorDataString := getDataFromUsb()
        sensorDataParsed := strings.Split(sensorDataString, ",")
/*
        sensorDataGps := strings.Split(sensorDataParsed[3], "%")

        sensorDataGps[3] = "dms1_1"
        sensorDataGps[4] = "dms1_2"
        sensorDataGps[5] = "dms2_1"
        sensorDataGps[6] = "dms2_2"

        gpsString := ""
        for a := 0; a < len(sensorDataGps); a++ {
            if(a != 0){
                gpsString += "%"
            }
            gpsString += sensorDataGps[a];
        }

        dbSensorDataQuery1 := "INSERT INTO " + dbTableSensorOut + " (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES (" + string(sensorDataString[0]) + "," + string(sensorDataString[1]) + "," + string(sensorDataString[2]) + "," + string(sensorDataString[3]) + "," + string(sensorDataString[4]) + "," + gpsString + "," + string(sensorDataString[6]) + ");"
        dbSensorDataQuery2 := "INSERT INTO " + dbTableActuatorOut + " (status, warning, battery1, battery2, airPump, parachuteServo, gasValve) VALUES (" + status + "," + warning + "," + string(sensorDataString[7]) + "," + string(sensorDataString[8]) + "," + string(sensorDataString[9]) + "," + string(sensorDataString[10]) + "," + string(sensorDataString[11]) + ");"
*/
        fmt.Println(sensorDataParsed)

        //dbSensorDataQuery1 := "INSERT INTO " + dbTableSensorOut + " (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES ('" + string(sensorDataParsed[0]) + "','" + string(sensorDataParsed[1]) + "','" + string(sensorDataParsed[2]) + "','','','','');"
        //dbSensorDataQuery2 := "INSERT INTO " + dbTableActuatorOut + " (status, warning, battery1, battery2, airPump, parachuteServo, gasValve) VALUES (" + status + "," + warning + "," + string(sensorDataString[7]) + "," + string(sensorDataString[8]) + "," + string(sensorDataString[9]) + "," + string(sensorDataString[10]) + "," + string(sensorDataString[11]) + ");"

        /*
        _, err1 := dbDataConn.Query(dbSensorDataQuery1)
        if err1 != nil {
            fmt.Println(err1.Error())
        }
        _, err2 := dbDataConn.Query(dbSensorDataQuery2)
        if err2 != nil {
            fmt.Println(err2.Error())
        }
        */
    }
}

func getDataFromUsb() string {
    options := serial.OpenOptions{
        PortName:        "/dev/cu.usbmodem21201",
        BaudRate:        9600,
        DataBits:        8,
        StopBits:        1,
        MinimumReadSize: 4,
    }
    port, err := serial.Open(options)
    if err != nil {
        fmt.Println("Error: Cant open serial connection\n - " + err.Error())
    }
    buf := make([]byte, 100)
    n, err := port.Read(buf)
    if err != nil {
        fmt.Println("Error: Cant read serial comunication\n - " + err.Error())
        fmt.Println(n)
    }
    return string(buf)
}
