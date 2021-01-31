import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        Content: "",
        Date: "",
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
                    Date: new Date(this.state.Date).getTime().toString(),
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
            "http://192.168.1.30:3000/api/main/investigation/addreport/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewInvestigation/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Add Brief Report to Investigation</h2>
                {this.state.message}

                <h1>Investigation ID - {this.state.ID}</h1>

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
                <TextField
                    type="date"
                    className="inputs"
                    label="Report Date"
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

                <Button onClick={this.onAddReport} variant="contained" color="primary">
                    Add Brief Report to Investigation
                </Button>
            </div>
        );
    }
}

export default App;
