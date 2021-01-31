const jwt = require("jsonwebtoken");
const config = require("./jwtConfig");

middleware = (req, res, next) => {
    res.setHeader("Access-Control-Allow-Origin", "*");

    let token = req.headers["x-access-token"];
    if (token) {
        jwt.verify(
            token,
            config.secretKey,
            {
                algorithm: config.algorithm,
            },
            function (err, decoded) {
                if (err) {
                    return res.status(401).send({
                        message: "Unauthorised access!",
                    });
                }
                req.user = decoded;
                next();
            }
        );
    } else {
        return res.status(403).send({
            message: "Forbidden access!",
        });
    }
};

module.exports = middleware;
