const express = require("express");
const router = express.Router();
const { Network } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let networks = await Network.ls();
        res.render("networks", {
            title: "Networks",
            user: req.user,
            networks: networks
        });
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;