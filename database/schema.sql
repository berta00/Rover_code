CREATE DATABASE ProbeBaloonData;
CREATE DATABASE ProbeBaloonWeb;

CREATE USER 'baseReader'@'%'   IDENTIFIED BY 'baseReader123';
CREATE USER 'webInterface'@'%' IDENTIFIED BY 'webInterface123';
CREATE USER 'EMMmantainer'@'%' IDENTIFIED BY 'EMMmantainer123';

GRANT INSERT, SELECT                 ON ProbeBaloonData.* TO 'baseReader'@'%';
GRANT INSERT, SELECT, UPDATE, DELETE ON ProbeBaloonData.* TO 'webInterface'@'%';
GRANT INSERT, SELECT, UPDATE, DELETE ON ProbeBaloonWeb.*  TO 'webInterface'@'%';
GRANT ALL                            ON ProbeBaloonData.* TO 'EMMmantainer'@'%' WITH GRANT OPTION;
GRANT ALL                            ON ProbeBaloonWeb.*  TO 'EMMmantainer'@'%'  WITH GRANT OPTION;

FLUSH PRIVILEGES;

USE ProbeBaloonData;

CREATE TABLE BaloonSensorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    sensorFailures  varchar(255),
    temperature     varchar(255),
    humidity        varchar(255),
    accelerometer   varchar(255),
    barometer       varchar(255),
    gps             varchar(255),
    gyroscope       varchar(255),
    PRIMARY KEY (id)
);
CREATE TABLE BaloonActuatorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    status          varchar(255),
    warning         varchar(255),
    battery1        varchar(255),
    battery2        varchar(255),
    airPump         varchar(255),
    parachuteServo  varchar(255),
    gasValve        varchar(255),
    PRIMARY KEY(id)
);
CREATE TABLE BaloonActuatorDataIn (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    status          int,
    warning         int,
    emergency       int,
    gasValveStatus  int,
    pumpStatus      int,
    parachuteStatus int,
    PRIMARY KEY (id)
);

USE ProbeBaloonWeb;

CREATE TABLE User (
    id              int          auto_increment,
    name            varchar(255) not null,
    email           varchar(255) not null       unique,
    password        varchar(255) not null,
    PRIMARY KEY (id)
);


# insert admin user
INSERT INTO User (name, email, password) VALUES ("admin", "bertagnollimarco999@gmail.com", "admin123");


USE ProbeBaloonData;

# insert baloon data
INSERT INTO BaloonSensorDataOut (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES ("none", 25.6, 45.3, "0.71%0.28%9.43", 101325, "05%46.416987%11.179133%N%46°25'01.2%E%11°10'44.9%1129.0", "6.95%86.58%-1.20");
INSERT INTO BaloonActuatorDataOut (status, warning, battery1, battery2, airPump, parachuteServo, gasValve) VALUES ("ok", "none", "8.6%/%none", "8.8%/%none", "4.6%none", "4.7%0%none", "3.6%34%24%34%57%none");
