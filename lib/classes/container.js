var util = require("../util/util");

class Container {
    constructor(id, name, image, size, run, status, ports) {
        this.id = id;
        this.name = name;
        this.image = image;
        this.size = size;
        this.run = run;
        this.status = status;
        this.ports = ports;
        this.logs = [];
    }

    async ls() {
        try {
            let containers = [];
            let command = "docker ps -a --format '{{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}\t{{.Ports}}'";
            let lines = await util.execCommand(command);

            lines.forEach((items) => {
                let item = items.split("\t");
                containers.push(new Container(item[0], item[1], item[2], item[3], item[4], item[5], item[6], item[7]));
            });

            return containers;
        } catch (e) {
            throw new Error(e);
        }
    }

    async logs(containerID, lineCount) {
        try {
            let containers = [];
            let command = `docker logs --tail ${lineCount} ${containerID}`;
            let lines = await util.execCommand(command);

            return lines;
        } catch (e) {
            throw new Error(e);
        }
    }
}

module.exports = Container;
