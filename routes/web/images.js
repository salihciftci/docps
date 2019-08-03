const express = require("express");
const router = express.Router();
const { Image } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let images = await Image.ls();
        res.render("images", {
            title: "Images",
            user: req.user,
            images: images
        });
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;