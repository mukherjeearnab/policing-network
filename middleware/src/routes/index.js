const { Router } = require("express");
const fs = require("fs");
const path = require("path");

const router = Router();

/**
 * @summary Read all files from the current directory and use then as routes.
 */
fs.readdirSync(`${path.resolve()}/src/routes/`).forEach((file) => {
    console.log(file);
    if (file.match(/\.js$/) !== null && file !== "index.js")
        // eslint-disable-next-line
        require(`./${file.replace(".js", "")}`)(router);
});

console.log(router);

module.exports = router;
