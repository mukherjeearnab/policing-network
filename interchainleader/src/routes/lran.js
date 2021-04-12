const express = require("express");
const fs = require("fs");
const fetch = require("node-fetch");

const rsaEncrypt = require("../helpers/rsaEncrypt");
const rsaDecryptMiddleware = require("../helpers/rsaDecryptMiddleware");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");

require("dotenv").config();
const router = new express.Router();

/**********************************
 *  VTAN ROUTES
 **********************************/

router.post("/api/icn/vtan/get/license/", rsaDecryptMiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.VTAN_ADDRESS, {
        icl: process.env.NETWORK_NAME,
        body: req.body,
    });

    const response = await fetch(`http://${process.env.VTAN_ADDRESS}/icn/vtan/get/license`, {
        method: "POST",
        body: JSON.stringify({ payload }),
        headers: { "Content-Type": "application/json" },
    });

    let resp = await response.json();

    // Debugging Response
    console.log("API/ICN/TEST-S", resp);

    res.status(200).send({
        response: resp,
    });
});

router.post("/api/icn/vtan/get/vehicle/", rsaDecryptMiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.VTAN_ADDRESS, {
        icl: process.env.NETWORK_NAME,
        body: req.body,
    });

    const response = await fetch(`http://${process.env.VTAN_ADDRESS}/icn/vtan/get/vehicle`, {
        method: "POST",
        body: JSON.stringify({ payload }),
        headers: { "Content-Type": "application/json" },
    });

    let resp = await response.json();

    // Debugging Response
    console.log("API/ICN/TEST-S", resp);

    res.status(200).send({
        response: resp,
    });
});

module.exports = router;
