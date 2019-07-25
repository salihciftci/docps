const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let containers = await Docker.Container.ls();
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
    try {
        let logs = await Docker.Container.logs(containerName, lineCount);
        res.json(logs);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;