const AddInvestigation = require("./addInvestigation");
const AddReport = require("./addReport");
const AddArrest = require("./addArrest");
const ReadInvestigation = require("./readInvestigation");
const QueryInvestigation = require("./queryInvestigation");
const UpdateInvestigation = require("./updateInvestigation");

const payload = {
    AddInvestigation,
    AddReport,
    AddArrest,
    ReadInvestigation,
    QueryInvestigation,
    UpdateInvestigation,
};

module.exports = payload;
