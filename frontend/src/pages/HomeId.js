import React, { Component } from "react";
import { Button } from "@material-ui/core";
import { Link, Redirect } from "react-router-dom";

class App extends Component {
    state = {};

    logout = () => {
        localStorage.removeItem("session");
        localStorage.removeItem("user");
        this.setState({ redirect: <Redirect to="/" /> });
    };

    render() {
        return (
            <div>
                <h2>Identity Provider's Dashboard</h2>
                <h2>
                    {this.state.redirect}Welcome, {localStorage.getItem("user")}!
                </h2>
                <Link to="/viewProfile/0">Check Citizen Profile</Link> <br />
                <Link to="/newCitizen">Create New Citizen Profile</Link> <br />
                <Link to="/updateCitizen">Update Citizen Profile</Link> <br />
                <Button m={1} onClick={this.logout} variant="contained" color="primary">
                    Log Out
                </Button>
            </div>
        );
    }
}

export default App;
