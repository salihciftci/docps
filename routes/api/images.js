const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    let image = new Docker.Image();
    try {
        let images = await image.ls();
        res.json(images);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;