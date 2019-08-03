const express = require("express");
const router = express.Router();
const { Docker } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let stats = await Docker.stats();
        res.render("stats", {
            title: "Stats",
            user: req.user,
            stats: stats
        });
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;