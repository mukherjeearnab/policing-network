const express = require("express");
const sha256 = require("sha256");
const jwt = require("jsonwebtoken");
const User = require("../models/user");
const JWTConfig = require("../helpers/jwtConfig");
const JWTmiddleware = require("../helpers/jwtVerifyMiddleware");
const router = new express.Router();

router.post("/api/auth/login", (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

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
    } catch (error) {
        console.log(error);
        res.status(500).send({
            message: "Server Error!",
        });
    }
});

router.post("/api/auth/signup", async (req, res) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    const username = req.body.username;
    const passhash = sha256(req.body.password);
    const group = req.body.group;
    try {
        let newUser = {
            username,
            passhash,
            group,
        };

        // Create Wallet Identity for the Username
        const regUser = require(`../../fabric/reg_user/reg-${newUser.group}`);
        await regUser(newUser);

        // Add username & passhash to the MongoDB Auth Database
        User.create(newUser, function (err, doc) {
            console.log(err);
            res.status(200).send({
                message: "Sign Up Successful!",
            });
        });
    } catch (error) {
        console.log(error);
        res.status(500).send({
            message: "Server Error!",
        });
    }
});

router.get("/api/auth/verify", (req, res) => {
    let token = req.headers["x-access-token"];
    if (token) {
        jwt.verify(
            token,
            JWTConfig.secretKey,
            {
                algorithm: JWTConfig.algorithm,
            },
            function (err, decoded) {
                if (err) {
                    res.status(401).send({
                        status: 0,
                    });
                }
                res.status(200).send({
                    status: 1,
                    username: decoded.username,
                    group: decoded.group,
                });
            }
        );
    } else {
        res.status(401).send({
            status: 0,
        });
    }
});

module.exports = router;
