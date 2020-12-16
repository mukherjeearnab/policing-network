"use strict";

module.exports.info = "Create and Read Evidences";

const contractID = "evidence_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
    for (let i = 0; i < clientArgs.assets; i++) {
        try {
            const assetID = `${clientIdx}_${i}`;
            console.log(`Client ${clientIdx}: Creating evidence ${assetID}`);
            const myArgs = {
                chaincodeFunction: "addEvidence",
                invokerIdentity: "Admin@org1.example.com",
                chaincodeArguments: [assetID, "jpeg", ".jpg", "An image.", "10000", "af34bce1"],
            };
            await bc.bcObj.invokeSmartContract(ctx, contractID, version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    }
};

module.exports.run = function () {
    const randomId = Math.floor(Math.random() * clientArgs.assets);
    const myArgs = {
        chaincodeFunction: "readEvidence",
        invokerIdentity: "Admin@org1.example.com",
        chaincodeArguments: [`${clientIdx}_${randomId}`],
    };
    return bc.bcObj.querySmartContract(ctx, contractID, version, myArgs);
};

module.exports.end = async function () {};
