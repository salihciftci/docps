const express = require("express");
const router = express.Router();
const Docker = require("../../lib/docker");

router.get("/", async (req, res) => {
    res.json("sa");
    // let container = new Docker.Container();
    // try {
    //     let containers = await container.ls();
    //     res.json(containers);
    // } catch (e) {
    //     console.log(e);
    // }
});

// router.get("/:id", function (req, res) {
//     console.log("2");
// });

// router.get("/:id/logs", function (req, res) {
//     console.log("3");
// });

module.exports = router;