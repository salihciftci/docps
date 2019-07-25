const util = require("../util/util");

class Network {
    constructor(name, driver, scope) {
        this.name = name;
        this.driver = driver;
        this.scope = scope;
    }

    static async ls() {
        try {
            let command = "docker network ls --format '{{.Name}}\t{{.Driver}}\t{{.Scope}}'";
            let lines = await util.execCommand(command);

            let networks = [];
            lines.forEach((items) => {
                let item = items.split("\t");
                networks.push(new Network(item[0], item[1], item[2]));
            });

            return networks;
        } catch (e) {
            throw new Error(e);
        }
    }
}

module.exports = Network;