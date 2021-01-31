import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        Content: "",
        Document: "XXXXX",
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
                    Content: this.state.Content,
                    Document: this.state.Document,
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
            "http://192.168.1.30:3000/api/main/chargesheet/addreport/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewChargeSheet/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Add Brief Report to Charge Sheet</h2>
                {this.state.message}

                <h1>Charge Sheet ID - {this.state.ID}</h1>

                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Report Content"
                    variant="outlined"
                    multiline
                    rows={10}
                    value={this.state.Content}
                    onChange={(event) => {
                        this.setState({
                            Content: event.target.value,
                        });
                    }}
                />
                <br />
                <br />
                <Button onClick={this.onAddReport} variant="contained" color="primary">
                    Add Brief Report
                </Button>
            </div>
        );
    }
}

export default App;
