// Baloon
#include <string.h>

#include <SPI.h>
#include <RF24.h>

#include <DHT.h>

#include <SoftwareSerial.h>
#include <TinyGPS++.h>

#include <Adafruit_BMP085.h>

#include <Adafruit_MPU6050.h>
#include <Adafruit_Sensor.h>
#include <Wire.h>

// setup vars
bool debugMode = false;

// pins
int pinGpsRX = 0;
int pinGpsTX = 1;
int pinTemperatureHumidity = 2;
int pinAntennaCE = 7;
int pinAntennaCNS = 8;
int pinStepperINT1 = 9;
int pinStepperINT2 = 10;
int pinStepperINT3 = 11;
int pinStepperINT4 = 12;
int pinAirPump = 49;
int pinParachuteServo = 22;
int pinAntennaMISO = 50;
int pinAntennaMOSI = 51;
int pinAntennaSCK = 52;
int pinAntennaSS = 53;
int pinBarometerSCL = 21;
int pinBarometerSDA = 22;
int pinAccelerometerSCL = A0;
int pinAccelerometerSDA = A1;
int pinBattery1V = A4;
int pinBattery2V = A5;
int pinStepperV = A6;
int pinServoV = A7;
int pinAirPumpV = A8;

// antenna variables and object
RF24 radio(pinAntennaCE, pinAntennaCNS);
byte node_A_address[6] = "NodeA";
byte node_B_address[6] = "NodeB";

// motor, airPump, parachute
int actuatorsData[12] = {0, 0, 0, 0};
// temperature, humidity, gps, barometer, airPump, accelerometer, stepper,  parachute, stepperV, servoV, airPumpV
char sensorDataString[64] = "";

// actuator variables
int currentDataIntegrity;
int currentMotor;
boolean airPumpEnable = false;
int currentParachute;

// sensorValiables
float currentTemperature;
char currentTemperatureS[32];
float currentHumidity;
char currentHumidityS[16];
double currentCoordinatesLat;
char currentCoordinatesLatS[12];
double currentCoordinatesLng;
char currentCoordinatesLngS[12];
int currentSatelitesNumber;
char currentSatelitesNumberS[8];
int currentSpeed;
char currentSpeedS[16];
int currentSeaAltitude;
int currentGroundAltitude;
int currentRelativeAltitude;
char currentGroundAltitudeS[16];
char currentCoordinatesS[16];
int currentPressure;
int currentAltitude;
char currentBarometerS[16];
char currentXAccelerationS[8];
char currentYAccelerationS[8];
char currentZAccelerationS[8];
char currentAccelerometerS[16];
char currentXGyroscopeS[8];
char currentYGyroscopeS[8];
char currentZGyroscopeS[8];
char currentGyroscopeS[16];
char battery1S[16];
char battery2S[16];
char currentAirpumpS[16];
char currentParachuteServoS[16];
char currentGasValveS[16];

// temperature humidity object
DHT dht(pinTemperatureHumidity, DHT11);
// gps object
SoftwareSerial gpsSerial(pinGpsRX, pinGpsTX);
TinyGPSPlus gpsParser;
// barometer object
Adafruit_BMP085 barometer;
// accelerometer gyroscope object
Adafruit_MPU6050 accGyro;

void setup() {
  // serial setup
  Serial.begin(115200);
  // antenna setup
  radio.begin();
  radio.setPALevel(RF24_PA_MAX); //RF24_PA_MAX  RF24_PA_LOW
  radio.openWritingPipe(node_A_address);
  radio.openReadingPipe(1, node_B_address);
  radio.setChannel(1);
  radio.startListening();
  // temperature humidity setup
  dht.begin();
  // barometer setup
  barometer.begin();
  // accelerometer gyroscope scope
  accGyro.begin();
  accGyro.setAccelerometerRange(MPU6050_RANGE_8_G);
  accGyro.setGyroRange(MPU6050_RANGE_500_DEG);
  accGyro.setFilterBandwidth(MPU6050_BAND_5_HZ);
  // air pump setup
  pinMode(pinAirPump, OUTPUT);
}

void loop() {
  // temperature and humidity
  currentTemperature = dht.readTemperature();
  currentHumidity = dht.readHumidity();
  dtostrf(currentTemperature,3, 2, currentTemperatureS);
  dtostrf(currentHumidity, 3, 2, currentHumidityS);

  // gps
  if(gpsSerial.available() > 0){
    gpsParser.encode(gpsSerial.read());
    currentCoordinatesLat = gpsParser.location.lat();
    currentCoordinatesLng = gpsParser.location.lng();
    currentGroundAltitude = gpsParser.altitude.meters();
    currentSatelitesNumber = gpsParser.satellites.value();
    currentSpeed = gpsParser.speed.kmph();
    dtostrf(currentCoordinatesLat, 3, 6, currentCoordinatesLatS);
    dtostrf(currentCoordinatesLng, 3, 6, currentCoordinatesLngS);
    dtostrf(currentGroundAltitude, 3, 2, currentGroundAltitudeS);
    dtostrf(currentSatelitesNumber, 3, 0, currentSatelitesNumberS);
    dtostrf(currentSpeed, 3, 1, currentSpeedS);
    strcpy(currentCoordinatesS, "");
    strcat(currentCoordinatesS, currentSatelitesNumberS);
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, currentCoordinatesLatS);
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, currentCoordinatesLngS);
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, currentSpeed);
    strcat(currentCoordinatesS, "%");
    strcat(currentCoordinatesS, currentGroundAltitudeS);
  } else {
    strcpy(currentCoordinatesS, "");
    strcat(currentCoordinatesS, "disconnected");
  }

  // barometer
  currentPressure = barometer.readPressure();
  currentSeaAltitude = barometer.readAltitude();
  currentRelativeAltitude = currentSeaAltitude - currentGroundAltitude;
  strcpy(currentBarometerS, "");
  strcat(currentBarometerS, currentPressure);
  strcat(currentBarometerS, "%");
  strcat(currentBarometerS, currentSeaAltitude);
  strcat(currentBarometerS, "%");
  strcat(currentBarometerS, currentRelativeAltitude);

  // accelerometer gyroscope
  sensors_event_t accelerometer, gyroscope, temp;
  accGyro.getEvent(&accelerometer, &gyroscope, &temp);
  float currentXAcceleration = accelerometer.acceleration.x;
  float currentYAcceleration = accelerometer.acceleration.y;
  float currentZAcceleration = accelerometer.acceleration.z;
  dtostrf(currentXAcceleration, 3, 2, currentXAccelerationS);
  dtostrf(currentYAcceleration, 3, 2, currentYAccelerationS);
  dtostrf(currentZAcceleration, 3, 2, currentZAccelerationS);
  strcpy(currentAccelerometerS, "");
  strcat(currentAccelerometerS, currentXAccelerationS);
  strcat(currentAccelerometerS, "%");
  strcat(currentAccelerometerS, currentYAccelerationS);
  strcat(currentAccelerometerS, "%");
  strcat(currentAccelerometerS, currentZAccelerationS);
  float currentXGyroscope = gyroscope.gyro.x;
  float currentYGyroscope = gyroscope.gyro.y;
  float currentZGyroscope = gyroscope.gyro.z;
  dtostrf(currentXGyroscope, 3, 2, currentXGyroscopeS);
  dtostrf(currentYGyroscope, 3, 2, currentYGyroscopeS);
  dtostrf(currentZGyroscope, 3, 2, currentZGyroscopeS);
  strcpy(currentGyroscopeS, "");
  strcat(currentGyroscopeS, currentXGyroscopeS);
  strcat(currentGyroscopeS, "%");
  strcat(currentGyroscopeS, currentYGyroscopeS);
  strcat(currentGyroscopeS, "%");
  strcat(currentGyroscopeS, currentZGyroscopeS);
/*
  // battery1
  strcpy(battery1S, "");
  strcat(battery1S, "none");
  strcat(battery1S, "%");
  strcat(battery1S, "0.0");
  strcat(battery1S, "%");
  strcat(battery1S, "0.0");

  // battery2
  strcpy(battery2S, "");
  strcat(battery2S, "none");
  strcat(battery2S, "%");
  strcat(battery2S, "0.0");
  strcat(battery2S, "%");
  strcat(battery2S, "0.0");

  // air pump
  airPumpEnable = false;
  if(airPumpEnable){
    digitalWrite(pinAirPump, HIGH);
  } else {
    digitalWrite(pinAirPump, LOW);
  }

  float airPumpVin = (float)analogRead(pinAirPumpV) * 5.0 / 1023.0;
  char airPumpVinS[8];
  dtostrf(airPumpVin, 2, 2, airPumpVinS);
  strcpy(currentAirpumpS, "");
  strcat(currentAirpumpS, "none");
  strcat(currentAirpumpS, "%");
  strcat(currentAirpumpS, airPumpVinS);

  // parachute servo
*/
  // asign data to array for stansmition
  strcpy(sensorDataString, "");
  strcat(sensorDataString, "none");
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentTemperatureS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentHumidityS);

  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentAccelerometerS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentBarometerS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentCoordinatesS);

  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentGyroscopeS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, battery1S);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, battery2S);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentAirpumpS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentParachuteServoS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentGasValveS);

  // antenna comunication
  if(radio.available()){
    while (radio.available()) {
      radio.read(&actuatorsData, sizeof(actuatorsData));
    }

    radio.stopListening();

    radio.write(&sensorDataString, sizeof(sensorDataString));
    radio.startListening();
  }
  delay(20);

  Serial.println(currentGyroscopeS);
}
