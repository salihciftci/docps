const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let images = await Docker.Image.ls();
        res.json(images);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

module.exports = router;