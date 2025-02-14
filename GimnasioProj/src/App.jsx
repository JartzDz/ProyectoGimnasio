import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './assets/pages/loginPage'
import Register from './assets/pages/registerPage'
function App() {
  
  return (
    <>
      <Router>
        <Routes>
          <Route path='/' element = {<Login/>} />
          <Route path='/Registro' element = {<Register/>} />
        </Routes>
      </Router>
    </>
  )
}

export default App
