const express = require("express");
const fs = require("fs");
const fetch = require("node-fetch");

const rsaEncrypt = require("../helpers/rsaEncrypt");
const rsaDecryptMiddleware = require("../helpers/rsaDecryptMiddleware");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");

require("dotenv").config();
const router = new express.Router();

/**********************************
 *  IAN ROUTES
 **********************************/

router.get("/api/icn/ian/get/citizen/:id", JWTmiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.ICL_ADDRESS, { user: req.user, body: { citizen_id: req.params.id } });

    const response = await fetch(`http://${process.env.ICL_ADDRESS}/icn/ian/get/citizen`, {
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

/**********************************
 *  VTAN ROUTES
 **********************************/

router.get("/api/icn/vtan/get/license/:id", JWTmiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.ICL_ADDRESS, { user: req.user, body: { license_id: req.params.id } });

    const response = await fetch(`http://${process.env.ICL_ADDRESS}/icn/vtan/get/license`, {
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

router.get("/api/icn/vtan/get/vehicle/:id", JWTmiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.ICL_ADDRESS, { user: req.user, body: { vehicle_id: req.params.id } });

    const response = await fetch(`http://${process.env.ICL_ADDRESS}/icn/vtan/get/vehicle`, {
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

/**********************************
 *  LRAN ROUTES
 **********************************/

router.get("/api/icn/lran/get/land/:id", JWTmiddleware, async (req, res) => {
    let payload = await rsaEncrypt(process.env.ICL_ADDRESS, { user: req.user, body: { land_id: req.params.id } });

    const response = await fetch(`http://${process.env.ICL_ADDRESS}/icn/lran/get/land`, {
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
