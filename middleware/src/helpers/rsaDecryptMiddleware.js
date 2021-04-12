const fs = require("fs");
const NodeRSA = require("node-rsa");

middleware = (req, res, next) => {
    // Import Sender's PubKey
    const keyDataS = req.body.payload.pubkey;
    const keyS = new NodeRSA();
    keyS.importKey(keyDataS, "pkcs8-public-pem");

    // Import Private Key of Client SDK
    const keyDataR = fs.readFileSync("./keys/private.pem");
    const keyR = new NodeRSA();
    keyR.importKey(keyDataR, "pkcs8-pem");

    // Decrypt Payload
    const decrypted = keyR.decrypt(req.body.payload.body, "base64");

    // Convert payload from Base64 to UTF-8 to JS Object
    const message = Buffer.from(decrypted, "base64").toString("utf8");
    req.body.message = JSON.parse(message);

    // Verify Payload Signature
    const signatureResult = keyS.verify(decrypted, req.body.payload.signature, "base64", "base64");
    req.body.verified = signatureResult;

    next();
};

module.exports = middleware;
