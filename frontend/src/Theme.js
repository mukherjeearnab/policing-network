import { createMuiTheme } from "@material-ui/core";

//import "typeface-roboto";

const Theme = createMuiTheme({
    palette: {
        type: "light",
        spacing: 10,
        common: {
            black: "#282C33",
            white: "#f5f5f5",
        },
        background: {
            paper: "#f5f5f5",
            default: "#fff",
        },
        primary: {
            light: "#2f89fc",
            main: "#2f89fc",
            dark: "#2f89fc",
            contrastText: "#f5f5f5",
        },
        secondary: {
            light: "#30e3ca",
            main: "#30e3ca",
            dark: "#30e3ca",
            contrastText: "#f5f5f5",
        },
        error: {
            light: "#e57373",
            main: "#f44336",
            dark: "#d32f2f",
            contrastText: "#fff",
        },
        text: {
            primary: "#282C33",
            secondary: "#282C33",
            disabled: "#282C33",
            hint: "#282C33",
        },
    },
});

export default Theme;
