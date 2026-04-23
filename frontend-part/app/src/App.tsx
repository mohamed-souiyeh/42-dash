import { Routes, Route } from 'react-router-dom'
import Games from './pages/games/main'
import { Link } from "react-router-dom";
import Status from './pages/status/main';
import About from './pages/about/main';

function NavBar() {
  return (
    <nav className="flex items-center justify-center gap-4 p-4 text-2xl font-mono">

      <Link className="hover:underline" to="/">
        Games
      </Link>

      <Link className="hover:underline" to="/status">
        Health
      </Link>

      <Link className="hover:underline" to="/about">
        About
      </Link>
    </nav>
  );
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<>
        <NavBar /><Games />
      </>} />
      <Route path="/status" element={<>
        <NavBar /><Status />
      </>} />
      <Route path="/about" element={<>
        <NavBar /><About />
      </>} />
    </Routes>
  )
}

export default App