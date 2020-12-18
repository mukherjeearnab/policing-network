"use strict";

module.exports.info = "Create and Read FIR";

const contractID = "fir_cc";
const version = "1.0";

let bc, ctx, clientArgs, clientIdx;

module.exports.init = async function (blockchain, context, args) {
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    clientIdx = context.clientIdx.toString();
    for (let i = 0; i < clientArgs.assets; i++) {
        try {
            const assetID = `FIR_${clientIdx}_${i}_${clientArgs.seed}`;
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
            await bc.bcObj.invokeSmartContract(ctx, contractID, version, myArgs);
        } catch (error) {
            console.log(`Client ${clientIdx}: Smart Contract threw with error: ${error}`);
        }
    }
};

module.exports.run = function () {
    const randomId = Math.floor(Math.random() * clientArgs.assets);
    const myArgs = {
        chaincodeFunction: "readFIR",
        invokerIdentity: "Admin@citizen.example.com",
        chaincodeArguments: [`FIR_${clientIdx}_${randomId}_${clientArgs.seed}`],
    };
    return bc.bcObj.querySmartContract(ctx, contractID, version, myArgs);
};

module.exports.end = async function () {};
