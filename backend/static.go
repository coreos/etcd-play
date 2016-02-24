// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend

// updated at 2016-02-24 11:45:56.426345718 -0800 PST

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

func staticLocalHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	fmt.Fprintln(w, htmlSourceFileLocal)
	return nil
}

func staticRemoteHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	fmt.Fprintln(w, htmlSourceFileRemote)
	return nil
}

var htmlSourceFileLocal = `<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="author" content="Gyu-Ho Lee">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.1.1/js/tether.min.js"></script>
    <script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.2/css/bootstrap.min.css" integrity="sha384-y3tfxAZXuh4HwSYylfB+J125MxIs6mR5FOHamPBG064zB+AFeWH94NdvaCBm8qnd" crossorigin="anonymous">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.2/js/bootstrap.min.js" integrity="sha384-vZ2WRJMwsjRMW/8U7i6PWi6AlO1L79snBrmgiDpgIWJ82z8eA5lenwvxbMV1PAh7" crossorigin="anonymous"></script>
    <script src="https://code.highcharts.com/highcharts.js"></script>
    <script src="https://code.highcharts.com/modules/exporting.js"></script>
    <link type="text/css" rel="stylesheet" href="http://fonts.googleapis.com/css?family=Inconsolata">
    <link rel='shortcut icon' type='image/x-icon' href='https://storage.googleapis.com/play-etcd/favicon.ico' />
    <script>
    $(document).ready(function() {
        function split(val) {
            return val.split(/,\s*/);
        }

        function extractLast(term) {
            return split(term).pop();
        }

        function appendLog(msg) {
            msg = $("<div />").html(msg);
            $("#log_output").append(msg);
            $("#log_box").scrollTop($("#log_box")[0].scrollHeight);
        }

        function runWithTimeout(callback, delay, times) {
            var internalCallback = function(t) {
                return function() {
                    if (--t > 0) {
                        window.setTimeout(internalCallback, delay);
                        callback();
                    }
                }
            }(times);
            window.setTimeout(internalCallback, delay);
        };

        // TODO: ajax is not working with display none
        //
        // var popoverOption = {
        //     trigger: 'hover',
        //     delay: {
        //         show: 1,
        //         hide: 3000
        //     },
        //     html: true,
        //     content: function() {
        //         return $('#etcd1_Contents').html();
        //     }
        // }
        // $('#etcd1_State_circle').popover(popoverOption).click(function(e) {
        //     e.preventDefault();
        // });

        // $('#etcd1_State_circle').mouseover(function(e) {
        //     e.preventDefault();
        //     $('#etcd1_Contents').modal('show');
        // }).mouseleave(function(e) {
        //     e.preventDefault();
        //     setTimeout(function() {
        //         $('#etcd1_Contents').modal('hide');
        //     }, 3000);
        // });
        // $('#etcd2_State_circle').mouseover(function(e) {
        //     e.preventDefault();
        //     $('#etcd2_Contents').modal('show');
        // }).mouseleave(function(e) {
        //     e.preventDefault();
        //     setTimeout(function() {
        //         $('#etcd2_Contents').modal('hide');
        //     }, 3000);
        // });
        // $('#etcd3_State_circle').mouseover(function(e) {
        //     e.preventDefault();
        //     $('#etcd3_Contents').modal('show');
        // }).mouseleave(function(e) {
        //     e.preventDefault();
        //     setTimeout(function() {
        //         $('#etcd3_Contents').modal('hide');
        //     }, 3000);
        // });
        // $('#etcd4_State_circle').mouseover(function(e) {
        //     e.preventDefault();
        //     $('#etcd4_Contents').modal('show');
        // }).mouseleave(function(e) {
        //     e.preventDefault();
        //     setTimeout(function() {
        //         $('#etcd4_Contents').modal('hide');
        //     }, 3000);
        // });
        // $('#etcd5_State_circle').mouseover(function(e) {
        //     e.preventDefault();
        //     $('#etcd5_Contents').modal('show');
        // }).mouseleave(function(e) {
        //     e.preventDefault();
        //     setTimeout(function() {
        //         $('#etcd5_Contents').modal('hide');
        //     }, 3000);
        // });

        $('#put_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px;";
            // document.getElementById('result').style = "display: none;"
            document.getElementById('result').style = ""
        });
        $('#get_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px; display: none;";
            document.getElementById('result').style = ""
        });
        $('#delete_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px; display: none;";
            document.getElementById('result').style = ""
        });

        $('#kill_1').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_1",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_1').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_1",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_2').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_2",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_2').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_2",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_3').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_3",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_3').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_3",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_4').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_4",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_4').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_4",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_5').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_5",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_5').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_5",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });

        var receiveStreams = function() {
            $.ajax({
                url: '/stream',
                async: true,
                dataType: "json",
                success: function(dataObj) {
                    if (dataObj.Size > 0) {
                        appendLog(dataObj.Logs);
                    }
                }
            });
        }

        var queue_StorageKeysTotal = new Array;
        var receiveServerStatus = function() {
            $.ajax({
                type: "GET",
                url: '/server_status',
                async: true,
                dataType: "json",
                success: function(dataObj) {
                    queue_StorageKeysTotal.push(dataObj);

                    document.getElementById('active_user_number').innerHTML = "(active user: " + dataObj.ActiveUserNumber + ")";
                    document.getElementById('active_user_list').innerHTML = dataObj.ActiveUserList;

                    document.getElementById('etcd1_ID').innerHTML = "ID: <b>" + dataObj.Etcd1_ID + "</b>";
                    document.getElementById('etcd1_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd1_Endpoint + "</b>";
                    document.getElementById('etcd1_State').innerHTML = "State: <b>" + dataObj.Etcd1_State + "</b>";
                    document.getElementById('etcd1_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd1_NumberOfKeys + "</b>";
                    document.getElementById('etcd1_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd1_Hash + "</b>";
                    document.getElementById('etcd1_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd1_Hash + ")";
                    if (dataObj.Etcd1_State == "Leader") {
                        document.getElementById('etcd1_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd1_State == "Follower") {
                        document.getElementById('etcd1_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd1_State == "unreachable") {
                        document.getElementById('etcd1_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd1_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd2_ID').innerHTML = "ID: <b>" + dataObj.Etcd2_ID + "</b>";
                    document.getElementById('etcd2_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd2_Endpoint + "</b>";
                    document.getElementById('etcd2_State').innerHTML = "State: <b>" + dataObj.Etcd2_State + "</b>";
                    document.getElementById('etcd2_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd2_NumberOfKeys + "</b>";
                    document.getElementById('etcd2_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd2_Hash + "</b>";
                    document.getElementById('etcd2_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd2_Hash + ")";
                    if (dataObj.Etcd2_State == "Leader") {
                        document.getElementById('etcd2_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd2_State == "Follower") {
                        document.getElementById('etcd2_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd2_State == "unreachable") {
                        document.getElementById('etcd2_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd2_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd3_ID').innerHTML = "ID: <b>" + dataObj.Etcd3_ID + "</b>";
                    document.getElementById('etcd3_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd3_Endpoint + "</b>";
                    document.getElementById('etcd3_State').innerHTML = "State: <b>" + dataObj.Etcd3_State + "</b>";
                    document.getElementById('etcd3_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd3_NumberOfKeys + "</b>";
                    document.getElementById('etcd3_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd3_Hash + "</b>";
                    document.getElementById('etcd3_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd3_Hash + ")";
                    if (dataObj.Etcd3_State == "Leader") {
                        document.getElementById('etcd3_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd3_State == "Follower") {
                        document.getElementById('etcd3_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd3_State == "unreachable") {
                        document.getElementById('etcd3_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd3_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd4_ID').innerHTML = "ID: <b>" + dataObj.Etcd4_ID + "</b>";
                    document.getElementById('etcd4_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd4_Endpoint + "</b>";
                    document.getElementById('etcd4_State').innerHTML = "State: <b>" + dataObj.Etcd4_State + "</b>";
                    document.getElementById('etcd4_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd4_NumberOfKeys + "</b>";
                    document.getElementById('etcd4_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd4_Hash + "</b>";
                    document.getElementById('etcd4_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd4_Hash + ")";
                    if (dataObj.Etcd4_State == "Leader") {
                        document.getElementById('etcd4_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd4_State == "Follower") {
                        document.getElementById('etcd4_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd4_State == "unreachable") {
                        document.getElementById('etcd4_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd4_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd5_ID').innerHTML = "ID: <b>" + dataObj.Etcd5_ID + "</b>";
                    document.getElementById('etcd5_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd5_Endpoint + "</b>";
                    document.getElementById('etcd5_State').innerHTML = "State: <b>" + dataObj.Etcd5_State + "</b>";
                    document.getElementById('etcd5_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd5_NumberOfKeys + "</b>";
                    document.getElementById('etcd5_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd5_Hash + "</b>";
                    document.getElementById('etcd5_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd5_Hash + ")";
                    if (dataObj.Etcd5_State == "Leader") {
                        document.getElementById('etcd5_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd5_State == "Follower") {
                        document.getElementById('etcd5_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd5_State == "unreachable") {
                        document.getElementById('etcd5_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd5_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }
                }
            });
        }

        var wsConn;
        $('#start_cluster').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/start_cluster",
                async: true,
                dataType: "json",
                success: function(dataObj) {
                    appendLog(dataObj.Message);

                    var wsURL = "ws://" + window.location.hostname + dataObj.PlayWebPort + "/ws";
                    // var wsURL = "ws://" +  window.location.hostname + "/ws";
                    console.log("connecting " + wsURL);
                    if (window["WebSocket"]) {
                        wsConn = new WebSocket(wsURL);
                        wsConn.onopen = function() {
                            appendLog("Successfully connected to " + wsURL);
                        }
                        wsConn.onclose = function(ev) {
                            appendLog($("<div><b>connection closed</b></div>"));
                        }
                        wsConn.onmessage = function(ev) {
                            appendLog(ev.data);
                        }
                        wsConn.onerror = function(ev) {
                            appendLog("ERROR: " + ev.data);
                        }
                    } else {
                        appendLog($("<div><b>browser does not support WebSocket</b></div>"))
                    }
                }
            });
        });
        $('#start_cluster').click(function(e) {
            runWithTimeout(receiveStreams, 100, 1);
            setInterval(receiveStreams, 1000);
        });
        $('#start_cluster').click(function(e) {
            runWithTimeout(receiveStreams, 100, 1);
            setInterval(receiveServerStatus, 1000);
        });

        Highcharts.setOptions({
            global: {
                useUTC: false
            }
        });
        var chartRefreshRate = 2000;
        $('#chart_StorageKeysTotal').highcharts({
            chart: {
                type: 'spline',
                animation: Highcharts.svg, // don't animate in old IE
                marginRight: 10,
                events: {
                    load: function() {
                        var series1 = this.series[0];
                        var series2 = this.series[1];
                        var series3 = this.series[2];
                        var series4 = this.series[3];
                        var series5 = this.series[4];
                        setInterval(function() {
                            // var copiedQueue = queue_StorageKeysTotal.slice();
                            if (queue_StorageKeysTotal.length > 0) {
                                var dataObj = queue_StorageKeysTotal.shift();
                                var x = (new Date()).getTime();
                                series1.addPoint({
                                    x: x,
                                    y: dataObj.Etcd1_NumberOfKeys,
                                    name: dataObj.Etcd1_Name,
                                }, true, true);
                                series2.addPoint({
                                    x: x,
                                    y: dataObj.Etcd2_NumberOfKeys,
                                    name: dataObj.Etcd2_Name,
                                }, true, true);
                                series3.addPoint({
                                    x: x,
                                    y: dataObj.Etcd3_NumberOfKeys,
                                    name: dataObj.Etcd3_Name,
                                }, true, true);
                                series4.addPoint({
                                    x: x,
                                    y: dataObj.Etcd4_NumberOfKeys,
                                    name: dataObj.Etcd4_Name,
                                }, true, true);
                                series5.addPoint({
                                    x: x,
                                    y: dataObj.Etcd5_NumberOfKeys,
                                    name: dataObj.Etcd5_Name,
                                }, true, true);
                            }
                        }, chartRefreshRate);
                    }
                }
            },
            title: {
                text: 'Number of keys'
            },
            xAxis: {
                type: 'datetime',
                tickPixelInterval: 150
            },
            yAxis: {
                title: {
                    text: 'Count'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                formatter: function() {
                    return '<b>' + this.point.name + '</b><br/>' +
                        Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                        Highcharts.numberFormat(this.y, 2);
                }
            },
            legend: {
                enabled: false
            },
            exporting: {
                enabled: false
            },
            credits: {
                enabled: false
            },
            series: [{
                name: 'etcd1',
                color: "#FFCC00", // orange
                data: (function() {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;
                    for (i = -19; i <= 0; i += 1) {
                        data.push({
                            x: time + i * chartRefreshRate,
                            y: 0,
                            name: "unknown"
                        });
                    }
                    return data;
                }())
            }, {
                name: 'etcd2',
                color: "#04B4E3", // blue
                data: (function() {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;
                    for (i = -19; i <= 0; i += 1) {
                        data.push({
                            x: time + i * chartRefreshRate,
                            y: 0,
                            name: "unknown"
                        });
                    }
                    return data;
                }())
            }, {
                name: 'etcd3',
                color: "#00695C", // teal
                data: (function() {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;
                    for (i = -19; i <= 0; i += 1) {
                        data.push({
                            x: time + i * chartRefreshRate,
                            y: 0,
                            name: "unknown"
                        });
                    }
                    return data;
                }())
            }, {
                name: 'etcd4',
                color: "#ff9933", // yellow
                data: (function() {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;
                    for (i = -19; i <= 0; i += 1) {
                        data.push({
                            x: time + i * chartRefreshRate,
                            y: 0,
                            name: "unknown"
                        });
                    }
                    return data;
                }())
            }, {
                name: 'etcd5',
                color: "#9E9E9E", // gray
                data: (function() {
                    var data = [],
                        time = (new Date()).getTime(),
                        i;
                    for (i = -19; i <= 0; i += 1) {
                        data.push({
                            x: time + i * chartRefreshRate,
                            y: 0,
                            name: "unknown"
                        });
                    }
                    return data;
                }())
            }]
        });

        $("#key_input")
            .on("click", function(event, ui) {
                if (event.keyCode === $.ui.keyCode.TAB && $(this).autocomplete("instance").menu.active) {
                    event.preventDefault();
                }
            })
            .autocomplete({
                minLength: 0,
                source: function(request, response) {
                    $.ajax({
                        type: "GET",
                        url: "/key_history",
                        dataType: "json",
                        data: {
                            q: request.term
                        },
                        success: function(dataObj) {
                            response($.ui.autocomplete.filter(dataObj.Values, extractLast(request.term)));
                        }
                    });
                },
                focus: function() {
                    // prevent value inserted on focus
                    return false;
                }
            });

        $('#key_submit').click(function(e) {
            var selectedOperation = $('#operations .active')[0].innerText;
            var selectedNodeName = $('#node_names .active')[0].innerText;
            var key = $('input#key_input').val();
            var value = $('textarea#value_input').val();
            var dataToSend = $(this).serializeArray()
            dataToSend.push({
                name: "selected_operation",
                value: selectedOperation
            })
            dataToSend.push({
                name: "selected_node_name",
                value: selectedNodeName
            })
            dataToSend.push({
                name: "key_input",
                value: key
            })
            dataToSend.push({
                name: "value_input",
                value: value
            })
            e.preventDefault();
            $.ajax({
                type: "POST",
                url: "/key_value",
                data: dataToSend,
                success: function(dataObj) {
                    $.ajax({
                        type: "GET",
                        url: "/key_value",
                        async: true,
                        dataType: "json",
                        success: function(dataObj) {
                            // appendLog(dataObj.Message);
                            wsConn.send(dataObj.Message);

                            document.getElementById('result').innerHTML = dataObj.Result
                        }
                    });
                }
            });
        });

        $('#stress_submit').click(function(e) {
            var selectedNodeName = $('#node_names .active')[0].innerText;
            var dataToSend = $(this).serializeArray()
            dataToSend.push({
                name: "selected_node_name",
                value: selectedNodeName
            })
            e.preventDefault();
            $.ajax({
                type: "POST",
                url: "/stress",
                data: dataToSend,
                success: function(dataObj) {
                    $.ajax({
                        type: "GET",
                        url: "/stress",
                        async: true,
                        dataType: "html",
                        success: function(dataObj) {
                            appendLog(dataObj);
                            // wsConn.send(dataObj);
                        }
                    });
                }
            });
        });
    });
    </script>
    <style>
    svg a:hover {
        cursor: pointer;
    }
    /*
 * Base structure
 */
    /* Move down content because we have a fixed navbar that is 50px tall */
    
    body {
        padding-top: 50px;
    }
    /*
 * Global add-ons
 */
    
    .sub-header {
        padding-bottom: 10px;
        border-bottom: 1px solid #eee;
    }
    /*
 * Top navigation
 * Hide default border to remove 1px line.
 */
    
    .navbar-fixed-top {
        border: 0;
    }
    /*
 * Main content
 */
    
    .main {
        padding: 20px;
        margin-left: 35px;
        margin-right: 35px;
    }
    
    @media (min-width: 768px) {
        .main {
            padding-right: 20px;
            padding-left: 0px;
        }
    }
    
    .main .page-header {
        margin-top: 0;
    }
    /*
 * Placeholder dashboard ideas
 */
    
    .placeholders {
        margin-bottom: 30px;
        text-align: center;
    }
    
    .placeholders h4 {
        margin-bottom: 0;
    }
    
    .placeholder {
        margin-bottom: 20px;
    }
    
    .placeholder img {
        display: inline-block;
        border-radius: 50%;
    }
    
    .wrapper {
        text-align: center;
    }
    
    .bg-inverse {
        background-color: rgb(0, 153, 255);
    }
    
    .ui-menu-item {
        width: 100%;
        height: 100%;
        font-family: "Lucida Console", Courier, monospace;
        font-size: 12px;
    }
    
    div .modal-body {
        font-family: 'Inconsolata', Courier;
        font-size: 17px;
    }
    
    #result {
        /*font-family: "Lucida Console", Courier, monospace;*/
        /*font-size: 11px;*/
        font-family: "Inconsolata", Courier, monospace;
        font-size: 13px;
        font-style: normal;
        font-variant: normal;
    }
    
    #log_box {
        /*font-family: "Lucida Console", Courier, monospace;*/
        /*font-size: 11px;*/
        font-family: "Inconsolata", Courier, monospace;
        font-size: 13px;
        font-style: normal;
        font-variant: normal;
    }
    
    ul.ui-autocomplete.ui-menu {
        z-index: 1000;
    }
    </style>
    <title>Play etcd</title>
</head>

<body>
    <div class="modal fade" id="active_user_Contents" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-sm" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Active Users</h5>
                </div>
                <div class="modal-body">
                    <div id="active_user_list">...</div>
                </div>
            </div>
        </div>
    </div>
    <nav class="navbar navbar-dark bg-primary navbar-fixed-top">
        <a class="navbar-brand" href="https://github.com/coreos/etcd" target="_blank">Play etcd</a>
        <ul class="nav navbar-nav">
            <li class="nav-item active">
                <a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="https://github.com/coreos/etcd-play/issues/new">Report bug</a>
            </li>
            <li class="nav-item pull-xs-right">
                <div class="nav-link" id="active_user_number" data-toggle="modal" data-target="#active_user_Contents">(active user: 0)</div>
            </li>
        </ul>
    </nav>
    <div class="container-fluid">
        <div class="row">
            <div class="main">
                <!--                 <div id="etcd1_Contents" style="display: none;">
                    <div class="btn-group" role="group">
                        <button type="button" class="btn btn-danger btn-sm" id="kill_1">Kill</button>
                        <button type="button" class="btn btn-success btn-sm" id="restart_1">Restart</button>
                    </div>
                </div> -->
                <div class="modal fade" id="etcd1_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd1</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_1">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_1">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd1_ID">ID: n/a</div>
                                <div id="etcd1_Endpoint">Endpoint: n/a</div>
                                <div id="etcd1_State">State: n/a</div>
                                <div id="etcd1_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd1_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd2_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd2</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_2">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_2">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd2_ID">ID: n/a</div>
                                <div id="etcd2_Endpoint">Endpoint: n/a</div>
                                <div id="etcd2_State">State: n/a</div>
                                <div id="etcd2_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd2_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd3_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd3</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_3">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_3">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd3_ID">ID: n/a</div>
                                <div id="etcd3_Endpoint">Endpoint: n/a</div>
                                <div id="etcd3_State">State: n/a</div>
                                <div id="etcd3_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd3_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd4_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd4</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_4">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_4">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd4_ID">ID: n/a</div>
                                <div id="etcd4_Endpoint">Endpoint: n/a</div>
                                <div id="etcd4_State">State: n/a</div>
                                <div id="etcd4_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd4_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd5_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd5</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_5">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_5">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd5_ID">ID: n/a</div>
                                <div id="etcd5_Endpoint">Endpoint: n/a</div>
                                <div id="etcd5_State">State: n/a</div>
                                <div id="etcd5_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd5_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="show_graph" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-lg" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">Graphs</h5>
                            </div>
                            <div class="modal-body">
                                <div id="chart_StorageKeysTotal" style="margin: 0 auto; width: 800px; height: 350px;"></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-5">
                        <svg xmlns="http://www.w3.org/2000/svg" version="1.1" style="border: 0px solid black" width="500" height="400">
                            <circle id="ring" cx="270" cy="210" r="110" style="fill: #F5F5F5"></circle>
                            <g id="servers">
                                <g id="etcd1_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd1_Contents" title="etcd1">
                                    <!-- <g id="etcd1_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="popover" title="etcd1"> -->
                                    <text x="255" y="50">etcd1</text>
                                    <text x="250" y="67" font-size="11" id="etcd1_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="270" cy="100" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd2_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd2_Contents" title="etcd2">
                                    <text x="395" y="160">etcd2</text>
                                    <text x="400" y="177" font-size="11" id="etcd2_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="370" cy="180" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd3_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd3_Contents" title="etcd3">
                                    <text x="360" y="320">etcd3</text>
                                    <text x="342" y="337" font-size="11" id="etcd3_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="335" cy="295" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd4_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd4_Contents" title="etcd4">
                                    <text x="130" y="320">etcd4</text>
                                    <text x="120" y="337" font-size="11" id="etcd4_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="200" cy="295" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd5_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd5_Contents" title="etcd5">
                                    <text x="95" y="173">etcd5</text>
                                    <text x="87" y="150" font-size="11" id="etcd5_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="165" cy="180" r="25"></circle>
                                    </a>
                                </g>
                            </g>
                        </svg>
                    </div>
                    <div class="col-md-7">
                        <br>
                        <br>
                        <div class="btn-toolbar">
                            <div class="btn-group" data-toggle="buttons">
                                <input type="submit" class="btn btn-sm btn-secondary" id="start_cluster" value="Start">
                            </div>
                            <div class="btn-group" data-toggle="buttons">
                                <input type="submit" class="btn btn-sm btn-secondary" data-toggle="modal" data-target="#show_graph" value="Graph">
                            </div>
                            <div class="btn-group" data-toggle="buttons" id="node_names">
                                <label class="btn btn-sm btn-info-outline active">
                                    <input type="radio" id="etcd1" autocomplete="off" checked>etcd1
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd2" autocomplete="off">etcd2
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd3" autocomplete="off">etcd3
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd4" autocomplete="off">etcd4
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd5" autocomplete="off">etcd5
                                </label>
                            </div>
                        </div>
                        <div class="btn-toolbar">
                            <div class="btn-group" data-toggle="buttons" id="operations">
                                <label id="put_label" class="btn btn-sm btn-success-outline active">
                                    <input type="radio" id="put" autocomplete="off" checked>PUT
                                </label>
                                <label id="get_label" class="btn btn-sm btn-primary-outline">
                                    <input type="radio" id="get" autocomplete="off">GET
                                </label>
                                <label id="delete_label" class="btn btn-sm btn-warning-outline">
                                    <input type="radio" id="delete" autocomplete="off">DELETE
                                </label>
                            </div>
                            <div class="btn-group" data-toggle="buttons">
                                <input type="submit" class="btn btn-sm btn-default" id="key_submit" value="Submit">
                            </div>
                            <div class="btn-group" data-toggle="buttons">
                                <input type="submit" class="btn btn-sm btn-secondary" id="stress_submit" value="Stress">
                            </div>
                        </div>
                        <div class="input-group">
                            <!--                             <form method="POST" id="key_history">
                                <input id="key_input" type="text" style="min-width: 370px;" class="form-control" placeholder="Type your key...">
                            </form> -->
                            <input id="key_input" type="text" style="min-width: 370px;" class="form-control" placeholder="Type your key...">
                        </div>
                        <div class="input-group">
                            <textarea id="value_input" type="text" style="min-width: 370px;" class="form-control" placeholder="Type your value..." rows="7"></textarea>
                        </div>
                        <br>
                        <div id="result"></div>
                    </div>
                </div>
                <h2 class="sub-header" id="log">Log</h2>
                <div id="log_box" style="padding-top: 10px; padding-left: 5px; border:1px solid black; height:500px; overflow: scroll;">
                    <div id="log_output">
                    </div>
                </div>
                <br>
            </div>
        </div>
    </div>
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="https://code.jquery.com/ui/1.12.0-beta.1/jquery-ui.js"></script>
    <link rel="stylesheet" type="text/css" href="https://code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css">
</body>

</html>
`

var htmlSourceFileRemote = `<html lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="author" content="Gyu-Ho Lee">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.1.1/js/tether.min.js"></script>
    <script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.2/css/bootstrap.min.css" integrity="sha384-y3tfxAZXuh4HwSYylfB+J125MxIs6mR5FOHamPBG064zB+AFeWH94NdvaCBm8qnd" crossorigin="anonymous">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.2/js/bootstrap.min.js" integrity="sha384-vZ2WRJMwsjRMW/8U7i6PWi6AlO1L79snBrmgiDpgIWJ82z8eA5lenwvxbMV1PAh7" crossorigin="anonymous"></script>
    <link type="text/css" rel="stylesheet" href="http://fonts.googleapis.com/css?family=Inconsolata">
    <link rel='shortcut icon' type='image/x-icon' href='https://storage.googleapis.com/play-etcd/favicon.ico' />
    <script>
    $(document).ready(function() {
        function split(val) {
            return val.split(/,\s*/);
        }

        function extractLast(term) {
            return split(term).pop();
        }

        function appendLog(msg) {
            msg = $("<div />").html(msg);
            $("#log_output").append(msg);
            $("#log_box").scrollTop($("#log_box")[0].scrollHeight);
        }

        function runWithTimeout(callback, delay, times) {
            var internalCallback = function(t) {
                return function() {
                    if (--t > 0) {
                        window.setTimeout(internalCallback, delay);
                        callback();
                    }
                }
            }(times);
            window.setTimeout(internalCallback, delay);
        };

        $('#put_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px;";
            // document.getElementById('result').style = "display: none;"
            document.getElementById('result').style = ""
        });
        $('#get_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px; display: none;";
            document.getElementById('result').style = ""
        });
        $('#delete_label').click(function(e) {
            document.getElementById('value_input').style = "min-width: 370px; display: none;";
            document.getElementById('result').style = ""
        });

        $('#kill_1').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_1",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_1').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_1",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_2').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_2",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_2').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_2",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_3').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_3",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_3').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_3",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_4').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_4",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_4').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_4",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#kill_5').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/kill_5",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });
        $('#restart_5').click(function(e) {
            e.preventDefault();
            $.ajax({
                type: "GET",
                url: "/restart_5",
                async: true,
                dataType: "html",
                success: function(dataObj) {
                    appendLog(dataObj);
                }
            });
        });

        var receiveStreams = function() {
            $.ajax({
                url: '/stream',
                async: true,
                dataType: "json",
                success: function(dataObj) {
                    if (dataObj.Size > 0) {
                        appendLog(dataObj.Logs);
                    }
                }
            });
        }

        var queue_StorageKeysTotal = new Array;
        var receiveServerStatus = function() {
            $.ajax({
                type: "GET",
                url: '/server_status',
                async: true,
                dataType: "json",
                success: function(dataObj) {
                    queue_StorageKeysTotal.push(dataObj);

                    document.getElementById('active_user_number').innerHTML = "(active user: " + dataObj.ActiveUserNumber + ")";
                    document.getElementById('active_user_list').innerHTML = dataObj.ActiveUserList;

                    document.getElementById('etcd1_ID').innerHTML = "ID: <b>" + dataObj.Etcd1_ID + "</b>";
                    document.getElementById('etcd1_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd1_Endpoint + "</b>";
                    document.getElementById('etcd1_State').innerHTML = "State: <b>" + dataObj.Etcd1_State + "</b>";
                    document.getElementById('etcd1_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd1_NumberOfKeys + "</b>";
                    document.getElementById('etcd1_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd1_Hash + "</b>";
                    document.getElementById('etcd1_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd1_Hash + ")";
                    if (dataObj.Etcd1_State == "Leader") {
                        document.getElementById('etcd1_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd1_State == "Follower") {
                        document.getElementById('etcd1_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd1_State == "unreachable") {
                        document.getElementById('etcd1_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd1_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd2_ID').innerHTML = "ID: <b>" + dataObj.Etcd2_ID + "</b>";
                    document.getElementById('etcd2_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd2_Endpoint + "</b>";
                    document.getElementById('etcd2_State').innerHTML = "State: <b>" + dataObj.Etcd2_State + "</b>";
                    document.getElementById('etcd2_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd2_NumberOfKeys + "</b>";
                    document.getElementById('etcd2_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd2_Hash + "</b>";
                    document.getElementById('etcd2_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd2_Hash + ")";
                    if (dataObj.Etcd2_State == "Leader") {
                        document.getElementById('etcd2_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd2_State == "Follower") {
                        document.getElementById('etcd2_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd2_State == "unreachable") {
                        document.getElementById('etcd2_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd2_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd3_ID').innerHTML = "ID: <b>" + dataObj.Etcd3_ID + "</b>";
                    document.getElementById('etcd3_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd3_Endpoint + "</b>";
                    document.getElementById('etcd3_State').innerHTML = "State: <b>" + dataObj.Etcd3_State + "</b>";
                    document.getElementById('etcd3_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd3_NumberOfKeys + "</b>";
                    document.getElementById('etcd3_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd3_Hash + "</b>";
                    document.getElementById('etcd3_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd3_Hash + ")";
                    if (dataObj.Etcd3_State == "Leader") {
                        document.getElementById('etcd3_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd3_State == "Follower") {
                        document.getElementById('etcd3_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd3_State == "unreachable") {
                        document.getElementById('etcd3_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd3_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd4_ID').innerHTML = "ID: <b>" + dataObj.Etcd4_ID + "</b>";
                    document.getElementById('etcd4_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd4_Endpoint + "</b>";
                    document.getElementById('etcd4_State').innerHTML = "State: <b>" + dataObj.Etcd4_State + "</b>";
                    document.getElementById('etcd4_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd4_NumberOfKeys + "</b>";
                    document.getElementById('etcd4_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd4_Hash + "</b>";
                    document.getElementById('etcd4_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd4_Hash + ")";
                    if (dataObj.Etcd4_State == "Leader") {
                        document.getElementById('etcd4_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd4_State == "Follower") {
                        document.getElementById('etcd4_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd4_State == "unreachable") {
                        document.getElementById('etcd4_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd4_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }

                    document.getElementById('etcd5_ID').innerHTML = "ID: <b>" + dataObj.Etcd5_ID + "</b>";
                    document.getElementById('etcd5_Endpoint').innerHTML = "Endpoint: <b>" + dataObj.Etcd5_Endpoint + "</b>";
                    document.getElementById('etcd5_State').innerHTML = "State: <b>" + dataObj.Etcd5_State + "</b>";
                    document.getElementById('etcd5_NumberOfKeys').innerHTML = "Number of Keys: <b>" + dataObj.Etcd5_NumberOfKeys + "</b>";
                    document.getElementById('etcd5_Hash').innerHTML = "Hash: <b>" + dataObj.Etcd5_Hash + "</b>";
                    document.getElementById('etcd5_Hash_circle').innerHTML = "(Hash: " + dataObj.Etcd5_Hash + ")";
                    if (dataObj.Etcd5_State == "Leader") {
                        document.getElementById('etcd5_State_circle').style = "fill: #2E7D32"; // green
                    } else if (dataObj.Etcd5_State == "Follower") {
                        document.getElementById('etcd5_State_circle').style = "fill: #40C4FF"; // blue
                    } else if (dataObj.Etcd5_State == "unreachable") {
                        document.getElementById('etcd5_State_circle').style = "fill: #F44336"; // red
                    } else {
                        document.getElementById('etcd5_State_circle').style = "fill: #B3E5FC"; // light-blue
                    }
                }
            });
        }

        var wsConn;
        $.ajax({
            type: "GET",
            url: "/start_cluster",
            async: true,
            dataType: "json",
            success: function(dataObj) {
                appendLog(dataObj.Message);

                var wsURL = "ws://" + window.location.hostname + dataObj.PlayWebPort + "/ws";
                // var wsURL = "ws://" +  window.location.hostname + "/ws";
                console.log("connecting " + wsURL);
                if (window["WebSocket"]) {
                    wsConn = new WebSocket(wsURL);
                    wsConn.onopen = function() {
                        appendLog("Successfully connected to " + wsURL);
                    }
                    wsConn.onclose = function(ev) {
                        appendLog($("<div><b>connection closed</b></div>"));
                    }
                    wsConn.onmessage = function(ev) {
                        appendLog(ev.data);
                    }
                    wsConn.onerror = function(ev) {
                        appendLog("ERROR: " + ev.data);
                    }
                } else {
                    appendLog($("<div><b>browser does not support WebSocket</b></div>"))
                }
            }
        });
        runWithTimeout(receiveStreams, 100, 1);
        setInterval(receiveStreams, 1000);
        runWithTimeout(receiveStreams, 100, 1);
        setInterval(receiveServerStatus, 1000);

        $("#key_input")
            .on("click", function(event, ui) {
                if (event.keyCode === $.ui.keyCode.TAB && $(this).autocomplete("instance").menu.active) {
                    event.preventDefault();
                }
            })
            .autocomplete({
                minLength: 0,
                source: function(request, response) {
                    $.ajax({
                        type: "GET",
                        url: "/key_history",
                        dataType: "json",
                        data: {
                            q: request.term
                        },
                        success: function(dataObj) {
                            response($.ui.autocomplete.filter(dataObj.Values, extractLast(request.term)));
                        }
                    });
                },
                focus: function() {
                    // prevent value inserted on focus
                    return false;
                }
            });

        $('#key_submit').click(function(e) {
            var selectedOperation = $('#operations .active')[0].innerText;
            var selectedNodeName = $('#node_names .active')[0].innerText;
            var key = $('input#key_input').val();
            var value = $('textarea#value_input').val();
            var dataToSend = $(this).serializeArray()
            dataToSend.push({
                name: "selected_operation",
                value: selectedOperation
            })
            dataToSend.push({
                name: "selected_node_name",
                value: selectedNodeName
            })
            dataToSend.push({
                name: "key_input",
                value: key
            })
            dataToSend.push({
                name: "value_input",
                value: value
            })
            e.preventDefault();
            $.ajax({
                type: "POST",
                url: "/key_value",
                data: dataToSend,
                success: function(dataObj) {
                    $.ajax({
                        type: "GET",
                        url: "/key_value",
                        async: true,
                        dataType: "json",
                        success: function(dataObj) {
                            // appendLog(dataObj.Message);
                            wsConn.send(dataObj.Message);

                            document.getElementById('result').innerHTML = dataObj.Result
                        }
                    });
                }
            });
        });
    });
    </script>
    <style>
    svg a:hover {
        cursor: pointer;
    }
    /*
 * Base structure
 */
    /* Move down content because we have a fixed navbar that is 50px tall */
    
    body {
        padding-top: 50px;
    }
    /*
 * Global add-ons
 */
    
    .sub-header {
        padding-bottom: 10px;
        border-bottom: 1px solid #eee;
    }
    /*
 * Top navigation
 * Hide default border to remove 1px line.
 */
    
    .navbar-fixed-top {
        border: 0;
    }
    /*
 * Main content
 */
    
    .main {
        padding: 20px;
        margin-left: 35px;
        margin-right: 35px;
    }
    
    @media (min-width: 768px) {
        .main {
            padding-right: 20px;
            padding-left: 0px;
        }
    }
    
    .main .page-header {
        margin-top: 0;
    }
    /*
 * Placeholder dashboard ideas
 */
    
    .placeholders {
        margin-bottom: 30px;
        text-align: center;
    }
    
    .placeholders h4 {
        margin-bottom: 0;
    }
    
    .placeholder {
        margin-bottom: 20px;
    }
    
    .placeholder img {
        display: inline-block;
        border-radius: 50%;
    }
    
    .wrapper {
        text-align: center;
    }
    
    .ui-menu-item {
        width: 100%;
        height: 100%;
        font-family: "Lucida Console", Courier, monospace;
        font-size: 12px;
    }
    
    div .modal-body {
        font-family: 'Inconsolata', Courier;
        font-size: 17px;
    }
    
    #result {
        /*font-family: "Lucida Console", Courier, monospace;*/
        /*font-size: 11px;*/
        font-family: "Inconsolata", Courier, monospace;
        font-size: 13px;
        font-style: normal;
        font-variant: normal;
    }
    
    #log_box {
        /*font-family: "Lucida Console", Courier, monospace;*/
        /*font-size: 11px;*/
        font-family: "Inconsolata", Courier, monospace;
        font-size: 13px;
        font-style: normal;
        font-variant: normal;
    }
    
    ul.ui-autocomplete.ui-menu {
        z-index: 1000;
    }
    </style>
    <title>Play etcd</title>
</head>

<body>
    <div class="modal fade" id="active_user_Contents" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-sm" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Active Users</h5>
                </div>
                <div class="modal-body">
                    <div id="active_user_list">...</div>
                </div>
            </div>
        </div>
    </div>
    <nav class="navbar navbar-dark bg-primary navbar-fixed-top">
        <a class="navbar-brand" href="https://github.com/coreos/etcd" target="_blank">Play etcd</a>
        <ul class="nav navbar-nav">
            <li class="nav-item active">
                <a class="nav-link" href="/">Home <span class="sr-only">(current)</span></a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="https://github.com/coreos/etcd-play/issues/new">Report bug</a>
            </li>
            <li class="nav-item pull-xs-right">
                <div class="nav-link" id="active_user_number" data-toggle="modal" data-target="#active_user_Contents">(active user: 0)</div>
            </li>
        </ul>
    </nav>
    <div class="container-fluid">
        <div class="row">
            <div class="main">
                <!--                 <div id="etcd1_Contents" style="display: none;">
                    <div class="btn-group" role="group">
                        <button type="button" class="btn btn-danger btn-sm" id="kill_1">Kill</button>
                        <button type="button" class="btn btn-success btn-sm" id="restart_1">Restart</button>
                    </div>
                </div> -->
                <div class="modal fade" id="etcd1_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd1</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_1">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_1">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd1_ID">ID: n/a</div>
                                <div id="etcd1_Endpoint">Endpoint: n/a</div>
                                <div id="etcd1_State">State: n/a</div>
                                <div id="etcd1_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd1_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd2_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd2</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_2">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_2">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd2_ID">ID: n/a</div>
                                <div id="etcd2_Endpoint">Endpoint: n/a</div>
                                <div id="etcd2_State">State: n/a</div>
                                <div id="etcd2_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd2_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd3_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd3</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_3">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_3">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd3_ID">ID: n/a</div>
                                <div id="etcd3_Endpoint">Endpoint: n/a</div>
                                <div id="etcd3_State">State: n/a</div>
                                <div id="etcd3_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd3_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd4_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd4</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_4">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_4">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd4_ID">ID: n/a</div>
                                <div id="etcd4_Endpoint">Endpoint: n/a</div>
                                <div id="etcd4_State">State: n/a</div>
                                <div id="etcd4_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd4_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="etcd5_Contents" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-sm" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">etcd5</h5>
                            </div>
                            <div class="modal-body">
                                <div class="wrapper">
                                    <div class="btn-group" role="group">
                                        <button type="button" class="btn btn-danger" id="kill_5">Kill</button>
                                        <button type="button" class="btn btn-danger-outline" id="restart_5">Restart</button>
                                    </div>
                                </div>
                                <br>
                                <div id="etcd5_ID">ID: n/a</div>
                                <div id="etcd5_Endpoint">Endpoint: n/a</div>
                                <div id="etcd5_State">State: n/a</div>
                                <div id="etcd5_NumberOfKeys">Number of Keys: 0</div>
                                <div id="etcd5_Hash">Hash: 0</div>
                                <br>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal fade" id="show_graph" tabindex="-1" role="dialog" aria-hidden="true">
                    <div class="modal-dialog modal-lg" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                    <span aria-hidden="true">&times;</span>
                                </button>
                                <h5 class="modal-title">Graphs</h5>
                            </div>
                            <div class="modal-body">
                                <div id="chart_StorageKeysTotal" style="margin: 0 auto; width: 800px; height: 350px;"></div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="row">
                    <div class="col-md-5">
                        <svg xmlns="http://www.w3.org/2000/svg" version="1.1" style="border: 0px solid black" width="500" height="400">
                            <circle id="ring" cx="270" cy="210" r="110" style="fill: #F5F5F5"></circle>
                            <g id="servers">
                                <g id="etcd1_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd1_Contents" title="etcd1">
                                    <!-- <g id="etcd1_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="popover" title="etcd1"> -->
                                    <text x="255" y="50">etcd1</text>
                                    <text x="250" y="67" font-size="11" id="etcd1_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="270" cy="100" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd2_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd2_Contents" title="etcd2">
                                    <text x="395" y="160">etcd2</text>
                                    <text x="400" y="177" font-size="11" id="etcd2_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="370" cy="180" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd3_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd3_Contents" title="etcd3">
                                    <text x="360" y="320">etcd3</text>
                                    <text x="342" y="337" font-size="11" id="etcd3_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="335" cy="295" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd4_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd4_Contents" title="etcd4">
                                    <text x="130" y="320">etcd4</text>
                                    <text x="120" y="337" font-size="11" id="etcd4_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="200" cy="295" r="25"></circle>
                                    </a>
                                </g>
                                <g id="etcd5_State_circle" style="fill: #B3E5FC" tabindex="0" data-toggle="modal" data-target="#etcd5_Contents" title="etcd5">
                                    <text x="95" y="173">etcd5</text>
                                    <text x="87" y="150" font-size="11" id="etcd5_Hash_circle">(Hash: 0)</text>
                                    <a>
                                        <circle cx="165" cy="180" r="25"></circle>
                                    </a>
                                </g>
                            </g>
                        </svg>
                    </div>
                    <div class="col-md-7">
                        <br>
                        <br>
                        <div class="btn-toolbar">
                            <div class="btn-group" data-toggle="buttons" id="node_names">
                                <label class="btn btn-sm btn-info-outline active">
                                    <input type="radio" id="etcd1" autocomplete="off" checked>etcd1
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd2" autocomplete="off">etcd2
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd3" autocomplete="off">etcd3
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd4" autocomplete="off">etcd4
                                </label>
                                <label class="btn btn-sm btn-info-outline">
                                    <input type="radio" id="etcd5" autocomplete="off">etcd5
                                </label>
                            </div>
                        </div>
                        <div class="btn-toolbar">
                            <div class="btn-group" data-toggle="buttons" id="operations">
                                <label id="put_label" class="btn btn-sm btn-success-outline active">
                                    <input type="radio" id="put" autocomplete="off" checked>PUT
                                </label>
                                <label id="get_label" class="btn btn-sm btn-primary-outline">
                                    <input type="radio" id="get" autocomplete="off">GET
                                </label>
                                <label id="delete_label" class="btn btn-sm btn-warning-outline">
                                    <input type="radio" id="delete" autocomplete="off">DELETE
                                </label>
                            </div>
                            <div class="btn-group" data-toggle="buttons">
                                <input type="submit" class="btn btn-sm btn-default" id="key_submit" value="Submit">
                            </div>
                        </div>
                        <div class="input-group">
                            <input id="key_input" type="text" style="min-width: 370px;" class="form-control" placeholder="Type your key...">
                        </div>
                        <div class="input-group">
                            <textarea id="value_input" type="text" style="min-width: 370px;" class="form-control" placeholder="Type your value..." rows="7"></textarea>
                        </div>
                        <br>
                        <div id="result"></div>
                    </div>
                </div>
                <h2 class="sub-header" id="log">Log</h2>
                <div id="log_box" style="padding-top: 10px; padding-left: 5px; border:1px solid black; height:500px; overflow: scroll;">
                    <div id="log_output">
                    </div>
                </div>
                <br>
            </div>
        </div>
    </div>
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="https://code.jquery.com/ui/1.12.0-beta.1/jquery-ui.js"></script>
    <link rel="stylesheet" type="text/css" href="https://code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css">
</body>

</html>
`
