package rk_grpc

import (
	"net/http"
	"strings"
)

var apiHTML = "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <title>RK TV</title>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link rel=\"icon\" type=\"image/png\" href=\"https://www.flaticon.com/svg/static/icons/svg/2944/2944070.svg\"/>\n    <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.5.3/cerulean/bootstrap.min.css\">\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.3/js/bootstrap.min.js\"></script>\n</head>\n<body>\n<nav class=\"navbar navbar-expand-sm bg-dark navbar-dark\">\n    <a class=\"navbar-brand\" href=\".\">RK TV</a>\n    <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#collapsibleNavbar\">\n        <span class=\"navbar-toggler-icon\"></span>\n    </button>\n    <div class=\"collapse navbar-collapse\" id=\"collapsibleNavbar\">\n        <ul class=\"navbar-nav\">\n            <li class=\"nav-item\">\n                <a class=\"nav-link\" href=\"/v1/rk/tv/dashboard\">Dashboard</a>\n            </li>\n            <li class=\"nav-item\">\n                <a class=\"nav-link\" href=\"/v1/rk/tv/api\">API</a>\n            </li>\n            <li class=\"nav-item\">\n                <a class=\"nav-link\" href=\"/v1/rk/tv/info\">Info</a>\n            </li>\n        </ul>\n    </div>\n</nav>\n<!-- GRPC -->\n<div class=\"container\" style=\"margin-top:30px\">\n    <div class=\"card bg-primary text-white\">\n        <div class=\"card-body\">GRPC</div>\n    </div>\n</div>\n<div class=\"container\" style=\"margin-top:30px\">\n    <table class=\"table\">\n        <thead class=\"thead-light\">\n        <tr>\n            <th scope=\"col\">Name</th>\n            <th scope=\"col\">Service</th>\n            <th scope=\"col\">Method</th>\n            <th scope=\"col\">Type</th>\n            <th scope=\"col\">Port</th>\n        </tr>\n        </thead>\n        <tbody id=\"grpc\">\n        </tbody>\n    </table>\n</div>\n\n<!-- Gateway -->\n<div class=\"container\" style=\"margin-top:30px\">\n    <div class=\"card bg-success text-white\">\n        <div class=\"card-body\">Gateway</div>\n    </div>\n</div>\n<div class=\"container\" style=\"margin-top:30px\">\n    <table class=\"table\">\n        <thead class=\"thead-light\">\n        <tr>\n            <th scope=\"col\">Name</th>\n            <th scope=\"col\">Method</th>\n            <th scope=\"col\">Path</th>\n            <th scope=\"col\">Port</th>\n            <th scope=\"col\">SW</th>\n        </tr>\n        </thead>\n        <tbody id=\"gw\">\n        </tbody>\n    </table>\n</div>\n\n<script type=\"text/javascript\">\n    $(document).ready( function () {\n        $.ajax({\n            url : 'http://localhost:8080/v1/rk/apis',\n            type : 'GET',\n            dataType : 'json',\n            success : function(obj) {\n                let api = obj.api\n                for (let i = 0; i < api.length; i++) {\n                    let grpc = api[i].grpc\n                    let grpcTable = document.getElementById(\"grpc\")\n                    for (let j = 0; j < grpc.length; j++) {\n                        let row = `<tr>\n                                        <td>${api[i].name}</td>\n                                        <td>${grpc[j].service}</td>\n                                        <td>${grpc[j].method}</td>\n                                        <td>${grpc[j].type}</td>\n                                        <td>${grpc[j].port}</td>\n                                   </tr>`\n                        grpcTable.innerHTML += row\n                    }\n\n                    let gw = api[i].gw\n                    let gwTable = document.getElementById(\"gw\")\n                    for (let j = 0; j < gw.length; j++) {\n                        let row = `<tr>\n                                        <td>${api[i].name}</td>\n                                        <td>${gw[j].method}</td>\n                                        <td>${gw[j].path}</td>\n                                        <td>${gw[j].port}</td>\n                                        <td>${gw[j].sw}</td>\n                                   </tr>`\n                        gwTable.innerHTML += row\n                    }\n                }\n            }\n        });\n    });\n</script>\n\n</body>\n</html>\n"
var infoHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>RK TV</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.5.3/cerulean/bootstrap.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.3/js/bootstrap.min.js"></script>
</head>
<body>
<nav class="navbar navbar-expand-sm bg-dark navbar-dark">
    <a class="navbar-brand" href=".">RK TV</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="collapsibleNavbar">
        <ul class="navbar-nav">
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/dashboard">Dashboard</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/api">API</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/info">Info</a>
            </li>
        </ul>
    </div>
</nav>

<div class="container">
    <div class="row my-2">
        <div class="col-lg-8 order-lg-2">
            <ul class="nav nav-tabs">
                <li class="nav-item">
                    <a href="" data-target="#profile" data-toggle="tab" class="nav-link active">Basic</a>
                </li>
            </ul>
            <div class="tab-content py-4">
                <div class="tab-pane active" id="profile">
                    <div class="row">
                        <div class="col-md-6">
                            <h5>Application</h5>
                            <p id="application"></p>
                            <h5>Start time</h5>
                            <p id="start_time"></p>
                            <h5>Up time</h5>
                            <p id="up_time"></p>
                            <h5>Username</h5>
                            <p id="username"></p>
                        </div>
                        <div class="col-md-6">
                            <h5>GID</h5>
                            <p id="gid"></p>
                            <h5>UID</h5>
                            <p id="uid"></p>
                            <h5>Realm</h5>
                            <p id="realm"></p>
                            <h5>Region</h5>
                            <p id="region"></p>
                            <h5>AZ</h5>
                            <p id="az"></p>
                            <h5>Domain</h5>
                            <p id="domain"></p>
                        </div>
                    </div>
                    <!--/row-->
                </div>
            </div>
        </div>

        <script type="text/javascript">
            $(document).ready( function () {
                $.ajax({
                    url : 'http://localhost:8080/v1/rk/info',
                    type : 'GET',
                    dataType : 'json',
                    success : function(data) {
                        $("#application").text(data.info.application);
                        $("#start_time").text(data.info.start_time);
                        $("#up_time").text(data.info.up_time_str);
                        $("#username").text(data.info.username);
                        $("#gid").text(data.info.gid);
                        $("#uid").text(data.info.uid);
                        $("#realm").text(data.info.realm);
                        $("#region").text(data.info.region);
                        $("#az").text(data.info.az);
                        $("#domain").text(data.info.domain);
                    }
                });
            });
        </script>

        <!-- icon -->
        <div class="col-lg-4 order-lg-1 text-center">
            <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="150" height="150" class="mx-auto img-fluid img-circle d-block" alt="avatar">
        </div>
    </div>
</div>

</body>
</html>
`
var dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>Bootstrap 4 Website Example</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="icon" type="image/png" href="https://www.flaticon.com/svg/static/icons/svg/2944/2944070.svg"/>
    <link rel="stylesheet" type="text/css" href="https://pixinvent.com/stack-responsive-bootstrap-4-admin-template/app-assets/css/bootstrap-extended.min.css">
    <link rel="stylesheet" type="text/css" href="https://pixinvent.com/stack-responsive-bootstrap-4-admin-template/app-assets/fonts/simple-line-icons/style.min.css">
    <link rel="stylesheet" type="text/css" href="https://pixinvent.com/stack-responsive-bootstrap-4-admin-template/app-assets/css/colors.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.1/css/all.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.5.3/cerulean/bootstrap.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.3/js/bootstrap.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.18.1/moment.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.0/Chart.bundle.min.js"></script>
</head>
<body>
<nav class="navbar navbar-expand-sm bg-dark navbar-dark">
    <a class="navbar-brand" href=".">RK TV</a>

    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="collapsibleNavbar">
        <ul class="navbar-nav">
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/dashboard">Dashboard</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/api">API</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/info">Info</a>
            </li>
        </ul>
    </div>
</nav>

<!-- Header cards -->
<div class="container" style="margin-top:30px">
    <div class="grey-bg container-fluid">
        <section id="minimal-statistics">
            <div class="row">
                <div class="col-xl-3 col-sm-6 col-12">
                    <div class="card">
                        <div class="card-content">
                            <div class="card-body">
                                <div class="media d-flex">
                                    <div class="align-self-center">
                                        <i class="fas fa-microchip primary font-large-2 float-left"></i>
                                    </div>
                                    <div class="media-body text-right">
                                        <h3 id="cpu_curr">0 %</h3>
                                        <span>CPU</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-xl-3 col-sm-6 col-12">
                    <div class="card">
                        <div class="card-content">
                            <div class="card-body">
                                <div class="media d-flex">
                                    <div class="align-self-center">
                                        <i class="fas fa-memory warning font-large-2 float-left"></i>
                                    </div>
                                    <div class="media-body text-right">
                                        <h3 id="mem_curr">0 MB</h3>
                                        <span>Memory</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-xl-3 col-sm-6 col-12">
                    <div class="card">
                        <div class="card-content">
                            <div class="card-body">
                                <div class="media d-flex">
                                    <div class="align-self-center">
                                        <i class="fas fa-bullseye success font-large-2 float-left"></i>
                                    </div>
                                    <div class="media-body text-right">
                                        <h3 id="req_total">0</h3>
                                        <span>Requests Total</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-xl-3 col-sm-6 col-12">
                    <div class="card">
                        <div class="card-content">
                            <div class="card-body">
                                <div class="media d-flex">
                                    <div class="align-self-center">
                                        <i class="fas fa-history danger font-large-2 float-left"></i>
                                    </div>
                                    <div class="media-body text-right">
                                        <h3 id="up_time">0</h3>
                                        <span>Up time</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</div>

<!-- DROPDOWN -->
<div class="container">
    <div class="grey-bg container-fluid">
        <section id="minimal-statistics">
            <div class="row">
                <!-- DROPDOWN -->
                <div class="col-md-12">
                    <select id="API" class="selectpicker browser-default custom-select custom-select-lg mb-3">
                        <option selected>Select API</option>
                    </select>
                </div>
            </div>
        </section>
    </div>
</div>

<!-- API -->
<div class="container">
    <div class="grey-bg container-fluid">
        <section id="minimal-statistics">
            <div class="row">
                <!-- REQ_PER_SEC -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="REQ_PER_SEC"></canvas>
                        </div>
                    </div>
                </div>
                <!-- RES_CODE -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="RES_CODE"></canvas>
                        </div>
                    </div>
                </div>
                <!-- REQ_ELAPSED -->
                <!-- P50 -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="REQ_ELAPSED_P50"></canvas>
                        </div>
                    </div>
                </div>
                <!-- P90 -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="REQ_ELAPSED_P90"></canvas>
                        </div>
                    </div>
                </div>
                <!-- P99 -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="REQ_ELAPSED_P99"></canvas>
                        </div>
                    </div>
                </div>
                <!-- P999 -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="REQ_ELAPSED_P999"></canvas>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</div>

<!-- MEM & CPU -->
<div class="container">
    <div class="grey-bg container-fluid">
        <section id="minimal-statistics">
            <div class="row">
                <!-- CPU -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="CPU"></canvas>
                        </div>
                    </div>
                </div>
                <!-- MEM -->
                <div class="col-md-6">
                    <div class="card">
                        <div class="card-body">
                            <canvas id="MEM"></canvas>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</div>

<script>
    let intervalMS = 2000
    let maxLength = 10
    let cpuStat = [{
        x: new Date().toISOString(),
        y: 0,
    }];
    let memStat = [{
        x: new Date().toISOString(),
        y: 0,
    }];
    let emptyStat = [{
        x: new Date().toISOString(),
        y: 0,
    }];
    let reqPerSecMap = new Map();
    let reqElapsedP50 = new Map();
    let reqElapsedP90 = new Map();
    let reqElapsedP99 = new Map();
    let reqElapsedP999 = new Map();
    let reqPrevMap = new Map();
    let reqResCodeLabel = new Map();
    let reqResCodeData = new Map();
    let reqResCodeColor = new Map();

    let colors = ['#007bff','#28a745','#333333','#c3e6cb','#dc3545','#6c757d'];
    let CPU = new Chart(document.getElementById("CPU"), {
        type: 'line',
        data: {
            datasets: [{
                data: cpuStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'CPU'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: '%',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true
                    }
                }]
            },
            legend: false
        }
    });
    let MEM = new Chart(document.getElementById("MEM"), {
        type: 'line',
        data: {
            datasets: [{
                data: memStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'MEM'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'MB',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: false
        }
    });
    let REQ_PER_SEC = new Chart(document.getElementById("REQ_PER_SEC"), {
        type: 'line',
        data: {
            datasets: [{
                label: 'api',
                data: emptyStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'Request Per Sec'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'Count',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: {
                position: "bottom"
            },
        }
    });
    let REQ_ELAPSED_P50 = new Chart(document.getElementById("REQ_ELAPSED_P50"), {
        type: 'line',
        data: {
            datasets: [{
                label: 'api',
                data: emptyStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'Elapsed P50'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'ms',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: {
                position: "bottom"
            },
        }
    });
    let REQ_ELAPSED_P90 = new Chart(document.getElementById("REQ_ELAPSED_P90"), {
        type: 'line',
        data: {
            datasets: [{
                label: 'api',
                data: emptyStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'Elapsed P90'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'ms',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: {
                position: "bottom"
            },
        }
    });
    let REQ_ELAPSED_P99 = new Chart(document.getElementById("REQ_ELAPSED_P99"), {
        type: 'line',
        data: {
            datasets: [{
                label: 'api',
                data: emptyStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'Elapsed P99'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'ms',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: {
                position: "bottom"
            },
        }
    });
    let REQ_ELAPSED_P999 = new Chart(document.getElementById("REQ_ELAPSED_P999"), {
        type: 'line',
        data: {
            datasets: [{
                label: 'api',
                data: emptyStat,
                backgroundColor: colors[3],
                borderColor: colors[1],
                borderWidth: 1,
                pointBackgroundColor: colors[1]
            }]
        },
        options: {
            title: {
                display: true,
                text: 'Elapsed P99.9'
            },
            scales: {
                xAxes: [{
                    type: 'time',
                    time: {
                        unit: 'second',
                        displayFormats: {
                            quarter: 'HH:MM:SS'
                        }
                    }
                }],
                yAxes: [{
                    scaleLabel: {
                        display: true,
                        labelString: 'ms',
                    },
                    ticks: {
                        suggestedMin: 0,
                        beginAtZero: true,
                    }
                }]
            },
            legend: {
                position: "bottom"
            },
        }
    });
    let RES_CODE = new Chart(document.getElementById("RES_CODE"), {
        type: 'doughnut',
        data: {
            labels: [],
            datasets: [
                {
                    label: "Response Code",
                    backgroundColor: [],
                    data: []
                }
            ]
        },
        options: {
            title: {
                display: true,
                text: 'Response Code'
            }
        }
    });

    $(document).ready(function() {
        $.ajax({
            url : 'http://localhost:8080/v1/rk/apis',
            type : 'GET',
            dataType : 'json',
            success : function(data) {
                reloadAPI(data)
            },
            error: function() {}
        });
        setInterval(function() {
            $.ajax({
                url : 'http://localhost:8080/v1/rk/sys',
                type : 'GET',
                dataType : 'json',
                success : function(data) {
                    reloadSys(data.cpu_percentage, data.mem_usage_mb, data.up_time)
                    CPU.update()
                    MEM.update()
                },
                error: function() {
                    reloadSys(0, 0, 0)
                    CPU.update()
                    MEM.update()
                }
            });
        }, intervalMS);
        setInterval(function() {
            $.ajax({
                url : 'http://localhost:8080/v1/rk/req',
                type : 'GET',
                dataType : 'json',
                success : function(data) {
                    processResponse(data)
                    REQ_PER_SEC.update()
                    REQ_ELAPSED_P50.update()
                    REQ_ELAPSED_P90.update()
                    REQ_ELAPSED_P99.update()
                    REQ_ELAPSED_P999.update()
                    RES_CODE.update()
                },
                error: function() {
                    processResponse({})
                    REQ_PER_SEC.update()
                    REQ_ELAPSED_P50.update()
                    REQ_ELAPSED_P90.update()
                    REQ_ELAPSED_P99.update()
                    REQ_ELAPSED_P999.update()
                    RES_CODE.update()
                }
            });
        }, intervalMS);
    });

    // on selection
    $(function() {
        $('select').on('change', function(e){
            // clear it first
            REQ_PER_SEC.data.datasets[0].data = emptyStat
            REQ_ELAPSED_P50.data.datasets[0].data = emptyStat
            REQ_ELAPSED_P90.data.datasets[0].data = emptyStat
            REQ_ELAPSED_P99.data.datasets[0].data = emptyStat
            REQ_ELAPSED_P999.data.datasets[0].data = emptyStat

            // select datasets from map
            if (reqPerSecMap.has(this.value)) {
                REQ_PER_SEC.data.datasets[0].label = this.value
                REQ_PER_SEC.data.datasets[0].data = reqPerSecMap.get(this.value)

                REQ_ELAPSED_P50.data.datasets[0].label = this.value
                REQ_ELAPSED_P50.data.datasets[0].data = reqElapsedP50.get(this.value)

                REQ_ELAPSED_P90.data.datasets[0].label = this.value
                REQ_ELAPSED_P90.data.datasets[0].data = reqElapsedP90.get(this.value)

                REQ_ELAPSED_P99.data.datasets[0].label = this.value
                REQ_ELAPSED_P99.data.datasets[0].data = reqElapsedP99.get(this.value)

                REQ_ELAPSED_P999.data.datasets[0].label = this.value
                REQ_ELAPSED_P999.data.datasets[0].data = reqElapsedP999.get(this.value)

                RES_CODE.data.labels = reqResCodeLabel.get(this.value)
                RES_CODE.data.datasets[0].data = reqResCodeData.get(this.value)
                RES_CODE.data.datasets[0].backgroundColor = reqResCodeColor.get(this.value)
            }
        });
    });

    function reloadSys(cpu, mem, up) {
        if (cpuStat.length > maxLength) {
            cpuStat.shift()
        }

        if (memStat.length > maxLength) {
            memStat.shift()
        }

        var now = new Date();
        cpuStat.push({
            x: now.toISOString(),
            y: cpu,
        })

        memStat.push({
            x: now.toISOString(),
            y: mem,
        })
        document.getElementById("cpu_curr").innerText = cpu+" %"
        document.getElementById("mem_curr").innerText = mem+" MB"
        document.getElementById("up_time").innerText = up
    }

    function processResponse(input) {
        let reqTotal = 0
        let resp = input.metrics
        for (let i = 0; i < resp.length; i++) {
            let metric = resp[i];
            reloadReqRate(metric)
            reloadResCode(metric)
            reloadElapsed(reqElapsedP50, metric.path, metric.elapsed_nano_p50)
            reloadElapsed(reqElapsedP90, metric.path, metric.elapsed_nano_p90)
            reloadElapsed(reqElapsedP99, metric.path, metric.elapsed_nano_p99)
            reloadElapsed(reqElapsedP999, metric.path, metric.elapsed_nano_p999)
            reqTotal += metric.count
        }
        document.getElementById("req_total").innerText = reqTotal
    }

    function getRandomColor(size) {
        let res = []
        for (let j = 0; j < size; j++) {
            let letters = '0123456789ABCDEF'.split('');
            let color = '#';
            for (let i = 0; i < 6; i++ ) {
                color += letters[Math.floor(Math.random() * 16)];
            }
            res.push(color)
        }

        return res
    }

    function reloadAPI(input) {
        let api = input.api
        for (let i = 0; i < api.length; i++) {
            let grpc = api[i].grpc
            for (let j = 0; j < grpc.length; j++) {
                let api = document.getElementById("API");
                let op = document.createElement("option");
                let linkText = document.createTextNode(grpc[j].method);
                op.appendChild(linkText);
                api.appendChild(op);
            }
        }
    }

    function reloadElapsed(map, path, elapsedQuantile) {
        let now = new Date()
        if (map.has(path)) {
            let array = map.get(path)
            if (array.length > maxLength) {
                array.shift();
            }

            array.push({
                x: now.toISOString(),
                y: elapsedQuantile/1e6,
            })
        } else {
            let elapsed = [{
                x: now.toISOString(),
                y: elapsedQuantile/1e6,
            }]
            map.set(path, elapsed)
        }
    }

    function reloadResCode(metric) {
        if (reqResCodeLabel.has(metric.path)) {
            for (let f = 0; f < reqResCodeLabel.get(metric.path).length; f++) {
                reqResCodeLabel.get(metric.path).shift()
            }
        } else {
            reqResCodeLabel.set(metric.path, new Array())
        }

        if (reqResCodeData.has(metric.path)) {
            for (let f = 0; f < reqResCodeData.get(metric.path).length; f++) {
                reqResCodeData.get(metric.path).shift()
            }
            reqResCodeData.get(metric.path).length = 0
        } else {
            reqResCodeData.set(metric.path, new Array())
        }

        if (!reqResCodeColor.has(metric.path)) {
            reqResCodeColor.set(metric.path, new Array())
        }

        for (let j = 0; j < metric.res_code.length; j++) {
            let element = metric.res_code[j]
            reqResCodeLabel.get(metric.path).push(element.res_code + "[" + element.count + "]")
            reqResCodeData.get(metric.path).push(element.count)
        }

        if (reqResCodeColor.get(metric.path).length != reqResCodeLabel.get(metric.path).length) {
            let colors = getRandomColor(reqResCodeLabel.get(metric.path).length - reqResCodeColor.get(metric.path).length)
            for (let k = 0; k < colors.length; k++) {
                reqResCodeColor.get(metric.path).push(colors[k])
            }
        }
    }

    function reloadReqRate(metric) {
        if (reqPerSecMap.has(metric.path)) {
            let array = reqPerSecMap.get(metric.path)
            if (array.length > maxLength) {
                array.shift();
            }
            // get prev value
            let prevMetric = reqPrevMap.get(metric.path)
            let prevValue = prevMetric[prevMetric.length-1].count
            let currValue = metric.count

            array.push({
                x: new Date().toISOString(),
                y: (currValue-prevValue)/(intervalMS/1000),
            })
            prevMetric.push(metric)
        } else {
            let perSec = [{
                x: new Date().toISOString(),
                y: metric.count,
            }]
            reqPerSecMap.set(metric.path, perSec)
            reqPrevMap.set(metric.path, [metric])
        }
    }
</script>
</body>
</html>
`
var notFoundHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>RK TV</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/4.5.3/cerulean/bootstrap.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.3/js/bootstrap.min.js"></script>
</head>
<body>
<nav class="navbar navbar-expand-sm bg-dark navbar-dark">
    <a class="navbar-brand" href=".">RK TV</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="collapsibleNavbar">
        <ul class="navbar-nav">
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/dashboard">Dashboard</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/api">API</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="/v1/rk/tv/info">Info</a>
            </li>
        </ul>
    </div>
</nav>

<div class="container">
    <div class="grey-bg container-fluid">
        <section id="minimal-statistics">
            <div class="row">
                <div class="col-md-12">
                    <div class="error-template" style="padding: 40px 15px;text-align: center;">
                        <h1>Oops!</h1>
                        <h2>404 Not Found</h2>
                        <div class="error-details" style="margin-top:15px;margin-bottom:15px;">
                            Sorry, an error has occured, Requested page not found!<br>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</div>

</body>
</html>
`

func tv(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/"), "/")

	println(path)

	switch path {
	case "v1/rk/tv":
		w.Write([]byte(dashboardHTML))
	case "v1/rk/tv/api":
		w.Write([]byte(apiHTML))
	case "v1/rk/tv/dashboard":
		w.Write([]byte(dashboardHTML))
	case "v1/rk/tv/info":
		w.Write([]byte(infoHTML))
	default:
		w.Write([]byte(notFoundHTML))
	}
}