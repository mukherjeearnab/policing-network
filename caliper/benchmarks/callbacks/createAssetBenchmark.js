"use strict";

module.exports.info = "Create FIRs";

const contractID = "fir_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
};

module.exports.run = async function () {
    for (let i = 0; i < clientArgs.assets; i++) {
        try {
            const assetID = `FIR_${clientIdx}_${i}_${Date.now()}`;
            console.log(`Client ${clientIdx}: Creating FIR ${assetID}`);
            const myArgs = {
                chaincodeFunction: "createNewFIR",
                invokerIdentity: "Admin@citizen.example.com",
                chaincodeArguments: [
                    assetID,
                    "citizen1",
                    "Bosnia",
                    "ZZZ",
                    "DUDEX",
                    "1000",
                    "XXX",
                    "ZZZ",
                    "XYZ",
                    "XCZ",
                    "XCV",
                ],
            };
            return await bc.bcObj.invokeSmartContract(ctx, contractID, version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    }
};

module.exports.end = async function () {};
