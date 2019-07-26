const util = require("../util/util");

class Volume {
    constructor(name, driver) {
        this.name = name;
        this.driver = driver;
    }

    static async ls() {
        try {
            let command = "docker volume ls --format '{{.Name}}\t{{.Driver}}'";
            let lines = await util.execCommand(command);

            let volumes = [];
            lines.forEach((items) => {
                let item = items.split("\t");
                volumes.push(new Volume(item[0], item[1]));
            });

            return volumes;
        } catch (e) {
            throw new Error(e);
        }
    }

    static async count() {
        let command = "docker volume ls";
        let lines = await util.execCommand(command);

        lines.shift();
        return Object.keys(lines).length;
    }
}

module.exports = Volume;