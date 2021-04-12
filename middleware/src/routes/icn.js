const express = require("express");
const router = new express.Router();
const fs = require("fs");

const fetch = require("node-fetch");

const rsaEncrypt = require("../helpers/rsaEncrypt");
const rsaDecryptMiddleware = require("../helpers/rsaDecryptMiddleware");

router.get("/api/icn/pubkey", (req, res) => {
    // Import Public Key
    // Import Private Key of Client SDK
    const keyDataR = fs.readFileSync("./keys/public.pem");

    res.status(200).send({
        pubKey: keyDataR.toString(),
    });
});

router.get("/api/icn/testS", async (req, res) => {
    let payload = await rsaEncrypt("localhost:3000", { name: "Arnab", age: "21" });

    const response = await fetch(`http://localhost:3000/api/icn/testR`, {
        method: "POST",
        body: JSON.stringify({ payload }),
        headers: { "Content-Type": "application/json" },
    });

    let resp = await response.json();
    console.log("API/ICN/TEST-S", resp);

    res.status(200).send({
        response: resp,
    });
});

router.post("/api/icn/testR", rsaDecryptMiddleware, async (req, res) => {
    console.log("API/ICN/TEST-T", req.body);

    res.status(200).send({
        message: req.body,
    });
});

module.exports = router;
