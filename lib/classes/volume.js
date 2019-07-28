const util = require("../util/cmd");

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
        let volumes = await Volume.ls();
        return volumes.length;
    }
}

module.exports = Volume;