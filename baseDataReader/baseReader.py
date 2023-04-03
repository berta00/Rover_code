import mysql.connector
import serial
import os

dataArr = []
index = 0

#mydb = mysql.connector.connect(
#  host = os.environ['MYSQL_IP'],
#  user = os.environ['MYSQL_USER'],
#  password = os.environ['MYSQL_PASS'],
#  database = os.environ['MYSQL_DB_DATA']
#)

ser = serial.Serial('/dev/cu.usbmodem21201', 9600)
#mycursor = mydb.cursor()

#query1 = "INSERT INTO " + os.environ['MYSQL_TB_SEN_OUT'] + " (sensorFailures, temperature, humidity, accelerometer, barometer, gps, gyroscope) VALUES (%s, %s, %s, %s, %s, %s, %s)"
#query2 = "INSERT INTO " + os.environ['MYSQL_TB_ACT_OUT'] + " (status, warning, battery1, battery2, airPump, parachuteServo, gasValve) VALUES (%s, %s, %s, %s, %s, %s, %s)"

while True:
    readedLine = ser.readLine()
    parsedLine = readLine.split(",")
    print(dataArr)

#    for a in parsedLine:
#        dataArr.push(a)
#    if index % 2 == 0:
#        print(dataArr)
#
#        val1 = (dataArr[0], dataArr[1], dataArr[2], dataArr[3], dataArr[4], dataArr[5], dataArr[6])
#        val2 = ("ok", "none", dataArr[7], dataArr[8], dataArr[9], dataArr[10], dataArr[11])
#
#        mycursor.execute(query1, val1)
#        mycursor.execute(query2, val2)
#
#        mydb.commit()
#        dataArr.clear()
#        index = 0
#    index += 1
#
