import React, { Component } from "react";
import { Link } from "react-router-dom";
import { Table, TableBody, TableContainer, TableHead, TableCell, TableRow } from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        judgement: {},
        message: "",
        ID: "",
    };

    async componentDidMount() {
        const { id } = this.props.match.params;
        if (id !== "0") {
            this.setState({
                ID: id,
                message: (
                    <p>
                        Press <b>LOAD INVESTIGATION</b> to View Investigation
                    </p>
                ),
            });
        }
    }

    onLoadInvestigation = async () => {
        console.log(this.state.judgement);
        const requestOptions = {
            method: "GET",
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

        let response = await fetch("http://192.168.1.30:3000/api/main/judgement/read/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ judgement: res });

        var output = (
            <div>
                <Link to={"/addEvidenceID/" + this.state.judgement.ID}>Add Evidence ID</Link>
                <br />
                <Link to={"/addSentence/" + this.state.judgement.ID}>Add Sentence</Link>
                <br />
                <Link to={"/setComplete/" + this.state.judgement.ID}>Set Complete</Link>
                <br />
                {this.createContent()}
            </div>
        );

        this.setState({ message: output });
    };

    createContent = () => {
        return (
            <div>
                <h3>Judgement ID: {this.state.judgement.ID}</h3>
                <h2>{this.state.judgement.Complete ? "Judgement is Complete!" : "Judgement is Under Process."}</h2>
                <TextField
                    className="inputs"
                    label="Aggravating Mitigating Circumstances"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Deliberations.AggravatingMitigatingCircumstances}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Statutory Laws"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.ApplicableLaw.StatutoryLaws}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Summary Of Defendants Case"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Introduction.SummaryOfDefendantsCase}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Issues To Be Determined"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Introduction.IssuesToBeDetermined}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Summary Of Prosecutions Case"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Introduction.SummaryOfProsecutionsCase}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Case Laws"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.ApplicableLaw.CaseLaws}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Preliminary Issues"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Introduction.PreliminaryIssues}
                    readOnly
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Guilt"
                    variant="outlined"
                    multiline
                    rows={5}
                    value={this.state.judgement.Deliberations.Guilt}
                    readOnly
                />
                <h3>
                    Evidence:{" "}
                    {this.state.judgement.Evidence
                        ? this.state.judgement.Evidence.map((content, index) => {
                              return <span>{content + "; "}</span>;
                          })
                        : ""}
                </h3>
                <TableContainer component={Paper}>
                    <Table aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="left">
                                    <b>Citizen ID</b>
                                </TableCell>
                                <TableCell align="left">
                                    <b>Sentence Statement</b>
                                </TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.state.judgement.Deliberations.Sentence
                                ? this.state.judgement.Deliberations.Sentence.map((content, index) => {
                                      return (
                                          <TableRow key={content.CitizenID}>
                                              <TableCell align="left">{content.CitizenID}</TableCell>
                                              <TableCell align="left">{content.Statement}</TableCell>
                                          </TableRow>
                                      );
                                  })
                                : null}
                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>View / Edit Judgement</h2>
                <TextField
                    className="inputs"
                    label="Judgement ID"
                    variant="outlined"
                    value={this.state.ID}
                    onChange={(event) => {
                        this.setState({
                            ID: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button onClick={this.onLoadInvestigation} variant="contained" color="primary">
                    Load Judgement
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
