const express = require("express");
const path = require("path");
const router = require("./routes");

const app = express();
app.use(express.json());
app.use(router);
app.listen(3000);
