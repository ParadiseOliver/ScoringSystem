"use strict"



$( document ).ready(function() {
    console.log("ready")
    const eventID = getEventIDFromURL()
    getEvent(eventID);
});

function getEventIDFromURL () {
    const pathList = window.location.pathname.split("/")
    return pathList[pathList.length - 1]
}

function getEvent(eventID) {
    var block = "";
    var data = $.getJSON("http://127.0.0.1:8080/api/v1/events/" + eventID, function( data ){
        console.log(data.name)
        block +='<div id="' + data.id + '"><h5>' + data.name + '</h5><h6>' + data.start_date + ' to ' + data.end_date + '</h6></div>'
        $("#event").html(block); 
    });
} 