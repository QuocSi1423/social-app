import { useEffect, useState } from 'react'
import Home from './views/Home';
import "./styles/App.scss";
import NavigationBar from './components/NavigationBar';
import { Routes, Route } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux/es/hooks/useSelector';
import Me from './views/Me';
import Notification from './views/Notification';
import Group from './views/Group';
import Loading from './components/Loading';
import { getBriefUserInformation } from './service/service-user';

function App() {

  
  return (
      <div className="app">
        <NavigationBar></NavigationBar>

        <Routes>
          <Route path="/" element={<Home/>} ></Route>
          <Route path="/me" exact element={<Me/>} ></Route>
          <Route path="/notification" exact element={<Notification/>} ></Route>
          <Route path="/group" exact element={<Group/>}></Route>
      </Routes>
      
      </div>
  )
}

export default App
