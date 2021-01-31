const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const ChargeSheet = require("../../fabric/chargesheet_cc");

const router = new express.Router();

router.get("/api/main/chargesheet/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await ChargeSheet.ReadChargeSheet(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.DateTime = Math.floor(new Date() / 1000).toString();
        ChargeSheetData.ID = md5(JSON.stringify(ChargeSheetData) + new Date().toString());
        await ChargeSheet.AddChargeSheet(req.user, ChargeSheetData);
        res.status(200).send({
            message: "ChargeSheet has been successfully added!",
            id: ChargeSheetData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! ChargeSheet NOT Added!" });
    }
});

router.post("/api/main/chargesheet/addaccused/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddAccusedPerson(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addreport/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddBriefReport(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addcharged/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddChargedPerson(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addfirid/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddFIRID(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addlaw/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddSectionOfLaw(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addofficer/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddInvestigatingOfficer(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.post("/api/main/chargesheet/addinvestigation/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        await ChargeSheet.AddInvestigationID(req.user, ChargeSheetData);
        res.status(200).send(ChargeSheetData);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

module.exports = router;
