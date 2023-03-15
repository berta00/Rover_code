CREATE DATABASE probeBaloonData;
CREATE DATABASE probeBaloonWeb;

CREATE USER 'baseReader'@'%'   IDENTIFIED BY 'baseReader123';
CREATE USER 'webInterface'@'%' IDENTIFIED BY 'webInterface123';
CREATE USER 'EMMmantainer'@'%' IDENTIFIED BY 'EMMmantainer123';

GRANT INSERT, SELECT                 ON probeBaloonData.* TO 'baseReader'@'%';
GRANT INSERT, SELECT, UPDATE, DELETE ON probeBaloonData.* TO 'webInterface'@'%';
GRANT ALL                            ON probeBaloonData.* TO 'EMMmantainer'@'%' WITH GRANT OPTION;
GRANT ALL                            ON probeBaloonWeb.*  TO 'EMMmantainer'@'%'  WITH GRANT OPTION;

USE probeBaloonData;

CREATE TABLE baloonSensorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    sensorFailures  varchar(255),
    temperature     float,
    humidity        float,
    gyroscope       varchar(255),
    accelerometer   varchar(255),
    gps             varchar(255),
    barometer       float,
    PRIMARY KEY(id)
);
CREATE TABLE baloonActuatorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    warning         int,
    emergency       int,
    gasValveStatus  int,
    pumpStatus      int,
    parachuteStatus int,
    PRIMARY KEY(id)
)
CREATE TABLE baloonActuatorDataIn (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    warning         int,
    emergency       int,
    gasValveStatus  int,
    pumpStatus      int,
    parachuteStatus int,
    PRIMARY KEY(id)
)

USE probeBaloonWeb;

CREATE TABLE user (
    id              int          auto_increment,
    name            varchar(255) not null,
    email           varchar(255) not null,
    password        varchar(255) not null,
    PRIMARY KEY(id)
)
