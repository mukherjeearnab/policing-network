import React from "react";
import { ThemeProvider, CssBaseline } from "@material-ui/core";
import Theme from "./Theme";
import "./App.css";

import { Route, Link } from "react-router-dom";
import Login from "./pages/login";
import HomeCi from "./pages/HomeCi";
import HomePo from "./pages/HomePo";
import HomeFo from "./pages/HomeFo";
import HomeCo from "./pages/HomeCo";
import HomeId from "./pages/HomeId";

import viewFIRsCi from "./pages/citizen/viewFIRs";
import newFIR from "./pages/citizen/newFIR";
import firViewer from "./pages/common/FIRviewer";
import submitEvidence from "./pages/common/submitEvidence";
import newInvestigation from "./pages/police/addInvestigation";
import viewInvestigation from "./pages/police/viewInvestigation";
import updateInvestigation from "./pages/police/updateInvestigation";
import addArrest from "./pages/police/addArrest";
import addiReport from "./pages/police/addiReport";
import viewEvidence from "./pages/common/evidenceViewer";
import newCitizen from "./pages/identityprovider/createCitizenProfile";
import updateCitizen from "./pages/identityprovider/updateCitizenProfile";
import viewChargesheet from "./pages/police/viewChargeSheet";
import newChargesheet from "./pages/police/addChargeSheet";
import addAccusedPerson from "./pages/police/addAccusedPerson";
import addBriefReport from "./pages/police/addBriefReport";
import addChargedPerson from "./pages/police/addChargedPerson";
import addFIRIDs from "./pages/police/addFIRIDs";
import addInvestigatingOfficer from "./pages/police/addInvestigatingOfficer";
import addInvestigationID from "./pages/police/addInvestigationID";
import addSectionOfLaw from "./pages/police/addSectionOfLaw";
import createJudgement from "./pages/court/createJudgement";
import viewJudgement from "./pages/court/viewJudgement";
import addEvidenceID from "./pages/court/addEvidenceID";
import addSentence from "./pages/court/addSentence";
import setComplete from "./pages/court/setComplete";
import viewProfile from "./pages/identityprovider/viewProfile";

function App() {
    return (
        <div className="App">
            <ThemeProvider theme={Theme}>
                <CssBaseline />
                <Link to="/">
                    <h1>Policing Platform</h1>
                </Link>
                <Route exact path="/" component={Login}></Route>
                <Route exact path="/HomeCi" component={HomeCi}></Route>
                <Route exact path="/HomePo" component={HomePo}></Route>
                <Route exact path="/HomeFo" component={HomeFo}></Route>
                <Route exact path="/HomeCo" component={HomeCo}></Route>
                <Route exact path="/HomeId" component={HomeId}></Route>
                <Route exact path="/viewFIRs" component={viewFIRsCi}></Route>
                <Route exact path="/newFIR" component={newFIR}></Route>
                <Route exact path="/firViewer/:id" component={firViewer}></Route>
                <Route exact path="/submitEvidence" component={submitEvidence}></Route>
                <Route exact path="/newInvestigation" component={newInvestigation}></Route>
                <Route exact path="/viewInvestigation/:id" component={viewInvestigation}></Route>
                <Route exact path="/updateInvestigation/:id" component={updateInvestigation}></Route>
                <Route exact path="/addArrest/:id" component={addArrest}></Route>
                <Route exact path="/addiReport/:id" component={addiReport}></Route>
                <Route exact path="/evidenceViewer" component={viewEvidence}></Route>
                <Route exact path="/newCitizen" component={newCitizen}></Route>
                <Route exact path="/updateCitizen" component={updateCitizen}></Route>
                <Route exact path="/newChargeSheet" component={newChargesheet}></Route>
                <Route exact path="/viewChargeSheet/:id" component={viewChargesheet}></Route>
                <Route exact path="/addFIRIDs/:id" component={addFIRIDs}></Route>
                <Route exact path="/addSectionOFLaw/:id" component={addSectionOfLaw}></Route>
                <Route exact path="/addInvestigatingOfficer/:id" component={addInvestigatingOfficer}></Route>
                <Route exact path="/addInvestigationID/:id" component={addInvestigationID}></Route>
                <Route exact path="/addAccusedPerson/:id" component={addAccusedPerson}></Route>
                <Route exact path="/addBriefReport/:id" component={addBriefReport}></Route>
                <Route exact path="/addChargedPerson/:id" component={addChargedPerson}></Route>
                <Route exact path="/createJudgement" component={createJudgement}></Route>
                <Route exact path="/viewJudgement/:id" component={viewJudgement}></Route>
                <Route exact path="/addEvidenceID/:id" component={addEvidenceID}></Route>
                <Route exact path="/addSentence/:id" component={addSentence}></Route>
                <Route exact path="/setComplete/:id" component={setComplete}></Route>
                <Route exact path="/viewProfile/:id" component={viewProfile}></Route>
            </ThemeProvider>
        </div>
    );
}

export default App;
