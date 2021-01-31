import React, { Component } from "react";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        judgement: {
            AggravatingMitigatingCircumstances: "",
            StatutoryLaws: "",
            SummaryOfDefendantsCase: "",
            IssuesToBeDetermined: "",
            SummaryOfProsecutionsCase: "",
            CaseLaws: "",
            PreliminaryIssues: "",
            Guilt: "",
        },
        message: "",
    };

    onAddFIR = async () => {
        console.log(this.state.judgement);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify(this.state.judgement),
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

        let response = await fetch("http://192.168.1.30:3000/api/main/judgement/add", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: "Judgement ID : " + res.id });
    };

    render() {
        return (
            <div>
                <h2>New Investigation</h2>
                {this.state.message}
                <br />
                <TextField
                    className="inputs"
                    label="Aggravating Mitigating Circumstances"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.AggravatingMitigatingCircumstances}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.AggravatingMitigatingCircumstances = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Statutory Laws"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.StatutoryLaws}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.StatutoryLaws = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Summary Of Defendants Case"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.SummaryOfDefendantsCase}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.SummaryOfDefendantsCase = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Issues To Be Determined"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.IssuesToBeDetermined}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.IssuesToBeDetermined = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Summary Of Prosecutions Case"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.SummaryOfProsecutionsCase}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.SummaryOfProsecutionsCase = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Case Laws"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.CaseLaws}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.CaseLaws = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Preliminary Issues"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.PreliminaryIssues}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.PreliminaryIssues = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Guilt"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Guilt}
                    onChange={(event) => {
                        let judgement = this.state.judgement;
                        judgement.Guilt = event.target.value;
                        this.setState({
                            judgement,
                        });
                    }}
                />
                <br />
                <br />
                <Button onClick={this.onAddFIR} variant="contained" color="primary">
                    Init. Judgement
                </Button>
            </div>
        );
    }
}

export default App;
