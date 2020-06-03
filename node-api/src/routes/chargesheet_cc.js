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
        ChargeSheetData.ID = md5(ChargeSheetData);
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

router.get("/api/main/chargesheet/addaccused/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddAccusedPerson(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addreport/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddBriefReport(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addcharged/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddChargedPerson(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addfirid/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddFIRID(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addlaw/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddSectionOfLaw(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addofficer/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddInvestigatingOfficer(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

router.get("/api/main/chargesheet/addinvestigation/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        ChargeSheetData = JSON.parse(req.body.payload);
        ChargeSheetData.ID = ID;
        let data = await ChargeSheet.AddInvestigationID(req.user, ChargeSheetData);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "ChargeSheet NOT found!" });
    }
});

module.exports = router;
