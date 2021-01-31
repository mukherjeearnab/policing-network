import React, { Component } from "react";
import { Link } from "react-router-dom";
import { Table, TableBody, TableContainer, TableHead, TableCell, TableRow } from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        investigation: {},
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
        console.log(this.state.investigation);
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

        let response = await fetch(
            "http://192.168.1.30:3000/api/main/investigation/read/" + this.state.ID,
            requestOptions
        );
        let res = await response.json();
        console.log(res);

        this.setState({ investigation: res });

        var output = (
            <div>
                <Link to={"/updateInvestigation/" + this.state.investigation.ID}>Update Investigation</Link> <br />
                <Link to={"/addArrest/" + this.state.investigation.ID}>Add Arrest to Investigation</Link> <br />
                <Link to={"/addiReport/" + this.state.investigation.ID}>Add Report to ID Investigation</Link> <br />
                {this.createContent()}
            </div>
        );

        this.setState({ message: output });
    };

    createContent = () => {
        let reports, arrests, evidence;

        if (this.state.investigation.Reports) {
            reports = (
                <div>
                    <TableContainer component={Paper}>
                        <Table aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="left">
                                        <b>Date & Time</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Report Content</b>
                                    </TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.investigation.Reports.map((content, index) => {
                                    let date = new Date(content.DateTime).toString();

                                    return (
                                        <TableRow key={content.DateTime}>
                                            <TableCell align="left">{date}</TableCell>
                                            <TableCell align="left">{content.Content}</TableCell>
                                        </TableRow>
                                    );
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </div>
            );
        }

        if (this.state.investigation.Arrests) {
            arrests = (
                <div>
                    <TableContainer component={Paper}>
                        <Table aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="left">
                                        <b>Date & Time</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Citizen ID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Cause</b>
                                    </TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.investigation.Arrests.map((content, index) => {
                                    let date = new Date(content.Date).toString();

                                    return (
                                        <TableRow key={content.Date}>
                                            <TableCell align="left">{date}</TableCell>
                                            <TableCell align="left">{content.CitizenID}</TableCell>
                                            <TableCell align="left">{content.Cause}</TableCell>
                                        </TableRow>
                                    );
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </div>
            );
        }

        if (this.state.investigation.Evidence) {
            evidence = (
                <div>
                    <TableContainer component={Paper}>
                        <Table aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="left">
                                        <b>Evidence ID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Action / View</b>
                                    </TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.investigation.Evidence.map((content, index) => {
                                    return (
                                        <TableRow key={content}>
                                            <TableCell align="left">{content}</TableCell>
                                            <TableCell align="left">
                                                <a target="blank" href={"https://ipfs.infura.io/ipfs/" + content}>
                                                    View Evidence
                                                </a>
                                            </TableCell>
                                            <TableCell align="left">{content.Cause}</TableCell>
                                        </TableRow>
                                    );
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </div>
            );
        }

        return (
            <div>
                <h3>Investigation ID: {this.state.investigation.ID}</h3>
                <h3>FIR ID: {this.state.investigation.FIRID}</h3>
                <h3>Investigation ID: {this.state.investigation.ID}</h3>
                <h3>Investigating Officer: {this.state.investigation.Officer}</h3>
                <h2>Investigation Reports</h2>
                {reports}
                {"e7de1ea556ef166b04dec961496ce613"}
                <h2>Investigation Arrests</h2>
                {arrests}
                <h2>Investigation Evidence</h2>
                {evidence}
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>New Investigation</h2>
                <TextField
                    className="inputs"
                    label="Investigation ID"
                    variant="outlined"
                    value={this.state.ID}
                    onChange={(event) => {
                        this.setState({
                            ID: event.target.value,
                        });
                    }}
                />
                <br /> <br />
                <Button
                    ref={(button) => {
                        this.LoadBT = button;
                    }}
                    onClick={this.onLoadInvestigation}
                    variant="contained"
                    color="primary"
                >
                    Load Investigation
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
