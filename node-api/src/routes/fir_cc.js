const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const FIR = require("../../fabric/fir_cc");

const router = new express.Router();

router.get("/api/main/fir/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await FIR.ReadFIR(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "FIR NOT found!" });
    }
});

router.post("/api/main/fir/query", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        FIRData = JSON.parse(req.body.payload);
        let data = await FIR.QueryFIR(req.user, FIRData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "FIR NOT found!" });
    }
});

router.post("/api/main/fir/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        FIRData = JSON.parse(req.body.payload);
        FIRData.CitizenID = req.user.username;
        FIRData.ID = md5(JSON.stringify(FIRData) + new Date().toString());
        await FIR.AddFIR(req.user, FIRData);
        res.status(200).send({
            message: "FIR has been successfully added!",
            id: FIRData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! FIR NOT Added!" });
    }
});

module.exports = router;
