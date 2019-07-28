const express = require("express");
const router = express.Router();
const { Container } = require("../../lib/docker");

router.get("/", async (req, res) => {
    try {
        let containers = await Container.ls();

        for (let [i, container] of containers.entries()) {
            containers[i] = await Container.inspect(container.id);
            containers[i].name = container.name.charAt(0).toUpperCase() + container.name.slice(1);
        }

        console.log(containers[0]);

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