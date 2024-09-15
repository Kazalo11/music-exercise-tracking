import "./App.css";
import { Heading } from "@chakra-ui/react";
import LoginPage from "./components/login/LoginPage";
import { MainPage } from "./components/MainPage";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Heading>Tracking your music listened during Strava</Heading>
      <Router>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/" element={<MainPage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
