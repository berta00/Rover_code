// Baloon
#include <string.h>

#include <SPI.h>
#include <RF24.h>

#include <Servo.h>

#include <DHT.h>

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
int pinAirPump = 13;
int pinParachuteServo = 22;
int pinAntennaMISO = 50;
int pinAntennaMOSI = 51;
int pinAntennaSCK = 52;
int pinAntennaSS = 53;
int pinAccelerometerSCL = 0;
int pinAccelerometerSDA = 1;
int pinBarometerSCL = 2;
int pinBarometerSDA = 3;
int pinBattery1V = 4;
int pinBattery2V = 5;
int pinStepperV = 6;
int pinServoV = 7;
int pinAirPumpV = 8;

// antenna variables and object
RF24 radio(7, 8);
byte node_A_address[6] = "NodeA";
byte node_B_address[6] = "NodeB";

// motor, airPump, parachute
int actuatorsData[12] = {0, 0, 0, 0};
// temperature, humidity, gps, barometer, airPump, accelerometer, stepper,  parachute, stepperV, servoV, airPumpV
char sensorsData[12][32] = {"101", "", "", "", "", "", "", "",  "", "", "", ""};
char sensorDataString[32] = "";

// actuator variables
int currentDataIntegrity;
int currentMotor;
int currentAirPump;
int currentParachute;

// sensorValiables
float currentTemperature;
char currentTemperatureS[32];
float currentHumidity;
char currentHumidityS[32];
char currentCoordinatesS[32];
char currentBarometerS[32];
char currentAirpumpS[32];
char currentAccelerometerS[32];
char currentStepperS[32];
char currentParachuteS[32];
char currentStepperVS[32];
char currentServoVS[32];
char currentAirPumpVS[32];

// temperature humidity object
DHT dht(pinTemperatureHumidity, DHT11);

void setup() {
  // serial setup
  Serial.begin(9600);
  // antenna setup
  radio.begin();
  radio.setPALevel(RF24_PA_LOW); //RF24_PA_MAX 
  radio.openWritingPipe(node_A_address);
  radio.openReadingPipe(1, node_B_address);
  radio.startListening();
  // temperature humidity setup
  dht.begin();
}

void loop() {
  // temperature and humidity
  currentTemperature = dht.readTemperature();
  currentHumidity = dht.readHumidity();
  dtostrf(currentTemperature,3, 2, currentTemperatureS);
  dtostrf(currentHumidity, 3, 2, currentHumidityS);

  // asign data to array for stansmition
/*strcpy(sensorsData[0], "101");
  strcpy(sensorsData[1], currentTemperatureS);
  strcpy(sensorsData[2], currentHumidityS);
  strcpy(sensorsData[3], currentCoordinatesS);
  strcpy(sensorsData[4], currentBarometerS);
  strcpy(sensorsData[5], currentAirpumpS);
  strcpy(sensorsData[6], currentAccelerometerS);
  strcpy(sensorsData[7], currentStepperS);
  strcpy(sensorsData[8], currentParachuteS);
  strcpy(sensorsData[9], currentStepperVS);
  strcpy(sensorsData[10], currentServoVS);
  strcpy(sensorsData[11], currentAirPumpVS);*/

  strcpy(sensorDataString, "");
  strcat(sensorDataString, currentTemperatureS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentHumidityS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentCoordinatesS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentBarometerS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentAirpumpS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentAccelerometerS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentStepperS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentParachuteS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentStepperVS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentServoVS);
  strcat(sensorDataString, ",");
  strcat(sensorDataString, currentAirPumpVS);
  
  Serial.println(sensorDataString);

  // antenna comunication
  if (radio.available()){
    while (radio.available()) {
      radio.read(&actuatorsData, sizeof(actuatorsData));
    }

    radio.stopListening();

    radio.write(&sensorDataString, sizeof(sensorDataString));
    radio.startListening();

    Serial.println(actuatorsData[0]);
  }

  delay(100);
}
