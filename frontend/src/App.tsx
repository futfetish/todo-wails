import "./styles/App.scss";
import { GetTodos, AddTodo } from "../wailsjs/go/main/App";
import { useEffect } from "react";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { Home } from "./pages/home";
import { Nav } from "./components/nav";
import { Completed } from "./pages/completed";

function App() {
  useEffect(() => {
    GetTodos(null)
      .then((data) => console.log(111, data[0]))
      .catch(console.error);
  }, []);

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
