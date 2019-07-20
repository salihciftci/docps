"use strict";

const util = require("./util/util");

const Container = require("./classes/container");
const Image = require("./classes/image");
const Volume = require("./classes/volume");
const Network = require("./classes/network");

class Docker {
    async info() {
        try {
            let command = "docker info --format '{{.Name}}\t{{.ServerVersion}}\t{{.NCPU}}\t{{.MemTotal}}'";
            let lines = util.execCommand(command);

            let items = lines[0].split("\t");

            return {
                "Name": items[0],
                "Version": items[1],
                "CPU": items[2],
                "MemTotal": items[3]
            };
        } catch (e) {
            throw new Error(e);
        }
    }

    async stats() {
        try {
            let command = "docker stats --no-stream --format '{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}'";
            let lines = await util.execCommand(command);

            let stats = [];
            lines.forEach((items) => {
                let item = items.split("\t");
                stats.push({
                    "Name": item[0],
                    "CPUPerc": item[1],
                    "MemUsage": item[2].split("/")[0].trimRight(),
                    "MemPerc": item[3],
                    "NetIO": item[4],
                    "BlockIO": item[5],
                });
            });

            return stats;
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