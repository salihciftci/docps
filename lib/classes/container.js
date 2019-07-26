var util = require("../util/util");

class Container {
    constructor(id, name, image, size, run, status, ports) {
        this.id = id;
        this.name = name;
        this.image = image;
        this.size = size;
        this.run = run;
        this.status = status;
        this.ports = ports || [];
    }

    static async ls() {
        try {
            let containers = [];
            let command = "docker ps -a --format '{{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.RunningFor}}\t{{.Status}}'";
            let lines = await util.execCommand(command);
            lines.forEach((items) => {
                let item = items.split("\t");

                item[5][0] === "U" ? item[5] = true : item[5] = false;

                containers.push(new Container(item[0], item[1], item[2], item[3], item[4], item[5]));
            });

            return containers;
        } catch (e) {
            throw new Error(e);
        }
    }

    static async logs(containerID, lineCount) {
        try {
            let command = `docker logs --tail ${lineCount} ${containerID}`;
            let lines = await util.execCommand(command);

            //todo remove \r in every line

            return lines;
        } catch (e) {
            throw new Error(e);
        }
    }

    static async count() {
        let command = "docker ps -a";
        let lines = await util.execCommand(command);
        lines.shift();
        return Object.keys(lines).length;
    }

    static async inspect() {
        // todo what we need to return here?
        /* 
        DONT FORGET PORTS
        IP address?
        Restart Pocliy?
        Mounts?
        RestartCount? 
        ..
        */
    }
}

module.exports = Container;