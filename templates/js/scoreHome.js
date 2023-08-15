"use strict"

$( document ).ready(function() {
    getEvents();
});

function getEvents() {
    var block = "";
    var data = $.getJSON("http://127.0.0.1:8080/api/v1/events", function( data ){
        $.each(data.events, function(key, item) {
            block +='<div id="' + item.id + '"><a href="/pages/score/' + item.id + '"><h5>*Event name*</h5><h6>' + item.start_date + ' to ' + item.end_date + '</h6></div></a>"'
        });
        $("#eventList").html(block);
    });  
}
