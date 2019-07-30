const util = require("../util/cmd");

class Image {
    constructor(repo, tag, id, created, size) {
        this.repository = repo;
        this.tag = tag;
        this.id = id;
        this.created = created;
        this.size = size;
    }

    static async ls() {
        try {
            let images = [];
            let command = "docker image ls --format '{{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.CreatedSince}}\t{{.Size}}'";
            let stdOut = await util.execCommand(command);

            stdOut.forEach((line) => {
                let e = line.split("\t");
                images.push(new Image(e[0], e[1], e[2], e[3], e[4]));
            });

            return images;
        } catch (e) {
            throw new Error(e);
        }
    }

    static async count() {
        let images = await Image.ls();
        return images.length;
    }
}

module.exports = Image;