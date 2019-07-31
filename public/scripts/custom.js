function renderSidebar() {
    let pageName = document.getElementById("pageName").innerHTML.toLowerCase();
    if (pageName === "logs") {
        pageName = "containers";
    }
    let sidebarElement = document.getElementById(pageName);
    sidebarElement.style.backgroundColor = "#e9ecef";
}

// eslint-disable-next-line
function containerSearch() {
    var input, filter, table, tr, td, i;
    input = document.getElementById("containerInput");
    filter = input.value.toUpperCase();

    table = document.getElementById("containerTable");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[0];
        if (td) {
            if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

// eslint-disable-next-line
function showLiveContainers() {
    var filter, table, tr, td, i;
    filter = "fa-check";
    table = document.getElementById("containerTable");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[4];
        if (td) {
            if (td.innerHTML.indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
    document.getElementById("liveButton").style.backgroundColor = "#e9ecef";
    document.getElementById("stoppedButton").style.backgroundColor = "inherit";
    document.getElementById("allButton").style.backgroundColor = "inherit";
}

// eslint-disable-next-line
function showStoppedContainers() {
    var filter, table, tr, td, i;
    filter = "fa-times";
    table = document.getElementById("containerTable");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[4];
        if (td) {
            if (td.innerHTML.indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }

    document.getElementById("liveButton").style.backgroundColor = "inherit";
    document.getElementById("stoppedButton").style.backgroundColor = "#e9ecef";
    document.getElementById("allButton").style.backgroundColor = "inherit";
}

// eslint-disable-next-line
function showAllContainers() {
    var table, tr, i;
    table = document.getElementById("containerTable");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        tr[i].style.display = "";
    }

    document.getElementById("liveButton").style.backgroundColor = "inherit";
    document.getElementById("stoppedButton").style.backgroundColor = "inherit";
    document.getElementById("allButton").style.backgroundColor = "#e9ecef";
}


renderSidebar();