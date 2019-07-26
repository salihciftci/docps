const express = require("express");
const router = express.Router();
const { Container } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let containers = await Container.ls();
        res.render("containers", {
            title: "Containers",
            containers: containers
        });
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;