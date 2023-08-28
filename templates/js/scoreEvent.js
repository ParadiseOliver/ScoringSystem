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
    var cnt = 0;
    var data = $.getJSON("http://127.0.0.1:8080/api/v1/events/" + eventID, function( data ){
        $.each(JSON.parse(data.disciplines), function(key, item) {
            if (cnt % 3 == 0) {
                if (cnt != 0) {
                    block += `</div><br/><br/>`
                }
                block += `<div class="row">`
            }
            block += `<div class="col-sm-4"><div id="` + item.id + `" class="card text-center">
                <div class="card-body">
                    <h5 class="card-title">` + item.discipline + `</h5>
                    <p class="card-text">FIG Senior men</p>
                    <p class="card-text">FIG Senior men</p>
                </div></div></div>`;
          cnt += 1;
        })
        $("#discipline").html(block);
    });
} 