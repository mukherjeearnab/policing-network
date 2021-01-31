const { FileSystemWallet, Gateway } = require("fabric-network");
const path = require("path");

UpdateCitizen = async (user, payload) => {
    const ccp = require(`../ccp/connection-${user.group}.json`);
    const walletPath = path.join(process.cwd(), `wallet_${user.group}`);
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: user.username,
        discovery: { enabled: true, asLocalhost: true },
    });

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork("mainchannel");

    // Get the contract from the network.
    const contract = network.getContract("citizenprofile_cc");

    // Evaluate the specified transaction.
    await contract.submitTransaction(
        "updateCitizenProfile",
        payload.ID,
        payload.Photo,
        payload.Name,
        payload.Email,
        payload.Phone,
        payload.DOB,
        payload.Gender,
        payload.BloodGroup,
        payload.EyeColor,
        payload.Nationality,
        payload.Address,
        payload.FathersName,
        payload.MothersName,
        payload.Religion,
        payload.Occupation,
        payload.Fingerprint[0],
        payload.Fingerprint[1],
        payload.Fingerprint[2],
        payload.Fingerprint[3],
        payload.Fingerprint[4],
        payload.Fingerprint[5],
        payload.Fingerprint[6],
        payload.Fingerprint[7],
        payload.Fingerprint[8],
        payload.Fingerprint[9]
    );
};

module.exports = UpdateCitizen;
