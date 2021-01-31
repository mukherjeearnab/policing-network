import React, { Component } from "react";
import { Button, TextField } from "@material-ui/core";
import { Link } from "react-router-dom";
import { Table, TableBody, TableContainer, TableHead, TableCell, TableRow, Paper } from "@material-ui/core";

class App extends Component {
    state = {
        redirect: "",
        ID: "",
        profile: {},
        output: null,
    };

    onComponentDidMount() {
        const { id } = this.props.match.params;
        if (id !== "0") {
            this.setState({
                ID: id,
                message: (
                    <p>
                        Press <b>LOAD Profile</b> to View Profile.
                    </p>
                ),
            });
        }
    }

    loadProfile = async () => {
        const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };
        let response = await fetch("http://192.168.1.30:3000/api/main/citizen/get/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ profile: res });

        let table = (
            <TableContainer component={Paper}>
                <Table aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell align="left">
                                <b>Serial</b>
                            </TableCell>
                            <TableCell align="left">
                                <b>Judgement ID</b>
                            </TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {this.state.profile.VerdictRecord
                            ? this.state.profile.VerdictRecord.map((content, index) => {
                                  return (
                                      <TableRow key={index}>
                                          <TableCell align="left">{(index + 1).toString()}</TableCell>
                                          <TableCell align="left">
                                              <Link target="blank" to={"/viewJudgement/" + content}>
                                                  {content}
                                              </Link>
                                          </TableCell>
                                      </TableRow>
                                  );
                              })
                            : ""}
                    </TableBody>
                </Table>
            </TableContainer>
        );

        let output = (
            <div>
                <img
                    alt="profile-pic"
                    width="200"
                    src={"https://ipfs.infura.io/ipfs/" + this.state.profile.Photo}
                ></img>
                <h3>Citizen Name: {this.state.profile.Name}</h3>
                <h3>Father's Name: {this.state.profile.FathersName}</h3>
                <h3>Mother's Name: {this.state.profile.MothersName}</h3>
                <h3>Religion: {this.state.profile.Religion}</h3>
                <h3>Phone: {this.state.profile.Phone}</h3>
                <h3>DOB: {new Date(this.state.profile.DOB).toString()}</h3>
                <h3>Gender: {this.state.profile.Gender}</h3>
                <h3>Blood Group: {this.state.profile.BloodGroup}</h3>
                <h3>Address: {this.state.profile.Address}</h3>
                <h3>Email: {this.state.profile.Email}</h3>
                <h3>Eye Color: {this.state.profile.EyeColor}</h3>
                <h3>Occupation: {this.state.profile.Occupation}</h3>
                <h3>Verdict Record</h3>
                <div>{table}</div>
            </div>
        );
        this.setState({ output });
    };

    render() {
        return (
            <div>
                <h2>View Citizen Profile</h2>
                <TextField
                    className="inputs"
                    label="Citizen ID"
                    variant="outlined"
                    value={this.state.ID}
                    onChange={(event) => {
                        this.setState({
                            ID: event.target.value,
                        });
                    }}
                />
                <br />
                <br />
                <Button m={1} onClick={this.loadProfile} variant="contained" color="primary">
                    Load Citizen
                </Button>
                <hr />
                {this.state.output}
            </div>
        );
    }
}

export default App;
