const mongoose = require("mongoose");

const connectionString =
    "mongodb+srv://zulu:1234@cluster0-5jxp1.mongodb.net/psbc?retryWrites=true&w=majority";

mongoose.connect(connectionString, {
    useNewUrlParser: true,
    useCreateIndex: true,
    useFindAndModify: false,
    useUnifiedTopology: true,
});
