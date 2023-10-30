import { useState } from 'react'
import Home from './views/Home';
import "./styles/App.scss";
import NavigationBar from './components/NavigationBar';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Me from './views/Me';
import Notification from './views/Notification';
import Group from './views/Group';

function App() {

  return (
    <BrowserRouter>
      <div className="app">
        <NavigationBar></NavigationBar>

        <Routes>
          <Route path="/" element={<Home/>} ></Route>
          <Route path="/me" exact element={<Me/>} ></Route>
          <Route path="/notification" exact element={<Notification/>} ></Route>
          <Route path="/group" exact element={<Group/>}></Route>
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default App
