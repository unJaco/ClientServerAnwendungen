# Projektarbeit Ride-Hailing / Uber

Diese Projektarbeit beinhaltet ein Front- und Backend um einen simplen Ride-Hailing Dienst darzustellen.  
Die Funktionalität ist zwar teilweise eingeschränkt bildet jedoch alle grundlegenden Funktionen zum User-Management und buchen einer Fahrt ab.

## Before SetUp

Installieren von npm und GOLANG (https://go.dev/doc/install).   

## SetUp Backend

Nach dem Clonen muss im geclonten Ordner die File ".env" angelegt werden. Die File muss folgenden Inhalt (versch. Attribute) enthalten:

user_name = "YOUR_MYSSQL_USERNAME",          // example: "root"  
password = "YOUR_MYSSQL_PASSWORD",          // example: "root"  
address = "YOUR_MYSSQL_SERVER_ADDRESS",     // example:"127.0.0.1:3306"  
db_name = "YOUR_MYSSQL_DB_NAME"             // example: "myUber"  


## SetUp Frontend
Um das Frontend zu starten muss man in den frontend folder navigieren und npm start ausführen.  
Falls einige Dependencies nicht vorhanden sind muss "npm i --legacy-peer-deps" ausgeführt werden.  
Danach kann erneut "npm start" eingegeben werden.   


## Start Application

1. starten des MySqlServer

Entweder :

2. ausführen von start.bsh

oder :

2. cd backend
3. go run main.go

Dabei werden alle go dependencies installiert. 
Danach sollte auf localhost:8080 das Backend laufen




