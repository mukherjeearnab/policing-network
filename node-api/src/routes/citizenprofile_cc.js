const ipfsAPI = require("ipfs-api");
const express = require("express");
const multer = require("multer");
const fs = require("fs");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const Citizen = require("../../fabric/citizenprofile_cc");

const router = new express.Router();
const ipfs = ipfsAPI("ipfs.infura.io", "5001", { protocol: "https" });
const uploadPath = path.join(process.cwd(), "uploads");
var upload = multer({ dest: uploadPath });

router.get("/api/main/citizen/get/:id", JWTmiddleware, async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const ID = req.params.id;
    try {
        let data = await Citizen.ReadCitizen(req.user, ID);
        res.status(200).send(data);
    } catch {
        res.status(404).send({ message: "Citizen NOT found!" });
    }
});

router.post("/api/main/citizen/add", upload.single("file"), (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    try {
        const oldname = "uploads/" + req.file.filename;
        const newname = "uploads/" + req.file.filename + "." + req.file.originalname.split(".").pop();
        fs.renameSync(oldname, newname, console.log);

        let file = fs.readFileSync(newname);
        let fileBuffer = new Buffer(file);

        ipfs.files.add(fileBuffer, async (err, file) => {
            if (err) {
                console.log(err);
            }
            CitizenData = JSON.parse(req.body.payload);
            await Citizen.AddCitizen(req.user, CitizenData);
            fs.unlinkSync(newname);
            res.status(200).send({
                message: "Citizen has been successfully added!",
                payload: CitizenData,
            });
        });
    } catch {
        res.status(500).send({ message: "Citizen NOT Added!" });
    }
});

module.exports = router;
