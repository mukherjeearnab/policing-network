const ipfsAPI = require("ipfs-api");
const express = require("express");
const path = require("path");
const multer = require("multer");
const fs = require("fs");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Evidence = require("../../fabric/evidence_cc");

const router = new express.Router();
const ipfs = ipfsAPI("ipfs.infura.io", "5001", { protocol: "https" });
const uploadPath = path.join(process.cwd(), "uploads");
var upload = multer({ dest: uploadPath });

router.get("/api/main/evidence/read/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Evidence.ReadEvidence(req.user, ID);
        res.status(200).send(data);
    } catch (error) {
        console.log(error);
        res.status(404).send({ message: "Evidence NOT found!" });
    }
});

router.post("/api/main/evidence/add", upload.single("file"), JWTmiddleware, (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const oldname = "uploads/" + req.file.filename;
        const newname = "uploads/" + req.file.filename + "." + req.file.originalname.split(".").pop();
        fs.renameSync(oldname, newname, console.log);

        let file = fs.readFileSync(newname);
        let fileBuffer = new Buffer(file);

        ipfs.files.add(fileBuffer, (err, file) => {
            if (err) {
                console.log(err);
            }
            evidenceData = JSON.parse(req.body.payload);
            evidenceData.ID = file[0].path;
            evidenceData.MimeType = req.file.mimetype;
            evidenceData.Extention = req.file.originalname.split(".").pop();
            evidenceData.DateTime = Math.floor(new Date() / 1000).toString();
            Evidence.AddEvidence(req.user, evidenceData).then(() => {
                fs.unlinkSync(newname);
                res.status(200).send({
                    message: "Evidence has been successfully added!",
                    hash: evidenceData,
                });
            });
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({ message: "Evidence NOT Added!" });
    }
});

module.exports = router;
