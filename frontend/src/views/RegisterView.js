import React, { useState } from 'react';
import BookingPage from './BookingPage';
import '../styles/RegisterView.css';
import Button from '../components/Button';
import UserService from '../services/UserService';
import User from '../model/user';
import LandingPage from './LandingPage';

function RegisterView({ onRegister }) {
  const [firstname, setFirstname] = useState('');
  const [lastname, setLastname] = useState('');
  const [telephoneNumber, setTelephoneNumber] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [registered, setRegistered] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (firstname === "" || lastname === "" || email === "" || telephoneNumber === "" || password === "") {
      alert('Please fill in all fields')
    } else {
      if (password === confirmPassword) {
        console.log('Submitting registration:', { firstname, lastname, email, password });
        const response = await UserService.register(new User(1, firstname, lastname, "", telephoneNumber, email, password));
        console.log(response);
        if (response instanceof User) {
          setRegistered(true);
          onRegister(response, [], [])
        }else {
          alert('Registration failed');
        }
      }else {
        alert("Password not the same!")
        return <BookingPage />;
      }
    }


  };

  if (registered) {
    return <LandingPage />;
  }

  return (
    <form className='register-view'>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="firstname-input">Firstname:</label>
        <input
          id="firstname-input"
          type="text"
          value={firstname}
          onChange={(event) => setFirstname(event.target.value)}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="lastname-input">Lastname:</label>
        <input
          id="lastname-input"
          type="text"
          value={lastname}
          onChange={(event) => setLastname(event.target.value)}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="telephoneNumber-input">Telephone Number:</label>
        <input
          id="telephoneNumber-input"
          type="long"
          value={telephoneNumber}
          onChange={(event) => setTelephoneNumber(event.target.value)}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column' }}>
        <label htmlFor="email-input">Email:</label>
        <input
          id="email-input"
          type="email"
          value={email}
          onChange={(event) => setEmail(event.target.value)}
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
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}>
        <label htmlFor="confirm-password-input">Confirm Password:</label>
        <input
          id="confirm-password-input"
          type="password"
          value={confirmPassword}
          onChange={(event) => setConfirmPassword(event.target.value)}
        />
      </div>
      <Button text="Register" onClick={handleSubmit}></Button>
    </form>
  );

}

export default RegisterView;
