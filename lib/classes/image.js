const util = require("../util/util");

class Image {
    constructor(repo, tag, created, size) {
        this.repo = repo;
        this.tag = tag;
        this.created = created;
        this.size = size;
    }

    static async ls() {
        try {
            let images = [];
            let command = "docker image ls --format '{{.Repository}}\t{{.Tag}}\t{{.CreatedSince}}\t{{.Size}}'";
            let lines = await util.execCommand(command);

            lines.forEach((items) => {
                let item = items.split("\t");
                images.push(new Image(item[0], item[1], item[2], item[3]));
            });

            return images;
        } catch (e) {
            throw new Error(e);
        }
    }
}

module.exports = Image;