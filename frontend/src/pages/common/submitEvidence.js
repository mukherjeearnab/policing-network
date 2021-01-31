import axios from "axios";
import { TextField, Button, CircularProgress } from "@material-ui/core";
import React, { Component } from "react";

class App extends Component {
    state = {
        // Initially, no file is selected
        selectedFile: null,
        description: "",
        InvestigationID: "",
        ID: "",
    };

    // On file select (from the pop up)
    onFileChange = (event) => {
        // Update the state
        this.setState({ selectedFile: event.target.files[0] });
    };

    onDescChange = (event) => {
        this.setState({ desc: event.target.value });
    };

    // On file upload (click the upload button)
    onFileUpload = async () => {
        this.setState({
            ID: (
                <span>
                    <CircularProgress />
                    <br></br> Loading.....
                </span>
            ),
        });

        // Create an object of formData
        const formData = new FormData();

        // Update the formData object
        formData.append("file", this.state.selectedFile, this.state.selectedFile.name);

        formData.append(
            "payload",
            JSON.stringify({ Description: this.state.description, InvestigationID: this.state.InvestigationID })
        );

        let config = {
            headers: { "Content-Type": "application/json", "x-access-token": localStorage.getItem("session") },
        };

        // Details of the uploaded file
        console.log(this.state.selectedFile);

        // Request made to the backend api
        // Send formData object
        var reply = await axios.post("http://192.168.1.30:3000/api/main/evidence/add", formData, config);
        console.log(reply);
        this.setState({ ID: "Evidence ID:" + reply.data.hash.ID });
    };

    // File content to be displayed after
    // file upload is complete
    fileData = () => {
        if (this.state.selectedFile) {
            return (
                <div>
                    <h2>File Details:</h2>
                    <p>File Name: {this.state.selectedFile.name}</p>
                    <p>File Type: {this.state.selectedFile.type}</p>
                    <h2>{this.state.ID}</h2>
                </div>
            );
        } else {
            return (
                <div>
                    <br />
                    <h4>Choose before Pressing the Upload button</h4>
                </div>
            );
        }
    };

    render() {
        return (
            <div>
                <h2>Submit Evidence</h2>
                <h1>{this.state.ID}</h1>
                <div>
                    <TextField variant="outlined" type="file" onChange={this.onFileChange} />
                    <br /> <br />
                    <TextField
                        label="Evidence Description"
                        variant="outlined"
                        type="text"
                        multiline
                        rows={8}
                        value={this.state.description}
                        onChange={(event) => {
                            this.setState({ description: event.target.value });
                        }}
                    />
                    <br /> <br />
                    <TextField
                        label="Investigation ID"
                        variant="outlined"
                        type="text"
                        value={this.state.InvestigationID}
                        onChange={(event) => {
                            this.setState({ InvestigationID: event.target.value });
                        }}
                    />
                    <br /> <br />
                    <Button variant="contained" color="primary" onClick={this.onFileUpload}>
                        Upload!
                    </Button>
                </div>
                {this.fileData()}
            </div>
        );
    }
}

export default App;
