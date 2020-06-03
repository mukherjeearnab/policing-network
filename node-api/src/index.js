const express = require("express");
const path = require("path");
//const router = require("./routes");

const app = express();
app.use(express.json());
app.use(require("./routes/auth"));
app.use(require("./routes/chargesheet_cc"));
app.use(require("./routes/citizenprofile_cc"));
app.use(require("./routes/evidence_cc"));
app.use(require("./routes/fir_cc"));
app.use(require("./routes/investigation_cc"));
app.use(require("./routes/judgement_cc"));
app.listen(3000);
