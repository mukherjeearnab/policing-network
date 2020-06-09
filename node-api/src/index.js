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

app.use(require("./routes/auth"));
app.use(require("./routes/chargesheet_cc"));
app.use(require("./routes/citizenprofile_cc"));
app.use(require("./routes/evidence_cc"));
app.use(require("./routes/fir_cc"));
app.use(require("./routes/investigation_cc"));
app.use(require("./routes/judgement_cc"));

app.listen(3000, console.log);
