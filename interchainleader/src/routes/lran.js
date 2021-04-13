const express = require("express");
const fetch = require("node-fetch");

const rsaEncryptSend = require("../helpers/rsaEncryptSend");
const rsaEncryptReturn = require("../helpers/rsaEncryptReturn");
const rsaDecrypt = require("../helpers/rsaDecrypt");
const rsaDecryptMiddleware = require("../helpers/rsaDecryptMiddleware");

require("dotenv").config();
const router = new express.Router();

/**********************************
 *  LRAN ROUTES
 **********************************/

router.post("/api/icn/lran/get/land", rsaDecryptMiddleware, async (req, res) => {
    if (!req.bod.message.verified) {
        res.status(400).send({ response: "ERROR! ICL Error!" });
    }

    let payload = await rsaEncryptSend(process.env.LRAN_ADDRESS, {
        icl: process.env.NETWORK_NAME,
        body: req.body.message.message,
    });

    const response = await fetch(`http://${process.env.LRAN_ADDRESS}/icn/lran/get/land`, {
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
