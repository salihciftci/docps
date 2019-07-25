const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let volumes = await Docker.Volume.ls();
        res.json(volumes);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;