import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        FIRID: "",
        Officer: "",
        ID: "",
        message: "",
    };

    async componentDidMount() {
        const { id } = this.props.match.params;

        console.log(localStorage.getItem("session"));
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };
        //console.log(requestOptions);
        let response = await fetch("http://192.168.1.30:3000/api/main/investigation/read/" + id, requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ FIRID: res.FIRID, ID: res.ID, Officer: res.Officer });
    }

    onAddFIR = async () => {
        console.log(this.state.investigation);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify({ FIRID: this.state.FIRID, Officer: this.state.Officer }),
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
            "http://192.168.1.30:3000/api/main/investigation/update/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to="/HomePo" /> });
    };

    render() {
        return (
            <div>
                <h2>New Investigation</h2>
                {this.state.message}

                <h1>Investigation ID - {this.state.ID}</h1>

                <TextField
                    className="inputs"
                    label="FIR ID"
                    variant="outlined"
                    value={this.state.FIRID}
                    onChange={(event) => {
                        this.setState({
                            FIRID: event.target.value,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Investigating Officer"
                    variant="outlined"
                    value={this.state.Officer}
                    onChange={(event) => {
                        this.setState({
                            Officer: event.target.value,
                        });
                    }}
                />

                <Button onClick={this.onAddFIR} variant="contained" color="primary">
                    Update Investigation
                </Button>
            </div>
        );
    }
}

export default App;
