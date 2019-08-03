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

        res.render("containers", {
            title: "Containers",
            user: req.user,
            containers: containers
        });
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

router.get("/:id/logs", async (req, res) => {
    try {
        let id = req.params.id;
        let logs = await Container.logs(id, 50);

        res.render("logs", {
            title: "Logs",
            user: req.user,
            logs: logs
        });
    } catch (e) {
        if (e.code === 404) {
            res.render("404", { title: "404" });
            return;
        }
        console.log(e);
        res.render("500", { title: "Error" });
    }
});



module.exports = router;