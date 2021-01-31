import React, { Component } from "react";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        investigation: {
            FIRID: "",
            Officer: "",
        },
        message: "",
    };

    onAddFIR = async () => {
        console.log(this.state.investigation);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify(this.state.investigation),
            }),
        };

        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        let response = await fetch("http://192.168.1.30:3000/api/main/investigation/add", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: "Investigation ID : " + res.id });
    };

    render() {
        return (
            <div>
                <h2>New Investigation</h2>
                {this.state.message}
                <br />
                <br />

                <TextField
                    className="inputs"
                    label="FIR ID"
                    variant="outlined"
                    value={this.state.investigation.FIRID}
                    onChange={(event) => {
                        let investigation = this.state.investigation;
                        investigation.FIRID = event.target.value;
                        this.setState({
                            investigation,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Investigating Officer"
                    variant="outlined"
                    value={this.state.investigation.Officer}
                    onChange={(event) => {
                        let investigation = this.state.investigation;
                        investigation.Officer = event.target.value;
                        this.setState({
                            investigation,
                        });
                    }}
                />
                <br />
                <br />

                <Button onClick={this.onAddFIR} variant="contained" color="primary">
                    Init. Investigation
                </Button>
            </div>
        );
    }
}

export default App;
