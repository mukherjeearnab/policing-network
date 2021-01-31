import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        CitizenID: "",
        Cause: "",
        Date: "",
        Mugshot: "XXXXX",
        ID: "",
        message: "",
    };

    componentDidMount() {
        const { id } = this.props.match.params;
        this.setState({ ID: id });
    }

    onAddArrest = async () => {
        console.log(this.state.investigation);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify({
                    CitizenID: this.state.CitizenID,
                    Cause: this.state.Cause,
                    Date: new Date(this.state.Date).getTime().toString(),
                    Mugshot: this.state.Mugshot,
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
            "http://192.168.1.30:3000/api/main/investigation/addarrest/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewInvestigation/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Add Arrest to Investigation</h2>
                {this.state.message}

                <h1>Investigation ID - {this.state.ID}</h1>

                <TextField
                    className="inputs"
                    label="Citizen ID"
                    variant="outlined"
                    value={this.state.CitizenID}
                    onChange={(event) => {
                        this.setState({
                            CitizenID: event.target.value,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Arrest Cause"
                    variant="outlined"
                    value={this.state.Cause}
                    onChange={(event) => {
                        this.setState({
                            Cause: event.target.value,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    type="date"
                    className="inputs"
                    label="Arrest Date"
                    variant="outlined"
                    value={this.state.Date}
                    onChange={(event) => {
                        this.setState({
                            Date: event.target.value,
                        });
                    }}
                />
                <br />
                <br />

                <Button onClick={this.onAddArrest} variant="contained" color="primary">
                    Add Arrest to Investigation
                </Button>
            </div>
        );
    }
}

export default App;
