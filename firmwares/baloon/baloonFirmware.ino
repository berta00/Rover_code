// Baloon

#include <SPI.h>
#include <nRF24L01.h>
#include <RF24.h>

#include <Servo.h>

#include <DHT.h>

#include <Adafruit_MPU6050.h>
#include <Adafruit_Sensor.h>
#include <Wire.h>

// global sensor and actuators data (101 is to ceck datas integrity)
int actuatorsData[] = {0, 0, 0, 0}; // motor, airPump, parachute
int sensorsData[30] = {101, 0, 0, 0}; // temperature, humidity, parachuteStatus

// current sensors vars
int currentDataIntegrity = 0;
int currentMotor = 0;
int currentAirPump = 0;
int currentParachute = 0;

// setup vars
bool sensorsDebugMode = false;
bool actuatorsDebugMode = true;

// radio object and vars (pong setup)
RF24 radio(7, 8); // CE, CSN
byte node_A_address[6] = "NodeA";
byte node_B_address[6] = "NodeB";

// buttons vars
int emergencyKillStatus = 0;
const int pinEmergencyKill = 22;
const int pinParachuteButton = 23;

// airPump vars
const int pinAirPump = 2;

// parachute abject and vars
Servo servoParachute;
const int pinParachute = 3;
int parachuteStatus = 0;

// gyrocope object and vars
Adafruit_MPU6050 gyrocope;

// temp sensor vars
const int pinTempSensor = 49;
DHT dht(pinTempSensor, DHT11);

// force actuators var
int forceMotorPos = 0;
int forceAirPump = 0;
int forceParachute = 0;

void setup() {
  // serial setup
  Serial.begin(9600);
  // radio setup
  radio.begin();
  radio.setPALevel(RF24_PA_LOW);
  radio.openWritingPipe(node_A_address);
  radio.openReadingPipe(1, node_B_address);
  radio.startListening();
  // buttons setup
  pinMode(pinEmergencyKill, INPUT);
  pinMode(pinParachuteButton, INPUT);
  // air pump setup
  pinMode(pinAirPump, OUTPUT);
  pinMode(pinParachute, OUTPUT);
  // parachute setup
  servoParachute.attach(pinParachute);
  servoParachute.write(0);
  // gyrocope setup
  gyrocope.setAccelerometerRange(MPU6050_RANGE_16_G);
  gyrocope.setGyroRange(MPU6050_RANGE_250_DEG);
  gyrocope.setFilterBandwidth(MPU6050_BAND_21_HZ);
  // temperature setup
  dht.begin();

  delay(800);
}

void loop() {
  int actuatorsData[] = {0, 0, 0, 0};

  // move actuators
  if(currentDataIntegrity == 101){
    // air pump
    digitalWrite(pinAirPump, HIGH);
    if(currentAirPump == 1 || forceAirPump == 1){
      digitalWrite(pinAirPump, HIGH);
    } else {
      digitalWrite(pinAirPump, LOW);
    }

    // parachute
    if((currentParachute == 1 || forceParachute == 1) && parachuteStatus == 0){
      servoParachute.write(100);
      parachuteStatus = 1;
    }
  }

  // gyroscope read
  sensors_event_t accelerometer, gyroscope, temperature;
  gyrocope.getEvent(&accelerometer, &gyroscope, &temperature);

  // temperature and humidity dht11
  int currentTemperature = int(dht.readTemperature());
  int currentHumidity = int(dht.readHumidity());

  // debug serial print
  if(sensorsDebugMode){
    Serial.print(" ");
    Serial.print(currentTemperature);
    Serial.print(" ");
    Serial.print(currentHumidity);
    Serial.print("\n");
  }
  if(actuatorsDebugMode){
    Serial.print(" ");
    Serial.print(currentMotor);
    Serial.print(" ");
    Serial.print(currentAirPump);
    Serial.print(" ");
    Serial.print(currentParachute);
    Serial.print("\n");
  }

  sensorsData[1] = currentTemperature;
  sensorsData[2] = currentHumidity;
  sensorsData[3] = parachuteStatus;
  
  // antenna comunication
  if(radio.available()){
    // antenna data recive
    while(radio.available()){
      radio.read(&actuatorsData, sizeof(actuatorsData));
    }
    currentDataIntegrity = actuatorsData[0];
    currentMotor = actuatorsData[1];
    currentAirPump = actuatorsData[2];
    currentParachute = actuatorsData[3];

    // antenna data transmition
    radio.stopListening();
    radio.write(&sensorsData, sizeof(sensorsData));
    radio.startListening();
  }
}
