module.exports.parsePorts = (str) => {
    if (typeof str === "undefined" || !str) {
        return str;
    }

    let ports = str.split(",");
    ports.forEach((port, i) => {
        port = port.trim();
        if (port.length > 8) {
            ports[i] = port.substring(8, port.length).replace("->", ":");
        } else {
            ports[i] = port.trim();
        }
    });

    return ports;
};