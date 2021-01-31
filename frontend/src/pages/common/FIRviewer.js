import React, { Component } from "react";
import { Redirect } from "react-router-dom";

import { TextField } from "@material-ui/core";

class App extends Component {
    state = {
        fir: {
            Offence: {},
        },
    };

    async componentDidMount() {
        const { id } = this.props.match.params;

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
            method: "GET",
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };
        //console.log(requestOptions);
        let response = await fetch("http://192.168.1.30:3000/api/main/fir/read/" + id, requestOptions);
        let res = await response.json();
        console.log(res);
        this.setState({ fir: res });
    }

    render() {
        return (
            <div>
                <h2>FIR Viewer</h2>
                <h3>FIR ID - {this.state.fir.ID}</h3>
                <p>Citizen ID - {this.state.fir.CitizenID}</p>
                <p>Police Station - {this.state.fir.PoliceStation}</p>
                <p>Date & Time - {this.state.fir.DateHour}</p>
                <p>District - {this.state.fir.District}</p>
                <p>Place Of Occurence - {this.state.fir.PlaceOfOccurence}</p>
                <p>Nature of Offence - {this.state.fir.Offence.Nature}</p>
                <p>Particulars of Offence - {this.state.fir.Offence.Particulars}</p>
                <TextField
                    className="inputs"
                    label="Description of Accused"
                    variant="outlined"
                    multiline
                    rows={8}
                    readOnly
                    value={this.state.fir.DescriptionOfAccused}
                ></TextField>
                <TextField
                    className="inputs"
                    label="Details of Witness"
                    variant="outlined"
                    multiline
                    rows={8}
                    readOnly
                    value={this.state.fir.DetailsOfWitness}
                ></TextField>
                <TextField
                    className="inputs"
                    label="Complaint"
                    variant="outlined"
                    multiline
                    rows={8}
                    readOnly
                    InputProps={{ readOnly: true }}
                    defaultValue={this.state.fir.Complaint}
                ></TextField>
                <p>
                    Investigation ID -{" "}
                    {this.state.fir.InvestigationID === ""
                        ? "Investigation NOT started!"
                        : this.state.fir.InvestigationID}
                </p>
            </div>
        );
    }
}

export default App;
