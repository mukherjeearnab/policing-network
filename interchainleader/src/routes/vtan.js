const express = require("express");
const fetch = require("node-fetch");

const rsaEncryptSend = require("../helpers/rsaEncryptSend");
const rsaEncryptReturn = require("../helpers/rsaEncryptReturn");
const rsaDecrypt = require("../helpers/rsaDecrypt");
const rsaDecryptMiddleware = require("../helpers/rsaDecryptMiddleware");

require("dotenv").config();
const router = new express.Router();

/**********************************
 *  VTAN ROUTES
 **********************************/

router.post("/api/icn/vtan/get/license", rsaDecryptMiddleware, async (req, res) => {
    if (!req.body.message.verified) {
        res.status(400).send({ response: "ERROR! ICL Error!" });
    }

    let payload = await rsaEncryptSend(process.env.VTAN_ADDRESS, {
        icl: process.env.NETWORK_NAME,
        body: req.body.message.message,
    });

    const response = await fetch(`http://${process.env.VTAN_ADDRESS}/icn/vtan/get/license`, {
        method: "POST",
        body: JSON.stringify({ payload }),
        headers: { "Content-Type": "application/json" },
    });

    let resp = await response.json();

    resp = rsaDecrypt(resp);

    payload = await rsaEncryptReturn(req.body.payload.pubkey, {
        icl: process.env.NETWORK_NAME,
        body: resp.message,
    });

    // Debugging Response
    console.log("API/ICN/TEST-S", resp);

    if (!resp.verified) {
        res.status(400).send({ response: "ERROR! ICN Error!" });
    }

    res.status(200).send({
        response: payload,
    });
});

router.post("/api/icn/vtan/get/vehicle", rsaDecryptMiddleware, async (req, res) => {
    if (!req.body.message.verified) {
        res.status(400).send({ response: "ERROR! ICL Error!" });
    }

    let payload = await rsaEncryptSend(process.env.VTAN_ADDRESS, {
        icl: process.env.NETWORK_NAME,
        body: req.body.message.message,
    });

    const response = await fetch(`http://${process.env.VTAN_ADDRESS}/icn/vtan/get/vehicle`, {
        method: "POST",
        body: JSON.stringify({ payload }),
        headers: { "Content-Type": "application/json" },
    });

    let resp = await response.json();

    resp = rsaDecrypt(resp);

    payload = await rsaEncryptReturn(req.body.payload.pubkey, {
        icl: process.env.NETWORK_NAME,
        body: resp.message,
    });

    // Debugging Response
    console.log("API/ICN/TEST-S", resp);

    if (!resp.verified) {
        res.status(400).send({ response: "ERROR! ICN Error!" });
    }

    res.status(200).send({
        response: payload,
    });
});

module.exports = router;
