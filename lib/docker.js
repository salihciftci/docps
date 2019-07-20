"use strict";

const util = require("./util/util");

const Container = require("./classes/container");
const Image = require("./classes/image");
const Volume = require("./classes/volume");
const Network = require("./classes/network");

class Docker {
    constructor() {
    }

    async info() {
        try {
            let info = {

            };
            const { stdout } = await exec(`
            docker info --format '{{.Containers}}\\t{{.Name}}\\t{{.ServerVersion}}\\t{{.NCPU}}\\t{{.MemTotal}}'`);

            /*let lines = stdout.split("\n");
            lines.pop();*/

            /*lines.forEach((items) => {
                let item = items.split("\t");
                containers.push(new Container(item[0], item[1], item[2], item[3], item[4], item[5], item[6], item[7]));
            });*/

            return stdout;
        } catch (e) {
            throw new Error(e);
        }
    }
}

module.exports = {
    Docker,
    Container,
    Image,
    Network,
    Volume,
};