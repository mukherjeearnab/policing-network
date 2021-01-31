import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        fir: {
            Complaint: "",
            DateHour: 0,
            DescriptionOfAccused: "",
            DetailsOfWitness: "",
            District: "",
            Nature: "",
            Particulars: "",
            PlaceOfOccurence: "",
            PoliceStation: "",
        },
        message: "",
    };

    onAddFIR = async () => {
        let fir = this.state.fir;
        fir.DateHour = new Date(this.state.fir.DateHour).getTime().toString();
        this.setState({ fir });
        console.log(this.state.fir);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
            body: JSON.stringify({
                payload: JSON.stringify(this.state.fir),
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

        let response = await fetch("http://192.168.1.30:3000/api/main/fir/add", requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ message: <Redirect to="/viewFIRs" /> });
    };

    render() {
        return (
            <div>
                <h2>New FIR</h2>
                {this.state.message}
                <TextField
                    className="inputs"
                    label="Police Station"
                    variant="outlined"
                    value={this.state.fir.PoliceStation}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.PoliceStation = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    type="date"
                    className="inputs"
                    label="Date"
                    variant="outlined"
                    value={this.state.fir.DateHour}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.DateHour = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="District"
                    variant="outlined"
                    value={this.state.fir.District}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.District = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Place Of Occurence"
                    variant="outlined"
                    value={this.state.fir.PlaceOfOccurence}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.PlaceOfOccurence = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Nature Of Offence"
                    variant="outlined"
                    value={this.state.fir.Nature}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.Nature = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Particulars Of Offence"
                    variant="outlined"
                    value={this.state.fir.Particulars}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.Particulars = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                />
                <TextField
                    className="inputs"
                    label="Description of Accused"
                    variant="outlined"
                    multiline
                    rows={8}
                    value={this.state.fir.DescriptionOfAccused}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.DescriptionOfAccused = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                ></TextField>
                <TextField
                    className="inputs"
                    label="Details Of Witness"
                    variant="outlined"
                    multiline
                    rows={8}
                    value={this.state.fir.DetailsOfWitness}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.DetailsOfWitness = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                ></TextField>
                <TextField
                    className="inputs"
                    label="Complaint"
                    variant="outlined"
                    multiline
                    rows={8}
                    value={this.state.fir.Complaint}
                    onChange={(event) => {
                        let fir = this.state.fir;
                        fir.Complaint = event.target.value;
                        this.setState({
                            fir,
                        });
                    }}
                ></TextField>
                <Button onClick={this.onAddFIR} variant="contained" color="primary">
                    File FIR
                </Button>
            </div>
        );
    }
}

export default App;
