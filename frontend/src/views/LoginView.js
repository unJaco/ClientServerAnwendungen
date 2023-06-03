import React, { useState } from 'react';
import '../styles/LoginView.css';
import Button from '../components/Button';
import UserService from '../services/UserService';
import RideService from '../services/RideService';
import User from '../model/user';
import LandingPage from './LandingPage';

function LoginView({ onLogin, declareDriver }) {
  const [loggedIn, setLoggedIn] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await UserService.login(username, password);
    const drivenRides = await RideService.getRideByCustomerID(response.id);
    const takenRides = await RideService.getRideByDriverID(response.id);

    console.log(response)

    if (response instanceof User) {
      setLoggedIn(true);
      onLogin(response, drivenRides, takenRides);
      if (UserService.getDriverById(response.id) instanceof User) {
        declareDriver();
      }
    } else {
      console.log('incorrect username or password');
    }    
  };

if (loggedIn) {
      return <LandingPage/>
    }

  return (
    <form className="login-view" onSubmit={handleSubmit}>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="username-input">Username:</label>
        <input
          id="username-input"
          type="text"
          value={username}
          onChange={(event) => setUsername(event.target.value)}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="password-input">Password:</label>
        <input
          id="password-input"
          type="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column', marginTop: "10px" }}>
        <Button text="Log In" onClick={handleSubmit}></Button>
      </div>

    </form>
  );
}

export default LoginView;