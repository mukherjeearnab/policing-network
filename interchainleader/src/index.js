const express = require("express");
const cors = require("cors");

const app = express();
app.use(express.json());
app.use(cors());
app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE");
    res.header("Access-Control-Allow-Headers", "Content-Type");
    next();
});

app.use(require("./routes/icn"));
app.use(require("./routes/ian"));
app.use(require("./routes/lran"));
app.use(require("./routes/vtan"));

app.listen(3001, console.log);
