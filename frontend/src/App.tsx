import "./styles/App.scss";
import { Route, HashRouter as Router, Routes } from "react-router-dom";
import { Home } from "./pages/home";
import { Nav } from "./components/nav";
import { Completed } from "./pages/completed";

function App() {
  return (
    <Router>
      <Nav />
      <div className="content">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/completed" element={<Completed />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
