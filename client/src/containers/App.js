import React from 'react';
import Navbar from "./Navbar";
import Main from "./Main";
import { Provider } from "react-redux";
import { configureStore } from "../store/index";
import { BrowserRouter as Router } from "react-router-dom";
import { setAuthToken, setCurrentUser } from '../store/actions/auth';
import jwtDecode from "jwt-decode";

const store = configureStore();

if (localStorage.jwtToken) {
  setAuthToken(localStorage.jwtToken);
  // prevent someone from manually tampering with the key of jwt Token in localStorage
  try {
    store.dispatch(setCurrentUser(jwtDecode(localStorage.jwtToken)))
  } catch (err) {
    store.dispatch(setCurrentUser({}));
  }
}

const App = () => (
  <Provider store={store}>
    <Router>
       <div>
         <Navbar />
         <Main />
       </div>
    </Router>
  </Provider>
);

export default App;
