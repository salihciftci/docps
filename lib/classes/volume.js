const util = require("../util/util");

class Volume {
    constructor(repo, tag, id, created, size) {
        this.repo = repo;
        this.tag = tag;
        this.id = id;
        this.created = created;
        this.size = size;
    }

    async ls() {
        try {
            command = "docker volume ls --format '{{.Driver}}\t{{.Name}}'";
            let lines = await util.execCommand(command);
            console.log(lines);
        } catch (e) {
            throw new Error(e);
        }
    }
}

module.exports = Image;