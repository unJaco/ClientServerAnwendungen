import React, { useState } from 'react';
import LandingPage from './views/LandingPage';
import LoginView from './views/LoginView';
import RegisterView from './views/RegisterView';
import './App.css';
import User from './model/user';
import UserView from './views/UserView';
import BookingPage from './views/BookingPage';
import RideView from './views/RideView';
import RideService from './services/RideService';

function App() {
  const [currentPage, setCurrentPage] = useState('landing');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState(new User());
  const [takenRides, setTakenRides] = useState([]);
  const [drivenRides, setDrivenRides] = useState([]);



  function handlePageChange(page) {
    setCurrentPage(page);
  }

  function handleLogin(activeUser, takenRidesUser, drivenRidesUser) {
    setIsLoggedIn(true);
    setUser(activeUser);
    setTakenRides(takenRidesUser);
    setDrivenRides(drivenRidesUser);
  }

  let page;
  if (currentPage === 'landing') {
    page = <LandingPage onPageChange={handlePageChange} />;
  } else if (currentPage === 'login') {
    page = <LoginView onLogin={handleLogin} onPageChange={handlePageChange} />;
  } else if (currentPage === 'register') {
    page = <RegisterView onRegister={handleLogin} onPageChange={handlePageChange} />;
  } else if (currentPage === 'booking') {
    page = <BookingPage activeUser={user} onPageChange={handlePageChange} />;
  } else if (currentPage === 'logout') {
    if (isLoggedIn) {
      setIsLoggedIn(false);
      setUser(new User());
    }
    page = <LandingPage onPageChange={handlePageChange} />;
  } else if (currentPage === 'profile') {
      page = <UserView activeUser={user} onRegister={handleLogin} onPageChange={handlePageChange} />;
  }else if (currentPage === 'rides') {
    page = <RideView takenRides={takenRides} drivenRides={drivenRides} onPageChange={handlePageChange} />;
}

  return (
    <div>
      <nav>
        <ul className='navigation-bar' style={{
          display: 'flex',
          flexDirection: 'row',
          listStyle: 'none',
          position: 'absolute',
        }}>
          <li style={{ marginRight: '1rem' }}>
            <button onClick={() => handlePageChange('landing')}>Home</button>
          </li>
          <li style={{ marginRight: '1rem' }}>
            <button disabled={isLoggedIn} onClick={() => handlePageChange('login')}>Log In</button>
          </li>
          <li style={{ marginRight: '1rem' }}>
            <button disabled={isLoggedIn} onClick={() => handlePageChange('register')}>Register</button>
          </li>
          <li style={{ marginRight: '1rem' }}>
            <button disabled={!isLoggedIn} onClick={() => handlePageChange('booking')}>Booking</button>
          </li>
          <li style={{ marginRight: '1rem' }}>
            <button disabled={!isLoggedIn} onClick={() => handlePageChange('rides')}>Rides</button>
          </li>
          <li style={{ marginRight: '1rem' }}>
            <button disabled={!isLoggedIn} onClick={() => handlePageChange('profile')}>Profile</button>
          </li>
          <li>
            <button disabled={!isLoggedIn} onClick={() => handlePageChange('logout')}>Log Out</button>
          </li>
        </ul>
      </nav>
      {page}
    </div>
  );
}

export default App;