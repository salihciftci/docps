const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        res.render("index", {
            title: "Home",
            user: req.user,
            containersCount: await Docker.Container.count(),
            imagesCount: await Docker.Image.count(),
            volumesCount: await Docker.Volume.count(),
            networksCount: await Docker.Network.count(),
            docker: await Docker.Docker.info()
        });
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;