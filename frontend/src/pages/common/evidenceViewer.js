import React, { Component } from "react";
import { TextField, Button, CircularProgress } from "@material-ui/core";

class App extends Component {
    state = {
        evidence: {},
        message: "",
        ID: "",
    };

    /*async componentDidMount() {
        const { id } = this.props.match.params;
        if (id !== "0") {
            this.setState({
                ID: id,
                message: (
                    <p>
                        Press <b>LOAD EVIDENCE</b> to View Evidence
                    </p>
                ),
            });
        }
    }*/

    onLoadInvestigation = async () => {
        console.log(this.state.evidence);
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

        let response = await fetch("http://192.168.1.30:3000/api/main/evidence/read/" + this.state.ID, requestOptions);
        let res = await response.json();
        console.log(res);

        this.setState({ evidence: res });

        var output = <div>{this.createContent()}</div>;

        this.setState({ message: output });
    };

    createContent = () => {
        return (
            <div>
                <h3>
                    Evidence ID:{" "}
                    <a target="blank" href={"https://ipfs.infura.io/ipfs/" + this.state.evidence.ID}>
                        {this.state.evidence.ID}
                    </a>
                </h3>
                <h3>Investigation ID: {this.state.evidence.InvestigationID}</h3>
                <h3>Date & Time: {this.state.evidence.DateTime}</h3>
                <h3>Mime Type: {this.state.evidence.MimeType}</h3>
                <TextField
                    className="inputs"
                    label="Description"
                    readOnly
                    multiline={true}
                    rows={8}
                    variant="outlined"
                    value={this.state.evidence.Description}
                />
            </div>
        );
    };

    render() {
        return (
            <div>
                <h2>New Investigation</h2>
                <TextField
                    className="inputs"
                    label="Evidence ID"
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
                    Load Evidence
                </Button>
                <br /> <br />
                {this.state.message}
            </div>
        );
    }
}

export default App;
