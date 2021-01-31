import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        CitizenID: "",
        Statement: "",
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
                    CitizenID: this.state.CitizenID,
                    Statement: this.state.Statement,
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
            "http://192.168.1.30:3000/api/main/judgement/addsentence/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewJudgement/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Add Sentence to Judgement</h2>
                {this.state.message}

                <h1>Judgement ID - {this.state.ID}</h1>

                <br />
                <br />
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
                    label="Statement"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.Statement}
                    onChange={(event) => {
                        this.setState({
                            Statement: event.target.value,
                        });
                    }}
                />
                <br />
                <br />

                <Button onClick={this.onAddReport} variant="contained" color="primary">
                    Add Sentence
                </Button>
            </div>
        );
    }
}

export default App;
