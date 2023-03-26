// clock on upper menu
function getTime(){
    let date = new Date();
    let hh = date.getHours();
    let mm = date.getMinutes();
    let ss = date.getSeconds();
    let session = "AM";

    if(hh === 0){
        hh = 12;
    }
    if(hh > 12){
        hh = hh - 12;
        session = "PM";
     }

     hh = (hh < 10) ? "0" + hh : hh;
     mm = (mm < 10) ? "0" + mm : mm;
     ss = (ss < 10) ? "0" + ss : ss;

     let time = hh + ":" + mm + ":" + ss + " " + session;

     return time;
}
setInterval(() => {
    document.querySelector(".upperMenu .central").innerText = getTime();
}, 1000);

// side menu mechanics
function sideMenu(side){
    switch(side){
        case "left":
            if(document.querySelector(".leftMenu").style.left == "-300px"){
                document.querySelector(".leftMenu").style.left = "0px";
                document.querySelector(".leftMenu .arrow img").style.transform = "rotate(-90deg)";
                if(document.querySelector(".rightMenu").style.right == "0px"){
                    sideMenu("right");
                }
            } else {
                document.querySelector(".leftMenu").style.left = "-300px";
                document.querySelector(".leftMenu .arrow img").style.transform = "rotate(90deg)";
            }
            break;
        case "right":
            if(document.querySelector(".rightMenu").style.right == "-300px"){
                document.querySelector(".rightMenu").style.right = "0px";
                document.querySelector(".rightMenu .arrow img").style.transform = "rotate(90deg)";
                if(document.querySelector(".leftMenu").style.left == "0px"){
                    sideMenu("left");
                }
            } else {
                document.querySelector(".rightMenu").style.right = "-300px";
                document.querySelector(".rightMenu .arrow img").style.transform = "rotate(-90deg)";
            }
            break;
    }
}

// data menu mechanics
let dataImgtextLabel = document.querySelector(".changeDataImg .currentDataImg a");
let img1Div = document.querySelector(".main .img1Div");
let img2Div = document.querySelector(".main .img2Div");
let img3Div = document.querySelector(".main .img3Div");
dataImgtextLabel.innerText = "Actuator out";
let currentSection = 0;
let sectionNum = 3;
function switchDataImg(direction) {
    if(direction == "forward"){
        currentSection++;
    } else if (direction == "backward"){
        currentSection--;
    }
    if(currentSection >= sectionNum){
        currentSection = 0;
    } else if(currentSection < 0){
        currentSection = sectionNum -1;
    }
    switch (currentSection) {
        case 0:
            dataImgtextLabel.innerText = "Actuator out";
            img1Div.style.display = "flex";
            img2Div.style.display = "none";
            img3Div.style.display = "none";
            break;
        case 1:
            dataImgtextLabel.innerText = "Sensor out";
            img1Div.style.display = "none";
            img2Div.style.display = "flex";
            img3Div.style.display = "none";
            break;
        case 2:
            dataImgtextLabel.innerText = "Data dash";
            img1Div.style.display = "none";
            img2Div.style.display = "none";
            img3Div.style.display = "flex";
    }
}
document.querySelector(".changeDataImg .arrowLeft").addEventListener("click", ()=>{
    switchDataImg("backward");

});
document.querySelector(".changeDataImg .arrowRight").addEventListener("click", ()=>{
    switchDataImg("forward");
});

// arrow information display
function infoArrow(x, y, height, width, textHeight, textWidth, thikness, ball, reverseY, reverseX, color, text, parent, img){
    imgOffsetX = img.offsetWidth;
    imgOffsetY = img.offsetHeight;
    parentOffsetX = parent.offsetWidth;
    parentOffsetY = parent.offsetHeight;

    let arrowDiv = document.createElement("div");
    arrowDiv.style.opacity = "0";
    arrowDiv.style.height = (height + thikness) + "px";
    arrowDiv.style.width = (width + thikness) + "px";
    arrowDiv.style.display = "flex";
    arrowDiv.style.flexDirection = "row";
    arrowDiv.style.justifyContent = "flex-end";
    arrowDiv.style.position = "absolute";
    arrowDiv.style.top = ((parentOffsetY - imgOffsetY)/2 + y - height) + "px";
    arrowDiv.style.left = ((parentOffsetX - imgOffsetX)/2 + x) + "px";
    arrowDiv.style.transition = "0.25s";

    let arrowDivSubV = document.createElement("div");
    arrowDivSubV.style.background = color;
    arrowDivSubV.style.height = thikness + "px";
    arrowDivSubV.style.width = Math.sqrt(height*height + height*height) + "px";
    arrowDivSubV.style.transformOrigin = "top right";
    arrowDivSubV.style.transform = "translateX(-" + (width - height) + "px) rotate(-45deg)";
    arrowDivSubV.style.position = "absolute";

    let arrowDivSubH = document.createElement("div");
    arrowDivSubH.style.background = color;
    arrowDivSubH.style.height = thikness + "px";
    arrowDivSubH.style.width = (width - height) + "px";
    arrowDivSubH.style.position = "relative";

    let arrowDivSubC = document.createElement("div");
    arrowDivSubC.style.background = color;
    arrowDivSubC.style.height = ball + "px";
    arrowDivSubC.style.width = ball + "px";
    arrowDivSubC.style.transform = "translateY(" + (height - thikness/2) + "px) translateX(-" + (width - thikness) + "px)";
    arrowDivSubC.style.borderRadius = "8px";
    arrowDivSubC.style.position = "absolute";

    let arrowDivSubB = document.createElement("div");
    arrowDivSubB.style.background = color;
    arrowDivSubB.style.height = textHeight + "px";
    arrowDivSubB.style.width = thikness + "px";
    arrowDivSubB.style.transform = "translateX(" + (width - height) + "px) translateY(-" + ((80 - thikness) / 2) + "px)"

    let arrowDivSubT = document.createElement("div");
    arrowDivSubT.style.background = "transparent";
    arrowDivSubT.style.transform = "translateY(-" + ((80 - thikness) / 2) + "px) translateX(" + (textWidth + 6) + "px)";
    arrowDivSubT.style.height = textHeight + "px";
    arrowDivSubT.style.width = textWidth + "px";
    arrowDivSubT.style.color = "white";
    arrowDivSubT.style.fontFamily = "'Roboto Mono', monospace";
    arrowDivSubT.style.fontSize = "11px";
    arrowDivSubT.style.fontWeight = "300";
    arrowDivSubT.style.position = "absolute";
    arrowDivSubT.innerHTML = text;

    if(reverseY && reverseX){
        arrowDiv.style.transform = "scale(-1, -1)";
        arrowDivSubT.style.transform += "scale(-1, -1)";
        arrowDivSubT.style.textAlign = "right";
    } else if(reverseX){
        arrowDiv.style.transform = "scale(1, -1)";
        arrowDivSubT.style.transform += "scale(1, -1)";
        arrowDivSubT.style.textAlign = "left";
    } else if(reverseY){
        arrowDiv.style.transform = "scale(-1, 1)";
        arrowDivSubT.style.transform += "scale(-1, 1)";
        arrowDivSubT.style.textAlign = "right";
    }

    arrowDiv.appendChild(arrowDivSubB);
    arrowDiv.appendChild(arrowDivSubT);
    arrowDiv.appendChild(arrowDivSubV);
    arrowDiv.appendChild(arrowDivSubH);
    arrowDiv.appendChild(arrowDivSubC);
    parent.appendChild(arrowDiv);

    setInterval(()=>{arrowDiv.style.opacity = "1";}, 600);

    return arrowDivSubT;
}
// air pump
let pumpTextString = "Current state: on<br>Current load: 10mL/s<br>Temperature: 45C<br>In voltage: 0.3V<br>Warning: none";
let pumpTextElement = infoArrow(710, 190, 140, 240, 80, 180, 4, 12, false, false, "rgb(32, 33, 36)", pumpTextString, document.querySelector(".main .img1Div"), document.querySelector(".main .img1Div .img1"));
// battery
let batteryTextString = "Current state: ok<br>Current load: 0.16A<br>Temperature: 45C<br>Out voltage: 8.4V<br>Warning: low voltage";
let batteryTextElement = infoArrow(355, 90, 50, 160, 80, 150, 4, 12, true, false, "rgb(32, 33, 36)", batteryTextString, document.querySelector(".main .img1Div"), document.querySelector(".main .img1Div .img1"));
// arduino mega
let arduinoTextString = "Firmware: v1.0<br>Temperature: 45C<br>In voltage: 8.0V<br>Warning: none";
let arduinoTextElement = infoArrow(240, 450, 240, 300, 62, 150, 4, 12, true, true, "rgb(32, 33, 36)", arduinoTextString, document.querySelector(".main .img1Div"), document.querySelector(".main .img1Div .img1"));

// data page information display
let batteryBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row1 .batteryBlock");
let tempHumBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row1 .tempHumBlock");
let airParBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row2 .airParBlock");
let accBarBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row2 .accBarBlock");
let gasValveBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row3 .gasValveBlock");
let gpsGyroBlock = document.querySelector(".main .img3Div .dataMain .leftBlock .row3 .gpsGyroBlock");
let gpsMapIframe = document.querySelector(".main .img3Div .dataMain .rightBlock .map .gmapsIframe");
let gyroscopeRollBlockTitle = document.querySelector(".main .img3Div .dataMain .rightBlock .gyroscope .roll .title");
let gyroscopePitchBlockTitle = document.querySelector(".main .img3Div .dataMain .rightBlock .gyroscope .pitch .title");
let gyroscopeYawBlockTitle = document.querySelector(".main .img3Div .dataMain .rightBlock .gyroscope .yaw .title");
function updateDataDash(element, value){
    switch (element) {
        case "battery": //value: [1[vOut,aOut,warning],2[vOut,aOut,warning]]
            batteryBlock.innerHTML = "<a style='font-weight: bold;'>Battery 1 (Arduino in):</a>";
            batteryBlock.innerHTML += "out voltage: " + value[0][0] + " V<br>";
            batteryBlock.innerHTML += "out current: " + value[0][1] + " A<br>";
            batteryBlock.innerHTML += "warning: " + value[0][2] + "<br>";
            batteryBlock.innerHTML += "<a style='font-weight: bold;'>Battery 2 (PCB in):</a>";
            batteryBlock.innerHTML += "out voltage: " + value[1][0] + " V<br>";
            batteryBlock.innerHTML += "out current: " + value[1][1] + " A<br>";
            batteryBlock.innerHTML += "warning: " + value[1][2] + "<br>";
            break;
        case "airpumpParachuteservo": //value: [p[vIn,warning],s[vIn,rotation,warning]]
            airParBlock.innerHTML = "<a style='font-weight: bold;'>Air pump:</a>";
            airParBlock.innerHTML += "in voltage: " + value[0][0] + " V<br>";
            airParBlock.innerHTML += "warning: " + value[0][1] + "<br>";
            airParBlock.innerHTML += "<a style='font-weight: bold;'>Parachute servo:</a>";
            airParBlock.innerHTML += "in voltage: " + value[1][0] + " V<br>";
            airParBlock.innerHTML += "rotation: " + value[1][1] + " deg<br>";
            airParBlock.innerHTML += "status: " + ((value[1][1] > 0) ? "open" : "closed") + "<br>";
            airParBlock.innerHTML += "warning: " + value[1][2] + "<br>";
            break;
        case "gasValve": //value: [vIn,steps,rotation,status,temp,warning]
            gasValveBlock.innerHTML = "<a style='font-weight: bold;'>Gas valve:</a>";
            gasValveBlock.innerHTML += "in voltage: " + value[0] + " V<br>";
            gasValveBlock.innerHTML += "steps: " + value[1] + "<br>";
            gasValveBlock.innerHTML += "rotation: " + value[2] + " deg<br>";
            gasValveBlock.innerHTML += "status: " + value[3] + " %<br>";
            gasValveBlock.innerHTML += "temperature: " + value[4] + " C<br>";
            gasValveBlock.innerHTML += "warning: " + value[5] + "<br>";
            break;
        case "temperatureHumidity": //value: [t[current,max,min],h[current,max,min]]
            tempHumBlock.innerHTML = "<a style='font-weight: bold;'>Temperature:</a>";
            tempHumBlock.innerHTML += "current temperature: " + value[0][0] + " C<br>";
            tempHumBlock.innerHTML += "max temperature: " + value[0][1] + " C<br>";
            tempHumBlock.innerHTML += "lowest temperature: " + value[0][2] + " C<br>";
            tempHumBlock.innerHTML += "<a style='font-weight: bold;'>Humidity:</a>";
            tempHumBlock.innerHTML += "current humidity: " + value[1][0] + " %<br>";
            tempHumBlock.innerHTML += "max humidity: " + value[1][1] + " %<br>";
            tempHumBlock.innerHTML += "lowest humidity: " + value[1][2] + " %<br>";
            break;
        case "accelerometerBarometer": //value: [a[x,y,z],b[pressure,altitudeS,altitudeG]]
            accBarBlock.innerHTML = "<a style='font-weight: bold;'>Accelerometer:</a>";
            accBarBlock.innerHTML += "x: " + value[0][0] + " m/s^2<br>";
            accBarBlock.innerHTML += "y: " + value[0][1] + " m/s^2<br>";
            accBarBlock.innerHTML += "z: " + value[0][2] + " m/s^2<br>";
            accBarBlock.innerHTML += "<a style='font-weight: bold;'>Barometer:</a>";
            accBarBlock.innerHTML += "pressure: " + value[1][0] + " Pa<br>";
            accBarBlock.innerHTML += "altitude (sea): " + value[1][1] + " m<br>";
            accBarBlock.innerHTML += "altitude (ground): " + value[1][2] + " m<br>";
            break;
        case "gpsGyroscope": //value: [g[satellites,dd1,dd2,dms1_1,dms1_2,dms2_1,dms2_2],g[x,y,z]]
            gpsGyroBlock.innerHTML = "<a style='font-weight: bold;'>GPS: (" + value[0][0] + " satellites)</a>";
            gpsGyroBlock.innerHTML += "DD: " + value[0][1] + " deg, " + value[0][2] + " deg<br>";
            gpsGyroBlock.innerHTML += "DMS (" + value[0][3] + "): " + value[0][4] + "\"<br>";
            gpsGyroBlock.innerHTML += "DMS (" + value[0][5] + "): " + value[0][6] + "\"<br>";
            gyroscopeRollBlockTitle.innerHTML = "<b>Roll:</b> " + value[1][0] + " deg";
            gyroscopePitchBlockTitle.innerHTML = "<b>Pitch:</b> " + value[1][1] + " deg";
            gyroscopeYawBlockTitle.innerHTML = "<b>Yaw:</b> " + value[1][2] + " deg";
            gpsGyroBlock.innerHTML += "<a style='font-weight: bold;'>Gyroscope:</a>";
            gpsGyroBlock.innerHTML += "x: " + value[1][0] + " deg<br>";
            gpsGyroBlock.innerHTML += "y: " + value[1][1] + " deg<br>";
            gpsGyroBlock.innerHTML += "z: " + value[1][2] + " deg<br>";
            let center = value[0][1] + "," + value[0][2];
            break;
        default:

    }
}
// default data
updateDataDash("battery", [[0,0,"none"],[0,0,"none"]]);
updateDataDash("temperatureHumidity", [[0,0,0],[0,0,0]]);
updateDataDash("airpumpParachuteservo", [[0,"none"],[0,0,"none"]]);
updateDataDash("accelerometerBarometer", [[0,0,0],[0,0,0]]);
updateDataDash("gasValve", [0, 0, 0, 0, 0, "none"]);
updateDataDash("gpsGyroscope", [[0,0,"/","0° 0' 0”","/","0° 0' 0”"],[0,0,0]]);

// logs (side menu)
let logMainDiv = document.querySelector(".rightMenu .logs");

function newLog(message){
    let logElementA = document.createElement("a");
    logElementA.innerText = getTime() + ": " + message;

    logMainDiv.appendChild(logElementA);
}
// autoscroll
setInterval(() => {logMainDiv.scrollTo(0, logMainDiv.scrollHeight);}, 100);

// websocket
let connectionStateOut = document.getElementById("connectionStateOut");

const websocket = new WebSocket("ws://" + document.location.host + "/dashboardWs");

websocket.onopen = (event) => {
    connectionStateOut.style.background = "red";
}
websocket.onclose = (event) => {
    connectionStateOut.style.background = "red";
    newLog("Websocket connection closed");
}
websocket.onerror = (event) => {
    connectionStateOut.style.background = "red";
    newLog("Error in websocket connection");
}

let websocketString = "";
let oldWebsocketString = "";
let websocketStringParsed = "";
let barometerData = [];
let newMessagesInDb = false;
let firstData = true;
websocket.onmessage = (event) => {
    connectionStateOut.style.background = "yellow";

    websocketString = event.data;

    if(oldWebsocketString == websocketString){
        if(newMessagesInDb){
            newLog("Old message from websocket");
        }
        newMessagesInDb = false;
        setTimeout(()=>{connectionStateOut.style.background = "red";}, 60);
    } else {
        if(!newMessagesInDb){
            newLog("New message from websocket");
        }
        newMessagesInDb = true;
        setTimeout(()=>{connectionStateOut.style.background = "green";}, 60);
    }

    websocketStringParsed = websocketString.split(",");

    if(firstData){
        gpsMapIframe.src = "https://maps.google.com/maps?q=" + websocketStringParsed[12].split("%")[1] + ",%20" + websocketStringParsed[12].split("%")[2] + "&amp;t=k&amp;z=13&amp;ie=UTF8&amp;iwloc=&amp;output=embed&output=embed";
        firstData = false;
    } else {
        barometerData = [websocketStringParsed[11]];
        barometerData.push(websocketStringParsed[12].split("%")[7]);
        barometerData.push(213.4);

        updateDataDash("battery", [websocketStringParsed[3].split("%"),websocketStringParsed[4].split("%")]);
        updateDataDash("temperatureHumidity", [websocketStringParsed[8].split("%"),websocketStringParsed[9].split("%")]);
        updateDataDash("airpumpParachuteservo", [websocketStringParsed[5].split("%"),websocketStringParsed[6].split("%")]);
        updateDataDash("accelerometerBarometer", [websocketStringParsed[10].split("%"),barometerData]);
        updateDataDash("gasValve", websocketStringParsed[7].split("%"));
        updateDataDash("gpsGyroscope", [websocketStringParsed[12].split("%"),websocketStringParsed[13].split("%")]);
    }

    oldWebsocketString = websocketString;
}
