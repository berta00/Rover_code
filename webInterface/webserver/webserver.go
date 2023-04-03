package main

import (
    "html/template"
    "database/sql"
    "net/http"
    "os/exec"
    "strconv"
    "time"
    "fmt"
    "os"

    "github.com/jacobsa/go-serial/serial"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/websocket"
)

// environment variables
var dbUsername         = os.Getenv("MYSQL_USER")
var dbPassword         = os.Getenv("MYSQL_PASS")
var dbIp               = os.Getenv("MYSQL_IP")
var dbPort             = os.Getenv("MYSQL_PORT")
var dbNameWeb          = os.Getenv("MYSQL_DB_WEB")
var dbNameData         = os.Getenv("MYSQL_DB_DATA")
var dbTableUser        = os.Getenv("MYSQL_TB_USER")
var dbTableActuatorOut = os.Getenv("MYSQL_TB_ACT_OUT")
var dbTableSensorOut   = os.Getenv("MYSQL_TB_SEN_OUT")/*
var dbTableActuatorIn  = os.Getenv("MYSQL_TB_ACT_IN")*/

// global variables
var domainName       = "http://127.0.0.1"
var dbDataConnString = dbUsername + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbNameData
var dbWebConnString = dbUsername + ":" + dbPassword + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbNameWeb

// websocket
var upgrader = websocket.Upgrader {
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// database query structures
type userTableQuery struct {
    Id       int
    Name     string
    Email    string
    Password string
}
type sensorDataOutTableQuery struct {
    Id              string
    DataTime        string
    SensorFailures  string
    Temperature     string
    Humidity        string
    Accelerometer   string
    Barometer       string
    Gps             string
    Gyroscope       string
}
type actuatorDataOutTableQuery struct {
    Id              string
    DataTime        string
    Status          string
    Warning         string
    Battery1        string
    Battery2        string
    AirPump         string
    GasValve        string
    ParachuteServo  string
}

// html pages structures
type dashboardData struct {
    Name               string
    Email              string
}

func main(){
    //cwd, _ := os.Getwd()

    // ceck if environment variables are setted
    if dbIp == "" {
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

    // static file handling
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static", fs))

    // route initialization
    http.HandleFunc("/",            redirectLoginPage)
    http.HandleFunc("/login",       loginPage)
    http.HandleFunc("/dashboard",   dashboardPage)
    http.HandleFunc("/dashboardWs", dashboardWebSocket)

    http.ListenAndServe(":80", nil)
}

func redirectLoginPage(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, domainName + "/login", http.StatusSeeOther)
}

func loginPage(w http.ResponseWriter, r *http.Request){
    cwd, _ := os.Getwd()

    htmlTemplate, _ := template.ParseFiles(cwd + "/static/pages/login.html")
    htmlTemplate.Execute(w, 0)
}

func dashboardPage(w http.ResponseWriter, r *http.Request){
    cwd, _ := os.Getwd()

    dbWebConn, err := sql.Open("mysql", dbWebConnString)
    if err != nil {
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }

    inputEmail := "";
    inputPassword := "";

    switch r.Method {
    case "POST":

        inputEmail = r.FormValue("email");
        inputPassword = r.FormValue("password");

        defer dbWebConn.Close()
        break

    default:
        http.Redirect(w, r, domainName + "/login", http.StatusSeeOther)
        break
    }

    emailQueryString := "SELECT * FROM " + dbTableUser + " WHERE email='" + inputEmail + "';"
    emailQuery, err := dbWebConn.Query(emailQueryString)
    if err != nil {
        fmt.Println("Error: Cant query the user table\n - " + err.Error())
    }

    emailQueryOut := new(userTableQuery)
    for emailQuery.Next() {
        err := emailQuery.Scan(&emailQueryOut.Id, &emailQueryOut.Name, &emailQueryOut.Email, &emailQueryOut.Password);
        if err != nil {
            fmt.Println("Error: Cant scan the query from user table\n - " + err.Error())
        }
    }

    if inputPassword != emailQueryOut.Password {
        http.Redirect(w, r, domainName + "/login", http.StatusSeeOther)
    }

    dashData := new(dashboardData)

    dashData.Name               = emailQueryOut.Name
    dashData.Email              = emailQueryOut.Email

    htmlTemplate, _ := template.ParseFiles(cwd + "/static/pages/dashboard.html")
    htmlTemplate.Execute(w, dashData)
}

func dashboardWebSocket(w http.ResponseWriter, r *http.Request){
    ws, err := upgrade(w, r);
    if err != nil {
        fmt.Println("Error: Cant upgrade to websocket conn\n - " + err.Error())
    }

    go writer(ws);
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    ws, err := upgrader.Upgrade(w, r, nil);
    if err != nil {
        return ws, err
    }

    return ws, nil
}

func writer(conn *websocket.Conn){
/*
    dbDataConn, err := sql.Open("mysql", dbDataConnString)
    if err != nil {
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }

    sensorOutQueryString := "SELECT * FROM " + dbTableSensorOut + " WHERE id=(SELECT MAX(id) FROM " + dbTableSensorOut + ");"
    actuatorOutQueryString := "SELECT * FROM " + dbTableActuatorOut + " WHERE id=(SELECT MAX(id) FROM " + dbTableActuatorOut + ");"

    webInterfaceDataString := "";

    maxTemp := 0.0
    minTemp := 999.0
    maxHum := 0.0
    minHum := 999.0

    ticket := time.NewTicker(150 * time.Millisecond)
    for range ticket.C {
        actuatorOutQuery, err1 := dbDataConn.Query(actuatorOutQueryString);
        sensorOutQuery, err2 := dbDataConn.Query(sensorOutQueryString);
        if err1 != nil { fmt.Println("Error: Cant query the actoruatorOut table\n - " + err1.Error()) }
        if err2 != nil { fmt.Println("Error: Cant query the actoruatorOut table\n - " + err2.Error()) }

        actuatorOutQueryOut := new(actuatorDataOutTableQuery)
        for actuatorOutQuery.Next(){
            err1 := actuatorOutQuery.Scan(
                &actuatorOutQueryOut.Id,
                &actuatorOutQueryOut.DataTime,
                &actuatorOutQueryOut.Status,
                &actuatorOutQueryOut.Warning,
                &actuatorOutQueryOut.Battery1,
                &actuatorOutQueryOut.Battery2,
                &actuatorOutQueryOut.AirPump,
                &actuatorOutQueryOut.ParachuteServo,
                &actuatorOutQueryOut.GasValve,
            )
            if err1 != nil {
                fmt.Println("Error: Cant scan the query from actoruatorOut table\n - " + err1.Error())
            }
        }

        sensorOutQueryOut := new(sensorDataOutTableQuery)
        for sensorOutQuery.Next(){
            err2 := sensorOutQuery.Scan(
                &sensorOutQueryOut.Id,
                &sensorOutQueryOut.DataTime,
                &sensorOutQueryOut.SensorFailures,
                &sensorOutQueryOut.Temperature,
                &sensorOutQueryOut.Humidity,
                &sensorOutQueryOut.Accelerometer,
                &sensorOutQueryOut.Barometer,
                &sensorOutQueryOut.Gps,
                &sensorOutQueryOut.Gyroscope,
            )
            if err2 != nil {
                fmt.Println("Error: Cant scan the query from sensorOut table\n - " + err2.Error())
            }
        }

        curTemp, _ := strconv.ParseFloat(sensorOutQueryOut.Temperature, 8)
        curHum, _ := strconv.ParseFloat(sensorOutQueryOut.Humidity, 8)

        if(curTemp > maxTemp){
            maxTemp = curTemp
        } else if(curTemp < minTemp){
            minTemp = curTemp
        }
        if(curHum > maxHum){
            maxHum = curHum
        } else if(curHum < minHum){
            minHum = curHum
        }

        webInterfaceDataString = actuatorOutQueryOut.DataTime + "," + actuatorOutQueryOut.Status + "," + actuatorOutQueryOut.Warning + "," + actuatorOutQueryOut.Battery1 + "," + actuatorOutQueryOut.Battery2 + "," + actuatorOutQueryOut.AirPump + "," + actuatorOutQueryOut.ParachuteServo + "," + actuatorOutQueryOut.GasValve
        webInterfaceDataString += "," + fmt.Sprintf("%v",curTemp) + "%" + fmt.Sprintf("%v",maxTemp) + "%" + fmt.Sprintf("%v",minTemp) + "," + fmt.Sprintf("%v",curHum) + "%" + fmt.Sprintf("%v",maxHum) + "%" + fmt.Sprintf("%v",minHum) + "," + sensorOutQueryOut.Accelerometer + "," + sensorOutQueryOut.Barometer + "," + sensorOutQueryOut.Gps + "," + sensorOutQueryOut.Gyroscope

        if err := conn.WriteMessage(websocket.TextMessage, []byte(webInterfaceDataString)); err != nil {
            fmt.Println("Error: Cant write message in websocket\n - " + err.Error())
            return
        }
    }
*/
    dbDataConn, err := sql.Open("mysql", dbDataConnString)
    if err != nil {
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }

    sensorInQueryString := ""
    actuatorInQueryString := ""
    webInterfaceDataString := ""

    maxTemp := 0.0
    minTemp := 999.0
    maxHum := 0.0
    minHum := 999.0

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

    ticket := time.NewTicker(150 * time.Millisecond)
    for range ticket.C {
        // get data from arduino
        dataBuf := make([]byte, 100)
        n, err := port.Read(dataBuf)
        if err != nil {
            fmt.Println("Error: Cant read serial comunication\n - " + err.Error())
            fmt.Println(n)
        }
        // parse and correct data
        parsedSensorData := strings.Split(sensorDataBuf, ",")

        curTemp, _ := strconv.ParseFloat(parsedSensorData[1], 8)
        curHum, _ := strconv.ParseFloat(parsedSensorData[2], 8)

        if(curTemp > maxTemp){
            maxTemp = curTemp
        } else if(curTemp < minTemp){
            minTemp = curTemp
        }
        if(curHum > maxHum){
            maxHum = curHum
        } else if(curHum < minHum){
            minHum = curHum
        }

        warning := "none" // ceck for warnings

        webInterfaceDataString = time.Now() + "," + parsedSensorData[0] + "," + warning + "," + parsedSensorData[7] + "," + parsedSensorData[8] + "," + parsedSensorData[9] + "," + parsedSensorData[10] + "," + parsedSensorData[11]
        webInterfaceDataString += "," + fmt.Sprintf("%v",curTemp) + "%" + fmt.Sprintf("%v",maxTemp) + "%" + fmt.Sprintf("%v",minTemp) + "," + fmt.Sprintf("%v",curHum) + "%" + fmt.Sprintf("%v",maxHum) + "%" + fmt.Sprintf("%v",minHum) + "," + parsedSensorData[3] + "," + parsedSensorData[4] + "," + parsedSensorData[5] + "," + parsedSensorData[6]
        // send data to websocket
        if err := conn.WriteMessage(websocket.TextMessage, []byte(dataBuf)); err != nil {
            fmt.Println("Error: Cant write message in websocket\n - " + err.Error())
            return
        }
    }
}
