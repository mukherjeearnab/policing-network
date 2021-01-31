const ReadChargeSheet = require("./readChargeSheet");
const AddChargeSheet = require("./addChargeSheet");
const AddAccusedPerson = require("./addAccusedPerson");
const AddBriefReport = require("./addBriefReport");
const AddChargedPerson = require("./addChargedPerson");
const AddFIRID = require("./addFIRID");
const AddInvestigatingOfficer = require("./addInvestigatingOfficer");
const AddInvestigationID = require("./addInvestigationID");
const AddSectionOfLaw = require("./addSectionOfLaw");

const payload = {
    ReadChargeSheet,
    AddChargeSheet,
    AddAccusedPerson,
    AddBriefReport,
    AddChargedPerson,
    AddFIRID,
    AddInvestigatingOfficer,
    AddInvestigationID,
    AddSectionOfLaw,
};

module.exports = payload;
