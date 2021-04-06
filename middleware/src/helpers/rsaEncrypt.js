const fetch = require("node-fetch");
const fs = require("fs");
const NodeRSA = require("node-rsa");

// Import Private Key of Client SDK
const keyDataS = fs.readFileSync("./icn/private.pem");
const keyS = new NodeRSA();
keyS.importKey(keyDataS, "pkcs8-pem");

encrypt = async (receipient, payload) => {
    // Import Public Key of Receipient
    const response = await fetch(`http://${receipient}/api/icn/pubkey`);
    const { pubKey } = await response.json();
    const keyR = new NodeRSA();
    keyR.importKey(pubKey, "pkcs8-public-pem");

    // Encrypt Payload
    const encrypted = keyR.encrypt(JSON.stringify(payload), "base64");

    // Generate Payload Signature
    const signature = keyS.sign(payload, "base64");

    // Create JSON output
    const encryptedPayload = { body: encrypted, signature: signature, pubkey: keyS.exportKey("pkcs8-public-pem") };

    return encryptedPayload;
};

module.exports = encrypt;
