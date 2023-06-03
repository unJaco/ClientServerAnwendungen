# Projektarbeit Ride-Hailing / Uber

Diese Projektarbeit beinhaltet ein Front- und Backend um einen simplen Ride-Hailing Dienst darzustellen.
Die Funktionalität ist zwar teilweise eingeschränkt bildet jedoch alle grundlegenden Funktionen zum User-Management und buchen einer Fahrt ab.

## Before SetUp

Installieren von npm und GOLANG (https://go.dev/doc/install). 

## SetUp Backend

Nach dem Clonen muss im geclonten Ordner die File ".env" angelegt werden. Die File muss folgenden Inhalt (versch. Attribute) enthalten:

user_name = "YOUR_MYSSQL_USERNAME",         // example: "root"
password = "YOUR_MYSSQL_PASSWORD",          // example: "root"
address = "YOUR_MYSSQL_SERVER_ADDRESS",     // example:"127.0.0.1:3306"
db_name = "YOUR_MYSSQL_DB_NAME"             // example: "myUber"


## SetUp Frontend
To start the frontend navigat to the frontend folder and run npm start  
you maybe need to install node.js   
you maybe need to install the router-dom package (npm install react-router-dom).
Wenn der Befehl "react-scripts" nicht ausgeführt werden kann, muss der Befehl "npm i --legacy-peer-deps" eingegeben werden. Danach kann erneut "npm start" eingegeben werden. 


## Start Application

1. starten des MySqlServer

Entweder :

2. ausführen von start.bsh

oder :

2. cd backend
3. go run main.go

Dabei werden alle go dependencies installiert. 
Danach sollte auf localhost:8080 das Backend laufen




