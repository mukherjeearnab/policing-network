const mongoose = require("mongoose");
const userJSON = require("./json/user-model.json");

const userInfo = new mongoose.Schema(userJSON);

const user = mongoose.model("userInfo", userInfo, "userinfo");

module.exports = user;
