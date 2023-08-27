"use strict"

import { formatDate } from "./utils.js";

$( document ).ready(function() {
    getEvents();
});

function getEvents() {
    var block = "";
    var cnt = 0;
    var data = $.getJSON("http://127.0.0.1:8080/api/v1/events", function( data ){
        $.each(data.events, function(key, item) {
            if (cnt % 3 == 0) {
                if (cnt != 0) {
                    block += `</div><br/><br/>`
                }
                block += `<div class="row">`
            }
            block += `<div class="col-sm-4"><a href="/pages/score/` + item.id + `"><div id="` + item.id + `" class="card text-center">
            <img class="card-img-top" src="` + item.image_url + `" alt="Card image cap" style="height:200px;width:auto">
            <div class="card-body">
              <h5 class="card-title">` + item.event + `</h5>
              <p class="card-text">` + formatDate(item.start_date) + ' to ' + formatDate(item.end_date) + `</p>
            </div>
          </div></div></a>`;
          cnt += 1;  
        });
        $("#eventList").html(block);
    });  
}