const express = require("express");
const router = express.Router();
const { Container } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let containers = await Container.ls();
        containers.forEach((e, i) => {
            delete containers[i].IP;
            delete containers[i].restartPolicy;
            delete containers[i].volumes;
        });
        res.json(containers);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

router.get("/:id", async (req, res) => {
    try {
        let id = req.params.id;
        let containers = await Container.inspect(id);
        res.json(containers);
    } catch (e) {
        if (e.code === 404) {
            res.sendStatus(404);
            return;
        }
        res.sendStatus(500);
    }
});

router.get("/:id/logs", async (req, res) => {
    let id = req.params.id;
    let lineCount = req.body.lines || 10;
    try {
        let logs = await Container.logs(id, lineCount);
        res.json(logs);
    } catch (e) {
        if (e.code === 404) {
            res.sendStatus(404);
            return;
        }
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;