const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    let network = new Docker.Network();
    try {
        let networks = await network.ls();
        res.json(networks);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;