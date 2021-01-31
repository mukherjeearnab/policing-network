import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        InvestigationID: "",
        ID: "",
        message: "",
    };

    componentDidMount() {
        const { id } = this.props.match.params;
        this.setState({ ID: id });
    }

    onAddReport = async () => {
        console.log(this.state.investigation);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify({
                    InvestigationID: this.state.InvestigationID,
                }),
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

        let response = await fetch(
            "http://192.168.1.30:3000/api/main/chargesheet/addinvestigation/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewChargeSheet/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Add Investigation ID to Charge Sheet</h2>
                {this.state.message}

                <h1>Charge Sheet ID - {this.state.ID}</h1>

                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Investigation ID"
                    variant="outlined"
                    value={this.state.InvestigationID}
                    onChange={(event) => {
                        this.setState({
                            InvestigationID: event.target.value,
                        });
                    }}
                />
                <br />
                <br />
                <Button onClick={this.onAddReport} variant="contained" color="primary">
                    Add Investigation ID
                </Button>
            </div>
        );
    }
}

export default App;
