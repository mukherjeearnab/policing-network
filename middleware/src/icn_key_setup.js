const fs = require("fs");
const NodeRSA = require("node-rsa");

// Generate New Key Pair
const key = new NodeRSA({ b: 512 });

// Export Key Pair
const publicPem = key.exportKey("pkcs8-public-pem");
const privatePem = key.exportKey("pkcs8-pem");

console.log("PubKey", publicPem);
console.log("PrivKey", privatePem);

// Save Key Pair as pem files
fs.writeFileSync(`./keys/private.pem`, privatePem);
fs.writeFileSync(`./keys/public.pem`, publicPem);
