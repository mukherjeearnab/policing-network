import React, { Component } from "react";
import axios from "axios";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        ID: "",
        profile: {},
        message: "",
    };

    onAddCitizen = async () => {
        this.setState({
            message: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        let profile = this.state.profile;
        profile.DOB = new Date(this.state.profile.DOB).getTime().toString();
        this.setState({ profile });
        console.log(this.state.profile);

        // Create an object of formData
        const formData = new FormData();

        // Update the formData object
        formData.append("file", this.state.selectedFile, this.state.selectedFile.name);

        formData.append("payload", JSON.stringify(this.state.profile));

        let config = {
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };

        // Details of the uploaded file
        console.log(this.state.selectedFile);

        // Request made to the backend api
        // Send formData object
        var reply = await axios.post(
            "http://192.168.1.30:3000/api/main/citizen/update/" + this.state.profile.ID,
            formData,
            config
        );
        console.log(reply);
        this.setState({ message: "Citizen with ID " + reply.data.hash.ID + " Saved!" });
    };

    loadContent = async () => {
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

        let response = await fetch("http://192.168.1.30:3000/api/main/citizen/get/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        var d = new Date(parseInt(res.DOB));
        res.DOB = d.getFullYear() + "-";

        res.DOB += d.getMonth() + 1 < 10 ? "0" + (d.getMonth() + 1) : d.getMonth() + 1;
        res.DOB += "-" + d.getDate();

        this.setState({ profile: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    createContent = () => {
        return (
            <div>
                <TextField
                    variant="outlined"
                    type="file"
                    label="Citizen Photo"
                    onChange={(event) => {
                        this.setState({ selectedFile: event.target.files[0] });
                    }}
                />
                <br />
                <TextField
                    className="inputs"
                    label="Nationality"
                    variant="outlined"
                    value={this.state.profile.Nationality}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Nationality = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Name"
                    variant="outlined"
                    value={this.state.profile.Name}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Name = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    type="date"
                    label="Date Of Birth"
                    variant="outlined"
                    value={this.state.profile.DOB}
                    onChange={(event) => {
                        console.log(this.state.profile.DOB);
                        let profile = this.state.profile;
                        profile.DOB = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Eye Color"
                    variant="outlined"
                    value={this.state.profile.EyeColor}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.EyeColor = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Gender"
                    variant="outlined"
                    value={this.state.profile.Gender}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Gender = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Blood Group"
                    variant="outlined"
                    value={this.state.profile.BloodGroup}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.BloodGroup = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                />
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="E-Mail"
                    variant="outlined"
                    value={this.state.profile.Email}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Email = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Phone"
                    variant="outlined"
                    value={this.state.profile.Phone}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Phone = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Address"
                    variant="outlined"
                    multiline
                    rows={4}
                    value={this.state.profile.Address}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Address = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Father's Name"
                    variant="outlined"
                    value={this.state.profile.FathersName}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.FathersName = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Mother's Name"
                    variant="outlined"
                    value={this.state.profile.MothersName}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.MothersName = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Religion"
                    variant="outlined"
                    value={this.state.profile.Religion}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Religion = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <TextField
                    className="inputs"
                    label="Occupation"
                    variant="outlined"
                    value={this.state.profile.Occupation}
                    onChange={(event) => {
                        let profile = this.state.profile;
                        profile.Occupation = event.target.value;
                        this.setState({
                            profile,
                        });
                    }}
                ></TextField>
                <br />
                <br />
                <Button onClick={this.onAddCitizen} variant="contained" color="primary">
                    Update Citizen Profile
                </Button>
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>Update Citizen Profile</h2>
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
                ></TextField>
                <br />
                <br />
                <Button onClick={this.loadContent} variant="contained" color="primary">
                    Load Citizen Profile
                </Button>
                <hr />
                {this.state.message}
            </div>
        );
    }
}

export default App;
