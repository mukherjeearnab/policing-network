import React, { Component } from "react";
import { withStyles } from "@material-ui/core/styles";
import { Button, Grid, Container, Paper } from "@material-ui/core";
import { Link, Redirect } from "react-router-dom";

const useStyles = (theme) => ({
    root: {
        flexGrow: 1,
    },
    paper: {
        padding: theme.spacing(2),
        textAlign: "center",
        color: theme.palette.common.white,
        backgroundColor: theme.palette.primary.light,
    },
    paper2: {
        padding: theme.spacing(2),
        textAlign: "center",
        color: theme.palette.common.white,
        backgroundColor: theme.palette.secondary.light,
    },
    paper3: {
        padding: theme.spacing(2),
        textAlign: "center",
        color: theme.palette.text.primary,
        backgroundColor: theme.palette.common.white,
    },
    link: {
        textDecoration: "none",
    },
    logout: {
        backgroundColor: theme.palette.error.light,
    },
});

class App extends Component {
    state = { redirect: "" };

    logout = () => {
        localStorage.removeItem("session");
        localStorage.removeItem("user");
        this.setState({ redirect: <Redirect to="/" /> });
    };

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
                <h2>Police Dashboard</h2>
                <h2>
                    {this.state.redirect}Welcome, {localStorage.getItem("user")}!
                </h2>
                <Container maxWidth="sm" spacing={10}>
                    <Grid container spacing={3}>
                        <Grid item xs>
                            <Link to={"/viewProfile/0/"} className={classes.link}>
                                <Paper className={classes.paper}>Check Citizen Profile</Paper>
                            </Link>
                        </Grid>
                        <Grid item xs>
                            <Link to={"/viewFIRs/"} className={classes.link}>
                                <Paper className={classes.paper}>View FIRs</Paper>
                            </Link>
                        </Grid>
                    </Grid>
                    <Grid container spacing={3}>
                        <Grid item xs>
                            <Link to="/viewInvestigation/0/" className={classes.link}>
                                <Paper className={classes.paper2}>View / Update Investigation</Paper>
                            </Link>
                        </Grid>
                        <Grid item xs>
                            <Link to="/newChargesheet/" className={classes.link}>
                                <Paper className={classes.paper2}>File New Charge-Sheet</Paper>
                            </Link>
                        </Grid>
                    </Grid>
                    <Grid container spacing={3}>
                        <Grid item xs>
                            <Link to="/submitEvidence" className={classes.link}>
                                <Paper className={classes.paper2}>Submit Evidence</Paper>
                            </Link>
                        </Grid>
                        <Grid item xs>
                            <Link to="/viewChargeSheet/0" className={classes.link}>
                                <Paper className={classes.paper2}>View / Update Charge-Sheet</Paper>
                            </Link>
                        </Grid>
                    </Grid>
                    <Grid container spacing={3}>
                        <Grid item xs>
                            <Button m={1} onClick={this.logout} variant="contained" className={classes.logout}>
                                Log Out
                            </Button>
                        </Grid>
                    </Grid>
                    <Grid container spacing={3} style={{ display: "none" }}>
                        <Grid item xs>
                            <Paper className={classes.paper3}></Paper>
                        </Grid>
                    </Grid>
                </Container>
            </div>
        );
    }
}

export default withStyles(useStyles)(App);
