import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        EvidenceID: "",
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
            "http://192.168.1.30:3000/api/main/judgement/setcomplete/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to={"/viewJudgement/" + this.state.ID} /> });
    };

    render() {
        return (
            <div>
                <h2>Set Complete Status to Judgement</h2>
                {this.state.message}

                <h1>Judgement ID - {this.state.ID}</h1>

                <br />
                <br />
                <Button onClick={this.onAddReport} variant="contained" color="primary">
                    Set Complete!
                </Button>
            </div>
        );
    }
}

export default App;
