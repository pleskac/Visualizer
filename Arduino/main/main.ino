

int firstSensor = 0;    // first analog sensor
int secondSensor = 0;   // second analog sensor
int thirdSensor = 0;    // digital sensor
int inByte = 0;    // incoming serial byte
int testLED = 13;

void setup()
{
  pinMode(testLED, OUTPUT);  
  Serial.begin(19200);
  while (!Serial) {
    ; // wait for serial port to connect. Needed for Leonardo only
  }
  pinMode(2, INPUT);   // digital sensor is on digital pin 2
  establishContact();  // send a byte to establish contact until receiver responds 
}

void loop()
{
  digitalWrite(testLED, HIGH);
  
  // if we get a valid byte, read analog ins:
  if (Serial.available() > 0) {
    digitalWrite(testLED, LOW); //CAUSES THE LED TO FLICKER WITH INCOMING DATA
    
    // get incoming byte:
    inByte = Serial.read();
    
    // print the data we just recieved
    Serial.print(inByte);         
  }
}

void establishContact() {
  while (Serial.available() <= 0) {
    Serial.println("0,0,0");   // send an initial string
    delay(300);
  }
}
