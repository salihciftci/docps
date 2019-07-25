const express = require("express");
const router = express.Router();
const { Docker } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let info = await Docker.info();
        res.json(info);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

router.get("/stats", async (req, res) => {
    try {
        let stats = await Docker.stats();
        res.json(stats);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;