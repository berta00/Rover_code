// Base

#include <SPI.h>
#include <RF24.h>

// antenna variables and object
RF24 radio(7, 8);
byte node_A_address[6] = "NodeA";
byte node_B_address[6] = "NodeB";

// global sensor and actuators data (101 is to ceck datas integrity)
// motor, airPump, parachute
int actuatorsData[12] = {101, 0, 0, 0};
// temperature, humidity, gps, barometer, airPump, accelerometer, stepper,  parachute, stepperV, servoV, airPumpV
char sensorDataString[64] = "";

void setup() {
  // antenna setup
  radio.begin();
  radio.setPALevel(RF24_PA_MAX); //RF24_PA_MAX  RF24_PA_LOW
  radio.openWritingPipe(node_B_address);
  radio.openReadingPipe(1, node_A_address);
  radio.startListening();

  // serial setup
  Serial.begin(9600);
}

void loop() {

  // antenna comunication
  radio.stopListening();

  unsigned long start_time = micros();
  if (!radio.write(&actuatorsData, sizeof(actuatorsData))) {
    Serial.println(F("failed"));
  }

  radio.startListening();

  unsigned long started_waiting_at = micros();
  boolean timeout = false;

  while (!radio.available()) {
    if (micros() - started_waiting_at > 200000) {
      timeout = true;
      break;
    }
  }

  if (timeout){
    Serial.println(F("Failed, response timed out."));
  } else {
    radio.read(&sensorDataString, sizeof(sensorDataString));
    unsigned long end_time = micros();
    Serial.println(sensorDataString);
  }
  delay(20);
}
