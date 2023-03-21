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
    temperature     float,
    humidity        float,
    accelerometer   varchar(255),
    barometer       float,
    gps             varchar(255),
    gyroscope       varchar(255),
    PRIMARY KEY (id)
);
CREATE TABLE BaloonActuatorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    status          int,
    warning         int,
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

# insert sensorDataOut
INSERT INTO BaloonSensorDataOut (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES ("none", 25.6, 45.3, "0.71%0.28%9.43", 101325, "05%41%008%N%XX.XXXXX%W%31.54761%1129.0", "6.95%86.58%-1.20");
