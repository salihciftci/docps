function Search() {
    var input, filter, table, tr, td, i;
    input = document.getElementById("search");
    filter = input.value.toUpperCase();
    table = document.getElementById("table");
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

function Live() {
    var filter, table, tr, td, i;
    filter = "fa-check";
    table = document.getElementById("table");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[6];
        if (td) {
            if (td.innerHTML.indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

function Stopped() {
    var filter, table, tr, td, i;
    filter = "fa-times";
    table = document.getElementById("table");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[6];
        if (td) {
            if (td.innerHTML.indexOf(filter) > -1) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

function All() {
    var table, tr, i;
    table = document.getElementById("table");
    tr = table.getElementsByTagName("tr");
    for (i = 0; i < tr.length; i++) {
        tr[i].style.display = "";
    }
}

function Copy() {
    var copyText = document.getElementById("key");
    copyText.select();
    document.execCommand("copy");
    document.getSelection().removeAllRanges();
}

function Root() {
    document.getElementById('d').checked = false;
    document.getElementById('c').checked = false;
    document.getElementById('s').checked = false;
    document.getElementById('i').checked = false;
    document.getElementById('v').checked = false;
    document.getElementById('n').checked = false;
    document.getElementById('l').checked = false;
    document.getElementById('b').checked = false;
}

function NotRoot() {
    document.getElementById('r').checked = false;
}
