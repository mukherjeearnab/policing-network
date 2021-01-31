const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Investigation = require("../../fabric/investigation_cc");

const router = new express.Router();

router.get("/api/main/investigation/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Investigation.ReadInvestigation(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Investigation NOT found!" });
    }
});

router.post("/api/main/investigation/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        InvestigationData = JSON.parse(req.body.payload);
        InvestigationData.ID = md5(JSON.stringify(InvestigationData) + new Date().toString());
        await Investigation.AddInvestigation(req.user, InvestigationData);
        res.status(200).send({
            message: "Investigation has been successfully added!",
            id: InvestigationData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Investigation NOT Added!" });
    }
});

router.post("/api/main/investigation/update/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        InvestigationData = JSON.parse(req.body.payload);
        InvestigationData.ID = ID;
        await Investigation.UpdateInvestigation(req.user, InvestigationData);
        res.status(200).send({
            message: "Investigation has been successfully Updated!",
            id: InvestigationData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Investigation NOT Added!" });
    }
});

router.post("/api/main/investigation/addreport/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        InvestigationData = JSON.parse(req.body.payload);
        InvestigationData.ID = ID;
        await Investigation.AddReport(req.user, InvestigationData);
        res.status(200).send({
            message: "Investigation Report has been successfully Added!",
            id: InvestigationData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Investigation Report NOT Added!" });
    }
});

router.post("/api/main/investigation/addarrest/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        InvestigationData = JSON.parse(req.body.payload);
        InvestigationData.ID = ID;
        await Investigation.AddArrest(req.user, InvestigationData);
        res.status(200).send({
            message: "Investigation Arrest has been successfully Added!",
            id: InvestigationData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Investigation Arrest NOT Added!" });
    }
});

module.exports = router;
