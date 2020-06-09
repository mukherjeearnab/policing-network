const express = require("express");
const md5 = require("md5");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Judgement = require("../../fabric/judgement_cc");

const router = new express.Router();

router.get("/api/main/judgement/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Judgement.ReadJudgement(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Judgement NOT found!" });
    }
});

router.post("/api/main/judgement/add", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = md5(JSON.stringify(judgementData) + new Date().toString());
        await Judgement.AddJudgement(req.user, judgementData);
        res.status(200).send({
            message: "Judgement has been successfully added!",
            id: judgementData.ID,
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Error! Judgement NOT Added!" });
    }
});

router.post("/api/main/judgement/addevidence/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = ID;
        await Judgement.AddEvidence(req.user, judgementData);
        res.status(200).send({ message: "Evidence has been Successfully Added to the Judgement Report!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Judgement NOT found!" });
    }
});

router.post("/api/main/judgement/addsentence/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        judgementData = JSON.parse(req.body.payload);
        judgementData.ID = ID;
        await Judgement.AddSentence(req.user, judgementData);
        res.status(200).send({ message: "Sentence has been Successfully Added to the Judgement Report!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Judgement NOT found!" });
    }
});

router.post("/api/main/judgement/setcomplete/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const ID = req.params.id;
        await Judgement.SetComplete(req.user, ID);
        res.status(200).send({ message: "Judgement has been Successfully Set to Complete!" });
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Judgement NOT found!" });
    }
});

module.exports = router;
