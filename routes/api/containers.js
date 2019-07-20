const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    let container = new Docker.Container();
    try {
        let containers = await container.ls();
        res.json(containers);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

router.get("/:id", function (req, res) {

});

router.get("/:name/logs", async (req, res) => {
    //todo check id if name is []
    let containerName = req.params.name;
    let lineCount = req.body.lines || 10;
    let container = new Docker.Container();
    try {
        let logs = await container.logs(containerName, lineCount);
        res.json(logs);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;