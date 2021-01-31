import React, { Component } from "react";
import { Button } from "@material-ui/core";
import { Link, Redirect } from "react-router-dom";

class App extends Component {
    state = {
        redirect: "",
    };

    logout = () => {
        localStorage.removeItem("session");
        localStorage.removeItem("user");
        this.setState({ redirect: <Redirect to="/" /> });
    };

    render() {
        return (
            <div>
                <h2>Forensics Dashboard</h2>
                <h2>
                    {this.state.redirect}Welcome, {localStorage.getItem("user")}!
                </h2>
                <Link to="/submitEvidence">Add Evidence</Link> <br />
                <Link to="/evidenceViewer">View Evidence</Link> <br />
                <Link to="/viewInvestigation/0">View Investigation</Link> <br />
                <Button m={1} onClick={this.logout} variant="contained" color="primary">
                    Log Out
                </Button>
            </div>
        );
    }
}

export default App;
