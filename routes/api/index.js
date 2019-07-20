const express = require("express");
const router = express.Router();
const { Docker } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let docker = new Docker();
        let info = await docker.info();
        res.json(info);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

router.get("/stats", async (req, res) => {
    try {
        let docker = new Docker();
        let stats = await docker.stats();
        res.json(stats);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;