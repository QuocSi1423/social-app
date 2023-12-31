import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import { Provider } from 'react-redux'
import store from './redux/store.js'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Login from './Login.jsx'
import Register from './Register.jsx'

ReactDOM.createRoot(document.getElementById('root')).render(
  <Provider store={ store } >
    
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={ <Login /> }></Route>
        <Route exact path='/*' element={ <App /> }></Route>
        <Route path="/register" element={<Register/>}></Route>
      </Routes>
    </BrowserRouter>
  </Provider>,
)
