const util = require("../util/util");
const { parsePorts } = require("../util/ports");

class Container {
    constructor() {
        this.id = "";
        this.name = "";
        this.image = "";
        this.size = "";
        this.run = "";
        this.status = "";
        this.ports = [];
        this.IP = "";
        this.restartPolicy = {};
        this.volumes = {};
    }

    static async ls() {
        try {
            let containers = [];
            let command = "docker ps -a --format '{{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Size}}\t{{.Status}}\t{{.Ports}}'";
            let stdOut = await util.execCommand(command);
            stdOut.forEach((line) => {
                let e = line.split("\t");

                let container = new Container();
                container.id = e[0];
                container.name = e[1];
                container.image = e[2];
                container.size = e[3];
                container.run = e[4];
                container.status = e[4][0] === "U" ? true : false;
                container.ports = parsePorts(e[5]) || [];

                containers.push(container);
            });

            return containers;
        } catch (e) {
            throw new Error(e);
        }
    }

    static async inspect(id) {
        let containers = await Container.ls();
        let container;

        containers.find((e) => {
            if (e.id === id) {
                container = e;
            }
        });

        try {
            let command = `docker inspect ${id} --format '{{json .Mounts}}\t{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}\t{{json .HostConfig.RestartPolicy}}'`;
            let stdOut = await util.execCommand(command);

            stdOut.forEach(line => {
                let e = line.split("\t");
                container.IP = e[1];
                container.volumes = JSON.parse(e[0]);
                container.restartPolicy = JSON.parse(e[2]);
            });

            return container;
        } catch (e) {
            if (e.stderr.indexOf("No such object")) {
                e.code = 404;
            }

            throw e;
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
}

module.exports = Container;