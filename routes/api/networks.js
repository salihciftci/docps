const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let networks = await Docker.Network.ls();
        res.json(networks);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;