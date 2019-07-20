const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    let volume = new Docker.Volume();
    try {
        let volumes = await volume.ls();
        res.json(volumes);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;