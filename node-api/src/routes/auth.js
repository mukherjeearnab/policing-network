const express = require("express");
const sha256 = require("sha256");
const jwt = require("jsonwebtoken");
const User = require("../models/user");
const JWTConfig = require("../helpers/jwtConfig");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const router = new express.Router();

router.get("/api/auth/login", (req, res) => {
    const username = req.body.username;
    const passhash = sha256(req.body.password);
    const group = req.body.group;
    try {
        User.findOne({ username }, (err, doc) => {
            if (err || doc == null) return res.sendStatus(404);
            if (doc.passhash === passhash) {
                let userdata = {
                    username,
                    passhash,
                    group,
                };
                let token = jwt.sign(userdata, JWTConfig.secretKey, {
                    algorithm: JWTConfig.algorithm,
                    expiresIn: "1m",
                });
                res.status(200).send({
                    message: "Login Successful!",
                    jwtoken: token,
                });
            } else {
                res.status(401).send({
                    message: "Login Failed!",
                });
            }
        });
    } catch {
        res.status(500).send({
            message: "Server Error!",
        });
    }
});

router.get("/api/auth/signup", JWTmiddleware, (req, res) => {
    const username = req.body.username;
    const passhash = sha256("1234");
    const group = req.body.group;
    try {
        let newUser = {
            username,
            passhash,
            group,
        };
        User.create(newUser, function (err, res) {
            console.log(err);
            res.status(200).send({
                message: "Sign Up Successful!",
            });
        });
    } catch {
        res.status(500).send({
            message: "Server Error!",
        });
    }
});

module.exports = router;
