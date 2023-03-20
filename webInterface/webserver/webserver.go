package main

import (
    "html/template"
    "database/sql"
    "net/http"
    "os/exec"
    "time"
    "fmt"
    "os"

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
var dbTableActuatorOut = os.Getenv("MYSQL_TB_ACT_OUT")/*
var dbTableActuatorIn  = os.Getenv("MYSQL_TB_ACT_IN")
var dbTableSensorOut   = os.Getenv("MYSQL_TB_SEN_OUT")*/

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
type actuatorDataOutTableQuery struct {
    Id              string
    DataTime        string
    Status          string
    Warning         string
    Emergency       string
    GasValveStatus  string
    PumpStatus      string
    ParachuteStatus string
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
    dbDataConn, err := sql.Open("mysql", dbDataConnString)
    if err != nil {
        fmt.Println("Error: Cant connect to data database\n - " + err.Error())
    }

    actuatorOutQueryString := "SELECT * FROM " + dbTableActuatorOut + " WHERE id=(SELECT MAX(id) FROM " + dbTableActuatorOut + ");"

    ticket := time.NewTicker(2 * time.Second)
    for range ticket.C {
        actuatorOutQuery, err := dbDataConn.Query(actuatorOutQueryString);
        if err != nil {
            fmt.Println("Error: Cant query the actoruatorOut table\n - " + err.Error())
        }

        actuatorOutQueryOut := new(actuatorDataOutTableQuery)
        for actuatorOutQuery.Next(){
            err := actuatorOutQuery.Scan(
                &actuatorOutQueryOut.Id,
                &actuatorOutQueryOut.DataTime,
                &actuatorOutQueryOut.Status,
                &actuatorOutQueryOut.Warning,
                &actuatorOutQueryOut.Emergency,
                &actuatorOutQueryOut.GasValveStatus,
                &actuatorOutQueryOut.PumpStatus,
                &actuatorOutQueryOut.ParachuteStatus,
            )
            if err != nil {
                fmt.Println("Error: Cant scan the query from user table\n - " + err.Error())
            }
        }

        actuatorOutQueryOutString := actuatorOutQueryOut.Id + "," + actuatorOutQueryOut.DataTime + "," + actuatorOutQueryOut.Status + "," + actuatorOutQueryOut.Warning + "," + actuatorOutQueryOut.Emergency + "," + actuatorOutQueryOut.GasValveStatus + "," + actuatorOutQueryOut.PumpStatus + "," + actuatorOutQueryOut.ParachuteStatus

        if err := conn.WriteMessage(websocket.TextMessage, []byte(actuatorOutQueryOutString)); err != nil {
            fmt.Println("Error: Cant write message in websocket\n - " + err.Error())
            return
        }
    }
}
