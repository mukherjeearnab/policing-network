import React, { Component } from "react";
import { Table, TableBody, TableContainer, TableHead, TableCell, TableRow, CircularProgress } from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import { Link, Redirect } from "react-router-dom";

class App extends Component {
    state = {
        firs: [],
        message: "",
    };

    async componentDidMount() {
        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        if (!localStorage.getItem("session")) this.setState({ redirect: <Redirect to="/" /> });
        if (localStorage.getItem("session")) {
            const requestOptions = {
                method: "GET",
                headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            };
            let response = await fetch("http://192.168.1.30:3000/api/auth/verify/", requestOptions);
            let res = await response.json();
            if (res.status === 0) this.setState({ redirect: <Redirect to="/" /> });
        }
        console.log(localStorage.getItem("session"));
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify({
                    CitizenID: localStorage.getItem("user"),
                    PoliceStation: "",
                }),
            }),
        };
        console.log(requestOptions);
        let response = await fetch("http://192.168.1.30:3000/api/main/fir/query/", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ firs: res, message: "" });
    }

    render() {
        return (
            <div>
                <h1>Past FIRs</h1>
                <div>
                    <TableContainer component={Paper}>
                        <Table aria-label="simple table">
                            <TableHead>
                                <TableRow>
                                    <TableCell align="left">
                                        <b>FIR ID</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Police Station</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Place Of Occurence</b>
                                    </TableCell>
                                    <TableCell align="left">
                                        <b>Date & Time</b>
                                    </TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {this.state.firs
                                    .slice(0)
                                    .reverse()
                                    .map((content, index) => {
                                        content = content.Value;
                                        let date = new Date(content.DateHour).toString();

                                        return (
                                            <TableRow key={content.ID}>
                                                <TableCell align="left">
                                                    <Link to={`/firViewer/${content.ID}`}>{content.ID}</Link>
                                                </TableCell>
                                                <TableCell align="left">{content.PoliceStation}</TableCell>
                                                <TableCell align="left">{content.PlaceOfOccurence}</TableCell>
                                                <TableCell align="left">{date}</TableCell>
                                            </TableRow>
                                        );
                                    })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </div>
                <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
