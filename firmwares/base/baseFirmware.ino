// Base

#include <SPI.h>
#include <nRF24L01.h>
#include <RF24.h>

// radio object and vars (pong setup)
RF24 radio(7, 8); // CE, CSN
byte node_A_address[6] = "NodeA";
byte node_B_address[6] = "NodeB";

// buttons pin
int pinManualAirPump = 2;
int pinParachuteButton = 3;

// buttons vars
int currentParachuteButton = 0;
int currentManualAirPumpButton = 0;

// global sensor and actuators data (101 is to ceck datas integrity)
int actuatorsData[] = {101, 0, 1, 0}; // motor, airPump, parachute
int sensorsData[32] = {0, 0, 0}; // temperature, humidity, parachute

// current sensors vars
int currentDataIntegrity = 0;
int currentTemperature = 0;
int currentHumidity = 0;
int currentParachuteStatus = 0;

// setup vars
bool output = true;

void setup() {
  // serial setup
  Serial.begin(9600);
  // radio setup
  radio.begin();
  radio.setPALevel(RF24_PA_LOW);
  radio.openWritingPipe(node_B_address);
  radio.openReadingPipe(1, node_A_address);
  radio.startListening();
  // buttons setup
  pinMode(pinParachuteButton, INPUT_PULLUP);
  pinMode(pinManualAirPump, INPUT_PULLUP);
}

void loop() {
  // read button data
  currentManualAirPumpButton = digitalRead(pinManualAirPump);
  currentParachuteButton = digitalRead(pinParachuteButton);
  
  if(currentManualAirPumpButton == LOW){
    actuatorsData[2] = 1;
  } else{
    actuatorsData[2] = 0;
  }
  if(currentParachuteButton == LOW){
    actuatorsData[3] = 1;
  } else{
    actuatorsData[3] = 0;
  }

  if(output){
    Serial.print(currentTemperature);
    Serial.print(" ");
    Serial.print(currentHumidity);
    //Serial.print(" ");
    //Serial.print(currentParachuteStatus);
    Serial.print("\n");
  }

  // antenna connection data trasmition
  radio.stopListening();
  int currentTry = 0;
  if(!radio.write(&actuatorsData, sizeof(actuatorsData))){
    currentTry ++;
  }

  // antenna connection data recive
  int sensorsData[] = {0, 0, 0, 0};
  radio.startListening();
  while(!radio.available()){
  }
  radio.read(&sensorsData, sizeof(sensorsData));

  currentDataIntegrity = sensorsData[0];
  currentTemperature = sensorsData[1];
  currentHumidity = sensorsData[2];
  currentParachuteStatus = sensorsData[3];

  delay(400);
}
