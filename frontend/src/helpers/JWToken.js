var JWT = (function () {
    var token = "";

    var get = function () {
        return token; // Or pull this from cookie/localStorage
    };

    var set = function (t) {
        token = t;
        // Also set this in cookie/localStorage
    };

    return {
        get: get,
        set: set,
    };
})();

export default JWT;
