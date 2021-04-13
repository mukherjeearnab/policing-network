const fs = require("fs");
const NodeRSA = require("node-rsa");

decrypt = (response) => {
    // Import Sender's PubKey
    const keyDataS = response.pubkey;
    const keyS = new NodeRSA();
    keyS.importKey(keyDataS, "pkcs8-public-pem");

    // Import Private Key of Client SDK
    const keyDataR = fs.readFileSync("./keys/private.pem");
    const keyR = new NodeRSA();
    keyR.importKey(keyDataR, "pkcs8-pem");

    let res = { message: "", verified: false };

    // Decrypt Payload
    const decrypted = keyR.decrypt(response.body, "base64");

    // Convert payload from Base64 to UTF-8 to JS Object
    const message = Buffer.from(decrypted, "base64").toString("utf8");
    res.message = JSON.parse(message);

    // Verify Payload Signature
    const signatureResult = keyS.verify(decrypted, response.signature, "base64", "base64");
    res.verified = signatureResult;

    return res;
};

module.exports = decrypt;
