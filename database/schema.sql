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
    gyroscope       varchar(255),
    accelerometer   varchar(255),
    gps             varchar(255),
    barometer       float,
    PRIMARY KEY (id)
);
CREATE TABLE BaloonActuatorDataOut (
    id              int          auto_increment,
    dataTime        timestamp    default        current_timestamp,
    status          int,
    warning         int,
    emergency       int,
    gasValveStatus  int,
    pumpStatus      int,
    parachuteStatus int,
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
