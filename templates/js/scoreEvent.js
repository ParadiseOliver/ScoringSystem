"use strict"

import { formatDate } from "./utils.js";

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
        block +='<div id="' + data.id + '"><h5>' + data.event + '</h5><h6>' + formatDate(data.start_date) + ' to ' + formatDate(data.end_date) + '</h6></div>'
        $("#event").html(block); 
    });
} 