const express = require("express");
const router = express.Router();
const { Volume } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let volumes = await Volume.ls();
        res.render("volumes", {
            title: "Volumes",
            user: req.user,
            volumes: volumes
        });
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;