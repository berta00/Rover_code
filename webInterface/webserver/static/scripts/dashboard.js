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

// information display
function infoArrow(x, y, height, width, spessore, pallino, colore, parent, img){
    imgOffsetX = img.offsetWidth;
    imgOffsetY = img.offsetHeight;

    let arrowDiv = document.createElement("div");
    arrowDiv.style.height = height + "px";
    arrowDiv.style.width = width + "px";
    arrowDiv.style.display = "flex";
    arrowDiv.style.flexDirection = "row";
    arrowDiv.style.position = "absolute";
    arrowDiv.style.top = (imgOffsetY/2 - height + y) + "px";
    arrowDiv.style.left = (imgOffsetX/2 - width + x) + "px";

    let arrowDivSubV = document.createElement("div");
    arrowDivSubV.style.background = colore;
    arrowDivSubV.style.height = spessore + "px";
    arrowDivSubV.style.width = Math.sqrt(height*height + height*height) + "px";
    arrowDivSubV.style.transform = "translateY(" + (height/2 - spessore/4) + "px) rotate(-45deg)";

    let arrowDivSubH = document.createElement("div");
    arrowDivSubH.style.background = colore;
    arrowDivSubH.style.height = spessore + "px";
    arrowDivSubH.style.width = (width - height) + "px";

    let arrowDivSubC = document.createElement("div");
    arrowDivSubC.style.background = colore;
    arrowDivSubC.style.height = pallino + "px";
    arrowDivSubC.style.width = pallino + "px";
    arrowDivSubC.style.transform = "translateY(" + (height - 16/2 + pallino/4) + "px) translateX(-" + (16/2) + "px)";
    arrowDivSubC.style.borderRadius = "8px";
    arrowDivSubC.style.position = "absolute";

    arrowDiv.appendChild(arrowDivSubV);
    arrowDiv.appendChild(arrowDivSubH);
    arrowDiv.appendChild(arrowDivSubC);
    parent.appendChild(arrowDiv);
}
infoArrow(700, 80, 100, 200, 6, 16, "rgb(32, 33, 36)", document.querySelector(".main"), document.querySelector(".main .img1"));

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

let websocketString = "";
let oldWebsocketString = "";
let newMessagesInDb = false;
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

    oldWebsocketString = websocketString;
}
